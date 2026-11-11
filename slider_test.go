package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSliderSetsItsValue(t *testing.T) {
	value := 0.0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSlider(c, &value, 0, 100).Label("音量")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("音量")
	if !ok {
		t.Fatal("the slider is not on screen")
	}
	b := info.Bounds
	// A click on the track sets the value near where it lands.
	tt.ClickAt(b.X+b.W*0.75, b.Y+b.H/2)
	tt.Advance(time.Second)
	if value < 60 || value > 90 {
		t.Errorf("the value is %v, want it near 75", value)
	}
}

func TestSliderMovesWithTheKeys(t *testing.T) {
	value := 50.0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSlider(c, &value, 0, 100).Label("音量")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("音量")
	if !ok {
		t.Fatal("the slider is not on screen")
	}
	b := info.Bounds
	tt.ClickAt(b.X+b.W*0.5, b.Y+b.H/2)
	tt.Advance(time.Second)
	before := value
	tt.KeyDown(ui.KeyRight)
	tt.KeyUp(ui.KeyRight)
	tt.Advance(time.Second)
	if value <= before {
		t.Errorf("the value is %v, want it above %v", value, before)
	}
}

func TestSliderTakesNoClickWhenDisabled(t *testing.T) {
	value := 50.0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSlider(c, &value, 0, 100).Label("音量").Disabled(true)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("音量")
	if !ok {
		t.Fatal("the slider is not on screen")
	}
	b := info.Bounds
	tt.ClickAt(b.X+b.W*0.9, b.Y+b.H/2)
	tt.Advance(time.Second)
	if value != 50 {
		t.Errorf("the value is %v, want it left alone", value)
	}
}
