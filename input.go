package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// Input is a text input with variant, size and events:
//
//	myui.NewInput(c, &value).Outline().Placeholder("Email")
//	myui.NewInput(c, &value).Clearable().OnChanged(func(v string) { ... })
//
// The input is a field: a frame with the text in it, and the prefix,
// the suffix and the clear button beside it. The frame carries the look
// of the preset, so the text itself has no border of its own.
//
// The chain applies the presets from the current theme every frame, so
// it follows palette and appearance changes. Call it every frame, ahead
// of anything else the app styles on the field, as the view runs.
type Input struct {
	ui.Element // the field: the frame, and the message under it
	frame      ui.Element // the frame around the text
	in         ui.Element // the text input itself
	mid        ui.Element // the slot of the text
	pre        ui.Element // the slot before the text
	suf        ui.Element // the slot after the text
	clear      ui.Element // the slot of the clear button
	msg        ui.Element // the slot of the error message
	value      *string
	c          *ui.Context
	variant    Variant
	size       Size
	invalid    bool
	readOnly   bool
	cleared    bool
	onChanged  func(string)
	onBlurred  func()
	onCleared  func()
	onSubmit   func()
}

// NewInput creates a text input editing *value with the native preset.
func NewInput(c *ui.Context, value *string) *Input {
	in := &Input{value: value, c: c}

	// The field holds the frame and the message under it.
	field := ui.Column(c)
	field.Children(func() {
		in.frame = ui.Row(c).AlignItems(ui.Center)
		in.msg = ui.Column(c)
	})

	// The frame holds the slots: the text grows between them.
	in.frame.Children(func() {
		in.pre = ui.Row(c).AlignItems(ui.Center)
		in.mid = ui.Row(c).AlignItems(ui.Center).Grow(1)
		in.suf = ui.Row(c).AlignItems(ui.Center)
		in.clear = ui.Row(c).AlignItems(ui.Center)
	})
	in.frame = in.frame.Transition(stateTransition())

	// The text goes into the slot in the middle.
	in.mid.Children(func() {
		in.in = ui.TextInput(c, value)
	})

	in.Element = field
	return in
}

// inputLook is the colors one variant paints for an input.
type inputLook struct {
	native  bool
	bg      ui.Color
	border  ui.Color
	width   float32
	text    ui.Color
	hasText bool
}

func inputLookOf(v Variant, t *ui.Theme) inputLook {
	switch v {
	case Native:
		return inputLook{native: true}
	case Primary:
		return inputLook{bg: t.Background, border: t.Accent, width: 1}
	case Outline:
		return inputLook{bg: t.Background, border: t.Border, width: 1}
	case Ghost:
		return inputLook{bg: ui.Transparent, border: ui.Transparent, width: 0}
	case Destructive:
		return inputLook{bg: t.Background, border: t.Danger, width: 1}
	case Link:
		return inputLook{bg: ui.Transparent, border: ui.Transparent, width: 0, text: t.Accent, hasText: true}
	}
	return inputLook{native: true}
}

// Variant applies the look of v and returns the input, for chaining.
// The zero variant is Native, which keeps the built-in look of the
// text input.
func (in *Input) Variant(v Variant) *Input {
	in.variant = v
	in.apply()
	return in
}

// Primary is Variant(myui.Primary).
func (in *Input) Primary() *Input { return in.Variant(Primary) }

// Secondary is Variant(myui.Secondary), alias of Native.
func (in *Input) Secondary() *Input { return in.Variant(Secondary) }

// Native is Variant(myui.Native).
func (in *Input) Native() *Input { return in.Variant(Native) }

