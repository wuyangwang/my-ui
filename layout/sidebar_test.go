package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSidebarShowsItems(t *testing.T) {
	selected := "b"
	view := func(c *ui.Context) {
		NewSidebar(c, &selected, func() {
			ui.SidebarItem(c, "a", nil, "甲")
			ui.SidebarItem(c, "b", nil, "乙")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	for _, s := range []string{"甲", "乙"} {
		if !tt.HasText(s) {
			t.Errorf("the sidebar item %q is not on screen", s)
		}
	}
}
