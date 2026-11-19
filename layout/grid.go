package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// NewGrid builds a grid container with my-ui's spacing. It keeps mygo's
// ui.Grid signature, so a call written for the old re-export compiles, but
// the gap comes from the theme on every frame. Children fill cells row by
// row; ColumnStart, RowStart, ColumnSpan and RowSpan are mygo's own.
//
//	layout.NewGrid(c).Children(func() {
//		ui.Text(c, "第一个")
//		ui.Text(c, "第二个")
//	})
func NewGrid(c *ui.Context) ui.Element {
	t := c.Theme()
	return ui.Grid(c).Gap(t.Space(2))
}

// Grid is a themed grid with an optional card face, kept as a struct so a
// view can style it after the shared presets. NewGrid returns the bare
// mygo-compatible element; use NewGridPanel when the fluent form is wanted.
type Grid struct {
	ui.Element
	c *ui.Context
}

// NewGridPanel creates a Grid with the theme's gap. Variant paints a card
// face: Outline gives it a bordered surface, Secondary a muted one, the
// zero value leaves it bare.
func NewGridPanel(c *ui.Context, variant Variant) *Grid {
	g := &Grid{c: c}
	t := c.Theme()
	el := ui.Grid(c).Gap(t.Space(2))
	switch variant {
	case Outline:
		el = el.Background(t.Surface).Border(1, t.Border).
			Radius(Radius(t, RadiusLg)).Padding(t.Space(4))
	case Secondary:
		el = el.Background(t.Surface).Radius(Radius(t, RadiusLg)).Padding(t.Space(4))
	}
	g.Element = el
	return g
}

// AutoGrid creates a responsive grid of n items: as many equal columns as
// fit at minWidth or wider, rows height DIPs high, sharing what room there
// is. It wraps mygo's virtualized GridView, so thousands of items cost no
// more than a dozen. state keeps the scroll place and, when its Selected
// is set, the choice; a nil state is fine.
func AutoGrid(c *ui.Context, state *ui.GridState, n int, minWidth, height float32, item func(i int)) ui.Element {
	t := c.Theme()
	el := ui.GridView(c, state, n, minWidth, height, item).Gap(t.Space(2))
	return el.ScrollbarStyle(LobeScrollbar(t))
}
