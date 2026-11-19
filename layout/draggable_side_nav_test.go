package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestDraggableSideNavReportsItsState(t *testing.T) {
	width, expanded := float32(240), true
	var seen SideNavState
	view := func(c *ui.Context) {
		nav := NewDraggableSideNav(c, &width, &expanded, DraggableSideNavProps{}, func(s SideNavState) {
			seen = s
			ui.Text(c, "内容")
		})
		if !nav.Expanded() {
			t.Error("the nav should report expanded")
		}
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	if !tt.HasText("内容") {
		t.Fatal("the nav body is not on screen")
	}
	if seen.RailWidth <= 0 {
		t.Error("the nav did not hand a rail width to its body")
	}
}
