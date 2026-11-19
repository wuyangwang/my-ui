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

// PopoverOptions configures a Popover at construction time. The open flag
// is the app's. Every field is read again each frame from the options the
// view passes.
type PopoverOptions struct {
	// Open is the app's flag; the panel belongs to the app.
	Open *bool
	// Trigger builds the anchor. It runs before the overlay so its click
	// lands in the same frame and toggles the panel.
	Trigger func() ui.Element
	// Title is the bold heading and the panel's accessible name.
	Title string
	// Description is the muted line under the title.
	Description string
	// Body fills the panel. The library builds it inside panel.Children.
	Body func()
	// Footer builds the right-aligned action row.
	Footer func()
	// OnOpenChange runs when the panel opens or closes, from any source.
	// It is kept across frames, so attaching it late does not lose a
	// change.
	OnOpenChange func(open bool)
	// OnDismiss runs when a click outside or Escape dismisses the panel.
	OnDismiss func()
	// Side is the side of the anchor the panel shows on. A nil value
	// keeps the default, PopoverBottom.
	Side *PopoverSide
	// ShowClose hides the corner close button when false. Nil keeps it on.
	ShowClose *bool
}

// popoverState is what a Popover keeps in its Local between frames.
type popoverState struct {
	open  openState
	parts panelParts
	side  PopoverSide
}

// Popover shows a panel over its anchor while the app's flag is true.
//
//	overlays.NewPopover(c, overlays.PopoverOptions{
//		Open:    &app.moreOpen,
//		Trigger: func() ui.Element { return controls.NewButton(c, "更多").Element },
//		Body:    func() { ui.Text(c, "面板里的内容") },
//	})
//
// The library builds the trigger first, then the panel, and puts the Body
// inside the panel. A click on the anchor toggles the panel; a click
// outside or Escape closes it.
type Popover struct {
	ui.Element
	st     *popoverState
	open   *bool
	anchor ui.Element
}

// NewPopover shows a panel over its anchor while o.Open is true.
func NewPopover(c *ui.Context, o PopoverOptions) *Popover {
	st := stateFor(c, o.Open, "popover", func() popoverState {
		return popoverState{
			parts: panelParts{ShowClose: true},
			side:  PopoverBottom,
		}
	})
	if o.Title != "" {
		st.parts.Title = o.Title
	}
	if o.Description != "" {
		st.parts.Description = o.Description
	}
	if o.Body != nil {
		st.parts.Body = o.Body
	}
	if o.Footer != nil {
		st.parts.Footer = o.Footer
	}
	if o.ShowClose != nil {
		st.parts.ShowClose = *o.ShowClose
	}
	if o.Side != nil {
		st.side = *o.Side
	}
	if o.OnOpenChange != nil {
		st.open.onChange = o.OnOpenChange
	}
	if o.OnDismiss != nil {
		st.open.onDismiss = o.OnDismiss
	}
	st.open.flush()

	p := &Popover{st: st, open: st.open.openPtr(o.Open)}
	if o.Trigger != nil {
		a := o.Trigger()
		p.anchor = a
		// A click on the anchor toggles the panel, as a popover does on
		// the web: the framework treats the anchor as inside the panel,
		// so it does not close the panel for us.
		if a.Clicked() {
			next := !*p.open
			*p.open = next
			st.open.record(next)
		}
	}
	panel := ui.PopoverBase(c, p.anchor, p.open, func(panel ui.Element) {
		th := c.Theme()
		panel.Background(th.Background).Border(1, th.Border).Radius(th.Radius).
			Padding(th.Space(4)).Shadow(0, 4, 12, 0, ui.RGBA(0, 0, 0, 0.15)).
			MaxWidth(th.Space(60)).
			Transition(StateTransition())
		attachSide(panel, p.anchor, st.side)
		OverlayEnter(c, panel, "popover-enter")
		if st.parts.Title != "" {
			panel.Label(st.parts.Title)
		}
		buildPanelParts(c, panel, st.parts, func() {
			*p.open = false
			st.open.record(false)
		})
	})
	p.Element = panel
	if panel.Dismissed() {
		st.open.record(false)
		st.open.recordDismiss()
	}
	return p
}

// attachSide puts a panel on the side of its anchor that tooltipAnchors
// maps to. Both the constructor's build and the fluent Side use it.
func attachSide(panel, anchor ui.Element, side PopoverSide) {
	at, self := tooltipAnchors(side)
	panel.AttachTo(anchor, at, self)
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

// OnOpenChange registers fn for open changes and returns the popover. A
// change that happened before this call is delivered at once.
func (p *Popover) OnOpenChange(fn func(open bool)) *Popover {
	p.st.open.onChange = fn
	p.st.open.flush()
	return p
}

// OnDismiss registers fn for dismissals and returns the popover.
func (p *Popover) OnDismiss(fn func()) *Popover {
	p.st.open.onDismiss = fn
	p.st.open.flush()
	return p
}

// Side sets the side of the anchor the panel shows on, for chaining. It
// works while the panel is closed: the side is used when it next opens.
func (p *Popover) Side(s PopoverSide) *Popover {
	p.st.side = s
	if p.Element.Valid() {
		attachSide(p.Element, p.anchor, s)
	}
	return p
}

// Title sets the heading of the panel, for chaining.
func (p *Popover) Title(s string) *Popover {
	p.st.parts.Title = s
	return p
}

// Description sets the muted line under the title, for chaining.
func (p *Popover) Description(s string) *Popover {
	p.st.parts.Description = s
	return p
}

// Footer builds the right-aligned action row of the panel.
func (p *Popover) Footer(build func()) *Popover {
	p.st.parts.Footer = build
	return p
}

// ShowClose shows or hides the corner close button.
func (p *Popover) ShowClose(on bool) *Popover {
	p.st.parts.ShowClose = on
	return p
}