// Outline is Variant(myui.Outline).
func (in *Input) Outline() *Input { return in.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (in *Input) Ghost() *Input { return in.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (in *Input) Destructive() *Input { return in.Variant(Destructive) }

// Link is Variant(myui.Link).
func (in *Input) Link() *Input { return in.Variant(Link) }

// Size applies the sizing of s and returns the input, for chaining.
func (in *Input) Size(s Size) *Input {
	in.size = s
	in.apply()
	return in
}

// Sm sizes the input down.
func (in *Input) Sm() *Input { return in.Size(SizeSm) }

// Md gives the built-in sizing.
func (in *Input) Md() *Input { return in.Size(SizeMd) }

// Lg sizes the input up.
func (in *Input) Lg() *Input { return in.Size(SizeLg) }

// Placeholder shows s in the empty input and returns it, for chaining.
func (in *Input) Placeholder(s string) *Input {
	in.in = in.in.Placeholder(s)
	return in
}

// Label names the input for assistive technology and returns it, for
// chaining.
func (in *Input) Label(s string) *Input {
	in.in = in.in.Label(s)
	return in
}

// Invalid marks the input as not holding a good value and returns it,
// for chaining. The frame takes the danger color of the theme.
func (in *Input) Invalid(on bool) *Input {
	in.invalid = on
	in.in = in.in.Invalid(on)
	in.apply()
	return in
}

// ReadOnly keeps the text of the input from being edited and returns
// it, for chaining.
func (in *Input) ReadOnly(on bool) *Input {
	in.readOnly = on
	in.in = in.in.ReadOnly(on)
	return in
}

// Disabled takes the input out of use and returns it, for chaining.
func (in *Input) Disabled(on bool) *Input {
	in.in = in.in.Disabled(on)
	return in
}

// Password hides what the input holds and returns it, for chaining.
func (in *Input) Password() *Input {
	in.in = in.in.Password()
	return in
}

// Prefix builds the content before the text, an icon or a word, and
// returns the input, for chaining.
func (in *Input) Prefix(fn func()) *Input {
	if fn != nil {
		in.pre.Children(fn)
	}
	return in
}

// PrefixIcon puts an icon before the text of the input and returns the
// input, for chaining. The icon takes the color of the text and the
// size of it, as an icon in a button does.
func (in *Input) PrefixIcon(svg *ui.SVG) *Input {
	c := in.c
	in.Prefix(func() {
		ui.Icon(c, svg).Margin(0, c.Theme().Space(1.5), 0, 0)
	})
	return in
}

// Suffix builds the content after the text and returns the input, for
// chaining.
func (in *Input) Suffix(fn func()) *Input {
	if fn != nil {
		in.suf.Children(fn)
	}
	return in
}

// SuffixIcon puts an icon after the text of the input and returns the
// input, for chaining. The icon takes the color of the text and the
// size of it, as an icon in a button does.
func (in *Input) SuffixIcon(svg *ui.SVG) *Input {
	c := in.c
	in.Suffix(func() {
		ui.Icon(c, svg).Margin(0, 0, 0, c.Theme().Space(1.5))
	})
	return in
}

// Clearable puts a button that empties the input beside it, and returns
// the input, for chaining. The button shows while the input holds text
// and can be edited, and it keeps the focus in the input.
//
// The button is the icon alone, without the face a button has: it turns
// darker under the pointer.
func (in *Input) Clearable() *Input {
	c := in.c
	in.clear.Children(func() {
		if *in.value == "" || in.in.IsDisabled() || in.readOnly {
			return
		}
		d := float32(18)
		if in.size == SizeLg {
			d = 20
		}
		t := c.Theme()
		b := ui.ButtonBase(c).Size(d, d).KeepFocus().TabStop(false).Label("清除")
		b.Children(func() {
			ic := ui.Icon(c, icons.X()).Size(d-4, d-4)
			if b.Hovered() {
				ic = ic.TextColor(t.Text)
			} else {
				ic = ic.TextColor(t.TextMuted)
			}
		})
		// The button was not there before, so it fades in.
		fadeIn(b)
		if b.Clicked() {
			*in.value = ""
			in.cleared = true
		}
	})
	return in
}

// ErrorText puts the message of the error under the input and returns
// it, for chaining. The message is named for assistive technology, as
// the error of the input is.
func (in *Input) ErrorText(s string) *Input {
	c := in.c
	in.msg.Children(func() {
		msg := ui.Text(c, s).TextColor(c.Theme().Danger)
		// The message was not there before, so it fades in.
		fadeIn(msg)
		in.in = in.in.ErrorMessage(msg)
	})
	return in
}

// ErrorMessage builds the message of the error under the input and
// returns the input, for chaining.
func (in *Input) ErrorMessage(fn func()) *Input {
	if fn != nil {
		in.msg.Children(fn)
	}
	return in
}

// OnChanged calls fn with the new value of the input when it changed in
// this frame, and returns the input, for chaining.
func (in *Input) OnChanged(fn func(string)) *Input {
	in.onChanged = fn
	if fn != nil && in.Changed() {
		fn(*in.value)
	}
	return in
}

// OnBlurred calls fn when the focus left the input in this frame, and
// returns the input, for chaining.
func (in *Input) OnBlurred(fn func()) *Input {
	in.onBlurred = fn
	if fn != nil && in.FocusLeft() {
		fn()
	}
	return in
}

// OnCleared calls fn when the clear button emptied the input, and
// returns the input, for chaining.
func (in *Input) OnCleared(fn func()) *Input {
	in.onCleared = fn
	if fn != nil && in.cleared {
		fn()
	}
	return in
}

// OnSubmit calls fn when the user presses Enter in the input, and
// returns the input, for chaining.
func (in *Input) OnSubmit(fn func()) *Input {
	in.onSubmit = fn
	if fn != nil {
		in.in = in.in.OnTextKey(func(k ui.TextKey) bool {
			if k.Key == ui.KeyEnter {
				fn()
			}
			return false
		})
	}
	return in
}

// Changed reports whether the value of the input changed in this
// frame.
func (in *Input) Changed() bool { return in.in.Changed() }

// Focused reports whether the input has the keyboard focus.
func (in *Input) Focused() bool { return in.in.Focused() }

// FocusLeft reports whether the focus left the input in this frame.
func (in *Input) FocusLeft() bool { return in.in.FocusLeft() }

// Value returns the value of the input.
func (in *Input) Value() string { return *in.value }

// apply paints the look of the variant and the size on the field.
func (in *Input) apply() {
	t := in.c.Theme()
	l := inputLookOf(in.variant, t)

	if l.native {
		// The built-in look of the text input, with no frame of our own.
		in.frame = in.frame.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0)
	} else {
		border := l.border
		switch {
		case in.invalid:
			border = t.Danger
		case in.in.Focused():
			// The focused input takes the accent color of its border,
			// which the frame fades into.
			border = t.Accent
		}
		in.frame = in.frame.Background(l.bg).Border(l.width, border).Radius(t.Radius)
		in.frame = in.frame.Padding(in.padding())
		// The text has no border of its own: the frame has it.
		in.in = in.in.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0)
	}
	if l.hasText {
		in.in = in.in.TextColor(l.text)
	}
	in.in = in.in.FontSize(in.fontSize())
}

// padding is the inset of the frame of the input, on its size.
func (in *Input) padding() (v, h float32) {
	t := in.c.Theme()
	switch in.size {
	case SizeSm:
		return t.Space(1), t.Space(2)
	case SizeLg:
		return t.Space(2.5), t.Space(4)
	default:
		return t.Space(1.5), t.Space(3)
	}
}

// fontSize is the size of the text of the input, on its size.
func (in *Input) fontSize() float32 {
	t := in.c.Theme()
	switch in.size {
	case SizeSm:
		return TextSize(t, TextSm)
	case SizeLg:
		return TextSize(t, TextLg)
	default:
		return t.FontSize
	}
}
