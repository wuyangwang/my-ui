package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSplitShowsBothPanes(t *testing.T) {
	size := float32(120)
	view := func(c *ui.Context) {
		NewSplit(c, &size, func() {
			ui.Text(c, "左窗格")
		}, func() {
			ui.Text(c, "右窗格")
		}).Height(160)
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	for _, s := range []string{"左窗格", "右窗格"} {
		if !tt.HasText(s) {
			t.Errorf("the pane text %q is not on screen", s)
		}
	}
}
