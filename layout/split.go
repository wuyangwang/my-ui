package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// NewSplit builds mygo's resizable split with my-ui's frame. It keeps
// mygo's ui.Split signature, so a call written for the old re-export
// compiles. The size pointer is the width of the first pane, written by
// mygo as the divider is dragged; the frame's colors follow c.Theme().
//
//	layout.NewSplit(c, &app.sidebarWidth, app.files, app.editor).Fill()
func NewSplit(c *ui.Context, size *float32, first, second func()) ui.Element {
	t := c.Theme()
	return ui.Split(c, size, first, second).
		Background(t.Background).
		Border(1, t.Border).
		Radius(Radius(t, RadiusLg)).
		Clip().
		Transition(StateTransition())
}
