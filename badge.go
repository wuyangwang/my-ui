package myui

import "github.com/egoist/mygo/ui"

// Badge is a short label in a pill with variant and size presets:
//
//	myui.NewBadge(c, "New").Primary()
//	myui.NewBadge(c, "99+").Secondary().Sm()
//
// The chain applies presets from the current theme every frame, so it
// follows palette and appearance changes. Call it each frame before any
// app-specific styling. Badge embeds ui.Element, so other ui methods
// remain available.
type Badge struct {
	ui.Element
	c    *ui.Context
	size Size
}

// NewBadge creates a badge with the native preset.
func NewBadge(c *ui.Context, text string) *Badge {
	b := &Badge{Element: ui.Badge(c, text), c: c}
	b.AlignSelf(ui.Start)
	b.Element = b.Element.Transition(stateTransition())
	return b
}

type badgeLook struct {
	native  bool
	bg      ui.Color
	border  ui.Color
	width   float32
	text    ui.Color
	hasText bool
}

func badgeLookOf(v Variant, t *ui.Theme) badgeLook {
	switch v {
	case Native:
		return badgeLook{native: true}
	case Primary:
		return badgeLook{bg: t.Accent, border: t.Accent, width: 0, text: t.AccentText, hasText: true}
	case Outline:
		return badgeLook{bg: ui.Transparent, border: t.Border, width: 1}
	case Ghost:
		return badgeLook{bg: ui.Transparent, border: ui.Transparent, width: 0}
	case Destructive:
		_, _, ink := shades(t.Danger, t.Dark)
		return badgeLook{bg: t.Danger, border: t.Danger, width: 0, text: ink, hasText: true}
	case Link:
		return badgeLook{bg: ui.Transparent, border: ui.Transparent, width: 0, text: t.Accent, hasText: true}
	}
	return badgeLook{native: true}
}

// Variant applies the look of v and returns the badge, for chaining.
func (b *Badge) Variant(v Variant) *Badge {
	b.applySizing()
	if l := badgeLookOf(v, b.c.Theme()); !l.native {
		b.Background(l.bg).Border(l.width, l.border)
		if l.hasText {
			b.TextColor(l.text)
		}
	}
	return b
}

// Primary is Variant(myui.Primary).
func (b *Badge) Primary() *Badge { return b.Variant(Primary) }

// Secondary is Variant(myui.Secondary), alias of Native.
func (b *Badge) Secondary() *Badge { return b.Variant(Secondary) }

// Native is Variant(myui.Native).
func (b *Badge) Native() *Badge { return b.Variant(Native) }

// Outline is Variant(myui.Outline).
func (b *Badge) Outline() *Badge { return b.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (b *Badge) Ghost() *Badge { return b.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (b *Badge) Destructive() *Badge { return b.Variant(Destructive) }

// Link is Variant(myui.Link).
func (b *Badge) Link() *Badge { return b.Variant(Link) }

// Size applies the sizing of s and returns the badge, for chaining.
func (b *Badge) Size(s Size) *Badge {
	b.size = s
	b.applySizing()
	return b
}

// Sm sizes the badge down.
func (b *Badge) Sm() *Badge { return b.Size(SizeSm) }

// Md gives the built-in sizing.
func (b *Badge) Md() *Badge { return b.Size(SizeMd) }

// Lg sizes the badge up.
func (b *Badge) Lg() *Badge { return b.Size(SizeLg) }

func (b *Badge) applySizing() {
	t := b.c.Theme()
	switch b.size {
	case SizeSm:
		b.Padding(t.Space(0.5), t.Space(1.5)).FontSize(TextSize(t, TextXs))
	case SizeLg:
		b.Padding(t.Space(1), t.Space(2.5)).FontSize(TextSize(t, TextSm))
	default:
		b.Padding(t.Space(0.75), t.Space(2)).FontSize(TextSize(t, TextXs))
	}
}
