package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestScrollAndScrollAreaShowContent(t *testing.T) {
	view := func(c *ui.Context) {
		ui.Row(c).Gap(8).Children(func() {
			NewScroll(c).Children(func() {
				ui.Text(c, "普通滚动")
			})
			var st ui.ScrollState
			NewScrollArea(c, &st, func() {
				ui.Text(c, "淡入滚动")
			}).Width(140).Height(120)
		})
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	for _, s := range []string{"普通滚动", "淡入滚动"} {
		if !tt.HasText(s) {
			t.Errorf("the text %q is not on screen", s)
		}
	}
}
