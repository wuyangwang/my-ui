package myui

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Icon is an SVG icon with variant and size presets:
//
//	myui.NewIcon(c, icons.Save()).Md()
//	myui.NewIcon(c, icons.Loading()).Primary().Spin()
//
// The icon draws in the color of the text around it, so a variant
// paints that color and the rest keeps it. The chain applies the
// presets from the current theme every frame, so it follows palette and
// appearance changes. Call it each frame before any app-specific
// styling. Icon embeds ui.Element, so other ui methods remain
// available.
type Icon struct {
	ui.Element
	c    *ui.Context
	size Size
}

// NewIcon creates an icon drawing svg, at the size of the text around
// it: the size presets set the size of the icon itself.
func NewIcon(c *ui.Context, svg *ui.SVG) *Icon {
	return &Icon{Element: ui.Icon(c, svg), c: c}
}

// Variant applies the color of v and returns the icon, for chaining.
// The zero variant keeps the color of the text around the icon.
func (i *Icon) Variant(v Variant) *Icon {
	i.applySizing()
	t := i.c.Theme()
	switch v {
	case Primary, Link:
		i.TextColor(t.Accent)
	case Destructive:
		i.TextColor(t.Danger)
	}
	return i
}

// Primary is Variant(myui.Primary).
func (i *Icon) Primary() *Icon { return i.Variant(Primary) }

// Secondary is Variant(myui.Secondary), alias of Native.
func (i *Icon) Secondary() *Icon { return i.Variant(Secondary) }

// Native is Variant(myui.Native).
func (i *Icon) Native() *Icon { return i.Variant(Native) }

// Outline is Variant(myui.Outline); an icon has no border of its own.
func (i *Icon) Outline() *Icon { return i.Variant(Outline) }

// Ghost is Variant(myui.Ghost); an icon has no border of its own.
func (i *Icon) Ghost() *Icon { return i.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (i *Icon) Destructive() *Icon { return i.Variant(Destructive) }

// Link is Variant(myui.Link).
func (i *Icon) Link() *Icon { return i.Variant(Link) }

// Size applies the sizing of s and returns the icon, for chaining. The
// sizes are the ones of an icon on the desktop: 14, 20 and 24 DIPs.
func (i *Icon) Size(s Size) *Icon {
	i.size = s
	i.applySizing()
	return i
}

// Sm sizes the icon down.
func (i *Icon) Sm() *Icon { return i.Size(SizeSm) }

// Md gives the icon the size of the text around it.
func (i *Icon) Md() *Icon { return i.Size(SizeMd) }

// Lg sizes the icon up.
func (i *Icon) Lg() *Icon { return i.Size(SizeLg) }

// Label names the icon for assistive technology and returns it, for
// chaining. An icon that means something needs a name: the name of
// what the icon stands for.
func (i *Icon) Label(s string) *Icon {
	i.Element = i.Element.Label(s)
	return i
}

// Spin turns the icon, as a spinner does, and returns it, for
// chaining. The turn slows down with the desktop's reduce motion.
func (i *Icon) Spin() *Icon {
	i.Element = i.Element.Rotate(i.Loop("spin", time.Second, ui.Linear) * 360)
	return i
}

func (i *Icon) applySizing() {
	var d float32
	switch i.size {
	case SizeSm:
		d = 14
	case SizeLg:
		d = 24
	case SizeMd:
		d = 20
	default:
		return
	}
	i.Element = i.Element.Size(d, d)
}
