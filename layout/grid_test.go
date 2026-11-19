package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestGridShowsItsChildren(t *testing.T) {
	view := func(c *ui.Context) {
		NewGrid(c).Children(func() {
			ui.Text(c, "格子一")
			ui.Text(c, "格子二")
		})
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	for _, s := range []string{"格子一", "格子二"} {
		if _, ok := tt.Find(s); !ok {
			t.Errorf("the text %q is not on screen", s)
		}
	}
}

func TestAutoGridFitsItems(t *testing.T) {
	view := func(c *ui.Context) {
		AutoGrid(c, nil, 6, 100, 60, func(i int) {
			ui.Text(c, "项目")
		})
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	if n := len(tt.Texts()); n == 0 {
		t.Fatal("the auto grid drew no items")
	}
}
