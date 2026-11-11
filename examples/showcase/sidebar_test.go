package main

import (
	"strings"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// sidebarWidth is the width of the sidebar: what is left of it is the
// page.
const sidebarWidth = 190

// inSidebar tells whether the text s shows in the sidebar, and not in
// the page beside it.
func inSidebar(tt *ui.Tester, s string) bool {
	r, ok := tt.Find(s)
	return ok && r.X < sidebarWidth
}

// foldLabel is the title of the fold of that name: the name and the
// count of its pages.
func foldLabel(tt *ui.Tester, name string) (string, bool) {
	for _, s := range tt.Texts() {
		if strings.HasPrefix(s, name+" (") {
			return s, true
		}
	}
	return "", false
}

// TestTheSidebarFoldsThePages: the folds are closed one at a time, and
// a fold's title opens it and closes the other.
func TestTheSidebarFoldsThePages(t *testing.T) {
	s := &showcase{page: "button"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)

	// The fold of the page that is open shows its pages.
	if !inSidebar(tt, "Button") {
		t.Fatal("the page of the open fold is missing")
	}
	// Another fold does not.
	if inSidebar(tt, "外观") {
		t.Error("the pages of a closed fold show")
	}
	// The title of a fold opens it.
	label, ok := foldLabel(tt, "主题")
	if !ok {
		t.Fatal("the fold of the theme is missing")
	}
	if err := tt.Click(label); err != nil {
		t.Fatal(err)
	}
	tt.Advance(100 * 1000 * 1000)
	tt.Advance(100 * 1000 * 1000)
	if !inSidebar(tt, "外观") {
		t.Fatal("the pages of the opened fold do not show")
	}
	// And the fold of the page that was open closes.
	if inSidebar(tt, "Button") {
		t.Error("the pages of the closed fold still show")
	}
}

// TestTheSidebarSearchesThePages: a search leaves only what it finds,
// and opens the folds to show it.
func TestTheSidebarSearchesThePages(t *testing.T) {
	s := &showcase{page: "button"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)

	// The search field of the sidebar.
	search, ok := tt.Access("搜索组件")
	if !ok {
		t.Fatal("the search field of the sidebar is missing")
	}
	b := search.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Type("toast")
	tt.Advance(time.Second)
	tt.Advance(100 * 1000 * 1000)

	if !inSidebar(tt, "Toast") {
		t.Error("the page the search finds is missing")
	}
	if inSidebar(tt, "Button") {
		t.Error("a page the search does not find shows")
	}
	// And a search that finds nothing says so.
	tt.Type("xyz")
	tt.Advance(time.Second)
	tt.Advance(100 * 1000 * 1000)
	if !inSidebar(tt, "没有匹配的组件") {
		t.Error("the sidebar does not say that it found nothing")
	}
}

// TestTheSidebarScrollsInASmallWindow: the folds are more than a small
// window is tall, and the wheel over the sidebar brings the last of them
// up.
func TestTheSidebarScrollsInASmallWindow(t *testing.T) {
	s := &showcase{page: "button"}
	const height = 300
	tt := ui.NewTester(s.view, 960, height)
	tt.Advance(time.Second)

	label, ok := foldLabel(tt, "主题")
	if !ok {
		t.Fatal("the last fold of the sidebar is missing")
	}
	before, _ := tt.Find(label)
	if before.Y < height {
		t.Skipf("the sidebar fits in the window: the last fold sits at %v", before.Y)
	}
	tt.Wheel(95, 200, 0, 600)
	tt.Advance(time.Second)
	after, _ := tt.Find(label)
	if after.Y >= height {
		t.Errorf("the last fold sits at %v, still below the window", after.Y)
	}
}
