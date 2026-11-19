package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestFloatingSheetOpenAndClose(t *testing.T) {
	open := true
	view := func(c *ui.Context) {
		NewFloatingSheet(c, &open, FloatingSheetProps{
			Title: "笔记", SnapPoints: []float32{0.3, 0.6},
		}, func() {
			ui.Text(c, "第一条笔记")
		})
	}
	tt := ui.NewTester(view, 400, 300)
	tt.Advance(time.Second)
	if !tt.HasText("笔记") || !tt.HasText("第一条笔记") {
		t.Fatal("the open sheet did not draw its title and content")
	}
	open = false
	tt.Frame()
	if tt.HasText("笔记") {
		t.Error("the closed sheet is still on screen")
	}
}
