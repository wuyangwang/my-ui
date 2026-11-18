package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestPaginationPagesThroughTheCount(t *testing.T) {
	page := 1
	changes := []int{}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewPagination(c, &page, 9).OnChanged(func(p int) { changes = append(changes, p) })
	}
	tt := ui.NewTester(view, 480, 120)
	adv := func() { tt.Advance(time.Second); tt.Advance(time.Second) }
	if _, ok := tt.Find("1"); !ok {
		t.Fatal("the first page is missing")
	}
	// The page buttons hold their labels inside a fixed box; press the
	// box's center from its accessibility bounds.
	clickPage := func(name string) {
		t.Helper()
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		x, y := info.Bounds.X+info.Bounds.W/2, info.Bounds.Y+info.Bounds.H/2
		tt.Press(x, y)
		tt.Advance(time.Second)
		tt.Release(x, y)
	}
	clickPage("2")
	adv()
	if page != 2 {
		t.Errorf("the pointer is %d, want 2", page)
	}
	if len(changes) != 1 || changes[0] != 2 {
		t.Errorf("the change callback got %v, want [2]", changes)
	}
	if err := tt.Click("Next page"); err != nil {
		t.Fatal(err)
	}
	adv()
	if page != 3 {
		t.Errorf("next gave %d, want 3", page)
	}
	if err := tt.Click("Previous page"); err != nil {
		t.Fatal(err)
	}
	adv()
	if page != 2 {
		t.Errorf("previous gave %d, want 2", page)
	}
}

func TestPaginationSkipsPagesWithAnEllipsis(t *testing.T) {
	page := 6
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewPagination(c, &page, 20)
	}
	tt := ui.NewTester(view, 480, 120)
	if _, ok := tt.Find("…"); !ok {
		t.Fatal("a long pager needs an ellipsis")
	}
	if _, ok := tt.Find("20"); !ok {
		t.Fatal("the last page is missing")
	}
	if _, ok := tt.Find("6"); !ok {
		t.Fatal("the current page is missing")
	}
}
