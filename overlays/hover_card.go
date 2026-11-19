package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// HoverCardOptions configures a HoverCard at construction time. The open
// flag is the app's. The library wraps Body in panel.Children, so callers
// never have to know that the panel's parent is the overlay layer.
type HoverCardOptions struct {
	// Open is the app's flag. A nil flag makes the card keep its own
	// stable state, but distinct cards under one parent should still set
	// Key so they do not share an identity.
	Open *bool
	// Trigger builds the hover target. It runs before the overlay.
	Trigger func() ui.Element
	// Body fills the panel. The library builds it inside panel.Children.
	Body func()
	// OnOpenChange runs when the card opens or closes. It is kept across
	// frames, so attaching it late does not lose a change.
	OnOpenChange func(open bool)
	// Key scopes the card when several share a parent and no Open flag.
	Key string
}

// hoverState is what a HoverCard keeps in its Local between frames.
type hoverState struct {
	open openState
}

// HoverCard opens a panel while the pointer is on its anchor or on the
// panel itself. A few frames of grace let the pointer travel from the
// trigger into the panel without closing it — a frame count, never a
// timer.
type HoverCard struct {
	ui.Element
	st   *hoverState
	open *bool
}

// NewHoverCard builds a hover card. Trigger builds the target and Body
// fills the panel; both run every frame.
//
//	overlays.NewHoverCard(c, overlays.HoverCardOptions{
//		Open:    &app.cardOpen,
//		Trigger: func() ui.Element { return controls.NewButton(c, "悬停我").Element },
//		Body:    func() { ui.Text(c, "卡片内容") },
//	})
func NewHoverCard(c *ui.Context, o HoverCardOptions) *HoverCard {
	st := stateFor(c, o.Open, "hovercard:"+o.Key, func() hoverState {
		return hoverState{}
	})
	if o.OnOpenChange != nil {
		st.open.onChange = o.OnOpenChange
	}
	st.open.flush()
	open := st.open.openPtr(o.Open)

	t := c.Theme()
	var away *int
	var hostHover, panelHover bool

	h := &HoverCard{st: st, open: open}
	h.Element = ui.Column(c).Children(func() {
		var trigger ui.Element
		// The trigger is built before the overlay, so its input lands
		// in this frame and the panel gets the layer as its parent.
		host := ui.Row(c).Children(func() {
			if o.Trigger != nil {
				trigger = o.Trigger()
			} else {
				trigger = ui.Box(c)
			}
		})
		away = ui.Local(host, "hover-away", func() int { return 0 })
		hostHover = host.Hovered()

		panel := ui.PopoverBase(c, trigger, open, func(p ui.Element) {
			p = p.Background(t.Background).Border(1, t.Border).
				Radius(Radius(t, RadiusMd)).Padding(t.Space(3), t.Space(4)).
				Shadow(0, 4, 16, -4, t.Text.Alpha(0.18))
			if o.Body != nil {
				p.Children(o.Body)
			}
			panelHover = p.Hovered()
		})
		_ = panel
	})

	if away != nil {
		if hostHover || panelHover {
			if !*open {
				*open = true
				st.open.record(true)
			}
			*away = 0
		} else {
			*away++
			if *away > 3 && *open {
				*open = false
				st.open.record(false)
			}
		}
	}
	return h
}

// Open reports whether the card shows.
func (h *HoverCard) Open() bool { return h.open != nil && *h.open }

// OnOpenChange registers fn for open changes and returns the card. A
// change that happened before this call is delivered at once.
func (h *HoverCard) OnOpenChange(fn func(bool)) *HoverCard {
	h.st.open.onChange = fn
	h.st.open.flush()
	return h
}
