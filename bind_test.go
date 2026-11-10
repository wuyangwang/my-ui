package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestBindMakesTheSameComponents(t *testing.T) {
	clicked := false
	view := func(c *ui.Context) {
		h := Bind(c)
		h.Apply(Options{Appearance: AppearanceLight})
		if h.Button("Save").Primary().Clicked() {
			clicked = true
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Access("Save"); !ok {
		t.Fatal("the button is not on screen")
	}
	if err := tt.Click("Save"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !clicked {
		t.Error("the button did not report the click")
	}
}

func TestBindTakesTheThemeOfTheFrame(t *testing.T) {
	view := func(c *ui.Context) {
		h := Bind(c)
		if h.Theme() == nil {
			t.Error("the handle has no theme")
		}
		if h.Apply(Options{Appearance: AppearanceDark}) == nil {
			t.Error("the handle did not apply the appearance")
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
}

func TestBindKeepsTheKeyOfItsContext(t *testing.T) {
	view := func(c *ui.Context) {
		Bind(c.Key("row")).Button("Save").Primary()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Access("Save"); !ok {
		t.Fatal("the button is not on screen")
	}
}

// TestBindOutlivesTheFrame checks that a handle made in one frame still
// works in the next: the context of a window does not change from frame
// to frame, only what the frame builds does.
func TestBindOutlivesTheFrame(t *testing.T) {
	var h UI
	frames := 0
	view := func(c *ui.Context) {
		if h.Context() == nil {
			h = Bind(c)
		}
		frames++
		h.Button("Save").Primary()
		c.Invalidate()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	tt.Advance(time.Second)
	if frames < 2 {
		t.Fatalf("the view ran %d times, want 2", frames)
	}
	if _, ok := tt.Access("Save"); !ok {
		t.Fatal("the button is not on screen")
	}
}
