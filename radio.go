package myui

import (
	"github.com/egoist/mygo/ui"
)

// Radio is a circle that is selected or not, with a label: one of the
// choices of a group.
//
//	myui.NewRadio(c, &app.size, "large", "大")
//
// The circle is styled by the library, as shadcn's: a ring, and a dot
// in it when it is selected. The choices of a group share one value:
// the one that is selected holds it.
type Radio[T comparable] struct {
	ui.Element
	c        *ui.Context
	dot      ui.Element
	selected *T
	value    T
	size     Size
	invalid  bool
}

// NewRadio creates a radio button that holds value while it is
// selected, with a label.
func NewRadio[T comparable](c *ui.Context, selected *T, value T, label string) *Radio[T] {
	r := &Radio[T]{c: c, selected: selected, value: value}
	t := c.Theme()
	row := ui.RadioBase(c, selected, value).Gap(t.Space(2))
	row.Children(func() {
		r.dot = ui.Box(c)
		ui.Text(c, label)
	})
	r.Element = row
	r.applyLook()
	return r
}

// applyLook paints the circle as the state says.
func (r *Radio[T]) applyLook() {
	t := r.c.Theme()
	d := r.dotSize()
	border := t.Border
	fill := t.Accent
	if r.invalid {
		border, fill = t.Danger, t.Danger
	}
	r.dot.Size(d, d).Radius(d/2).Border(1, border)
	if r.selected != nil && *r.selected == r.value {
		r.dot.Border(1, fill)
		r.dot.Children(func() {
			ui.Box(r.c).Size(d-8, d-8).Radius((d - 8) / 2).Background(fill)
		})
	}
}

// dotSize is the size of the circle in DIPs.
func (r *Radio[T]) dotSize() float32 {
	switch r.size {
	case SizeXs:
		return 14
	case SizeLg:
		return 22
	case SizeSm:
		return 16
	}
	return 18
}

// Size applies the sizing of s and returns the button, for chaining.
func (r *Radio[T]) Size(s Size) *Radio[T] {
	r.size = s
	r.applyLook()
	return r
}

// Sm sizes the circle down.
func (r *Radio[T]) Sm() *Radio[T] { return r.Size(SizeSm) }

// Xs sizes the circle to the smallest size.
func (r *Radio[T]) Xs() *Radio[T] { return r.Size(SizeXs) }

// Md gives the circle the size of the text around it.
func (r *Radio[T]) Md() *Radio[T] { return r.Size(SizeMd) }

// Lg sizes the circle up.
func (r *Radio[T]) Lg() *Radio[T] { return r.Size(SizeLg) }

// OnChanged runs fn when the button is selected in this frame, and
// returns the button, for chaining.
func (r *Radio[T]) OnChanged(fn func(value T)) *Radio[T] {
	if fn != nil && r.Changed() {
		fn(r.value)
	}
	return r
}

// Disabled dims the button and takes its clicks away, and returns the
// button, for chaining.
func (r *Radio[T]) Disabled(disabled bool) *Radio[T] {
	r.Element = r.Element.Disabled(disabled)
	return r
}

// Invalid marks the radio as not holding a good value and returns it, for
// chaining. The ring takes the danger color of the theme.
func (r *Radio[T]) Invalid(on bool) *Radio[T] {
	r.invalid = on
	r.Element = r.Element.Invalid(on)
	r.applyLook()
	return r
}

// Control returns the element a field describes: the radio itself.
func (r *Radio[T]) Control() ui.Element { return r.Element }
