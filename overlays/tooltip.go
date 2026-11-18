package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// TooltipSide is the side of the anchor a tip shows on.
type TooltipSide uint8

const (
	TooltipTop TooltipSide = iota
	TooltipBottom
	TooltipLeft
	TooltipRight
)

// tooltipAnchors returns the anchor on the target and the one on the tip
// for a side. The framework flips the tip when there is no room.
func tooltipAnchors(s TooltipSide) (at, self ui.Anchor) {
	switch s {
	case TooltipBottom:
		return ui.AnchorBottom, ui.AnchorTop
	case TooltipLeft:
		return ui.AnchorLeft, ui.AnchorRight
	case TooltipRight:
		return ui.AnchorRight, ui.AnchorLeft
	default:
		return ui.AnchorTop, ui.AnchorBottom
	}
}

// Tooltip shows a tip over the anchor, when the pointer rests on it or
// it has the focus.
//
//	ui.NewTooltip(c, myui.NewButton(c, "帮助").Element, "这是帮助文字").Side(myui.TooltipBottom)
//
// The tip is styled by the library, as shadcn's: the theme turned over, a
// small text, rounded corners and a shadow. Its text becomes the anchor's
// description too, so assistive technology reads it with the anchor.
// Content() builds arbitrary content instead of a single text.
type Tooltip struct {
	ui.Element
	c       *ui.Context
	build   func(tip ui.Element)
	side    TooltipSide
	tooltip ui.Element // the base, rebuilt each frame by the caller
}

// NewTooltip shows a tip with a text over the anchor.
func NewTooltip(c *ui.Context, anchor ui.Element, text string) *Tooltip {
	return newTooltip(c, anchor, func(tip ui.Element) {
		ui.Text(c, text)
	}, text)
}

// NewTooltipContent shows a tip built by build over the anchor. The tip
// is built in the overlay, so build must put its content inside the tip
// element with tip.Children(func(){ ... }); content made outside lands at
// the window's corner.
func NewTooltipContent(c *ui.Context, anchor ui.Element, build func(tip ui.Element)) *Tooltip {
	return newTooltip(c, anchor, build, "")
}

func newTooltip(c *ui.Context, anchor ui.Element, build func(tip ui.Element), description string) *Tooltip {
	t := &Tooltip{c: c, build: build}
	if description != "" {
		anchor = anchor.Description(description)
	}
	t.Element = ui.TooltipBase(c, anchor, func(tip ui.Element) {
		th := c.Theme()
		fill, ink := InverseColors(th)
		at, self := tooltipAnchors(t.side)
		tip.Background(fill).TextColor(ink).
			Padding(th.Space(1.5), th.Space(3)).
			Radius(th.Radius).
			FontSize(TextSize(th, TextXs)).
			Shadow(0, 2, 8, 0, ui.RGBA(0, 0, 0, 0.2)).
			MaxWidth(th.Space(60)).
			AttachTo(anchor, at, self).
			Transition(StateTransition())
		OverlayEnter(c, tip, "tooltip-enter")
		if t.build != nil {
			tip.Children(func() { t.build(tip) })
		}
	})
	return t
}

// Side sets the side of the anchor the tip shows on, for chaining. The
// view runs each frame, so the next frame's tip picks up the side.
func (t *Tooltip) Side(s TooltipSide) *Tooltip {
	t.side = s
	return t
}

// Content replaces the content of the tip with what build puts in it,
// for chaining.
func (t *Tooltip) Content(build func(tip ui.Element)) *Tooltip {
	t.build = build
	return t
}
