package controls

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
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
	invalid bool
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
	border := t.Border
	if cb.invalid {
		border = t.Danger
	}
	cb.box.Size(d, d).Radius(Radius(t, RadiusSm)).Border(1, border)
	if cb.checked != nil && *cb.checked {
		fill := t.Accent
		if cb.invalid {
			fill = t.Danger
		}
		cb.box.Background(fill).Border(1, fill)
		cb.box.Children(func() {
			ui.Icon(cb.c, icons.Check()).Size(checkGlyph(d), checkGlyph(d)).TextColor(t.AccentText)
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

// Invalid marks the box as not holding a good value and returns it, for
// chaining. The square takes the danger color of the theme.
func (cb *Checkbox) Invalid(on bool) *Checkbox {
	cb.invalid = on
	cb.Element = cb.Element.Invalid(on)
	cb.applyLook()
	return cb
}

// Control returns the element a field describes: the check box itself.
func (cb *Checkbox) Control() ui.Element { return cb.Element }

// CheckboxState is a check box with three states — off, on and mixed — as
// shadcn's indeterminate checkbox and Element Plus's half-checked one.
type CheckboxState struct {
	ui.Element
	c     *ui.Context
	state *ui.CheckState
	box   ui.Element
	label string
	size  Size
}

// NewCheckboxState creates a three-state box editing *state, with a label.
func NewCheckboxState(c *ui.Context, state *ui.CheckState, label string) *CheckboxState {
	cs := &CheckboxState{c: c, state: state, label: label}
	t := c.Theme()
	row := ui.CheckboxStateBase(c, state).Gap(t.Space(2))
	row.Children(func() {
		cs.box = ui.Box(c)
		ui.Text(c, label)
	})
	cs.Element = row
	cs.applyLook()
	return cs
}

// applyLook paints the box for the state: a check when on, a dash when
// mixed, empty when off.
func (cs *CheckboxState) applyLook() {
	t := cs.c.Theme()
	d := 18
	if cs.size == SizeSm {
		d = 16
	} else if cs.size == SizeLg {
		d = 22
	}
	cs.box.Size(float32(d), float32(d)).Radius(Radius(t, RadiusSm)).Border(1, t.Border)
	if cs.state == nil {
		return
	}
	switch *cs.state {
	case ui.CheckOn:
		cs.box.Background(t.Accent).Border(1, t.Accent)
		cs.box.Children(func() {
			s := checkGlyph(float32(d))
			ui.Icon(cs.c, icons.Check()).Size(s, s).TextColor(t.AccentText)
		})
	case ui.CheckMixed:
		cs.box.Background(t.Accent).Border(1, t.Accent)
		cs.box.Children(func() {
			ui.Box(cs.c).Size(float32(d-8), 2).Radius(1).Background(t.AccentText)
		})
	}
}

// Size applies the sizing of s and returns the box, for chaining.
func (cs *CheckboxState) Size(s Size) *CheckboxState {
	cs.size = s
	cs.applyLook()
	return cs
}

// OnChanged runs fn with the new state when it changes in this frame.
func (cs *CheckboxState) OnChanged(fn func(ui.CheckState)) *CheckboxState {
	if fn != nil && cs.Changed() && cs.state != nil {
		fn(*cs.state)
	}
	return cs
}

// NewCheckboxGroup wraps check boxes in a group, as shadcn's checkbox
// group: assistive technology announces the group and its label.
func NewCheckboxGroup(c *ui.Context, label string, build func()) ui.Element {
	return ui.CheckboxGroup(c, label, build)
}
