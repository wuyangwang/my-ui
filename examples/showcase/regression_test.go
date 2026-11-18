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
	// Read the slider's own box instead of a baked pixel row: the
	// showcase's spacing changes as its layout evolves. A track shows
	// as a long horizontal run of non-background pixels, not a dot.
	slider, ok := tt.Access("音量")
	if !ok {
		t.Fatal("the volume slider is missing")
	}
	best := 0
	for y := int(slider.Bounds.Y); y < int(slider.Bounds.Y+slider.Bounds.H); y++ {
		run := 0
		for x := int(slider.Bounds.X); x < int(slider.Bounds.X+slider.Bounds.W); x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r>>8 < 235 || g>>8 < 235 || b>>8 < 235 {
				run++
			}
		}
		if run > best {
			best = run
		}
	}
	if best < 40 {
		t.Errorf("the longest slider row has %d painted pixels, want a track", best)
	}
}

// TestRadioDotIsCentered checks a selected radio's inner dot sits at the
// center of its ring, not at a corner.
func TestRadioDotIsCentered(t *testing.T) {
	s := &showcase{page: "radio", size: "medium"}
	tt := ui.NewTester(s.view, 1100, 700)
	tt.Advance(time.Second)
	img := tt.Image()
	darkAt := func(x, y int) bool {
		p := img.At(x, y)
		r, g, b, _ := p.RGBA()
		return r>>8 < 100 && g>>8 < 100 && b>>8 < 100
	}
	// The medium radio's ring: the dot is 18 DIPs across, 8 from the
	// label ("中"), and centers on the label's line. Find the label and
	// read the dot's center from it, so the check survives layout
	// changes.
	label, ok := tt.Find("中")
	if !ok {
		t.Fatal("the medium radio's label is missing")
	}
	cx := label.X - 8 - 9
	cy := label.Y + label.H/2
	if !darkAt(int(cx), int(cy)) {
		t.Errorf("the center of the selected radio at (%v, %v) is not dark: the inner dot is not centered", cx, cy)
	}
}
