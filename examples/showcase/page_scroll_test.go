package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// TestSwitchingPagesStartsAtTheTop guards the scroll handover: the
// content column's offset belongs to the open page and is reset when
// the app switches to another component.
func TestSwitchingPagesStartsAtTheTop(t *testing.T) {
	s := &showcase{page: "button"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)
	s.scroll.Y = 400
	tt.Frame()
	if s.scroll.Y == 0 {
		t.Fatal("the scroll state did not take the offset")
	}

	s.page = "input"
	tt.Frame()
	if s.scroll.Y != 0 {
		t.Errorf("the scroll stayed at %v after switching pages", s.scroll.Y)
	}
	if s.lastPage != "input" {
		t.Errorf("the scroll's page is %q, want input", s.lastPage)
	}
}
