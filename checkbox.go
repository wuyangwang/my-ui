package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// Checkbox is a box that is checked or not, with a label.
//
//	myui.NewCheckbox(c, &app.sync, "同步")
//
// The box is styled by the library, as shadcn's: a square with a
// border, and the accent with a check in it when it is checked. The
// label is the name of the box for assistive technology too.
type Checkbox struct {
	ui.Element
	c       *ui.Context
	box     ui.Element
	checked *bool
	size    Size
}

// NewCheckbox creates a check box editing *checked, with a label.
func NewCheckbox(c *ui.Context, checked *bool, label string) *Checkbox {
	cb := &Checkbox{c: c, checked: checked}
	t := c.Theme()
	row := ui.CheckboxBase(c, checked).Gap(t.Space(2))
	row.Children(func() {
		cb.box = ui.Box(c)
		ui.Text(c, label)
	})
	cb.Element = row
	cb.applyLook()
	return cb
}

// applyLook paints the box as the state says.
func (cb *Checkbox) applyLook() {
	t := cb.c.Theme()
	d := cb.boxSize()
	cb.box.Size(d, d).Radius(t.Radius/2).Border(1, t.Border)
	if cb.checked != nil && *cb.checked {
		cb.box.Background(t.Accent).Border(1, t.Accent)
		cb.box.Children(func() {
			ui.Icon(cb.c, icons.Check()).Size(d-4, d-4).TextColor(t.AccentText)
		})
	}
}

// boxSize is the size of the box in DIPs.
func (cb *Checkbox) boxSize() float32 {
	switch cb.size {
	case SizeXs:
		return 14
	case SizeLg:
		return 22
	case SizeSm:
		return 16
	}
	return 18
}

// Size applies the sizing of s and returns the box, for chaining: the
// box of the check box grows with it.
func (cb *Checkbox) Size(s Size) *Checkbox {
	cb.size = s
	cb.applyLook()
	return cb
}

// Sm sizes the box down.
func (cb *Checkbox) Sm() *Checkbox { return cb.Size(SizeSm) }

// Xs sizes the box to the smallest size.
func (cb *Checkbox) Xs() *Checkbox { return cb.Size(SizeXs) }

// Md gives the box the size of the text around it.
func (cb *Checkbox) Md() *Checkbox { return cb.Size(SizeMd) }

// Lg sizes the box up.
func (cb *Checkbox) Lg() *Checkbox { return cb.Size(SizeLg) }

// OnChanged runs fn when the box is checked or unchecked in this frame,
// and returns the box, for chaining.
func (cb *Checkbox) OnChanged(fn func(checked bool)) *Checkbox {
	if fn != nil && cb.Changed() {
		fn(*cb.checked)
	}
	return cb
}

// Disabled dims the box and takes its clicks away, and returns the box,
// for chaining.
func (cb *Checkbox) Disabled(disabled bool) *Checkbox {
	cb.Element = cb.Element.Disabled(disabled)
	return cb
}
