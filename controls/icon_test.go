package controls

import (
	. "github.com/wuyangwang/my-ui"
	testutil "github.com/wuyangwang/my-ui/internal/testutil"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// testIcon is a filled square that takes the color of the text around
// it, as an icon of the set does.
var testIcon = ui.MustParseSVG([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 4h16v16H4z" fill="currentColor"/></svg>`))

func TestIconSizes(t *testing.T) {
	bounds := func(tt *ui.Tester, name string) ui.Rect {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds
	}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewIcon(c, testIcon).Sm().Label("small")
			NewIcon(c, testIcon).Md().Label("medium")
			NewIcon(c, testIcon).Lg().Label("large")
		})
	}
	tt := ui.NewTester(view, 240, 120)
	tt.Advance(time.Second)
	if w := bounds(tt, "small").W; w != 14 {
		t.Errorf("the small icon is %v wide, want 14", w)
	}
	if w := bounds(tt, "medium").W; w != 20 {
		t.Errorf("the medium icon is %v wide, want 20", w)
	}
	if w := bounds(tt, "large").W; w != 24 {
		t.Errorf("the large icon is %v wide, want 24", w)
	}
}

func TestIconVariantPaintsItsColor(t *testing.T) {
	accent, _, _ := Palettes[PaletteBlue].Triplet(false)
	danger := ui.Hex("#dc2626")
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight, Palette: PaletteBlue})
		ui.Row(c).Gap(c.Theme().Space(4)).Children(func() {
			NewIcon(c, testIcon).Primary().Md().Label("accent")
			NewIcon(c, testIcon).Destructive().Md().Label("danger")
		})
	}
	tt := ui.NewTester(view, 240, 120)
	tt.Advance(time.Second)

	accentIcon, ok := tt.Access("accent")
	if !ok {
		t.Fatal("the accent icon is not on screen")
	}
	r, g, b := testutil.Pixel(tt, accentIcon.Bounds.X+accentIcon.Bounds.W/2, accentIcon.Bounds.Y+accentIcon.Bounds.H/2)
	testutil.SamePixel(t, r, g, b, accent, "the primary icon")

	dangerIcon, ok := tt.Access("danger")
	if !ok {
		t.Fatal("the danger icon is not on screen")
	}
	r, g, b = testutil.Pixel(tt, dangerIcon.Bounds.X+dangerIcon.Bounds.W/2, dangerIcon.Bounds.Y+dangerIcon.Bounds.H/2)
	testutil.SamePixel(t, r, g, b, danger, "the destructive icon")
}

func TestIconDarkTheme(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewIcon(c, testIcon).Md().Label("moon")
	}
	tt := ui.NewTester(view, 120, 120)
	tt.Advance(time.Second)
	if _, ok := tt.Access("moon"); !ok {
		t.Fatal("the icon is not on screen in the dark theme")
	}
}
