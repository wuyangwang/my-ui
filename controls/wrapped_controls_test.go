package controls

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestWrappedMygoControlsBuild(t *testing.T) {
	low, high, step, value := 20.0, 80.0, 3.0, 5.0
	rating := 3
	var levels ui.MeterLevels
	color := ui.Hex("#ff0000")
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(3)).Children(func() {
			NewRangeSlider(c, &low, &high, 0, 100, 1)
			NewStepSlider(c, &value, 0, 10, 1)
			NewStepper(c, &step, 0, 10, 1).Variant(Outline).Width(120)
			NewRating(c, &rating, 5)
			NewMeter(c, 60, 0, 100, &levels)
			NewColorPicker(c, &color)
		})
	}
	tt := ui.NewTester(view, 320, 360)
	tt.Advance(time.Second)
	if low != 20 || high != 80 || step != 3 || rating != 3 {
		t.Errorf("a wrapped control changed its bound value: %v %v %v %v", low, high, step, rating)
	}
}
