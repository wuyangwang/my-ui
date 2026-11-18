package controls

import (
	. "github.com/wuyangwang/my-ui"
	testutil "github.com/wuyangwang/my-ui/internal/testutil"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestVariantSecondaryIsAMutedFill(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewButton(c, "Save").Secondary()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	label, ok := tt.Find("Save")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	x, y := label.X-6, label.Y+label.H/2
	r, g, b := testutil.Pixel(tt, x, y)
	// The muted fill is the theme's surface tone.
	testutil.SamePixel(t, r, g, b, ui.Hex("#f4f4f5"), "the secondary fill")
}

func TestSemanticVariantsPaintTheirColor(t *testing.T) {
	success := SemanticColor(ui.LightTheme(), SemanticSuccess)
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewButton(c, "Save").Success()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	label, ok := tt.Find("Save")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	r, g, b := testutil.Pixel(tt, label.X-6, label.Y+label.H/2)
	testutil.SamePixel(t, r, g, b, success, "the success fill")
}

func TestShapeRoundsTheCorners(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewButton(c, "Save").Shape(ShapeRound)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	label, ok := tt.Find("Save")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	// A pill's corner is rounded all the way: the pixel at the corner of
	// the button is the background, not the fill.
	button, ok := tt.Access("Save")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	x := button.Bounds.X + 1
	y := button.Bounds.Y + 1
	r, g, b := testutil.Pixel(tt, x, y)
	if r == 0xfa && g == 0xfa && b == 0xfa {
		t.Error("the corner of the button is not rounded")
	}
	_ = label
}

func TestSizeXsIsSmaller(t *testing.T) {
	height := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds.H
	}
	var a, b string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewButton(c, "Small").Xs().Label("xs")
			NewButton(c, "Middle").Label("md")
		})
	}
	_ = a
	_ = b
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if h := height(tt, "xs"); h >= height(tt, "md") {
		t.Errorf("the smallest button is not shorter: %v against %v", h, height(tt, "md"))
	}
}
