package myui

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// Progress is a bar that shows how far along something is.
//
//	myui.NewProgress(c, 0.4)
//
// The bar is styled by the library, as shadcn's: a thin rounded line,
// the accent up to where the work is.
type Progress struct {
	ui.Element
}

// NewProgress creates a progress bar at value: a number between 0 and
// 1.
func NewProgress(c *ui.Context, value float64) *Progress {
	t := c.Theme()
	p := &Progress{}
	bar := ui.Box(c).Height(t.Space(2)).Radius(t.Space(1)).Background(t.SurfaceHover).Clip()
	bar.Draw(func(painter *ui.Painter, r ui.Rect) {
		at := float32(math.Min(math.Max(value, 0), 1))
		painter.Fill(ui.Rect{X: r.X, Y: r.Y, W: r.W * at, H: r.H}, t.Accent, r.H/2)
	})
	p.Element = bar
	return p
}

// Label names the bar for assistive technology and returns it, for
// chaining: a bar says how far along, a name says what of.
func (p *Progress) Label(label string) *Progress {
	p.Element = p.Element.Label(label)
	return p
}
