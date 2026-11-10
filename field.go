package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// field is the frame of a text field: the frame that carries the look
// of the preset, the text inside it, the slots beside the text (the
// prefix, the suffix and the clear button) and the message under it.
//
// The type parameter is the widget that wraps the field, so that the
// presets return that widget and the chain goes on: an Input's chain
// ends in *Input, a TextArea's in *TextArea.
type field[T any] struct {
	ui.Element
	self     T
	c        *ui.Context
	inner    ui.Element // the text itself
	frame    ui.Element // the frame around the text
	mid      ui.Element // the slot of the text
	pre      ui.Element // the slot before the text
	suf      ui.Element // the slot after the text
	clear    ui.Element // the slot of the clear button
	msg      ui.Element // the slot of the error message
	value    *string
	variant  Variant
	size     Size
	shape    Shape
	invalid  bool
	readOnly bool
	cleared  bool
}

// newField builds the frame of a text field. The text itself is made
// by make, so that a one-line input and a text area share the frame.
// The host is the slot the text goes in, for a control that keeps a
// state of its own there.
func newField[T any](c *ui.Context, value *string, make func(c *ui.Context, value *string, host ui.Element) ui.Element) *field[T] {
	f := &field[T]{c: c, value: value}
	// The field holds the frame and the message under it.
	column := ui.Column(c)
	column.Children(func() {
		f.frame = ui.Row(c).AlignItems(ui.Center)
		f.msg = ui.Column(c)
	})

	// The frame holds the slots: the text grows between them.
	f.frame.Children(func() {
		f.pre = ui.Row(c).AlignItems(ui.Center)
		f.mid = ui.Row(c).AlignItems(ui.Center).Grow(1)
		f.suf = ui.Row(c).AlignItems(ui.Center)
		f.clear = ui.Row(c).AlignItems(ui.Center)
	})
	f.frame = f.frame.Transition(stateTransition())

	// The text goes into the slot in the middle.
	f.mid.Children(func() {
		f.inner = make(c, value, f.mid)
	})

	f.Element = column
	return f
}

// Variant applies the look of v and returns the widget, for chaining.
// The zero variant is Native, which keeps the built-in look of the
// text.
func (f *field[T]) Variant(v Variant) T {
	f.variant = v
	f.apply()
	return f.self
}

// Primary is Variant(myui.Primary).
func (f *field[T]) Primary() T { return f.Variant(Primary) }

// Secondary is Variant(myui.Secondary), alias of Native.
func (f *field[T]) Secondary() T { return f.Variant(Secondary) }

