package display

import (
	"strconv"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// BadgeStatus is the state a status badge reports. Each status picks the
// dot's color; Processing also draws a pulsing ring. It is orthogonal to
// the Variant of a text badge.
type BadgeStatus uint8

const (
	// BadgeOnline reports that a service is up (the zero value).
	BadgeOnline BadgeStatus = iota
	// BadgeSyncing reports background work in progress.
	BadgeSyncing
	// BadgeDegraded reports reduced service.
	BadgeDegraded
	// BadgeOffline reports that a service is down.
	BadgeOffline
	// BadgeIdle reports that nothing is happening.
	BadgeIdle
	// BadgeProcessing reports active work: the dot pulses.
	BadgeProcessing
)

// badgePulse is the period of the Processing ring. The pulse itself is
// drawn with Element.Loop — there is no timer in the component, so the
// desktop's Reduce Motion comes free.
const badgePulse = 1200 * time.Millisecond

// Badge is a short label in a pill, with variant and size presets, or a
// status dot with a text beside it:
//
//	display.NewBadge(c, "New").Primary()
//	display.NewBadge(c, "").Status(BadgeProcessing).Text("同步中")
//
// A badge can also count: Count caps the number and ShowZero keeps a zero
// count visible. The chain applies its presets from the current theme on
// every frame, so it follows palette and appearance changes. Badge embeds
// ui.Element, so the rest of ui's methods remain available.
type Badge struct {
	ui.Element
	c *ui.Context

	text     string
	size     Size
	variant  Variant
	color    ui.Color // optional override of the pill or dot color
	status   BadgeStatus
	statusOn bool
	textOn   bool
	showZero bool
	small    bool
	count    int
	countSet bool
	max      int

	root      ui.Element
	pill      ui.Element
	textSlot  ui.Element
	statusRow ui.Element
	dot       ui.Element
	ring      ui.Element
	statusTxt ui.Element
}

// NewBadge creates a badge with the native preset.
func NewBadge(c *ui.Context, text string) *Badge {
	b := &Badge{c: c, text: text, size: SizeMd}
	t := c.Theme()
	root := ui.Row(c).AlignSelf(ui.Start).AlignItems(ui.Center).Shrink(0)
	root.Children(func() {
		b.pill = ui.Row(c).AlignItems(ui.Center)
		b.pill.Children(func() {
			b.textSlot = ui.Box(c)
		})

		b.statusRow = ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1.5)).Hidden(true)
		b.statusRow.Children(func() {
			b.dot = ui.Box(c).Size(t.Space(1.5), t.Space(1.5)).Radius(t.Space(0.75))
			b.dot.Children(func() {
				b.ring = ui.Box(c).Absolute().Left(-1).Top(-1).
					Size(t.Space(2), t.Space(2)).Radius(t.Space(1)).Hidden(true)
			})
			b.statusTxt = ui.Box(c)
		})
	})
	b.Element = root
	b.root = root
	b.renderText(text)
	b.apply()
	return b
}

// renderText puts text in the pill. The count replaces it when Count is
// called; a caller that changes text in one frame calls Text instead.
func (b *Badge) renderText(s string) {
	b.textSlot.Children(func() {
		ui.Text(b.c, s)
	})
}

// badgeLook is the colors a Variant paints on the pill, resolved from the
// theme of this frame.
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
		return badgeLook{bg: t.Surface, border: ui.Transparent}
	case Primary:
		return badgeLook{bg: t.Accent, border: t.Accent, text: t.AccentText, hasText: true}
	case Outline:
		return badgeLook{bg: ui.Transparent, border: t.Border, width: 1}
	case Ghost:
		return badgeLook{bg: ui.Transparent, border: ui.Transparent}
	case Destructive:
		_, _, ink := Shades(t.Danger, t.Dark)
		return badgeLook{bg: t.Danger, border: t.Danger, text: ink, hasText: true}
	case Link:
		return badgeLook{bg: ui.Transparent, border: ui.Transparent, text: t.Accent, hasText: true}
	case Success, Info, Warning:
		role, _ := SemanticOf(v)
		base := SemanticColor(t, role)
		_, _, ink := Shades(base, t.Dark)
		return badgeLook{bg: base, border: base, text: ink, hasText: true}
	}
	return badgeLook{native: true}
}

// Variant applies the look of v and returns the badge, for chaining.
func (b *Badge) Variant(v Variant) *Badge {
	b.variant = v
	b.statusOn = false
	b.apply()
	return b
}

