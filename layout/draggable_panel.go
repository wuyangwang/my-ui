package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// PanelPlacement is the edge of the container a DraggablePanel sits at.
type PanelPlacement uint8

const (
	// PanelRight puts the panel on the right, as an inspector.
	PanelRight PanelPlacement = iota
	// PanelLeft puts the panel on the left, as a session list.
	PanelLeft
)

// DraggablePanelProps configures NewDraggablePanel.
type DraggablePanelProps struct {
	// Placement picks the edge; zero is PanelRight.
	Placement PanelPlacement
	// MinSize and MaxSize clamp the dragged size in DIPs. Zero MinSize is
	// the theme's Space(20); zero MaxSize is unconstrained.
	MinSize, MaxSize float32
	// CollapseThreshold collapses the panel when a drag leaves it below
	// this size. Zero disables drag-to-collapse.
	CollapseThreshold float32
}

// DraggablePanel is a panel at an edge of its container that resizes by
// dragging its seam and collapses with a button on it. The app owns the
// size (*float32) and the state (*bool); the panel only writes them, so
// they survive from frame to frame and can be set from elsewhere.
//
// The body is a shadcn-style surface: t.Surface and a 1px t.Border, read
// from c.Theme() each frame. The seam lights up under the pointer and
// turns accent-colored while it is dragged. A collapsed panel leaves a
// slim grip so it can be reopened; the size it had is kept and restored on
// expand. The toggle is its own control, never swallowed by the seam's
// drag.
type DraggablePanel struct {
	ui.Element
	c        *ui.Context
	size     *float32
	expanded *bool
	p        DraggablePanelProps
}

// NewDraggablePanel builds the panel. size is the width of the panel when
// expanded, written as the seam is dragged; expanded says whether it shows
// at that width. Both may be nil, in which case the panel keeps its own
// state in a ui.Local of its root. build draws the panel's content.
func NewDraggablePanel(c *ui.Context, size *float32, expanded *bool, p DraggablePanelProps, build func()) *DraggablePanel {
	t := c.Theme()
	d := &DraggablePanel{c: c, p: p}
	root := ui.Row(c).AlignItems(ui.Stretch).FillHeight()
	if size == nil {
		size = ui.Local(root, "panel size", func() float32 { return t.Space(40) })
	}
	if expanded == nil {
		expanded = ui.Local(root, "panel expanded", func() bool { return true })
	}
	d.size, d.expanded = size, expanded
	last := ui.Local(root, "panel last size", func() float32 { return *size })
	min := p.MinSize
	if min == 0 {
		min = t.Space(20)
	}
	if *size < min {
		*size = min
	}
	if p.MaxSize > 0 && *size > p.MaxSize {
		*size = p.MaxSize
	}
	if *expanded && *size >= min {
		*last = *size
	}

	// buildSeam returns the seam and writes the drag into the app's
	// pointer state. Reading Dragged is the only thing that moves the
	// size: no timer and no goroutine is involved.
	buildSeam := func() ui.Element {
		seam := ui.Box(c).Width(t.Space(1.5)).Cursor(ui.CursorResizeEW).
			Background(t.Border).Transition(StateTransition())
		dx, _, held := seam.Dragged()
		if held && dx != 0 {
			sign := float32(1)
			if p.Placement == PanelRight {
				sign = -1
			}
			v := *size + sign*dx
			if v < min {
				v = min
			}
			if p.MaxSize > 0 && v > p.MaxSize {
				v = p.MaxSize
			}
			*size, *expanded, *last = v, true, v
		}
		if _, _, ok := seam.DragEnded(); ok && p.CollapseThreshold > 0 && *size < p.CollapseThreshold {
			*expanded = false
			*last = *size
		}
		if seam.Hovered() || seam.Pressed() {
			seam = seam.Background(t.Accent)
		}
		return seam
	}

	// buildPanel returns the panel body. A collapsed body holds only the
	// toggle grip; the content is built only while expanded.
	buildPanel := func() ui.Element {
		panelWidth := *size
		if !*expanded {
			panelWidth = t.Space(4)
		}
		body := ui.Column(c).Width(panelWidth).Clip().
			Background(t.Surface).
			BorderWidth(0, 0, 0, 1).BorderColor(t.Border).
			Transition(StateTransition())
		body.Children(func() {
			if *expanded {
				if build != nil {
					build()
				}
				return
			}
			ui.Row(c).Grow(1).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
				glyph := icons.ChevronsRight()
				if p.Placement == PanelRight {
					glyph = icons.ChevronsLeft()
				}
				tg := controls.NewIconAction(c, controls.IconActionOptions{
					SVG: glyph, Label: "展开面板", Size: t.Space(6),
				})
				if tg.Clicked() {
					*expanded = true
					if *size < min {
						*size = *last
					}
				}
			})
		})
		return body
	}

	root.Children(func() {
		if p.Placement == PanelLeft {
			buildPanel()
			buildSeam()
			ui.Box(c).Grow(1).Clip()
			return
		}
		ui.Box(c).Grow(1).Clip()
		buildSeam()
		buildPanel()
	})
	d.Element = root
	return d
}

// Expanded reports whether the panel is open.
func (d *DraggablePanel) Expanded() bool { return d.expanded != nil && *d.expanded }

// Size reports the panel's width in DIPs.
func (d *DraggablePanel) Size() float32 {
	if d.size == nil {
		return 0
	}
	return *d.size
}
