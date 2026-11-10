package myui

import (
	"image/color"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// pixel reads what the tester drew at a DIP of the window, as bytes.
func pixel(tt *ui.Tester, x, y float32) (r, g, b uint8) {
	p := color.NRGBAModel.Convert(tt.Image().At(int(x), int(y))).(color.NRGBA)
	return p.R, p.G, p.B
}

// samePixel fails if the tester drew not the wanted color.
func samePixel(t *testing.T, r, g, b uint8, want ui.Color, what string) {
	t.Helper()
	if r != want.R || g != want.G || b != want.B {
		t.Errorf("%s: drew (%d, %d, %d), want %v", what, r, g, b, want)
	}
}

func TestButtonPrimaryPaintsAndPressesTheAccent(t *testing.T) {
	accent, hover, pressed := palettes[PaletteBlue].Triplet(false)
	clicked := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight, Palette: PaletteBlue})
		if NewButton(c, "Save").Primary().Clicked() {
			clicked = true
		}
	}
	tt := ui.NewTester(view, 320, 240)
	label, ok := tt.Find("Save")
	if !ok {
		t.Fatal("the label is not on screen")
	}
	// A padding pixel beside the label, off any corner: the button
	// fills with the accent.
	x, y := label.X-6, label.Y+label.H/2
	tt.Advance(time.Second)
	r, g, b := pixel(tt, x, y)
	samePixel(t, r, g, b, accent, "the primary at rest")

	// Held down, it paints the pressed shade; released, the hover one
	// stays while the pointer does.
	tt.Press(x, y)
	tt.Advance(time.Second)
	r, g, b = pixel(tt, x, y)
	samePixel(t, r, g, b, pressed, "the primary pressed")
	tt.Release(x, y)
	tt.Advance(time.Second)
	r, g, b = pixel(tt, x, y)
	if !clicked {
		t.Error("the click did not report")
	}
	samePixel(t, r, g, b, hover, "the primary hovered after the click")
	tt.Move(2, 200)
	tt.Advance(time.Second)
	r, g, b = pixel(tt, x, y)
	samePixel(t, r, g, b, accent, "the primary at rest once the pointer leaves")
}

func TestButtonGhostIsBareAndHovers(t *testing.T) {
	hoverGhost := ui.Hex("#e9e9ec")
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewButton(c, "Cancel").Ghost()
	}
	tt := ui.NewTester(view, 320, 240)
	label, ok := tt.Find("Cancel")
	if !ok {
		t.Fatal("the label is not on screen")
	}
	x, y := label.X-6, label.Y+label.H/2
	tt.Advance(time.Second)
	r, g, b := pixel(tt, x, y)
	samePixel(t, r, g, b, ui.Hex("#ffffff"), "the ghost at rest, the background")
	tt.Move(x, y)
	tt.Advance(time.Second)
	r, g, b = pixel(tt, x, y)
	samePixel(t, r, g, b, hoverGhost, "the ghost hovered")
}

func TestButtonDestructivePaintsTheDanger(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewButton(c, "Delete").Destructive().Sm()
	}
	tt := ui.NewTester(view, 320, 240)
	label, ok := tt.Find("Delete")
	if !ok {
		t.Fatal("the label is not on screen")
	}
	x, y := label.X-4, label.Y+label.H/2
	tt.Advance(time.Second)
	r, g, b := pixel(tt, x, y)
	samePixel(t, r, g, b, ui.Hex("#dc2626"), "the destructive at rest")

	// Held down, Shades' pressed shades underline the pressed pixel;
	// released, the hover one holds while the pointer does.
	hover, pressed, _ := shades(ui.Hex("#dc2626"), false)
	tt.Press(x, y)
	tt.Advance(time.Second)
	r, g, b = pixel(tt, x, y)
	samePixel(t, r, g, b, pressed, "the destructive pressed")

	tt.Release(x, y)
	tt.Advance(time.Second)
	r, g, b = pixel(tt, x, y)
	samePixel(t, r, g, b, hover, "the destructive hovered after the release")
}

func TestButtonSizesGiveRoomierBounds(t *testing.T) {
	width := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds.W
	}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark, Palette: PaletteRose})
		NewButton(c, "Alpha").Sm()
		NewButton(c, "Beta").Md()
		NewButton(c, "Gamma").Lg()
	}
	tt := ui.NewTester(view, 480, 240)
	if w := (width(tt, "Alpha")); w >= (width(tt, "Beta")) {
		t.Errorf("the small button is not smaller than the middle's: %v against %v", w, width(tt, "Beta"))
	}
	if w := (width(tt, "Beta")); w >= (width(tt, "Gamma")) {
		t.Errorf("the middle button is not smaller than the large's: %v against %v", w, width(tt, "Gamma"))
	}
}

func TestButtonDisabledTakesNoClick(t *testing.T) {
	clicked := false
	view := func(c *ui.Context) {
		Apply(c, Options{})
		if NewButton(c, "Locked").Primary().Disabled(true).Clicked() {
			clicked = true
		}
	}
	tt := ui.NewTester(view, 320, 240)
	info, ok := tt.Access("Locked")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	if !info.Disabled {
		t.Error("the button does not read as disabled")
	}
	if err := tt.Click("Locked"); err != nil {
		t.Fatal(err)
	}
	if clicked {
		t.Error("a disabled button reported a click")
	}
}

func TestButtonIconSitsBeforeTheText(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewButton(c, "Save").Icon(testIcon).Primary()
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	button, ok := tt.Access("Save")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	text, ok := tt.Find("Save")
	if !ok {
		t.Fatal("the text of the button is not on screen")
	}
	if text.X <= button.Bounds.X {
		t.Errorf("the text sits at %v, the button at %v: the icon is not before it", text.X, button.Bounds.X)
	}
}

func TestButtonWithOnlyAnIconIsNamed(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewButton(c, "").Icon(testIcon).Label("保存")
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	info, ok := tt.Access("保存")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	if info.Name != "保存" {
		t.Errorf("the button reads %q, want %q", info.Name, "保存")
	}
}

func TestButtonLoadingTakesNoClick(t *testing.T) {
	clicked := false
	view := func(c *ui.Context) {
		Apply(c, Options{})
		if NewButton(c, "Save").Loading().Clicked() {
			clicked = true
		}
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if err := tt.Click("Save"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if clicked {
		t.Error("a busy button reported a click")
	}
}
