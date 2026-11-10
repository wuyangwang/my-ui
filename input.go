package myui

import "github.com/egoist/mygo/ui"

// Input is a single-line text input with variant and size presets:
//
//	myui.NewInput(c, &value).Placeholder("Email").Outline()
//	myui.NewInput(c, &value).Sm().Destructive()
//
// The chain applies presets from the current theme every frame, so it
// follows palette and appearance changes. Call it each frame before any
// app-specific styling. Input embeds ui.Element, so Placeholder,
// Disabled, ReadOnly, Invalid, Password and other ui text-input methods
// remain available.
type Input struct {
	ui.Element
	c    *ui.Context
	size Size
}

// NewInput creates a text input editing *value with the native preset.
func NewInput(c *ui.Context, value *string) *Input {
	in := &Input{Element: ui.TextInput(c, value), c: c}
	// Size to content by default; a column stretches children across.
	in.AlignSelf(ui.Start)
	return in
}

// look is the colors one variant paints for an input on a theme.
type inputLook struct {
	native bool
	bg     ui.Color
	border ui.Color
	width  float32
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
		return inputLook{native: true}
	}
	return inputLook{native: true}
}

// Variant applies the look of v and returns the input, for chaining.
func (in *Input) Variant(v Variant) *Input {
	in.applySizing()
	if l := inputLookOf(v, in.c.Theme()); !l.native {
		in.Background(l.bg).Border(l.width, l.border)
	}
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

// Link is Variant(myui.Link); for inputs it behaves as Native.
func (in *Input) Link() *Input { return in.Variant(Link) }

// Size applies the sizing of s and returns the input, for chaining.
func (in *Input) Size(s Size) *Input {
	in.size = s
	in.applySizing()
	return in
}

// Sm sizes the input down.
func (in *Input) Sm() *Input { return in.Size(SizeSm) }

// Md gives the built-in sizing.
func (in *Input) Md() *Input { return in.Size(SizeMd) }

// Lg sizes the input up.
func (in *Input) Lg() *Input { return in.Size(SizeLg) }

func (in *Input) applySizing() {
	t := in.c.Theme()
	switch in.size {
	case SizeSm:
		in.Padding(t.Space(1), t.Space(2)).FontSize(TextSize(t, TextSm))
	case SizeLg:
		in.Padding(t.Space(2.5), t.Space(4)).FontSize(TextSize(t, TextLg))
	default:
		in.Padding(t.Space(1.5), t.Space(3)).FontSize(t.FontSize)
	}
}
