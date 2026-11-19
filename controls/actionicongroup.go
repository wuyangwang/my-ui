package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// ActionIconGroupItem is one icon of an ActionIconGroup.
type ActionIconGroupItem struct {
	// Key is what the group writes to its chosen pointer when clicked.
	Key string
	// Icon is the glyph; Label names the button for assistive technology.
	Icon  *ui.SVG
	Label string
	// Danger paints the glyph in the theme's danger color.
	Danger bool
	// Disabled keeps this icon out of use.
	Disabled bool
}

// ActionIconGroupVariant picks the surface a group sits on.
type ActionIconGroupVariant uint8

const (
	// ActionIconGroupFilled keeps a muted fill of the theme's surface.
	ActionIconGroupFilled ActionIconGroupVariant = iota
	// ActionIconGroupOutlined borders the group on the window's background.
	ActionIconGroupOutlined
	// ActionIconGroupBorderless shows a fill only while hovered.
	ActionIconGroupBorderless
)

// ActionIconGroup gathers icon-only actions on one surface, as the
// actions under a message are. Clicking an item writes its Key into the
// chosen pointer; Chosen() reads the same value in the frame of the
// click, so the next statement can use it.
type ActionIconGroup struct {
	ui.Element
	c       *ui.Context
	variant ActionIconGroupVariant
	size    Size
	items   []ui.Element
	chosen  string
}

// NewActionIconGroup builds a row of item icons. chosen receives the Key
// of the item clicked in the frame, or keeps its value when none was.
func NewActionIconGroup(c *ui.Context, chosen *string, items []ActionIconGroupItem) *ActionIconGroup {
	t := c.Theme()
	g := &ActionIconGroup{c: c}
	g.Element = ui.Row(c).AlignItems(ui.Center).AlignSelf(ui.Start).
		Gap(t.Space(0.5)).Padding(t.Space(0.5)).
		Radius(Radius(t, RadiusMd)).Transition(StateTransition())
	g.Element.Children(func() {
		for _, it := range items {
			it := it
			box, glyph := actionIconMetrics(g.size)
			b := NewIconAction(c.Key(it.Key), IconActionOptions{
				SVG:      it.Icon,
				Label:    it.Label,
				Size:     box,
				IconSize: glyph,
				Danger:   it.Danger,
				OnClick: func() {
					if !it.Disabled && chosen != nil {
						*chosen = it.Key
						g.chosen = it.Key
					}
				},
			})
			if it.Disabled {
				b = b.Disabled(true)
			}
			g.items = append(g.items, b)
		}
	})
	g.apply()
	return g
}

// Variant applies the surface of v and returns the group, for chaining.
func (g *ActionIconGroup) Variant(v ActionIconGroupVariant) *ActionIconGroup {
	g.variant = v
	g.apply()
	return g
}

// Filled is Variant(ActionIconGroupFilled).
func (g *ActionIconGroup) Filled() *ActionIconGroup { return g.Variant(ActionIconGroupFilled) }

// Outlined is Variant(ActionIconGroupOutlined).
func (g *ActionIconGroup) Outlined() *ActionIconGroup { return g.Variant(ActionIconGroupOutlined) }

// Borderless is Variant(ActionIconGroupBorderless).
func (g *ActionIconGroup) Borderless() *ActionIconGroup { return g.Variant(ActionIconGroupBorderless) }

// Size applies the shared size of the icons and returns the group, for
// chaining.
func (g *ActionIconGroup) Size(s Size) *ActionIconGroup {
	g.size = s
	box, _ := actionIconMetrics(s)
	for _, item := range g.items {
		item.Size(box, box)
	}
	return g
}

// Xs sizes the icons down to the smallest shared size.
func (g *ActionIconGroup) Xs() *ActionIconGroup { return g.Size(SizeXs) }

// Sm sizes the icons down.
func (g *ActionIconGroup) Sm() *ActionIconGroup { return g.Size(SizeSm) }

// Md gives the built-in sizing.
func (g *ActionIconGroup) Md() *ActionIconGroup { return g.Size(SizeMd) }

// Lg sizes the icons up.
func (g *ActionIconGroup) Lg() *ActionIconGroup { return g.Size(SizeLg) }

// Vertical stacks the icons in a column instead of a row and returns the
// group, for chaining.
func (g *ActionIconGroup) Vertical(on bool) *ActionIconGroup {
	if on {
		g.Element = g.Element.Column()
	} else {
		g.Element = g.Element.Row()
	}
	return g
}

// Disabled takes the whole group out of use and returns it, for chaining.
func (g *ActionIconGroup) Disabled(on bool) *ActionIconGroup {
	g.Element = g.Element.Disabled(on)
	if on {
		g.Element = g.Element.Cursor(ui.CursorNotAllowed).Opacity(0.5)
	} else {
		g.Element = g.Element.Cursor(ui.CursorDefault)
	}
	return g
}

// Chosen returns the Key clicked in this frame, or "" when none was.
func (g *ActionIconGroup) Chosen() string { return g.chosen }

// apply paints the group's surface on this frame's theme.
func (g *ActionIconGroup) apply() {
	t := g.c.Theme()
	bg, border := t.Surface, ui.Transparent
	width := float32(0)
	switch g.variant {
	case ActionIconGroupOutlined:
		bg, border, width = t.Background, t.Border, 1
	case ActionIconGroupBorderless:
		bg = ui.Transparent
		if g.Element.Hovered() && !g.Element.IsDisabled() {
			bg = t.SurfaceHover
		}
	}
	g.Element = g.Element.Background(bg).Border(width, border)
}
