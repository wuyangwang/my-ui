package chart

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	mchart "github.com/egoist/mygo/ui/chart"
	myui "github.com/wuyangwang/my-ui"
	testutil "github.com/wuyangwang/my-ui/internal/testutil"
)

func TestNewPaletteFollowsTheme(t *testing.T) {
	th := myui.ShadcnLight()
	th.Accent = ui.Hex("#123456")
	p := NewPalette(th)
	if p.Series[0] != th.Accent {
		t.Errorf("first series color is %v, want the theme's accent %v", p.Series[0], th.Accent)
	}
	if p.Grid != th.Border || p.Axis != th.TextMuted || p.Surface != th.Surface {
		t.Error("grid, axis and surface colors do not come from the theme")
	}
	if p.Color(0) != th.Accent || p.Color(len(p.Series)) != th.Accent || p.Color(-1) != p.Series[len(p.Series)-1] {
		t.Error("Color does not cycle through the series")
	}
	if got := NewPalette(nil).Series; len(got) == 0 {
		t.Error("a nil theme should still give a usable palette")
	}
}

func TestPresetsPaint(t *testing.T) {
	view := func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
		values := []float64{3, 7, 5, 9}
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewBars(c, values)
			NewLine(c, values)
			NewArea(c, values)
			NewPie(c, values)
		})
	}
	tt := ui.NewTester(view, 500, 900)
	tt.Advance(time.Second)
	painted := false
	for y := 0; y < 850 && !painted; y++ {
		for x := 0; x < 450; x++ {
			r, g, b := testutil.Pixel(tt, float32(x), float32(y))
			if !(r == 255 && g == 255 && b == 255) {
				painted = true
				break
			}
		}
	}
	if !painted {
		t.Error("no chart painted anything")
	}
}

func TestEmptyPresets(t *testing.T) {
	view := func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceDark})
		NewBars(c, nil)
		NewLine(c, []float64{})
		NewArea(c, nil)
		NewPie(c, nil)
	}
	tt := ui.NewTester(view, 400, 400)
	tt.Frame()
}

func TestLabelsAndCurve(t *testing.T) {
	view := func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			Bars(c, []float64{1, 2, 3}, Options{Labels: []string{"甲", "乙", "丙"}})
			Line(c, []float64{1, 2, 3}, Options{Curve: mchart.CurveMonotone, Label: "趋势"})
		})
	}
	tt := ui.NewTester(view, 500, 500)
	for _, s := range []string{"甲", "乙", "丙", "趋势"} {
		if !tt.HasText(s) {
			t.Errorf("no %q in %q", s, tt.Texts())
		}
	}
}

func TestSelectionPointer(t *testing.T) {
	selected := 99
	view := func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
		Bars(c, []float64{2, 4, 6, 8}, Options{
			Labels:   []string{"a", "b", "c", "d"},
			Selected: &selected,
		})
	}
	tt := ui.NewTester(view, 600, 400)
	if selected != -1 {
		t.Fatalf("an out-of-range selection was not reset: %d", selected)
	}
	r, ok := tt.Find("b")
	if !ok {
		t.Fatal("no bar label b")
	}
	// The plot area sits above the bottom axis label; click inside it to
	// focus the area, then move the choice with the keyboard.
	tt.ClickAt(r.X+r.W/2, r.Y-40)
	tt.Key(0, ui.KeyRight)
	if selected < 0 || selected >= 4 {
		t.Errorf("the keyboard did not move the selection: %d", selected)
	}
}

func TestThemeSwitchesNextFrame(t *testing.T) {
	appearance := myui.AppearanceLight
	var got ui.Color
	view := func(c *ui.Context) {
		th := myui.Apply(c, myui.Options{Appearance: appearance})
		got = NewPalette(th).Series[0]
		NewBars(c, []float64{1, 2})
	}
	tt := ui.NewTester(view, 400, 300)
	light := got
	appearance = myui.AppearanceDark
	tt.Frame()
	dark := got
	if light == dark {
		t.Error("the palette did not follow the theme on the next frame")
	}
}
