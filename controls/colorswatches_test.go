package controls

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// TestColorSwatchesChoosesAndShowsChecks checks the binding, the
// accessible names and the check on the chosen swatch.
func TestColorSwatchesChoosesAndShowsChecks(t *testing.T) {
	red, green := ui.Hex("#ff0000"), ui.Hex("#00ff00")
	value := red
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewColorSwatches(c, &value,
			ColorSwatch{Color: red, Label: "Red"},
			ColorSwatch{Color: green, Label: "Green"},
		)
	}
	tt := ui.NewTester(view, 320, 200)
	tt.Frame()
	for _, name := range []string{"Red", "Green"} {
		if _, ok := tt.Access(name); !ok {
			t.Errorf("%s is not on screen", name)
		}
	}
	if err := tt.Click("Green"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if value != green {
		t.Errorf("clicking Green left the value at %v", value)
	}
}

// TestColorSwatchesSquareSizes checks that Square and Size take effect
// on the swatches' bounds.
func TestColorSwatchesSquareSizes(t *testing.T) {
	built := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		v := ui.Transparent
		sw := NewColorSwatches(c, &v, ColorSwatch{Color: ui.Hex("#0000ff"), Label: "Blue"}).
			Square(true).Size(32)
		built = sw != nil
	}
	tt := ui.NewTester(view, 200, 160)
	tt.Frame()
	info, ok := tt.Access("Blue")
	if !ok {
		t.Fatal("Blue is not on screen")
	}
	if info.Bounds.W != 32 || info.Bounds.H != 32 {
		t.Errorf("the square swatch is %vx%v, want 32", info.Bounds.W, info.Bounds.H)
	}
	if !built {
		t.Fatal("the swatches were not built")
	}
}
