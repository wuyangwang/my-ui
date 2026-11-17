package myui

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

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
	case Default:
		return badgeLook{native: true}
	case Secondary:
		return badgeLook{bg: t.Surface, border: ui.Transparent, width: 0}
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
	case Success, Info, Warning:
		role, _ := semanticOf(v)
		base := semanticColor(t, role)
		_, _, ink := shades(base, t.Dark)
		return badgeLook{bg: base, border: base, width: 0, text: ink, hasText: true}
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

// Secondary is Variant(myui.Secondary), a muted fill.
func (b *Badge) Secondary() *Badge { return b.Variant(Secondary) }

// Default is Variant(myui.Default), the built-in look.
func (b *Badge) Default() *Badge { return b.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (b *Badge) Native() *Badge { return b.Variant(Native) }

// Outline is Variant(myui.Outline).
func (b *Badge) Outline() *Badge { return b.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (b *Badge) Ghost() *Badge { return b.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (b *Badge) Destructive() *Badge { return b.Variant(Destructive) }

// Link is Variant(myui.Link).
func (b *Badge) Link() *Badge { return b.Variant(Link) }

// Success is Variant(myui.Success).
func (b *Badge) Success() *Badge { return b.Variant(Success) }

// Info is Variant(myui.Info).
func (b *Badge) Info() *Badge { return b.Variant(Info) }

// Warning is Variant(myui.Warning).
func (b *Badge) Warning() *Badge { return b.Variant(Warning) }

// Size applies the sizing of s and returns the badge, for chaining.
func (b *Badge) Size(s Size) *Badge {
	b.size = s
	b.applySizing()
	return b
}

// Xs sizes the badge to the smallest size.
func (b *Badge) Xs() *Badge { return b.Size(SizeXs) }

// Sm sizes the badge down.
func (b *Badge) Sm() *Badge { return b.Size(SizeSm) }

// Md gives the built-in sizing.
func (b *Badge) Md() *Badge { return b.Size(SizeMd) }

// Lg sizes the badge up.
func (b *Badge) Lg() *Badge { return b.Size(SizeLg) }

func (b *Badge) applySizing() {
	t := b.c.Theme()
	switch b.size {
	case SizeXs:
		b.Padding(t.Space(0.25), t.Space(1)).FontSize(TextSize(t, TextXs))
	case SizeSm:
		b.Padding(t.Space(0.5), t.Space(1.5)).FontSize(TextSize(t, TextXs))
	case SizeLg:
		b.Padding(t.Space(1), t.Space(2.5)).FontSize(TextSize(t, TextSm))
	default:
		b.Padding(t.Space(0.75), t.Space(2)).FontSize(TextSize(t, TextXs))
	}
}

// Count turns the badge into a count: it shows n, capped at max with a
// "+" when max is above zero, and hides the badge when n is zero. It
// returns the badge for chaining, and reads n from the value the app
// holds each frame.
func (b *Badge) Count(n, max int) *Badge {
	c := b.c
	if n == 0 {
		b.Hidden(true)
		return b
	}
	b.Hidden(false)
	text := strconv.Itoa(n)
	if max > 0 && n > max {
		text = strconv.Itoa(max) + "+"
	}
	b.Element.Children(func() {
		ui.Text(c, text)
	})
	return b
}

// Closable puts a small close button in the badge and returns it, for
// chaining. Clicking it sets *closed to true.
func (b *Badge) Closable(closed *bool) *Badge {
	b.Children(func() {
		btn := ui.ButtonBase(b.c).Size(12, 12).KeepFocus().Margin(0, 0, 0, b.c.Theme().Space(1)).
			Label(Message(b.c, "close"))
		btn.Children(func() {
			ic := ui.Icon(b.c, icons.X()).Size(10, 10)
			if btn.Hovered() {
				ic = ic.TextColor(b.c.Theme().Text)
			}
		})
		if btn.Clicked() && closed != nil {
			*closed = true
		}
	})
	return b
}

// BadgeOn pins a badge to the corner of an anchor, as Element Plus's
// el-badge does. o holds the content, whether to show only a dot, the cap
// for a count, whether it is hidden and the badge's variant.
type BadgeOptions struct {
	Content string
	Dot     bool
	Count   int
	Max     int
	Hidden  bool
	Variant Variant
}

// NewBadgeOn builds an anchor and pins a badge to its top-right corner.
// The badge shows Content, or a count capped at Max, or just a dot.
func NewBadgeOn(c *ui.Context, anchor func() ui.Element, o BadgeOptions) *Badge {
	wrap := ui.Box(c)
	wrap.Children(func() {
		if anchor != nil {
			anchor()
		}
		if o.Hidden {
			return
		}
		bd := NewBadge(c, o.Content)
		if o.Count != 0 {
			bd.Count(o.Count, o.Max)
		} else if o.Dot {
			bd.Element = bd.Element.Size(c.Theme().Space(2), c.Theme().Space(2))
			bd.Element.Children(func() {})
		}
		if o.Variant != Native {
			bd.Variant(o.Variant)
		} else {
			bd.Destructive()
		}
		bd.Element = bd.Element.Absolute().Top(-c.Theme().Space(1)).Right(-c.Theme().Space(1))
	})
	return &Badge{Element: wrap, c: c}
}
