package myui

import (
	"github.com/egoist/mygo/ui"
)

// ButtonVariant names what a Button looks like, the shadcn set.
type ButtonVariant uint8

const (
	// Secondary is the built-in look: the theme's face, which the button
	// itself hovers and presses. The zero value, so an unchained Button
	// is a secondary one.
	Secondary ButtonVariant = iota
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

// ButtonSize names how roomy a Button is.
type ButtonSize uint8

const (
	// SizeMd is the built-in sizing. The zero value.
	SizeMd ButtonSize = iota
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
	c       *ui.Context
	variant ButtonVariant
	size    ButtonSize
}

// NewButton creates a button with the secondary preset: myui.NewButton(c,
// "Save"). Clicked() reports the click on it as on ui.Button.
func NewButton(c *ui.Context, label string) *Button {
	b := &Button{Element: ui.Button(c, label), c: c}
	// Size to the content as a button does: a column stretches its
	// children across, which a button never asks for.
	b.AlignSelf(ui.Start)
	return b
}

// Variant applies the look of v and returns the button, for chaining.
// The zero variant is the secondary.
func (b *Button) Variant(v ButtonVariant) *Button {
	b.variant = v
	b.applySizing()
	t := b.c.Theme()
	switch b.variant {
	case Secondary:
		// The look the built-in button already has.
	case Primary:
		b.Background(b.shade(t.Accent, t.AccentHover, t.AccentPressed)).
			Border(0, t.Border).
			TextColor(t.AccentText)
	case Outline:
		b.Background(t.Background).Border(1, t.Border)
	case Ghost:
		b.Background(ui.Transparent).Border(0, t.Border)
	case Destructive:
		b.Background(b.shade(t.Danger, darken(t.Danger, 0.12), darken(t.Danger, 0.24))).
			Border(0, t.Border).
			TextColor(OnAccent(t.Danger))
	case Link:
		b.Background(ui.Transparent).Border(0, t.Border).TextColor(t.Accent)
	}
	b.applyVariantState(t)
	return b
}

// Primary is Variant(myui.Primary).
func (b *Button) Primary() *Button { return b.Variant(Primary) }

// Secondary is Variant(myui.Secondary), the look the button already has
// from the theme.
func (b *Button) Secondary() *Button { return b.Variant(Secondary) }

// Outline is Variant(myui.Outline).
func (b *Button) Outline() *Button { return b.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (b *Button) Ghost() *Button { return b.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (b *Button) Destructive() *Button { return b.Variant(Destructive) }

// Link is Variant(myui.Link).
func (b *Button) Link() *Button { return b.Variant(Link) }

// Size applies the sizing of s and returns the button, for chaining.
func (b *Button) Size(s ButtonSize) *Button {
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

// applyVariantState paints the pointer's colors of the variant over its
// base ones, each frame: hover and pressed, where the built-in button's
// own state colors do not reach a rebuilt look. MyGo's transitions fade
// the changes as the pointer comes and goes.
func (b *Button) applyVariantState(t *ui.Theme) {
	if b.IsDisabled() || b.variant == Secondary {
		return
	}
	a := b.c.Theme().Accent
	var hover, pressed ui.Color
	switch b.variant {
	case Primary:
		hover, pressed = t.AccentHover, t.AccentPressed
	case Ghost, Outline:
		hover, pressed = t.SurfaceHover, t.SurfacePressed
	case Destructive:
		hover, pressed = darken(t.Danger, 0.12), darken(t.Danger, 0.24)
	case Link:
		hover, pressed = a.Alpha(0.08), a.Alpha(0.16)
	default:
		return
	}
	switch {
	case b.Pressed():
		b.Background(pressed)
	case b.Hovered():
		b.Background(hover)
	}
}

// shade returns the colors a filled button paints at rest, hovered and
// pressed.
func (b *Button) shade(rest, hover, pressed ui.Color) ui.Color {
	switch {
	case b.Pressed():
		return pressed
	case b.Hovered():
		return hover
	}
	return rest
}

// darken returns the color c a little darker, by t of the way toward
// black in Oklab.
func darken(c ui.Color, t float32) ui.Color {
	return c.MixOklab(ui.RGB(0, 0, 0), t)
}
