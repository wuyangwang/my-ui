package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// PopoverSide is the side of the anchor a panel shows on. It shares the
// tooltip side values and their anchor mapping, where the framework flips
// the panel when there is no room on that side.
type PopoverSide = TooltipSide

const (
	PopoverTop    = TooltipTop
	PopoverBottom = TooltipBottom
	PopoverLeft   = TooltipLeft
	PopoverRight  = TooltipRight
)

// PopoverOptions configures a panel at construction time. The same
// setting is available as the Side method on the returned Popover.
type PopoverOptions struct {
	Side PopoverSide
}

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
	open   *bool
	side   PopoverSide
	anchor ui.Element
	*Parts
}

// NewPopover shows a panel over its anchor while open is true. The flag
// is the app's: it opens and closes the panel. The panel shows below the
// anchor by default, as the framework's own popover does; Side changes it.
func NewPopover(c *ui.Context, open *bool, anchor func() ui.Element, build func(panel ui.Element), options ...PopoverOptions) *Popover {
	o := PopoverOptions{Side: PopoverBottom}
	if len(options) > 0 {
		o = options[0]
	}
	p := &Popover{open: open, side: o.Side}
	var a ui.Element
	has := false
	if anchor != nil {
		a = anchor()
		has = true
	}
	p.anchor = a
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
		attachSide(panel, a, p.side)
		OverlayEnter(c, panel, "popover-enter")
		p.Parts = BuildSlots(c, build, panel, open)
	})
	p.Element = panel
	return p
}

// attachSide puts a panel on the side of its anchor that tooltipAnchors
// maps to. Both the constructor's build and the fluent Side use it.
func attachSide(panel, anchor ui.Element, side PopoverSide) {
	at, self := tooltipAnchors(side)
	panel.AttachTo(anchor, at, self)
}

// Side sets the side of the anchor the panel shows on, for chaining. The
// view builds the panel before this runs, so it re-attaches the live panel.
func (p *Popover) Side(s PopoverSide) *Popover {
	p.side = s
	if p.Element.Valid() {
		attachSide(p.Element, p.anchor, s)
	}
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
