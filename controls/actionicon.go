package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// ActionIconVariant picks the surface an ActionIcon sits on. The zero
// value, ActionIconBorderless, shows a fill only while hovered.
type ActionIconVariant uint8

const (
	// ActionIconBorderless shows no fill at rest, only under the pointer.
	ActionIconBorderless ActionIconVariant = iota
	// ActionIconFilled keeps a muted fill of the theme's surface.
	ActionIconFilled
	// ActionIconOutlined sits on the window's background with a border.
	ActionIconOutlined
)

// ActionIcon is an icon-only button with a variant's surface: a quiet
// affordance for one action, as the icons under a chat message are.
// Colors come from the theme of every frame, and state changes fade
// through StateTransition.
//
//	controls.NewActionIcon(c, icons.Copy(), "Copy").Filled().Sm()
type ActionIcon struct {
	ui.Element
	c       *ui.Context
	variant ActionIconVariant
	size    Size
	shape   Shape
	active  bool
	danger  bool
	glyph   ui.Element
}

// NewActionIcon builds an icon-only button named title for assistive
// technology. Clicked() reports a click as on any button.
func NewActionIcon(c *ui.Context, svg *ui.SVG, title string) *ActionIcon {
	a := &ActionIcon{c: c}
	a.Element = ui.Box(c).AlignItems(ui.Center).Justify(ui.Center).
		Role(ui.RoleButton).TabStop(true).Cursor(ui.CursorPointer).
		Label(title).Unselectable().Transition(StateTransition())
	a.Element.Children(func() {
		a.glyph = ui.Icon(c, svg)
	})
	a.apply()
	return a
}

// Variant applies the surface of v and returns the icon, for chaining.
func (a *ActionIcon) Variant(v ActionIconVariant) *ActionIcon {
	a.variant = v
	a.apply()
	return a
}

// Borderless is Variant(ActionIconBorderless).
func (a *ActionIcon) Borderless() *ActionIcon { return a.Variant(ActionIconBorderless) }

// Filled is Variant(ActionIconFilled).
func (a *ActionIcon) Filled() *ActionIcon { return a.Variant(ActionIconFilled) }

// Outlined is Variant(ActionIconOutlined).
func (a *ActionIcon) Outlined() *ActionIcon { return a.Variant(ActionIconOutlined) }

// Size applies the shared size s and returns the icon, for chaining.
func (a *ActionIcon) Size(s Size) *ActionIcon {
	a.size = s
	a.apply()
	return a
}

// Xs sizes the icon down to the smallest shared size.
func (a *ActionIcon) Xs() *ActionIcon { return a.Size(SizeXs) }

// Sm sizes the icon down.
func (a *ActionIcon) Sm() *ActionIcon { return a.Size(SizeSm) }

// Md gives the built-in sizing.
func (a *ActionIcon) Md() *ActionIcon { return a.Size(SizeMd) }

// Lg sizes the icon up.
func (a *ActionIcon) Lg() *ActionIcon { return a.Size(SizeLg) }

// Shape sets the corners of the surface and returns the icon, for
// chaining: a circle for a round action, the theme's radius otherwise.
func (a *ActionIcon) Shape(s Shape) *ActionIcon {
	a.shape = s
	a.apply()
	return a
}

// Active keeps the hover fill while the icon is not hovered, for a
// toggled or open state. It returns the icon for chaining.
func (a *ActionIcon) Active(on bool) *ActionIcon {
	a.active = on
	a.apply()
	return a
}

// Danger paints the glyph in the theme's danger color and returns the
// icon, for chaining.
func (a *ActionIcon) Danger(on bool) *ActionIcon {
	a.danger = on
	a.apply()
	return a
}

// Label names the icon for assistive technology and returns it, for
// chaining.
func (a *ActionIcon) Label(s string) *ActionIcon {
	a.Element = a.Element.Label(s)
	return a
}

// apply paints the surface and the glyph on this frame's theme.
func (a *ActionIcon) apply() {
	t := a.c.Theme()
	box, glyph := actionIconMetrics(a.size)
	a.Element = a.Element.Size(box, box)
	switch a.shape {
	case ShapeRound, ShapeCircle:
		a.Element = a.Element.Radius(1000)
	default:
		a.Element = a.Element.Radius(Radius(t, RadiusMd))
	}

	hovered, pressed := false, false
	if !a.Element.IsDisabled() {
		hovered, pressed = a.Element.Hovered(), a.Element.Pressed()
	}
	rest, hover, down := t.SurfaceHover, t.SurfaceHover, t.SurfacePressed
	switch a.variant {
	case ActionIconFilled:
		rest = t.Surface
	case ActionIconOutlined:
		rest, hover, down = t.Background, t.SurfaceHover, t.SurfacePressed
	default:
		rest, hover, down = ui.Transparent, t.SurfaceHover, t.SurfacePressed
	}
	if a.active {
		rest = hover
	}
	bg := rest
	switch {
	case pressed:
		bg = down
	case hovered:
		bg = hover
	}
	width := float32(0)
	border := ui.Transparent
	if a.variant == ActionIconOutlined {
		width, border = 1, t.Border
	}
	a.Element = a.Element.Background(bg).Border(width, border)

	ink := t.TextMuted
	if a.danger {
		h, d, _ := Shades(t.Danger, t.Dark)
		ink = t.Danger.Alpha(0.66)
		if hovered {
			ink = h
		}
		if pressed {
			ink = d
		}
	} else if hovered || pressed {
		ink = t.Text
	}
	a.glyph = a.glyph.Size(glyph, glyph).TextColor(ink)
}

// actionIconMetrics is the box side and the glyph size of a shared size.
func actionIconMetrics(s Size) (box, glyph float32) {
	switch s {
	case SizeXs:
		return 28, IconSizeXs
	case SizeSm:
		return 32, IconSizeSm
	case SizeLg:
		return 40, IconSizeLg
	default:
		return 36, IconSizeMd
	}
}
