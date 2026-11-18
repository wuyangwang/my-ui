package forms

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// InputNumber edits a float with step buttons, bounded by lo and hi.
// The value pointer is the app's, as in mygo's NumberInput.
type InputNumber struct {
	ui.Element
	c       *ui.Context
	value   *float64
	lo, hi  float64
	step    float64
	variant Variant
	size    Size
	width   float32
}

// NewInputNumber builds a numeric field. step at or below zero keeps
// mygo's default step.
func NewInputNumber(c *ui.Context, value *float64, lo, hi, step float64) *InputNumber {
	n := &InputNumber{c: c, value: value, lo: lo, hi: hi, step: step}
	n.Element = ui.NumberInput(c, value, lo, hi, step).Transition(StateTransition())
	n.apply()
	return n
}

// Variant applies the shared look of v.
func (n *InputNumber) Variant(v Variant) *InputNumber {
	n.variant = v
	n.apply()
	return n
}

// Size picks the text and padding scale.
func (n *InputNumber) Size(sz Size) *InputNumber {
	n.size = sz
	n.apply()
	return n
}

// Width pins the field's width in DIPs; zero fills the parent.
func (n *InputNumber) Width(w float32) *InputNumber {
	n.width = w
	n.apply()
	return n
}

func (n *InputNumber) apply() {
	t := n.c.Theme()
	look := FrameLookOf(n.variant, t)
	if !look.Native {
		n.Element = n.Element.Background(look.Bg).Border(look.Width, look.Border)
		if look.HasText {
			n.Element = n.Element.TextColor(look.Text)
		}
	}
	v, h := FramePadding(t, n.size)
	n.Element = n.Element.Radius(t.Radius).Padding(v, h).FontSize(FrameFontSize(t, n.size))
	if n.width > 0 {
		n.Element = n.Element.Width(n.width)
	} else {
		n.Element = n.Element.FillWidth()
	}
}
