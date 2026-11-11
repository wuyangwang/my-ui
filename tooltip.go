package myui

import (
	"github.com/egoist/mygo/ui"
)

// Tooltip shows a tip over the anchor, when the pointer rests on it or
// it has the focus.
//
//	myui.NewTooltip(c, myui.NewButton(c, "帮助").Element, "这是帮助文字")
//
// The tip is styled by the library, as shadcn's: the theme turned over
// (the ink of the window on its background, or the theme's own inverse),
// a small text, rounded corners and a shadow, above the anchor. The
// tip's text becomes the anchor's description too, so assistive
// technology reads it with the anchor.
type Tooltip struct {
	ui.Element
}

// NewTooltip shows a tip with a text over the anchor. The anchor is the
// element the tip belongs to: a button, an icon, a piece of text.
//
// The tip shows after the pointer has rested on the anchor for a while,
// and at once when the anchor takes the keyboard focus; the theme's
// TooltipDelay sets how long. It goes as the pointer leaves the anchor
// or presses it, and with Escape.
func NewTooltip(c *ui.Context, anchor ui.Element, text string) *Tooltip {
	t := &Tooltip{}
	// The tip's text is the anchor's description too: a screen reader
	// reads it with the anchor, as the tip itself is not read.
	anchor = anchor.Description(text)
	tip := ui.TooltipBase(c, anchor, func(tip ui.Element) {
		th := c.Theme()
		// The tip is the theme turned over, so that it stands out from
		// what it is over.
		fill, ink := InverseColors(th)
		tip.Background(fill).TextColor(ink).
			Padding(th.Space(1.5), th.Space(3)).
			Radius(th.Radius).
			FontSize(TextSize(th, TextXs)).
			Shadow(0, 2, 8, 0, ui.RGBA(0, 0, 0, 0.2)).
			MaxWidth(th.Space(60)).
			Margin(0, 0, th.Space(1.5), 0).
			Transition(stateTransition())
		// The text belongs to the tip: the tip is built in the overlay,
		// where a text made now would land at the window's corner.
		tip.Children(func() {
			ui.Text(c, text)
		})
	})
	t.Element = tip
	return t
}
