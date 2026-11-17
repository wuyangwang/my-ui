package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// TestComboEcho selects an option and checks the input echoes the label.
func TestComboEcho(t *testing.T) {
	s := &showcase{page: "combobox"}
	tt := ui.NewTester(s.view, 1100, 700)
	tt.Advance(time.Second)
	if err := tt.Click("语言"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("简体中文"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if s.query != "简体中文" {
		t.Errorf("the input holds %q after choosing, want the label 简体中文", s.query)
	}
}

// TestSliderHasTrack checks the slider draws a track wide enough for the
// knob to travel along, not just a collapsed knob.
func TestSliderHasTrack(t *testing.T) {
	s := &showcase{page: "slider", volume: 50}
	tt := ui.NewTester(s.view, 1100, 700)
	tt.Advance(time.Second)
	img := tt.Image()
	// The first slider sits near x 308..450 on the page; a track drawn
	// across it means many non-white pixels along that row, not only a
	// single dot. Scan a horizontal band for the dark track/knob pixels.
	y := 260
	dark := 0
	for x := 300; x < 460; x++ {
		p := img.At(x, y)
		r, g, b, _ := p.RGBA()
		if r>>8 < 120 && g>>8 < 120 && b>>8 < 120 {
			dark++
		}
	}
	if dark < 40 {
		t.Errorf("only %d dark pixels along the slider row, want a track (the knob alone would be ~16)", dark)
	}
}

// TestRadioDotIsCentered checks a selected radio's inner dot sits at the
// center of its ring, not at a corner.
func TestRadioDotIsCentered(t *testing.T) {
	s := &showcase{page: "radio", size: "medium"}
	tt := ui.NewTester(s.view, 1100, 700)
	tt.Advance(time.Second)
	// The medium radio ring is centered near (364, 259) with a radius of
	// about 9 DIPs. The inner dot (dark) must be at the center, not the
	// upper-left corner.
	img := tt.Image()
	darkAt := func(x, y int) bool {
		p := img.At(x, y)
		r, g, b, _ := p.RGBA()
		return r>>8 < 100 && g>>8 < 100 && b>>8 < 100
	}
	if !darkAt(364, 259) {
		t.Error("the center of the selected radio is not dark: the inner dot is not centered")
	}
}
