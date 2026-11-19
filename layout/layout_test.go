package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestLayoutSlots(t *testing.T) {
	view := func(c *ui.Context) {
		NewLayout(c, LayoutProps{
			Header:  func() { ui.Text(c, "页头") },
			Sidebar: func() { ui.Text(c, "侧栏") },
			Toc:     func() { ui.Text(c, "目录") },
			Footer:  func() { ui.Text(c, "页脚") },
		}, func() {
			ui.Text(c, "正文")
		})
	}
	tt := ui.NewTester(view, 400, 300)
	tt.Advance(time.Second)
	for _, s := range []string{"页头", "侧栏", "目录", "页脚", "正文"} {
		if !tt.HasText(s) {
			t.Errorf("the layout slot %q is not on screen", s)
		}
	}
}
