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

// TestPaginationSizePersistsAcrossFrames checks that a chained Size is
// staged for the next frame's build instead of being written to a pager
// object that is thrown away when the view runs again. The wide buttons
// must stay wide on later frames, and must still take a single click.
func TestPaginationSizePersistsAcrossFrames(t *testing.T) {
	page := 1
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewPagination(c, &page, 9).Size(SizeLg)
	}
	tt := ui.NewTester(view, 480, 120)
	tt.Advance(time.Second)
	tt.Advance(time.Second)
	// Press a page once: the click forces a rebuild, which is the frame
	// that reads the size staged by Size on the first build.
	clickPage := func(name string) {
		t.Helper()
		pi, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		x, y := pi.Bounds.X+pi.Bounds.W/2, pi.Bounds.Y+pi.Bounds.H/2
		tt.Press(x, y)
		tt.Advance(time.Second)
		tt.Release(x, y)
		tt.Advance(time.Second)
		tt.Advance(time.Second)
	}
	clickPage("2")
	if page != 2 {
		t.Fatalf("one click gave page %d, want 2", page)
	}

	info, ok := tt.Access("3")
	if !ok {
		t.Fatal("page 3 is not on screen")
	}
	// SizeLg gives each page button a 36 DIP box; the default is 28.
	// The buttons were rebuilt after Size ran, so this checks that the
	// choice survived into later frames.
	if info.Bounds.W < 34 {
		t.Fatalf("page button is %.0f wide after Size(SizeLg), want ~36", info.Bounds.W)
	}

	clickPage("3")
	if page != 3 {
		t.Errorf("one click gave page %d, want 3", page)
	}
}
