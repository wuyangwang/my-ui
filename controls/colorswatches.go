package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// ColorSwatch is one color of a ColorSwatches row.
type ColorSwatch struct {
	// Color is the swatch's color. A fully transparent one shows a
	// checkerboard made from the theme's surface.
	Color ui.Color
	// Label names the color for assistive technology.
	Label string
}

// swatchItem pairs a swatch's element with its color.
type swatchItem struct {
	element ui.Element
	color   ui.Color
}

// ColorSwatches chooses one of its colors into a bound pointer. The
// chosen swatch carries a check, and a pointer hover lifts it with a
// ring. Keyboard users move between the swatches with the arrow keys,
// as in a mygo radio group.
type ColorSwatches struct {
	ui.Element
	c       *ui.Context
	value   *ui.Color
	size    float32
	square  bool
	items   []swatchItem
	wrapped bool
}

// NewColorSwatches builds a row of swatches bound to *value. Give every
// swatch a Label so its color has a name to assistive technology.
func NewColorSwatches(c *ui.Context, value *ui.Color, swatches ...ColorSwatch) *ColorSwatches {
	cs := &ColorSwatches{c: c, value: value}
	cs.Element = ui.RadioGroup(c, func() {
		for i, sw := range swatches {
			i, sw := i, sw
			t := c.Theme()
			b := ui.RadioBase(c.Key(i), value, sw.Color).
				Center().Shrink(0).Cursor(ui.CursorPointer).Transition(StateTransition())
			if sw.Label != "" {
				b = b.Label(sw.Label)
			}
			if sw.Color.A == 0 {
				b = b.Background(t.Surface).Clip()
			} else {
				b = b.Background(sw.Color)
			}
			b.Children(func() {
				if sw.Color.A == 0 {
					checkerboard(c, t, swatchSide(cs.size))
				}
				if sw.Color == *value {
					ui.Icon(c, icons.Check()).Size(IconSizeXs, IconSizeXs).
						TextColor(checkInk(sw.Color, t))
				}
			})
			cs.items = append(cs.items, swatchItem{element: b, color: sw.Color})
		}
	}).Row().Gap(c.Theme().Space(1.5))
	cs.apply()
	return cs
}

// Size sets a swatch's side in DIPs and returns the swatches, for
// chaining. Zero (the default) is 24 DIPs.
func (cs *ColorSwatches) Size(d float32) *ColorSwatches {
	cs.size = d
	cs.apply()
	return cs
}

// Square rounds the swatches' corners instead of making them circles,
// and returns the swatches, for chaining.
func (cs *ColorSwatches) Square(on bool) *ColorSwatches {
	cs.square = on
	cs.apply()
	return cs
}

// Wrap lets the row break into lines and returns the swatches, for
// chaining.
func (cs *ColorSwatches) Wrap(on bool) *ColorSwatches {
	if on && !cs.wrapped {
		cs.Element = cs.Element.Wrap()
		cs.wrapped = true
	}
	return cs
}

// apply paints the swatches on this frame's theme.
func (cs *ColorSwatches) apply() {
	t := cs.c.Theme()
	side := swatchSide(cs.size)
	radius := side / 2
	if cs.square {
		radius = Radius(t, RadiusMd)
	}
	for _, item := range cs.items {
		ring := t.Border
		if item.color == *cs.value {
			ring = t.Text
		}
		item.element.Size(side, side).Radius(radius).InsetShadow(0, 0, 0, 1, ring)
		if item.element.Hovered() {
			item.element.Shadow(0, 0, 0, 2, t.Text)
		}
	}
}

// swatchSide is a swatch's side in DIPs.
func swatchSide(size float32) float32 {
	if size <= 0 {
		return 24
	}
	return size
}

// checkerboard paints the two tiles a transparent swatch shows under its
// color, in the theme's own surface ink.
func checkerboard(c *ui.Context, t *ui.Theme, size float32) {
	half := size / 2
	for _, at := range [][2]float32{{0, 0}, {half, half}} {
		ui.Box(c).Absolute().Left(at[0]).Top(at[1]).
			Size(half, half).Background(t.SurfaceHover).IgnoreInput(true)
	}
}

// checkInk is the ink of the check on a swatch, dark on a light color and
// light on a dark one.
func checkInk(color ui.Color, t *ui.Theme) ui.Color {
	if color.A == 0 {
		return t.Text.Alpha(0.33)
	}
	// OnAccent derives the readable ink from the swatch itself, so the
	// check mark follows the chosen color instead of a hardcoded black/white.
	return OnAccent(color).Alpha(0.33)
}
