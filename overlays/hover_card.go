package overlays

import "github.com/egoist/mygo/ui"

// HoverCard opens a panel while the pointer is on its anchor or on the
// panel itself. The open flag belongs to the app, like Popover's.
type HoverCard struct {
	ui.Element
}

// NewHoverCard builds a hover card. anchor builds the trigger element
// and build fills the panel; both run every frame. A few frames of
// grace let the pointer travel from the trigger into the panel without
// closing it — a frame count, never a timer.
func NewHoverCard(c *ui.Context, open *bool, anchor func() ui.Element, build func(panel ui.Element)) *HoverCard {
	if open == nil {
		open = new(bool)
	}
	t := c.Theme()
	var away *int
	var hostHover, panelHover bool

	h := &HoverCard{}
	h.Element = ui.Column(c).Children(func() {
		var trigger ui.Element
		// The trigger is built before the overlay, so its click lands
		// in this frame and the panel gets the layer as its parent.
		host := ui.Row(c.Key("hover card")).Children(func() {
			if anchor != nil {
				trigger = anchor()
			} else {
				trigger = ui.Box(c)
			}
		})
		away = ui.Local(host, "hover-away", func() int { return 0 })
		hostHover = host.Hovered()

		panel := ui.PopoverBase(c, trigger, open, func(p ui.Element) {
			p = p.Background(t.Background).Border(1, t.Border).
				Radius(t.Radius).Padding(t.Space(3), t.Space(4)).
				Shadow(0, 4, 16, -4, t.Text.Alpha(0.18))
			if build != nil {
				build(p)
			}
			panelHover = p.Hovered()
		})
		_ = panel
	})

	if away != nil {
		if hostHover || panelHover {
			*away = 0
			*open = true
		} else {
			*away++
			if *away > 3 {
				*open = false
			}
		}
	}
	return h
}
