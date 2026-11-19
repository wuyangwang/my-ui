package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// Toggle is a button that stays pressed while *on. It is my-ui's look over
// mygo's toggle button: a click or Space writes the app's bool, and the
// change is reported both by Changed and by the callbacks below.
//
//	controls.NewToggle(c, &app.bold, "Bold").Primary().Icon(icons.Bold())
//
// The size, radius and pressed look come from c.Theme() on every frame.
// Handle registers a callback that keeps working when it is attached after
// the constructor; OnChanged fires only when chained in the click's frame.
type Toggle struct {
	ui.Element
	c       *ui.Context
	on      *bool
	variant Variant
	size    Size
	changed bool
	cb      *func(bool)
	icon    ui.Element
	text    ui.Element
}

// NewToggle builds a toggle bound to *on, with label beside an optional
// icon. The label may be empty when an icon and Label are used instead.
func NewToggle(c *ui.Context, on *bool, label string) *Toggle {
	tg := &Toggle{c: c, on: on, size: SizeMd}
	t := c.Theme()
	el := ui.Toggle(c, on, "")
	tg.cb = ui.Local(el, "toggle-cb", func() func(bool) { return nil })
	el.Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1)).Children(func() {
			tg.icon = ui.Box(c).Hidden(true)
			tg.text = ui.Box(c)
			tg.text.Children(func() {
				if label != "" {
					ui.Text(c, label).SingleLine()
				}
			})
		})
	})
	tg.Element = el
	tg.changed = el.Changed()
	if tg.changed && *tg.cb != nil {
		(*tg.cb)(*on)
	}
	tg.apply()
	return tg
}

// apply paints the toggle's size and look from this frame's theme.
func (t *Toggle) apply() {
	th := t.c.Theme()
	v, h, font, height := toggleMetrics(th, t.size)
	bg, fg, border := th.Surface, th.Text, th.Border
	on := t.on != nil && *t.on
	if on {
		bg = th.Accent.Alpha(0.14)
		fg = th.Accent
		border = th.Accent.Alpha(0.55)
	}
	switch t.variant {
	case Primary:
		if on {
			bg = th.Accent.Alpha(0.20)
		}
	case Secondary:
		bg = th.SurfaceHover
		if on {
			bg = th.Accent.Alpha(0.12)
		}
	case Outline, Ghost:
		if on {
			bg = th.Accent.Alpha(0.10)
		}
	case Destructive:
		if on {
			bg = th.Danger.Alpha(0.16)
			fg = th.Danger
			border = th.Danger.Alpha(0.50)
		}
	}
	t.Element = t.Element.Padding(v, h).Height(height).FontSize(font).
		Radius(Radius(th, RadiusMd)).Background(bg).Border(1, border).
		TextColor(fg).Transition(StateTransition())
}

// toggleMetrics is the inset, font and height of a shared size.
func toggleMetrics(t *ui.Theme, s Size) (v, h, font, height float32) {
	switch s {
	case SizeXs:
		return t.Space(0.5), t.Space(2), TextSize(t, TextXs), 28
	case SizeSm:
		return t.Space(1), t.Space(2.5), TextSize(t, TextSm), 32
	case SizeLg:
		return t.Space(2), t.Space(4), TextSize(t, TextLg), 40
	default:
		return t.Space(1), t.Space(2.5), TextSize(t, TextMd), 36
	}
}

// Size applies a shared size and returns the toggle, for chaining.
func (t *Toggle) Size(sz Size) *Toggle {
	t.size = sz
	t.apply()
	return t
}

// Xs sizes the toggle down to the smallest shared size.
func (t *Toggle) Xs() *Toggle { return t.Size(SizeXs) }

// Sm sizes the toggle down.
func (t *Toggle) Sm() *Toggle { return t.Size(SizeSm) }

// Md gives the built-in sizing.
func (t *Toggle) Md() *Toggle { return t.Size(SizeMd) }

// Lg sizes the toggle up.
func (t *Toggle) Lg() *Toggle { return t.Size(SizeLg) }

// Variant applies a preset look and returns the toggle, for chaining.
func (t *Toggle) Variant(v Variant) *Toggle {
	t.variant = v
	t.apply()
	return t
}

// Primary is Variant(Primary).
func (t *Toggle) Primary() *Toggle { return t.Variant(Primary) }

// Secondary is Variant(Secondary), a muted fill.
func (t *Toggle) Secondary() *Toggle { return t.Variant(Secondary) }

// Outline is Variant(Outline).
func (t *Toggle) Outline() *Toggle { return t.Variant(Outline) }

// Ghost is Variant(Ghost).
func (t *Toggle) Ghost() *Toggle { return t.Variant(Ghost) }

// Destructive is Variant(Destructive).
func (t *Toggle) Destructive() *Toggle { return t.Variant(Destructive) }

// Icon puts svg before the label and returns the toggle, for chaining.
func (t *Toggle) Icon(svg *ui.SVG) *Toggle {
	if svg != nil {
		t.icon.Hidden(false)
		t.icon.Children(func() {
			ui.Icon(t.c, svg).Size(IconSizeSm, IconSizeSm)
		})
	}
	return t
}

// Label names the toggle for assistive technology and returns it.
func (t *Toggle) Label(s string) *Toggle {
	t.Element = t.Element.Label(s)
	return t
}

// Disabled takes the toggle out of use and returns it, for chaining.
func (t *Toggle) Disabled(on bool) *Toggle {
	t.Element = t.Element.Disabled(on)
	return t
}

// Block stretches the toggle across its row and returns it, for chaining.
func (t *Toggle) Block() *Toggle {
	t.Element = t.Element.AlignSelf(ui.Stretch)
	return t
}

// Changed reports whether the toggle was clicked in this frame.
func (t *Toggle) Changed() bool { return t.changed }

// OnChanged calls fn with the new state when the toggle changed in this
// frame. Chain it in the same frame as the constructor; use Handle when
// the callback is attached later.
func (t *Toggle) OnChanged(fn func(on bool)) *Toggle {
	if fn != nil && t.changed && t.on != nil {
		fn(*t.on)
	}
	return t
}

// Handle registers a callback that stays with the element, so it still
// fires when attached after the constructor or on a later frame. It
// returns the toggle, for chaining.
func (t *Toggle) Handle(fn func(on bool)) *Toggle {
	if t.cb != nil {
		*t.cb = fn
	}
	return t
}

// ColorWell is a swatch MyGo's native color panel edits.
type ColorWell struct {
	ui.Element
}

// NewColorWell builds a swatch for *color with my-ui's border and
// radius, on the theme's surface.
func NewColorWell(c *ui.Context, color *ui.Color) *ColorWell {
	t := c.Theme()
	return &ColorWell{Element: ui.ColorWell(c, color).
		Border(1, t.Border).Radius(Radius(t, RadiusMd)).
		Background(t.Surface).Padding(t.Space(1)).
		Transition(StateTransition())}
}
