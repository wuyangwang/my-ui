package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSideNavSlots(t *testing.T) {
	view := func(c *ui.Context) {
		NewSideNav(c, SideNavProps{
			Avatar:        func() { ui.Text(c, "头像") },
			TopActions:    func() { ui.Text(c, "顶部") },
			BottomActions: func() { ui.Text(c, "底部") },
		})
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	for _, s := range []string{"头像", "顶部", "底部"} {
		if !tt.HasText(s) {
			t.Errorf("the side nav slot %q is not on screen", s)
		}
	}
}
