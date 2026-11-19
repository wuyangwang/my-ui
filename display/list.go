package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// NewList builds mygo's virtual list in a shadcn container: the theme's
// surface, a hairline border, a rounded face and a little padding. The
// scrolling surface is the container itself, so the list keeps mygo's
// height rules and the caller can still size or grow the returned
// element.
//
// Everything the list does — virtual rows, selection, keys, reordering,
// keyboard and assistive technology — stays with mygo. The state pointer
// is passed through untouched: the app owns ListState.Selected, Key and
// ScrollTo as before.
//
//	display.NewList(c, &app.files, len(app.files), func(i int) {
//		ui.Text(c, app.files[i])
//	})
//
// The colors and metrics are read from c.Theme() on each build, and
// changes fade through StateTransition, so a theme change lands on the
// next frame. There are no timers or goroutines.
//
// NewList and NewSortableList are independent: the latter adds a drag
// grip to each row, while this one only supplies the container.
func NewList(c *ui.Context, state *ui.ListState, n int, row func(i int)) ui.Element {
	if c == nil {
		return ui.Element{}
	}
	t := c.Theme()
	return ui.List(c, state, n, row).
		Background(t.Surface).Border(1, t.Border).
		Radius(Radius(t, RadiusLg)).Padding(t.Space(2)).
		Transition(StateTransition())
}
