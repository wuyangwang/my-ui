package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestCheckboxTogglesItsState(t *testing.T) {
	checked := false
	changed := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewCheckbox(c, &checked, "同步").OnChanged(func(v bool) { changed = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("同步"); !ok {
		t.Fatal("the label of the check box is not on screen")
	}
	if err := tt.Click("同步"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !checked {
		t.Error("the check box is not checked after the click")
	}
	if !changed {
		t.Error("the change was not reported")
	}
	// And a second click unchecks it.
	if err := tt.Click("同步"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if checked {
		t.Error("the check box is still checked after the second click")
	}
}

func TestCheckboxTakesNoClickWhenDisabled(t *testing.T) {
	checked := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewCheckbox(c, &checked, "同步").Disabled(true)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("同步"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if checked {
		t.Error("the disabled check box took the click")
	}
}
