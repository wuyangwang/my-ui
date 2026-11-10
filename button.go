package myui

import (
	"github.com/egoist/mygo/ui"
)

// Variant names what a Button looks like, the shadcn set. The zero
// value is Native: no preset is painted, so the button keeps the
// built-in look of the underlying ui.Button (which follows the theme).
type Variant uint8

const (
	// Native leaves the button's built-in look untouched. Zero value.
	Native Variant = iota
	// Primary fills with the accent, as ui.PrimaryButton does.
	Primary
	// Outline is a bordered face over the background, muting as it
	// hovers.
	Outline
	// Ghost is bare until hovered.
	Ghost
	// Destructive fills with the theme's Danger, as a Delete button.
	Destructive
	// Link looks like a link: the accent's text over nothing, a faint
	// accent under the pointer.
	Link
)

// Secondary is the shadcn "secondary" variant: an alias of Native, so
// it keeps the button's built-in look.
const Secondary Variant = Native

// Size names how roomy a Button is.
type Size uint8

const (
	// SizeMd is the built-in sizing. The zero value.
	SizeMd Size = iota
	SizeSm
	SizeLg
)

// Button is a ui button with a variant's preset look:
//
//	myui.NewButton(c, "Save").Primary()
//	myui.NewButton(c, "Remove").Destructive().Sm()
//
// The chain applies the preset from the build's theme, so it keeps
// following palette and appearance changes; call it every frame, ahead
// of anything else the app styles on the button, as the view runs. What
// a variant does not set keeps the built-in button's look, and every
// method of ui.Element is on hand: Clicked, Disabled (which dims the
// button as the theme dims disabled ones), Children and the rest.
type Button struct {
	ui.Element
	c    *ui.Context
	size Size
}

// NewButton creates a button with the native preset: myui.NewButton(c,
// "Save"). Clicked() reports the click on it as on ui.Button.
func NewButton(c *ui.Context, label string) *Button {
	b := &Button{Element: ui.Button(c, label), c: c}
	// Size to the content as a button does: a column stretches its
	// children across, which a button never asks for.
	b.AlignSelf(ui.Start)
	return b
}

// look is the colors one variant paints on a build's theme: the fill at
// rest, while hovered and while pressed, and the outline and the ink
// beside them. native marks the built-in look, which the variant keeps
// and paints nothing over.
type look struct {
	native  bool
	rest    ui.Color
	hover   ui.Color
	pressed ui.Color
	width   float32
	border  ui.Color
	text    ui.Color
	hasText bool
}

// lookOf resolves one variant's look on a theme. The sizing of a
// variant is not here: it rides the Size scale, which runs beside the
// variants. The theme's triplet slots feed the primary straight, the
// surface's feed ghost and outline; the destructive and the link shades
// come from Shades, which is where those rules live.
func lookOf(v Variant, t *ui.Theme) look {
	switch v {
	case Native:
		return look{native: true}
	case Primary:
		return look{
			rest: t.Accent, hover: t.AccentHover, pressed: t.AccentPressed,
			width: 0, border: t.Border, text: t.AccentText, hasText: true,
		}
	case Outline:
		return look{
			rest: t.Background, hover: t.SurfaceHover, pressed: t.SurfacePressed,
			width: 1, border: t.Border,
		}
	case Ghost:
		return look{
			rest: ui.Transparent, hover: t.SurfaceHover, pressed: t.SurfacePressed,
			width: 0, border: t.Border,
		}
	case Destructive:
		hover, pressed, ink := shades(t.Danger, t.Dark)
		return look{
			rest: t.Danger, hover: hover, pressed: pressed,
			width: 0, border: t.Border, text: ink, hasText: true,
		}
	case Link:
		a := t.Accent
		return look{
			rest: ui.Transparent, hover: a.Alpha(0.08), pressed: a.Alpha(0.16),
			width: 0, border: t.Border, text: a, hasText: true,
		}
	}
	// An unknown variant reads as native.
	return look{native: true}
}

// Variant applies the look of v and returns the button, for chaining.
// The zero variant is Native.
func (b *Button) Variant(v Variant) *Button {
	b.applySizing()
	if l := lookOf(v, b.c.Theme()); !l.native {
		b.Background(l.rest).Border(l.width, l.border)
		if l.hasText {
			b.TextColor(l.text)
		}
		// Hover and pressed paint over the rest, each frame, where the
		// built-in button's own state colors do not reach a preset's
		// rebuilt look. MyGo's transitions fade the changes as the
		// pointer comes and goes, and a disabled button keeps its rest
		// colors as the theme dims it.
		if !b.IsDisabled() {
			switch {
			case b.Pressed():
				b.Background(l.pressed)
			case b.Hovered():
				b.Background(l.hover)
			}
		}
	}
	return b
}

// Primary is Variant(myui.Primary).
func (b *Button) Primary() *Button { return b.Variant(Primary) }

// Secondary is Variant(myui.Secondary), the shadcn secondary look: an
// alias of Native, so the button keeps the theme's built-in face.
func (b *Button) Secondary() *Button { return b.Variant(Secondary) }

// Native is Variant(myui.Native).
func (b *Button) Native() *Button { return b.Variant(Native) }

// Outline is Variant(myui.Outline).
func (b *Button) Outline() *Button { return b.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (b *Button) Ghost() *Button { return b.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (b *Button) Destructive() *Button { return b.Variant(Destructive) }

// Link is Variant(myui.Link).
func (b *Button) Link() *Button { return b.Variant(Link) }

// Size applies the sizing of s and returns the button, for chaining.
func (b *Button) Size(s Size) *Button {
	b.size = s
	b.applySizing()
	return b
}

// Sm sizes the button down.
func (b *Button) Sm() *Button { return b.Size(SizeSm) }

// Md gives the sizing of the built-in button.
func (b *Button) Md() *Button { return b.Size(SizeMd) }

// Lg sizes the button up.
func (b *Button) Lg() *Button { return b.Size(SizeLg) }

// applySizing sizes the button on the theme of its build. The paddings
// of SizeMd are the built-in button's; the others move along the
// spacing scale the theme's Space says.
func (b *Button) applySizing() {
	t := b.c.Theme()
	switch b.size {
	case SizeSm:
		b.Padding(t.Space(1), t.Space(2.5)).FontSize(TextSize(t, TextSm))
	case SizeLg:
		b.Padding(t.Space(2.5), t.Space(4.5)).FontSize(TextSize(t, TextLg))
	default:
		b.Padding(t.Space(1.5), t.Space(3.5)).FontSize(t.FontSize)
	}
}