// Color overrides the fill with c — the pill's background or the status
// dot's color. A zero Color clears the override and restores the preset.
func (b *Badge) Color(c ui.Color) *Badge {
	b.color = c
	b.apply()
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
	b.apply()
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

// Small makes a count pill 14 DIPs tall instead of 18, for a denser
// corner on a small icon.
func (b *Badge) Small(on bool) *Badge {
	b.small = on
	b.apply()
	return b
}

// Count turns the badge into a count: it shows n, capped at max with a
// "+" when max is above zero, and hides the pill when n is zero unless
// ShowZero(true) was called. It returns the badge, for chaining.
func (b *Badge) Count(n, max int) *Badge {
	b.count, b.max, b.countSet = n, max, true
	b.statusOn = false
	b.renderCount()
	b.apply()
	return b
}

func (b *Badge) renderCount() {
	text := strconv.Itoa(b.count)
	if b.max > 0 && b.count > b.max {
		text = strconv.Itoa(b.max) + "+"
	}
	b.textSlot.Children(func() {
		ui.Text(b.c, text)
	})
}

// ShowZero keeps a zero count visible instead of hiding the pill.
func (b *Badge) ShowZero(on bool) *Badge {
	b.showZero = on
	b.apply()
	return b
}

// Status switches the badge to a status dot and returns it. The dot's
// color follows s; BadgeProcessing pulses the dot with a ring drawn by
// Element.Loop, with no timer in the component.
func (b *Badge) Status(s BadgeStatus) *Badge {
	b.status = s
	b.statusOn = true
	b.apply()
	return b
}

// Text sets the label beside a status dot, or the pill's text. It returns
// the badge for chaining.
func (b *Badge) Text(s string) *Badge {
	b.text = s
	if b.statusOn {
		b.statusTxt.Children(func() {
			ui.Text(b.c, s).TextColor(b.c.Theme().Text)
		})
	} else {
		b.countSet = false
		b.renderText(s)
	}
	b.apply()
	return b
}

// statusColor is the hue of a status on this frame's theme.
func (b *Badge) statusColor(s BadgeStatus) ui.Color {
	if b.color != (ui.Color{}) {
		return b.color
	}
	t := b.c.Theme()
	switch s {
	case BadgeOnline:
		return SemanticColor(t, SemanticSuccess)
	case BadgeSyncing:
		return SemanticColor(t, SemanticInfo)
	case BadgeDegraded:
		return SemanticColor(t, SemanticWarning)
	case BadgeOffline:
		return t.Danger
	case BadgeIdle:
		return t.TextMuted
	case BadgeProcessing:
		return t.Accent
	}
	return t.TextMuted
}

// apply paints the current mode, variant and size on this frame's
// elements. It never rebuilds the row, so the controls keep their clicks.
func (b *Badge) apply() {
	t := b.c.Theme()
	b.root.Hidden(false)
	b.pill.Hidden(b.statusOn)
	b.statusRow.Hidden(!b.statusOn)

	// A count of zero hides the pill unless ShowZero is on.
	if !b.statusOn && b.countSet && b.count == 0 && !b.showZero {
		b.root.Hidden(true)
	}

	if b.statusOn {
		col := b.statusColor(b.status)
		b.dot = b.dot.Background(col)
		if b.status == BadgeProcessing {
			p := b.dot.Loop("badge-pulse", badgePulse, ui.EaseOut)
			scale := 1 + 1.4*p
			b.ring = b.ring.Border(1, col).Scale(scale, scale).Opacity(1 - p).Hidden(false)
		} else {
			b.ring.Hidden(true)
		}
		return
	}

	l := badgeLookOf(b.variant, t)
	if !l.native {
		if b.color != (ui.Color{}) {
			l.bg, l.border = b.color, b.color
		}
		b.pill = b.pill.Background(l.bg).Border(l.width, l.border)
		if l.hasText {
			b.pill = b.pill.TextColor(l.text)
		}
	}
	switch b.size {
	case SizeXs:
		b.pill.Padding(t.Space(0.25), t.Space(1)).FontSize(TextSize(t, TextXs))
	case SizeSm:
		b.pill.Padding(t.Space(0.5), t.Space(1.5)).FontSize(TextSize(t, TextXs))
	case SizeLg:
		b.pill.Padding(t.Space(1), t.Space(2.5)).FontSize(TextSize(t, TextSm))
	default:
		b.pill.Padding(t.Space(0.75), t.Space(2)).FontSize(TextSize(t, TextXs))
	}
	if b.small && b.countSet {
		b.pill.Height(t.Space(3.5)).Padding(0, t.Space(1.25))
	}
}

// Closable puts a small close button in the badge and returns it, for
// chaining. Clicking it sets *closed to true.
func (b *Badge) Closable(closed *bool) *Badge {
	b.pill.Children(func() {
		close := ctrlpkg.NewIconAction(b.c, ctrlpkg.IconActionOptions{
			SVG:      icons.X(),
			Label:    Message(b.c, "close"),
			Size:     16,
			IconSize: IconSizeXs,
			OnClick: func() {
				if closed != nil {
					*closed = true
				}
			},
		})
		close.Margin(0, 0, 0, b.c.Theme().Space(1))
	})
	return b
}

// BadgeOptions configures NewBadgeOn.
type BadgeOptions struct {
	Content string
	Dot     bool
	Count   int
	Max     int
	Hidden  bool
	Variant Variant
	// ShowZero keeps a count of zero visible.
	ShowZero bool
	// Small makes a count pill smaller.
	Small bool
}

// NewBadgeOn builds an anchor and pins a badge to its top-right corner.
// The badge shows Content, or a count capped at Max, or just a dot.
func NewBadgeOn(c *ui.Context, anchor func() ui.Element, o BadgeOptions) *Badge {
	wrap := ui.Box(c)
	var result *Badge
	wrap.Children(func() {
		if anchor != nil {
			anchor()
		}
		if o.Hidden {
			return
		}
		bd := NewBadge(c, o.Content)
		if o.ShowZero {
			bd.ShowZero(true)
		}
		if o.Small {
			bd.Small(true)
		}
		if o.Count != 0 || o.ShowZero {
			bd.Count(o.Count, o.Max)
		} else if o.Dot {
			bd.Element = bd.Element.Size(c.Theme().Space(2), c.Theme().Space(2))
		}
		if o.Variant != Native {
			bd.Variant(o.Variant)
		} else {
			bd.Destructive()
		}
		bd.Element = bd.Element.Absolute().Top(-c.Theme().Space(1)).Right(-c.Theme().Space(1))
		result = bd
	})
	if result == nil {
		result = &Badge{Element: wrap, c: c}
	}
	return result
}
