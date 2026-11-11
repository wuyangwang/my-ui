package myui

import (
	"github.com/egoist/mygo/ui"
)

// Popover shows a panel over its anchor, while the app's flag is true.
//
//	myui.NewPopover(c, &app.open, func() ui.Element {
//		return myui.NewButton(c, "更多").Element
//	}, func(panel ui.Element) {
//		ui.Text(c, "面板里的内容")
//	})
//
// The anchor is built by the app, and a click on it opens the panel; a
// click outside the panel closes it, and Escape closes it too.
//
// The panel is styled by the library, as shadcn's: a card of the surface
// tone over the anchor, with a border and a shadow, and the app fills
// it.
type Popover struct {
	ui.Element
	open *bool
}

// NewPopover shows a panel over its anchor while open is true. The flag
// is the app's: it opens and closes the panel.
//
// The anchor is built where the popover is, so it sits beside the panel
// it opens; a click on it opens the panel in the same frame.
func NewPopover(c *ui.Context, open *bool, anchor func() ui.Element, build func(panel ui.Element)) *Popover {
	p := &Popover{open: open}
	var a ui.Element
	has := false
	if anchor != nil {
		a = anchor()
		has = true
	}
	// A click on the anchor opens the panel.
	if has && a.Clicked() && p.open != nil {
		*p.open = true
	}
	panel := ui.PopoverBase(c, a, open, func(panel ui.Element) {
		th := c.Theme()
		panel.Background(th.Surface).Border(1, th.Border).Radius(th.Radius).
			Padding(th.Space(4)).Shadow(0, 4, 12, 0, ui.RGBA(0, 0, 0, 0.15)).
			MaxWidth(th.Space(60)).
			Transition(stateTransition())
		if build != nil {
			// The content belongs to the panel: the panel is built in
			// the overlay, where content made now would land at the
			// window's corner.
			panel.Children(func() { build(panel) })
		}
	})
	p.Element = panel
	return p
}

// Open reports whether the panel shows.
func (p *Popover) Open() bool { return p.open != nil && *p.open }

// SetOpen opens or closes the panel and returns it, for chaining.
func (p *Popover) SetOpen(open bool) *Popover {
	if p.open != nil {
		*p.open = open
	}
	return p
}
