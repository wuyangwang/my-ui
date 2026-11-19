package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// ToggleGroupVariant picks how the toggles of a group are joined.
type ToggleGroupVariant uint8

const (
	// ToggleGroupOutlined joins the toggles in one bordered face.
	ToggleGroupOutlined ToggleGroupVariant = iota
	// ToggleGroupBorderless spaces the toggles apart without a face.
	ToggleGroupBorderless
	// ToggleGroupSegmented puts the toggles on a segmented track.
	ToggleGroupSegmented
)

// ToggleGroup gathers the toggle buttons that fn builds into one
// control, as a row of text styles is gathered in a toolbar. It is
// my-ui's look over mygo's ToggleGroup: one Tab stop, arrows moving
// focus between the toggles and the role of a group all come from the
// framework. The signature matches mygo's ToggleGroup.
//
//	controls.NewToggleGroup(c, func() {
//		controls.NewToggle(c, &bold, "B")
//		controls.NewToggle(c, &italic, "I")
//	}).Label("Style")
type ToggleGroup struct {
	ui.Element
	c       *ui.Context
	variant ToggleGroupVariant
	size    Size
}

// NewToggleGroup builds a group from the toggles fn builds. Call Label
// on the returned element to name the group.
func NewToggleGroup(c *ui.Context, fn func()) *ToggleGroup {
	g := &ToggleGroup{c: c}
	g.Element = ui.ToggleGroup(c, fn)
	g.apply()
	return g
}

// Variant applies the look of v and returns the group, for chaining.
func (g *ToggleGroup) Variant(v ToggleGroupVariant) *ToggleGroup {
	g.variant = v
	g.apply()
	return g
}

// Outlined is Variant(ToggleGroupOutlined).
func (g *ToggleGroup) Outlined() *ToggleGroup { return g.Variant(ToggleGroupOutlined) }

// Borderless is Variant(ToggleGroupBorderless).
func (g *ToggleGroup) Borderless() *ToggleGroup { return g.Variant(ToggleGroupBorderless) }

// Segmented is Variant(ToggleGroupSegmented).
func (g *ToggleGroup) Segmented() *ToggleGroup { return g.Variant(ToggleGroupSegmented) }

// Label names the group for assistive technology and returns it, for
// chaining.
func (g *ToggleGroup) Label(name string) *ToggleGroup {
	g.Element = g.Element.Label(name)
	return g
}

// Size applies the shared size s and returns the group, for chaining.
// The group's toggles pick up the size from their own constructors.
func (g *ToggleGroup) Size(s Size) *ToggleGroup {
	g.size = s
	g.apply()
	return g
}

// apply paints the group on this frame's theme.
func (g *ToggleGroup) apply() {
	t := g.c.Theme()
	bg, border := ui.Transparent, ui.Transparent
	width := float32(0)
	gap := t.Space(0.5)
	pad := t.Space(0.5)
	switch g.variant {
	case ToggleGroupBorderless:
		pad = 0
	case ToggleGroupSegmented:
		bg, width, border = t.Surface, 1, t.Border
		pad, gap = t.Space(0.75), t.Space(0.5)
	default:
		width, border = 1, t.Border
	}
	g.Element = g.Element.Padding(pad).Gap(gap).Background(bg).
		Border(width, border).Radius(Radius(t, RadiusMd)).
		Transition(StateTransition())
	if g.variant == ToggleGroupOutlined {
		g.Element = g.Element.Clip()
	}
}
