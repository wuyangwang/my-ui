package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestDraggablePanelToggle(t *testing.T) {
	size, expanded := float32(220), true
	view := func(c *ui.Context) {
		NewDraggablePanel(c, &size, &expanded, DraggablePanelProps{}, func() {
			ui.Text(c, "面板内容")
		})
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	if !tt.HasText("面板内容") {
		t.Fatal("the expanded panel's content is not on screen")
	}
	// The collapse button lives only in the collapsed grip, so drive the
	// app's pointer directly: the panel writes only what the app owns.
	expanded = false
	tt.Frame()
	if tt.HasText("面板内容") {
		t.Error("the content of a collapsed panel is still on screen")
	}
}
