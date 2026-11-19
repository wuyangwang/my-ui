package forms

import (
	. "github.com/wuyangwang/my-ui"
	"strconv"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// field is the frame of a text field: the frame that carries the look
// of the preset, the text inside it, the slots beside the text (the
// prefix, the suffix and the clear button) and the message under it.
//
// The type parameter is the widget that wraps the field, so that the
// presets return that widget and the chain goes on: an Input's chain
// ends in *Input, a TextArea's in *TextArea.
type Frame[T any] struct {
	ui.Element
	Self     T
	c        *ui.Context
	Inner    ui.Element // the text itself
	frame    ui.Element // the frame around the text
	mid      ui.Element // the slot of the text
	pre      ui.Element // the slot before the text
	suf      ui.Element // the slot after the text
	clear    ui.Element // the slot of the clear button
	count    ui.Element // the slot of the count
	msg      ui.Element // the slot of the error message
	value    *string
	variant  Variant
	size     Size
	shape    Shape
	invalid  bool
	readOnly bool
	cleared  bool
	limit    int
	counted  bool
}

// NewFrame builds the frame of a text field. The text itself is made
// by make, so that a one-line input and a text area share the frame.
// The host is the slot the text goes in, for a control that keeps a
// state of its own there.
func NewFrame[T any](c *ui.Context, value *string, make func(c *ui.Context, value *string, host ui.Element) ui.Element) *Frame[T] {
	f := &Frame[T]{c: c, value: value}
	// The field holds the frame and the message under it.
	column := ui.Column(c)
	column.Children(func() {
		f.frame = ui.Row(c).AlignItems(ui.Center).FillWidth()
		f.msg = ui.Column(c)
	})

	// The frame holds the slots: the text grows between them.
	f.frame.Children(func() {
		f.pre = ui.Row(c).AlignItems(ui.Center)
		f.mid = ui.Row(c).AlignItems(ui.Center).Grow(1)
		f.count = ui.Row(c).AlignItems(ui.Center)
		f.suf = ui.Row(c).AlignItems(ui.Center)
		f.clear = ui.Row(c).AlignItems(ui.Center)
	})
	f.frame = f.frame.Transition(StateTransition())

	// The text goes into the slot in the middle.
	f.mid.Children(func() {
		f.Inner = make(c, value, f.mid)
	})

	f.Element = column
	// Apply the default look and padding even for the zero variant: a
	// bare field keeps shadcn's size, not the framework's built-in one.
	f.apply()
	return f
}

// Variant applies the look of v and returns the widget, for chaining.
// The zero variant is Native, which keeps the built-in look of the
// text.
func (f *Frame[T]) Variant(v Variant) T {
	f.variant = v
	f.apply()
	return f.Self
}

// Primary is Variant(myui.Primary).
func (f *Frame[T]) Primary() T { return f.Variant(Primary) }

// Secondary is Variant(myui.Secondary), alias of Native.
func (f *Frame[T]) Secondary() T { return f.Variant(Secondary) }

// Default is Variant(myui.Default), the built-in look.
func (f *Frame[T]) Default() T { return f.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (f *Frame[T]) Native() T { return f.Variant(Native) }

// Outline is Variant(myui.Outline).
func (f *Frame[T]) Outline() T { return f.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (f *Frame[T]) Ghost() T { return f.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (f *Frame[T]) Destructive() T { return f.Variant(Destructive) }

// Link is Variant(myui.Link).
func (f *Frame[T]) Link() T { return f.Variant(Link) }

// Success is Variant(myui.Success).
func (f *Frame[T]) Success() T { return f.Variant(Success) }

// Info is Variant(myui.Info).
func (f *Frame[T]) Info() T { return f.Variant(Info) }

// Warning is Variant(myui.Warning).
func (f *Frame[T]) Warning() T { return f.Variant(Warning) }

// Size applies the sizing of s and returns the widget, for chaining.
func (f *Frame[T]) Size(s Size) T {
	f.size = s
	f.apply()
	return f.Self
}

// Control returns the element a field describes: the text inside the
// frame, which is the one that takes the focus and the one the label,
// the description and the error speak of.
func (f *Frame[T]) Control() ui.Element { return f.Inner }

// MaxLength keeps the user from making the value longer than n, and
// returns the widget, for chaining. The count is in runes, as HTML's
// maxlength counts them is in code units — a rune is one character of
// the text as the user sees it.
func (f *Frame[T]) MaxLength(n int) T {
	f.limit = n
	f.Inner = f.Inner.MaxLength(n)
	return f.Self
}

// Count shows how long the value is, and how long it may be, and
// returns the widget, for chaining. The count is in runes.
func (f *Frame[T]) Count() T {
	f.counted = true
	c := f.c
	t := c.Theme()
	f.count.Children(func() {
		n := len([]rune(*f.value))
		text := ""
		if f.limit > 0 {
			text = itoa(n) + "/" + itoa(f.limit)
		} else {
			text = itoa(n)
		}
		ui.Text(c, text).TextColor(t.TextMuted).FontSize(TextSize(t, TextXs))
	})
	return f.Self
}

// itoa is the number as a string.
func itoa(n int) string { return strconv.Itoa(n) }

// Width sets the width of the field in DIPs and returns the widget,
// for chaining.
func (f *Frame[T]) Width(w float32) T {
	f.frame = f.frame.Width(w)
	f.Element = f.Element.Width(w)
	return f.Self
}

// Shape sets the shape of the field and returns the widget, for
// chaining: a pill or a circle, or the theme's own corners.
func (f *Frame[T]) Shape(s Shape) T {
	f.shape = s
	f.apply()
	return f.Self
}

// Sm sizes the field down.
func (f *Frame[T]) Sm() T { return f.Size(SizeSm) }

// Xs sizes the field to the smallest size.
func (f *Frame[T]) Xs() T { return f.Size(SizeXs) }

// Md gives the built-in sizing.
func (f *Frame[T]) Md() T { return f.Size(SizeMd) }

// Lg sizes the field up.
func (f *Frame[T]) Lg() T { return f.Size(SizeLg) }

// Placeholder shows s in the empty text and returns the widget, for
// chaining.
func (f *Frame[T]) Placeholder(s string) T {
	f.Inner = f.Inner.Placeholder(s)
	return f.Self
}

// Label names the field for assistive technology and returns it, for
// chaining.
func (f *Frame[T]) Label(s string) T {
	f.Inner = f.Inner.Label(s)
	return f.Self
}

// Invalid marks the text as not holding a good value and returns the
// widget, for chaining. The frame takes the danger color of the theme.
func (f *Frame[T]) Invalid(on bool) T {
	f.invalid = on
	f.Inner = f.Inner.Invalid(on)
	f.apply()
	return f.Self
}

// ReadOnly keeps the text from being edited and returns the widget,
// for chaining.
func (f *Frame[T]) ReadOnly(on bool) T {
	f.readOnly = on
	f.Inner = f.Inner.ReadOnly(on)
	return f.Self
}

// Disabled takes the field out of use and returns it, for chaining.
func (f *Frame[T]) Disabled(on bool) T {
	f.Inner = f.Inner.Disabled(on)
	return f.Self
}

// Prefix builds the content before the text, an icon or a word, and
// returns the widget, for chaining.
func (f *Frame[T]) Prefix(fn func()) T {
	if fn != nil {
		f.pre.Children(fn)
	}
	return f.Self
}

// PrefixIcon puts an icon before the text and returns the widget, for
// chaining. The icon takes the color of the text and the size of it.
func (f *Frame[T]) PrefixIcon(svg *ui.SVG) T {
	c := f.c
	f.Prefix(func() {
		ui.Icon(c, svg).Margin(0, c.Theme().Space(1.5), 0, 0)
	})
	return f.Self
}

// Suffix builds the content after the text and returns the widget, for
// chaining.
func (f *Frame[T]) Suffix(fn func()) T {
	if fn != nil {
		f.suf.Children(fn)
	}
	return f.Self
}

// SuffixIcon puts an icon after the text and returns the widget, for
// chaining.
func (f *Frame[T]) SuffixIcon(svg *ui.SVG) T {
	c := f.c
	f.Suffix(func() {
		ui.Icon(c, svg).Margin(0, 0, 0, c.Theme().Space(1.5))
	})
	return f.Self
}

// Clearable puts a button that empties the field beside it, and
// returns the widget, for chaining. The button shows while the field
// holds text and can be edited, and it keeps the focus in the field.
//
// The button is the icon alone, without the face a button has: it turns
// darker under the pointer.
func (f *Frame[T]) Clearable() T {
	c := f.c
	f.clear.Children(func() {
		if *f.value == "" || f.Inner.IsDisabled() || f.readOnly {
			return
		}
		d := float32(18)
		if f.size == SizeLg {
			d = 20
		}
		// The shared icon button carries the field's clear affordance:
		// the same glyph size and hover states as every other one.
		ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
			SVG:      icons.X(),
			Label:    Message(c, "clear"),
			Size:     d,
			IconSize: IconSizeXs,
			OnClick: func() {
				*f.value = ""
				f.cleared = true
			},
		})
	})
	return f.Self
}

// ErrorText puts the message of the error under the field and returns
// the widget, for chaining. The message is named for assistive
// technology, as the error of the field is.
func (f *Frame[T]) ErrorText(s string) T {
	c := f.c
	f.msg.Children(func() {
		msg := ui.Text(c, s).TextColor(c.Theme().Danger)
		f.Inner = f.Inner.ErrorMessage(msg)
	})
	return f.Self
}

// ErrorMessage builds the message of the error under the field and
// returns the widget, for chaining.
func (f *Frame[T]) ErrorMessage(fn func()) T {
	if fn != nil {
		f.msg.Children(fn)
	}
	return f.Self
}

// OnChanged calls fn with the new value when it changed in this frame,
// and returns the widget, for chaining.
func (f *Frame[T]) OnChanged(fn func(string)) T {
	if fn != nil && f.Changed() {
		fn(*f.value)
	}
	return f.Self
}

// OnBlurred calls fn when the focus left the field in this frame, and
// returns the widget, for chaining.
func (f *Frame[T]) OnBlurred(fn func()) T {
	if fn != nil && f.FocusLeft() {
		fn()
	}
	return f.Self
}

// OnCleared calls fn when the clear button emptied the field, and
// returns the widget, for chaining.
func (f *Frame[T]) OnCleared(fn func()) T {
	if fn != nil && f.cleared {
		fn()
	}
	return f.Self
}

// Changed reports whether the value of the field changed in this
// frame.
func (f *Frame[T]) Changed() bool { return f.Inner.Changed() }

// Focused reports whether the field has the keyboard focus.
func (f *Frame[T]) Focused() bool { return f.Inner.Focused() }

// FocusLeft reports whether the focus left the field in this frame.
func (f *Frame[T]) FocusLeft() bool { return f.Inner.FocusLeft() }

// Value returns the value of the field.
func (f *Frame[T]) Value() string { return *f.value }

// apply paints the look of the variant and the size on the field.
func (f *Frame[T]) apply() {
	t := f.c.Theme()
	l := FrameLookOf(f.variant, t)

	if l.Native {
		// The built-in look of the text, with no frame of our own.
		f.frame = f.frame.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0)
	} else {
		border := l.Border
		switch {
		case f.invalid:
			border = t.Danger
		case f.Inner.Focused():
			// The focused field takes the accent color of its border,
			// which the frame fades into.
			border = t.Accent
		}
		f.frame = f.frame.Background(l.Bg).Border(l.Width, border)
		f.frame = f.frame.Padding(f.padding())
		// The text has no border of its own: the frame has it.
		f.Inner = f.Inner.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0)
	}
	f.frame = f.frame.Radius(f.radius())
	if l.HasText {
		f.Inner = f.Inner.TextColor(l.Text)
	}
	f.Inner = f.Inner.FontSize(f.fontSize())
}

