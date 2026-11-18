package forms

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// DecimalInput edits an exact decimal string; mygo owns the arithmetic
// and stepping, the library owns the field look.
type DecimalInput struct {
	ui.Element
	c       *ui.Context
	variant Variant
	size    Size
	width   float32
}

// NewDecimalInput builds a decimal field editing *value; lo/hi/step
// are decimal strings, as mygo requires.
func NewDecimalInput(c *ui.Context, value *string, lo, hi, step string) *DecimalInput {
	d := &DecimalInput{c: c}
	d.Element = ui.DecimalInput(c, value, lo, hi, step).Transition(StateTransition())
	d.apply()
	return d
}

// Variant applies the shared look of v.
func (d *DecimalInput) Variant(v Variant) *DecimalInput {
	d.variant = v
	d.apply()
	return d
}

// Size picks the text scale.
func (d *DecimalInput) Size(sz Size) *DecimalInput {
	d.size = sz
	d.apply()
	return d
}

// Width pins the field's width in DIPs; zero fills the parent.
func (d *DecimalInput) Width(w float32) *DecimalInput {
	d.width = w
	d.apply()
	return d
}

func (d *DecimalInput) apply() {
	t := d.c.Theme()
	look := FrameLookOf(d.variant, t)
	if !look.Native {
		d.Element = d.Element.Background(look.Bg).Border(look.Width, look.Border)
		if look.HasText {
			d.Element = d.Element.TextColor(look.Text)
		}
	}
	v, h := FramePadding(t, d.size)
	d.Element = d.Element.Radius(Radius(t, RadiusMd)).
		Padding(v, h).FontSize(FrameFontSize(t, d.size))
	if d.width > 0 {
		d.Element = d.Element.Width(d.width)
	} else {
		d.Element = d.Element.FillWidth()
	}
}
