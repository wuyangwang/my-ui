package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// DividerOptions configures NewDivider.
type DividerOptions struct {
	// Vertical draws an upright line between inline items.
	Vertical bool
	// Dashed draws the line in dashes.
	Dashed bool
	// Label shows in the middle of a horizontal line.
	Label string
}

// Divider is a thin line between content, across its column or upright
// between the items of a row. Its color, text and metrics come from the
// theme on the frame it is built.
type Divider struct {
	ui.Element
	c *ui.Context
	o DividerOptions
}

// NewDivider creates a divider. The options are optional, so
// NewDivider(c) is a plain horizontal rule.
func NewDivider(c *ui.Context, opts ...DividerOptions) *Divider {
	d := &Divider{c: c}
	if len(opts) > 0 {
		d.o = opts[0]
	}
	d.build()
	return d
}

func (d *Divider) build() {
	// A fluent call may rebuild the line in the same frame. Hide the old
	// element so only one divider is laid out and painted.
	if d.Element.Valid() {
		d.Element.Hidden(true)
	}
	t := d.c.Theme()
	if d.o.Vertical {
		d.Element = dividerLine(d.c, t.Border, d.o.Dashed).
			Width(1).Height(TextSize(t, TextMd)).AlignSelf(ui.Center)
		return
	}
	if d.o.Label == "" {
		d.Element = dividerLine(d.c, t.Border, d.o.Dashed).
			Height(1).AlignSelf(ui.Stretch)
		return
	}
	row := ui.Row(d.c).Gap(t.Space(3)).AlignSelf(ui.Stretch).AlignItems(ui.Center)
	row.Children(func() {
		dividerLine(d.c, t.Border, d.o.Dashed).Height(1).Grow(1)
		// The label sits on a quiet Surface chip, so the hairline never
		// runs through the text; both colors come from this frame's theme.
		ui.Text(d.c, d.o.Label).FontSize(TextSize(t, TextXs)).
			TextColor(t.TextMuted).Background(t.Surface).
			Padding(0, t.Space(1)).SingleLine()
		dividerLine(d.c, t.Border, d.o.Dashed).Height(1).Grow(1)
	})
	d.Element = row
}

func dividerLine(c *ui.Context, color ui.Color, dashed bool) ui.Element {
	line := ui.Box(c).Shrink(0)
	if !dashed {
		return line.Background(color)
	}
	return line.Draw(func(p *ui.Painter, r ui.Rect) {
		const dash, gap = float32(4), float32(4)
		if r.W >= r.H {
			for x := r.X; x < r.X+r.W; x += dash + gap {
				p.Line(x, r.Y+r.H/2, min(x+dash, r.X+r.W), r.Y+r.H/2, 1, color)
			}
			return
		}
		for y := r.Y; y < r.Y+r.H; y += dash + gap {
			p.Line(r.X+r.W/2, y, r.X+r.W/2, min(y+dash, r.Y+r.H), 1, color)
		}
	})
}

// Vertical makes the divider upright.
func (d *Divider) Vertical() *Divider {
	d.o.Vertical = true
	d.build()
	return d
}

// Dashed draws the line in dashes.
func (d *Divider) Dashed() *Divider {
	d.o.Dashed = true
	d.build()
	return d
}

// Label puts s in the middle of a horizontal line.
func (d *Divider) Label(s string) *Divider {
	d.o.Label = s
	d.build()
	return d
}
