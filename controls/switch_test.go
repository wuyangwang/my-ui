package controls

import (
	. "github.com/wuyangwang/my-ui"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSwitchTogglesItsState(t *testing.T) {
	on := false
	changed := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSwitch(c, &on).Label("深色").OnChanged(func(v bool) { changed = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	// The switch has no text of its own: it is found by its name.
	info, ok := tt.Access("深色")
	if !ok {
		t.Fatal("the switch is not on screen")
	}
	b := info.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Advance(time.Second)
	if !on {
		t.Error("the switch is not on after the click")
	}
	if !changed {
		t.Error("the change was not reported")
	}
	// And a second click turns it off.
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Advance(time.Second)
	if on {
		t.Error("the switch is still on after the second click")
	}
}

func TestSwitchTakesNoClickWhenDisabled(t *testing.T) {
	on := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSwitch(c, &on).Label("深色").Disabled(true)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("深色")
	if !ok {
		t.Fatal("the switch is not on screen")
	}
	b := info.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Advance(time.Second)
	if on {
		t.Error("the disabled switch took the click")
	}
}
