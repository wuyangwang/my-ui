package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestWindowPageOffersTheOpenButton checks the page renders its open
// button and its notes. It never clicks the button: that would create
// a real window.
func TestWindowPageOffersTheOpenButton(t *testing.T) {
	s := &showcase{page: "window"}
	tt := ui.NewTester(s.view, 960, 620)
	for _, what := range []string{
		"打开自定义窗口",
		"一行建出自定义标题栏窗口，随明暗换色。",
	} {
		if _, ok := tt.Find(what); !ok {
			t.Fatalf("the block %q is missing", what)
		}
	}
}
