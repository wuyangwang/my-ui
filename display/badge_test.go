package display

import (
	. "github.com/wuyangwang/my-ui"
	testutil "github.com/wuyangwang/my-ui/internal/testutil"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestBadgeVariantsPaint(t *testing.T) {
	accent, _, _ := Palettes[PaletteBlue].Triplet(false)
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight, Palette: PaletteBlue})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewBadge(c, "Primary").Primary()
			NewBadge(c, "Outline").Outline()
			NewBadge(c, "Danger").Destructive()
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)

	p, ok := tt.Find("Primary")
	if !ok {
		t.Fatal("primary badge missing")
	}
	r, g, b := testutil.Pixel(tt, p.X-2, p.Y+p.H/2)
	testutil.SamePixel(t, r, g, b, accent, "primary badge background")

	d, ok := tt.Find("Danger")
	if !ok {
		t.Fatal("destructive badge missing")
	}
	r, g, b = testutil.Pixel(tt, d.X-2, d.Y+d.H/2)
	testutil.SamePixel(t, r, g, b, ui.Hex("#dc2626"), "destructive badge background")
}

func TestBadgeSizes(t *testing.T) {
	width := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s missing", name)
		}
		return info.Bounds.W
	}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewBadge(c, "99").Sm().Label("small")
			NewBadge(c, "99").Md().Label("medium")
			NewBadge(c, "99").Lg().Label("large")
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if !(width(tt, "small") < width(tt, "medium") && width(tt, "medium") < width(tt, "large")) {
		t.Errorf("badge sizes not increasing: %v %v %v", width(tt, "small"), width(tt, "medium"), width(tt, "large"))
	}
}

func TestBadgeDarkTheme(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark, Palette: PaletteRose})
		NewBadge(c, "Rose").Primary()
	}
	tt := ui.NewTester(view, 200, 80)
	tt.Advance(time.Second)
	if _, ok := tt.Access("Rose"); !ok {
		t.Fatal("badge missing in dark theme")
	}
}
