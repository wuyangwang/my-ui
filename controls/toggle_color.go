package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// Toggle is a button that stays pressed, with my-ui's shared look.
type Toggle struct {
	ui.Element
	c *ui.Context
}

// NewToggle builds a toggle bound to *on.
func NewToggle(c *ui.Context, on *bool, label string) *Toggle {
	t := c.Theme()
	tg := &Toggle{c: c}
	tg.Element = ui.Toggle(c, on, label).
		Padding(t.Space(1), t.Space(2.5)).Height(32).
		Radius(Radius(t, RadiusMd)).Transition(StateTransition())
	return tg
}

// Size picks the toggle's height on shadcn's scale.
func (t *Toggle) Size(sz Size) *Toggle {
	h := float32(32)
	switch sz {
	case SizeXs:
		h = 28
	case SizeLg:
		h = 40
	}
	t.Element = t.Element.Height(h)
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
