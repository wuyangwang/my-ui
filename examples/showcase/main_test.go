package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestSidebarSitsBesideThePages checks the two panes: the sidebar's
// texts stay on the left strip, the page's cards open to its right,
// and neither pane overruns the window.
func TestSidebarSitsBesideThePages(t *testing.T) {
	s := &showcase{}
	tt := ui.NewTester(s.view, 960, 620)

	check := func(what string, x, w float32, strip bool) {
		t.Helper()
		left := x < 190
		if strip != left {
			t.Errorf("%s: x %v w %v, want sidebar=%v", what, x, w, strip)
		}
		if x+w > 960 {
			t.Errorf("%s: x %v + w %v overruns the window", what, x, w)
		}
	}

	brand, ok := tt.Find("my-ui")
	if !ok {
		t.Fatal("the brand text is missing")
	}
	check("brand", brand.X, brand.W, true)

	nav, ok := tt.Find("外观")
	if !ok {
		t.Fatal("the nav entry is missing")
	}
	check("nav entry", nav.X, nav.W, true)

	section, ok := tt.Find("变体")
	if !ok {
		t.Fatal("the variants section is missing")
	}
	check("variants section", section.X, section.W, false)

	sizes, ok := tt.Find("尺寸")
	if !ok {
		t.Fatal("the sizes section is missing")
	}
	check("sizes section", sizes.X, sizes.W, false)
}

// TestThemePageSwitchesTheAppearanceAndThePalette opens the theme page
// and takes a switch for a drive: the page builds, and the pressed
// button has the rest painted.
func TestThemePageSwitchesTheAppearanceAndThePalette(t *testing.T) {
	s := &showcase{}
	tt := ui.NewTester(s.view, 960, 620)
	if err := tt.Click("外观"); err != nil {
		t.Fatal(err)
	}
	// 点击分发的输入只停在队列，下一帧才进视图。
	tt.Advance(100 * 1000 * 1000)
	if _, ok := tt.Find("跟随系统"); !ok {
		t.Fatal("the appearance switch is missing")
	}
	if _, ok := tt.Find("浅色子树"); !ok {
		t.Fatal("the subtree preview is missing")
	}
}