// radius is the corner radius of the frame of the field, on its shape.
func (f *Frame[T]) radius() float32 {
	t := f.c.Theme()
	switch f.shape {
	case ShapeRound, ShapeCircle:
		return 1000
	default:
		return Radius(t, RadiusMd)
	}
}

// FrameLook is the colors one variant paints for a field.
type FrameLook struct {
	Native  bool
	Bg      ui.Color
	Border  ui.Color
	Width   float32
	Text    ui.Color
	HasText bool
}

// FrameLookOf resolves one variant's look for the frame of a field.
func FrameLookOf(v Variant, t *ui.Theme) FrameLook {
	switch v {
	case Default:
		return FrameLook{Native: true}
	case Secondary:
		return FrameLook{Bg: t.Surface, Border: ui.Transparent, Width: 0}
	case Primary:
		return FrameLook{Bg: t.Background, Border: t.Accent, Width: 1}
	case Outline:
		return FrameLook{Bg: t.Background, Border: t.Border, Width: 1}
	case Ghost:
		return FrameLook{Bg: ui.Transparent, Border: ui.Transparent, Width: 0}
	case Destructive:
		return FrameLook{Bg: t.Background, Border: t.Danger, Width: 1}
	case Link:
		return FrameLook{Bg: ui.Transparent, Border: ui.Transparent, Width: 0, Text: t.Accent, HasText: true}
	case Success, Info, Warning:
		role, _ := SemanticOf(v)
		return FrameLook{Bg: t.Background, Border: SemanticColor(t, role), Width: 1}
	}
	return FrameLook{Native: true}
}

// padding is the inset of the frame of the field, on its size.
func (f *Frame[T]) padding() (v, h float32) { return FramePadding(f.c.Theme(), f.size) }

// fontSize is the size of the text of the field, on its size.
func (f *Frame[T]) fontSize() float32 { return FrameFontSize(f.c.Theme(), f.size) }

// FramePadding is the inset of the frame of a field, on its size.
func FramePadding(t *ui.Theme, size Size) (v, h float32) {
	switch size {
	case SizeXs:
		return t.Space(1), t.Space(2)
	case SizeSm:
		return t.Space(1.25), t.Space(2.5)
	case SizeLg:
		return t.Space(3), t.Space(4)
	default:
		return t.Space(2), t.Space(3)
	}
}

// FrameFontSize is the size of the text of a field, on its size.
func FrameFontSize(t *ui.Theme, size Size) float32 {
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

// Context returns the frame's context, so a builder that composes a field
// can pass it on to the controls it puts in the field.
func (f *Frame[T]) Context() *ui.Context { return f.c }
