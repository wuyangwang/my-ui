package myui

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// Progress is a bar that shows how far along something is.
//
//	myui.NewProgress(c, 0.4)
//	myui.NewProgress(c, -1).Label("加载中")  // unknown length
//
// The bar is styled by the library, as shadcn's: a thin rounded line,
// the accent up to where the work is. It carries the progressbar role,
// and a negative value shows activity of unknown length.
type Progress struct {
	ui.Element
	c       *ui.Context
	value   float64
	variant Variant
	height  float32
	label   string
}

// NewProgress creates a progress bar at value: a number between 0 and 1.
// A negative value shows activity of unknown length.
func NewProgress(c *ui.Context, value float64) *Progress {
	p := &Progress{c: c, value: value, height: 2}
	p.build()
	return p
}

// build paints the bar. It is called again when a preset changes, so the
// fill follows the theme and the variant.
func (p *Progress) build() {
	t := p.c.Theme()
	old := p.Element
	bar := ui.Box(p.c).Height(p.height).Radius(t.Space(1)).Background(t.SurfaceHover).
		Clip().Role(ui.RoleProgress)
	if p.label != "" {
		bar = bar.Label(p.label)
	}
	fill := t.Accent
	switch p.variant {
	case Success, Info, Warning:
		role, _ := semanticOf(p.variant)
		fill = semanticColor(t, role)
	case Destructive:
		fill = t.Danger
	}
	value := p.value
	bar.Draw(func(painter *ui.Painter, r ui.Rect) {
		if value >= 0 {
			at := float32(math.Min(math.Max(value, 0), 1))
			painter.Fill(ui.Rect{X: r.X, Y: r.Y, W: r.W * at, H: r.H}, fill, r.H/2)
			return
		}
		// Unknown length: a segment travels along the track, painted
		// again without rebuilding the view.
		painter.AnimationFrame()
		th := p.c.Theme()
		phase := float32(painter.Now().UnixMilli()%1400) / 1400
		w := r.W * 0.3
		x := r.X - w + (r.W+w)*phase
		painter.Clip(r, th.Space(1), func() {
			painter.Fill(ui.Rect{X: x, Y: r.Y, W: w, H: r.H}, fill, r.H/2)
		})
	})
	p.Element = bar
	_ = old
}

// Indeterminate makes the bar show activity of unknown length and returns
// it, for chaining.
func (p *Progress) Indeterminate() *Progress {
	p.value = -1
	p.build()
	return p
}

// Value sets the bar to current out of max and returns it, for chaining.
func (p *Progress) Value(current, max float64) *Progress {
	if max > 0 {
		p.value = current / max
	} else {
		p.value = 0
	}
	p.build()
	return p
}

// Height sets the thickness of the bar in DIPs and returns it, for
// chaining.
func (p *Progress) Height(h float32) *Progress {
	p.height = h
	p.build()
	return p
}

// Variant paints the bar with the color of v — a status color — and
// returns it, for chaining.
func (p *Progress) Variant(v Variant) *Progress {
	p.variant = v
	p.build()
	return p
}

// Label names the bar for assistive technology and returns it, for
// chaining: a bar says how far along, a name says what of.
func (p *Progress) Label(label string) *Progress {
	p.label = label
	p.Element = p.Element.Label(label)
	return p
}
