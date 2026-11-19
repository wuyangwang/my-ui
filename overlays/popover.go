package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// Popover shows a panel over its anchor, while the app's flag is true.
//
//	overlays.NewPopover(c, &app.open, func() ui.Element {
//		return controls.NewButton(c, "更多").Element
//	}, func(panel ui.Element) {
//		ui.Text(c, "面板里的内容")
//	}).Title("更多操作")
//
// The anchor is built by the app, and a click on it opens the panel; a
// click outside the panel closes it, and Escape closes it too.
//
// The panel is styled by the library, as shadcn's: a card over the anchor,
// with a border and a shadow. Title, Description and Footer add the
// structured parts of a shadcn popover.
type Popover struct {
	ui.Element
	open *bool
	*Parts
}

// NewPopover shows a panel over its anchor while open is true. The flag
// is the app's: it opens and closes the panel.
func NewPopover(c *ui.Context, open *bool, anchor func() ui.Element, build func(panel ui.Element)) *Popover {
	p := &Popover{open: open}
	var a ui.Element
	has := false
	if anchor != nil {
		a = anchor()
		has = true
	}
	// A click on the anchor toggles the panel, as a popover trigger does
	// on the web: the first opens it, the next closes it. The framework
	// treats the anchor as inside the panel, so it does not close it for
	// us.
	if has && a.Clicked() && p.open != nil {
		*p.open = !*p.open
	}
	panel := ui.PopoverBase(c, a, open, func(panel ui.Element) {
		th := c.Theme()
		panel.Background(th.Background).Border(1, th.Border).Radius(th.Radius).
			Padding(th.Space(4)).Shadow(0, 4, 12, 0, ui.RGBA(0, 0, 0, 0.15)).
			MaxWidth(th.Space(60)).
			Transition(StateTransition())
		OverlayEnter(c, panel, "popover-enter")
		p.Parts = BuildSlots(c, build, panel, open)
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

// Dismissed reports whether the panel was dismissed in this frame — by a
// click outside it or by Escape.
func (p *Popover) Dismissed() bool { return p.Element.Dismissed() }

// OnDismiss calls fn when the panel was dismissed in this frame.
func (p *Popover) OnDismiss(fn func()) *Popover {
	if fn != nil && p.Dismissed() {
		fn()
	}
	return p
}

// Title sets the heading of the panel, for chaining.
func (p *Popover) Title(s string) *Popover {
	if p.Parts != nil {
		p.Parts.Title(s)
	}
	return p
}

// Description sets the muted line under the title, for chaining.
func (p *Popover) Description(s string) *Popover {
	if p.Parts != nil {
		p.Parts.Description(s)
	}
	return p
}

// Footer builds the right-aligned action row of the panel.
func (p *Popover) Footer(build func()) *Popover {
	if p.Parts != nil {
		p.Parts.Footer(build)
	}
	return p
}

// ShowClose shows or hides the corner close button.
func (p *Popover) ShowClose(on bool) *Popover {
	if p.Parts != nil {
		p.Parts.ShowClose(on)
	}
	return p
}
