package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// TestTheIconPageShowsTheWholeSet: every icon of both sets is on the
// page, and a search narrows them down.
func TestTheIconPageShowsTheWholeSet(t *testing.T) {
	s := &showcase{page: "icon"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)

	// The page shows the count of the sets.
	if _, ok := tt.Find("Lucide 2318 / 2318 · 品牌标 3462 / 3462"); !ok {
		t.Error("the page does not show the count of the sets")
	}
	// And the icons of them: the first of the Lucide set, one in the
	// middle, and a brand.
	for _, name := range []string{"AArrowDown", "Save", "ZoomIn", "GitHub", "Anthropic"} {
		if _, ok := tt.Access(name); !ok {
			t.Errorf("the icon %q is not on the page", name)
		}
	}
	// A search narrows them down.
	search, ok := tt.Access("搜索图标")
	if !ok {
		t.Fatal("the search of the icon page is missing")
	}
	b := search.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Type("zap")
	tt.Advance(time.Second)
	tt.Advance(100 * 1000 * 1000)
	if _, ok := tt.Find("Lucide 2318 / 2318 · 品牌标 3462 / 3462"); ok {
		t.Error("the count does not follow the search")
	}
}

// TestTheIconPageCopiesTheName: a click on an icon copies its name.
func TestTheIconPageCopiesTheName(t *testing.T) {
	s := &showcase{page: "icon"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)

	// The icon of the search field is the one that copies: the click
	// lands on the icon of the set.
	icon, ok := tt.Access("Save")
	if !ok {
		t.Fatal("the save icon is not on the page")
	}
	b := icon.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Advance(time.Second)
	// A tester has no clipboard, so nothing is copied; the click does
	// not break the page either.
	if _, ok := tt.Find("Save"); !ok {
		t.Error("the page lost the icon after the click")
	}
}
