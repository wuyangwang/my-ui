package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestSpinSizesAndVariants(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(4)).Children(func() {
			NewSpin(c)
			NewSpin(c).Size(SizeSm).Variant(Success)
			NewSpin(c).Size(SizeLg).Variant(Destructive)
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	// A spinner draws under the framework's motion, without timers in
	// the component: this is a smoke test of the three presets.
	painted := false
	for y := 0; y < 40 && !painted; y++ {
		for x := 0; x < 120; x++ {
			r, g, b := pixelOfDisplay(tt, x, y)
			if !(r == 255 && g == 255 && b == 255) {
				painted = true
				break
			}
		}
	}
	if !painted {
		t.Error("no spinner was drawn")
	}
}