// Default is Variant(myui.Default), the built-in look.
func (f *field[T]) Default() T { return f.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (f *field[T]) Native() T { return f.Variant(Native) }

// Outline is Variant(myui.Outline).
func (f *field[T]) Outline() T { return f.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (f *field[T]) Ghost() T { return f.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (f *field[T]) Destructive() T { return f.Variant(Destructive) }

// Link is Variant(myui.Link).
func (f *field[T]) Link() T { return f.Variant(Link) }

// Success is Variant(myui.Success).
func (f *field[T]) Success() T { return f.Variant(Success) }

// Info is Variant(myui.Info).
func (f *field[T]) Info() T { return f.Variant(Info) }

// Warning is Variant(myui.Warning).
func (f *field[T]) Warning() T { return f.Variant(Warning) }

// Size applies the sizing of s and returns the widget, for chaining.
func (f *field[T]) Size(s Size) T {
	f.size = s
	f.apply()
	return f.self
}

// Control returns the element a field describes: the text inside the
// frame, which is the one that takes the focus and the one the label,
// the description and the error speak of.
func (f *field[T]) Control() ui.Element { return f.inner }

// Width sets the width of the field in DIPs and returns the widget,
// for chaining.
func (f *field[T]) Width(w float32) T {
	f.frame = f.frame.Width(w)
	return f.self
}

// Shape sets the shape of the field and returns the widget, for
// chaining: a pill or a circle, or the theme's own corners.
func (f *field[T]) Shape(s Shape) T {
	f.shape = s
	f.apply()
	return f.self
}

// Sm sizes the field down.
func (f *field[T]) Sm() T { return f.Size(SizeSm) }

// Xs sizes the field to the smallest size.
func (f *field[T]) Xs() T { return f.Size(SizeXs) }

// Md gives the built-in sizing.
func (f *field[T]) Md() T { return f.Size(SizeMd) }

// Lg sizes the field up.
func (f *field[T]) Lg() T { return f.Size(SizeLg) }

// Placeholder shows s in the empty text and returns the widget, for
// chaining.
func (f *field[T]) Placeholder(s string) T {
	f.inner = f.inner.Placeholder(s)
	return f.self
}

// Label names the field for assistive technology and returns it, for
// chaining.
func (f *field[T]) Label(s string) T {
	f.inner = f.inner.Label(s)
	return f.self
}

// Invalid marks the text as not holding a good value and returns the
// widget, for chaining. The frame takes the danger color of the theme.
func (f *field[T]) Invalid(on bool) T {
	f.invalid = on
	f.inner = f.inner.Invalid(on)
	f.apply()
	return f.self
}

// ReadOnly keeps the text from being edited and returns the widget,
// for chaining.
func (f *field[T]) ReadOnly(on bool) T {
	f.readOnly = on
	f.inner = f.inner.ReadOnly(on)
	return f.self
}

// Disabled takes the field out of use and returns it, for chaining.
func (f *field[T]) Disabled(on bool) T {
	f.inner = f.inner.Disabled(on)
	return f.self
}

// Prefix builds the content before the text, an icon or a word, and
// returns the widget, for chaining.
func (f *field[T]) Prefix(fn func()) T {
	if fn != nil {
		f.pre.Children(fn)
	}
	return f.self
}

// PrefixIcon puts an icon before the text and returns the widget, for
// chaining. The icon takes the color of the text and the size of it.
func (f *field[T]) PrefixIcon(svg *ui.SVG) T {
	c := f.c
	f.Prefix(func() {
		ui.Icon(c, svg).Margin(0, c.Theme().Space(1.5), 0, 0)
	})
	return f.self
}

// Suffix builds the content after the text and returns the widget, for
// chaining.
func (f *field[T]) Suffix(fn func()) T {
	if fn != nil {
		f.suf.Children(fn)
	}
	return f.self
}

// SuffixIcon puts an icon after the text and returns the widget, for
// chaining.
func (f *field[T]) SuffixIcon(svg *ui.SVG) T {
	c := f.c
	f.Suffix(func() {
		ui.Icon(c, svg).Margin(0, 0, 0, c.Theme().Space(1.5))
	})
	return f.self
}

// Clearable puts a button that empties the field beside it, and
// returns the widget, for chaining. The button shows while the field
// holds text and can be edited, and it keeps the focus in the field.
//
// The button is the icon alone, without the face a button has: it turns
// darker under the pointer.
func (f *field[T]) Clearable() T {
	c := f.c
	f.clear.Children(func() {
		if *f.value == "" || f.inner.IsDisabled() || f.readOnly {
			return
		}
		d := float32(18)
		if f.size == SizeLg {
			d = 20
		}
		t := c.Theme()
		// The button keeps the focus in the field while it is pressed, and
		// it is a stop of the Tab key, as a clear button is in Ant Design.
		b := ui.ButtonBase(c).Size(d, d).KeepFocus().Label("清除")
		b.Children(func() {
			ic := ui.Icon(c, icons.X()).Size(d-4, d-4)
			if b.Hovered() {
				ic = ic.TextColor(t.Text)
			} else {
				ic = ic.TextColor(t.TextMuted)
			}
		})
		if b.Clicked() {
			*f.value = ""
			f.cleared = true
		}
	})
	return f.self
}

// ErrorText puts the message of the error under the field and returns
// the widget, for chaining. The message is named for assistive
// technology, as the error of the field is.
func (f *field[T]) ErrorText(s string) T {
	c := f.c
	f.msg.Children(func() {
		msg := ui.Text(c, s).TextColor(c.Theme().Danger)
		f.inner = f.inner.ErrorMessage(msg)
	})
	return f.self
}

// ErrorMessage builds the message of the error under the field and
// returns the widget, for chaining.
func (f *field[T]) ErrorMessage(fn func()) T {
	if fn != nil {
		f.msg.Children(fn)
	}
	return f.self
}

// OnChanged calls fn with the new value when it changed in this frame,
// and returns the widget, for chaining.
func (f *field[T]) OnChanged(fn func(string)) T {
	if fn != nil && f.Changed() {
		fn(*f.value)
	}
	return f.self
}

// OnBlurred calls fn when the focus left the field in this frame, and
// returns the widget, for chaining.
func (f *field[T]) OnBlurred(fn func()) T {
	if fn != nil && f.FocusLeft() {
		fn()
	}
	return f.self
}

// OnCleared calls fn when the clear button emptied the field, and
// returns the widget, for chaining.
func (f *field[T]) OnCleared(fn func()) T {
	if fn != nil && f.cleared {
		fn()
	}
	return f.self
}

// Changed reports whether the value of the field changed in this
// frame.
func (f *field[T]) Changed() bool { return f.inner.Changed() }

// Focused reports whether the field has the keyboard focus.
func (f *field[T]) Focused() bool { return f.inner.Focused() }

// FocusLeft reports whether the focus left the field in this frame.
func (f *field[T]) FocusLeft() bool { return f.inner.FocusLeft() }

// Value returns the value of the field.
func (f *field[T]) Value() string { return *f.value }

// apply paints the look of the variant and the size on the field.
func (f *field[T]) apply() {
	t := f.c.Theme()
	l := fieldLookOf(f.variant, t)

	if l.native {
		// The built-in look of the text, with no frame of our own.
		f.frame = f.frame.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0)
	} else {
		border := l.border
		switch {
		case f.invalid:
			border = t.Danger
		case f.inner.Focused():
			// The focused field takes the accent color of its border,
			// which the frame fades into.
			border = t.Accent
		}
		f.frame = f.frame.Background(l.bg).Border(l.width, border)
		f.frame = f.frame.Padding(f.padding())
		// The text has no border of its own: the frame has it.
		f.inner = f.inner.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0)
	}
	f.frame = f.frame.Radius(f.radius())
	if l.hasText {
		f.inner = f.inner.TextColor(l.text)
	}
	f.inner = f.inner.FontSize(f.fontSize())
}

// radius is the corner radius of the frame of the field, on its shape.
func (f *field[T]) radius() float32 {
	t := f.c.Theme()
	switch f.shape {
	case ShapeRound, ShapeCircle:
		return 1000
	default:
		return t.Radius
	}
}

// fieldLook is the colors one variant paints for a field.
type fieldLook struct {
	native  bool
	bg      ui.Color
	border  ui.Color
	width   float32
	text    ui.Color
	hasText bool
}

// fieldLookOf resolves one variant's look for the frame of a field.
func fieldLookOf(v Variant, t *ui.Theme) fieldLook {
	switch v {
	case Default:
		return fieldLook{native: true}
	case Secondary:
		return fieldLook{bg: t.Surface, border: ui.Transparent, width: 0}
	case Primary:
		return fieldLook{bg: t.Background, border: t.Accent, width: 1}
	case Outline:
		return fieldLook{bg: t.Background, border: t.Border, width: 1}
	case Ghost:
		return fieldLook{bg: ui.Transparent, border: ui.Transparent, width: 0}
	case Destructive:
		return fieldLook{bg: t.Background, border: t.Danger, width: 1}
	case Link:
		return fieldLook{bg: ui.Transparent, border: ui.Transparent, width: 0, text: t.Accent, hasText: true}
	case Success, Info, Warning:
		role, _ := semanticOf(v)
		return fieldLook{bg: t.Background, border: semanticColor(t, role), width: 1}
	}
	return fieldLook{native: true}
}

// padding is the inset of the frame of the field, on its size.
func (f *field[T]) padding() (v, h float32) { return fieldPadding(f.c.Theme(), f.size) }

// fontSize is the size of the text of the field, on its size.
func (f *field[T]) fontSize() float32 { return fieldFontSize(f.c.Theme(), f.size) }

// fieldPadding is the inset of the frame of a field, on its size.
func fieldPadding(t *ui.Theme, size Size) (v, h float32) {
	switch size {
	case SizeXs:
		return t.Space(0.5), t.Space(2)
	case SizeSm:
		return t.Space(1), t.Space(2.5)
	case SizeLg:
		return t.Space(2.5), t.Space(4)
	default:
		return t.Space(1.5), t.Space(3)
	}
}

// fieldFontSize is the size of the text of a field, on its size.
func fieldFontSize(t *ui.Theme, size Size) float32 {
	switch size {
	case SizeXs:
		return TextSize(t, TextXs)
	case SizeSm:
		return TextSize(t, TextSm)
	case SizeLg:
		return TextSize(t, TextLg)
	default:
		return t.FontSize
	}
}
