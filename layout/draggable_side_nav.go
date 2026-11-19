package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// SideNavState is what a DraggableSideNav hands to its builders every
// frame, so they can decide what to show while the rail collapses.
type SideNavState struct {
	// Expanded says whether the nav is at its full width.
	Expanded bool
	// Collapse runs from 0 (full) to 1 (rail) with the width. Feeding it
	// to a list's slide lets rows turn into centered icons on the way.
	Collapse float32
	// RailWidth is the width of the collapsed rail in DIPs.
	RailWidth float32
}

// DraggableSideNavProps configures NewDraggableSideNav.
type DraggableSideNavProps struct {
	// MinWidth and MaxWidth clamp the dragged width. Defaults are the
	// theme's Space(14) and Space(80).
	MinWidth, MaxWidth float32
	// Header builds the workspace switcher above the body.
	Header func(nav SideNavState)
	// Footer builds the row pinned to the foot.
	Footer func(nav SideNavState)
}

// DraggableSideNav is a workspace-style rail that resizes by dragging its
// edge, collapses to a slim rail, and tells its content which way it
// shows. The app owns width (*float32) and expanded (*bool).
type DraggableSideNav struct {
	ui.Element
	c        *ui.Context
	width    *float32
	expanded *bool
}

// NewDraggableSideNav builds the rail. body draws the middle section and
// receives the current SideNavState; Header and Footer do the same when
// set. The edge tab toggles collapse; dragging the edge resizes, and
// leaving it under a third of the way from MinWidth to the full width
// collapses it.
func NewDraggableSideNav(c *ui.Context, width *float32, expanded *bool, p DraggableSideNavProps, body func(nav SideNavState)) *DraggableSideNav {
	t := c.Theme()
	s := &DraggableSideNav{c: c}
	root := ui.Row(c).AlignItems(ui.Stretch).FillHeight()
	if width == nil {
		width = ui.Local(root, "side nav width", func() float32 { return t.Space(75) })
	}
	if expanded == nil {
		expanded = ui.Local(root, "side nav expanded", func() bool { return true })
	}
	s.width, s.expanded = width, expanded
	rail := p.MinWidth
	if rail == 0 {
		rail = t.Space(14)
	}
	max := p.MaxWidth
	if max == 0 {
		max = t.Space(80)
	}
	if *width < rail {
		*width = rail
	}
	if *width > max {
		*width = max
	}
	last := ui.Local(root, "side nav last width", func() float32 { return *width })
	if *expanded {
		*last = *width
	}
	shown := *width
	if !*expanded {
		shown = rail
	}
	nav := SideNavState{Expanded: *expanded, RailWidth: rail}
	if *expanded && *width > rail {
		nav.Collapse = 0
	} else {
		nav.Collapse = 1
	}

	// The column of the rail and its resize edge.
	root.Children(func() {
		col := ui.Column(c).Width(shown).Clip().
			Background(t.Surface).
			BorderWidth(0, 1, 0, 0).BorderColor(t.Border).
			Transition(StateTransition())
		col.Children(func() {
			if p.Header != nil {
				p.Header(nav)
			}
			if body != nil {
				wrap := ui.Column(c).Grow(1).MinHeight(0)
				wrap.Children(func() { body(nav) })
			}
			if p.Footer != nil {
				p.Footer(nav)
			}
		})

		// The edge: a few DIPs of hit area. Dragging it resizes; its
		// toggle tab (a button of its own) collapses the rail.
		edge := ui.Box(c).Width(t.Space(1.5)).Cursor(ui.CursorResizeEW).
			Background(t.Border).Transition(StateTransition())
		dx, _, held := edge.Dragged()
		if held && dx != 0 {
			v := *width + dx
			if v < rail {
				v = rail
			}
			if v > max {
				v = max
			}
			*width, *expanded, *last = v, true, v
			shown = v
		}
		if _, _, ok := edge.DragEnded(); ok {
			threshold := rail + (*width-rail)/3
			if *width < threshold {
				*expanded = false
				*last = *width
				shown = rail
			}
		}
		if edge.Hovered() || edge.Pressed() {
			edge = edge.Background(t.Accent)
		}
		// The tab sits on the edge, shown while the pointer is near it
		// or when the rail is collapsed. Its click never feeds the drag.
		opacity := float32(0)
		if !*expanded || edge.Hovered() {
			opacity = 1
		}
		tab := ui.Row(c).AlignItems(ui.Center).Justify(ui.Center).
			Width(t.Space(5)).Height(t.Space(8)).
			Background(t.Surface).Border(1, t.Border).
			Radius(Radius(t, RadiusMd)).Opacity(opacity)
		tab.Children(func() {
			glyph := icons.ChevronsLeft()
			if !*expanded {
				glyph = icons.ChevronsRight()
			}
			tg := controls.NewIconAction(c, controls.IconActionOptions{
				SVG: glyph, Label: "折叠侧栏", Size: t.Space(5),
			})
			if tg.Clicked() {
				if *expanded {
					*last = *width
					*expanded = false
				} else {
					*expanded = true
					if *width < rail {
						*width = *last
					}
				}
			}
		})

		ui.Box(c).Grow(1).Clip()
	})
	s.Element = root
	return s
}

// Expanded reports whether the nav is at its full width.
func (s *DraggableSideNav) Expanded() bool { return s.expanded != nil && *s.expanded }

// Width reports the nav's expanded width in DIPs.
func (s *DraggableSideNav) Width() float32 {
	if s.width == nil {
		return 0
	}
	return *s.width
}
