package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestFloatingPanelOpenAndClose(t *testing.T) {
	open := true
	view := func(c *ui.Context) {
		if NewFloatingPanel(c, &open, FloatingPanelProps{Title: "任务面板"}, func() {
			ui.Text(c, "运行步骤")
		}).Open() {
			ui.Text(c, "已打开")
		}
	}
	tt := ui.NewTester(view, 400, 300)
	tt.Advance(time.Second)
	if !tt.HasText("任务面板") || !tt.HasText("运行步骤") {
		t.Fatal("the open panel did not draw its title and body")
	}
	open = false
	tt.Frame()
	if tt.HasText("任务面板") {
		t.Error("the closed panel is still on screen")
	}
}
