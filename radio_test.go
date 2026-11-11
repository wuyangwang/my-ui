package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestRadioSelectsItsValue(t *testing.T) {
	size := "small"
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewRadio(c, &size, "small", "小")
		NewRadio(c, &size, "large", "大")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("大"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if size != "large" {
		t.Errorf("the value is %q, want %q", size, "large")
	}
	// The other button of the group lets it go.
	if err := tt.Click("小"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if size != "small" {
		t.Errorf("the value is %q, want %q", size, "small")
	}
}

func TestRadioReportsItsChoice(t *testing.T) {
	size := ""
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewRadio(c, &size, "large", "大").OnChanged(func(v string) { chosen = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("大"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if chosen != "large" {
		t.Errorf("the choice reported %q, want %q", chosen, "large")
	}
}
