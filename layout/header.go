package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// HeaderProps are the slots of a Header. Every builder is optional; a nil
// slot leaves its place empty. Mobile moves Nav to the front, as a menu
// button, and gives the room left to Logo.
type HeaderProps struct {
	// Logo builds the brand at the start of the bar, or its middle on
	// mobile.
	Logo func()
	// Nav builds the links. On mobile this is where the menu button goes.
	Nav func()
	// Actions builds the trailing controls: search, theme, account.
	Actions func()
	// Mobile switches to the compact arrangement.
	Mobile bool
}

// Header is a site's top bar, 64 DIPs tall by default, with a hairline
// under it. It owns no state: the slots are built by the app, so their
// clicks and pointers work as anywhere else.
type Header struct {
	ui.Element
	c      *ui.Context
	height float32
	sticky bool
}

// NewHeader builds the bar with the slots in HeaderProps. The height
// defaults to the theme's Space(16) (64 DIPs); Height overrides it.
func NewHeader(c *ui.Context, p HeaderProps) *Header {
	t := c.Theme()
	h := &Header{c: c, height: t.Space(16)}
	row := ui.Row(c).FillWidth().AlignItems(ui.Center).
		Gap(t.Space(2)).Height(h.height).
		Padding(0, t.Space(6)).
		Background(t.Background).
		BorderWidth(0, 0, 1, 0).BorderColor(t.Border).
		Transition(StateTransition())
	row.Children(func() {
		build := func(fn func()) {
			if fn != nil {
				fn()
			}
		}
		if p.Mobile {
			// The menu button leads, the brand takes what room is left.
			build(p.Nav)
			center := ui.Row(c).Grow(1).Justify(ui.Center)
			center.Children(func() { build(p.Logo) })
			build(p.Actions)
			return
		}
		build(p.Logo)
		nav := ui.Row(c).MarginStart(t.Space(6)).Gap(t.Space(6))
		nav.Children(func() { build(p.Nav) })
		ui.Box(c).Grow(1)
		build(p.Actions)
	})
	h.Element = row
	return h
}

// Height sets the bar's height in DIPs and returns it, for chaining.
func (h *Header) Height(v float32) *Header {
	h.height = v
	h.Element = h.Element.Height(v)
	return h
}

// Sticky keeps the bar at the top while the page scrolls under it, as the
// docs layout wants. It returns the header for chaining.
func (h *Header) Sticky() *Header {
	h.sticky = true
	h.Element = h.Element.Sticky()
	return h
}
