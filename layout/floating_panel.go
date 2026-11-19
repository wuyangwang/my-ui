package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// FloatingPlacement is the corner of the window a FloatingPanel sits at.
type FloatingPlacement uint8

const (
	// FloatBottomRight is the lower right corner, the default.
	FloatBottomRight FloatingPlacement = iota
	// FloatBottomLeft is the lower left corner.
	FloatBottomLeft
	// FloatTopRight is the upper right corner.
	FloatTopRight
	// FloatTopLeft is the upper left corner.
	FloatTopLeft
)

// FloatingPanelProps configures NewFloatingPanel.
type FloatingPanelProps struct {
	// Title is the heading of the panel.
	Title string
	// Placement picks the corner; zero is FloatBottomRight.
	Placement FloatingPlacement
	// Width is the initial width in DIPs; zero means 360. The panel is
	// never wider than the window less 32 DIPs.
	Width float32
	// Offset is the distance from the corner in DIPs; zero uses the
	// theme's Space(4).
	Offset float32
	// MinWidth, MinHeight, MaxWidth and MaxHeight bound the free-edge
	// resize. Zero means no bound except the window and a small minimum.
	MinWidth, MinHeight, MaxWidth, MaxHeight float32
	// Actions builds controls at the right of the title row.
	Actions func()
	// Footer builds the action row at the foot of the panel.
	Footer func()
	// NoKeyboard leaves Escape engaged by default; set it to keep the
	// panel open on Escape.
	NoKeyboard bool
}

// FloatingPanel is a card floating at a corner of the window while the
// app's *bool is true. It is not modal: the page behind it stays usable.
// Escape and the close button write false to the flag; the app's trigger
// writes true, and the panel only reads the flag.
//
// The edges facing away from the corner resize the panel between the
// bounds in FloatingPanelProps. Their drag state lives on mygo elements;
// no timer or goroutine is involved.
type FloatingPanel struct {
	ui.Element
	open *bool
}

// NewFloatingPanel builds the panel over the window while open is true.
// A nil open builds nothing. build fills the panel's body; it is called
// inside the panel's own Children, so its elements land in the overlay
// layer and not in the window's corner.
func NewFloatingPanel(c *ui.Context, open *bool, p FloatingPanelProps, build func()) *FloatingPanel {
	f := &FloatingPanel{open: open}
	if open == nil || !*open {
		return f
	}
	t := c.Theme()
	// The size is remembered in a local of the panel's root. It is a
	// frame-stable value, not a cached color, and follows the drag.
	width := p.Width
	if width == 0 {
		width = t.Space(90)
	}
	offset := p.Offset
	if offset == 0 {
		offset = t.Space(4)
	}
	minW := p.MinWidth
	if minW == 0 {
		minW = t.Space(60)
	}
	minH := p.MinHeight
	if minH == 0 {
		minH = t.Space(40)
	}
	ui.Overlay(c, func() {
		root := ui.Box(c).Absolute().Left(0).Top(0).Right(0).Bottom(0)
		f.Element = root
		root.Children(func() {
			w := ui.Local(root, "floating panel width", func() float32 { return width })
			h := ui.Local(root, "floating panel height", func() float32 { return t.Space(80) })
			winW, _ := c.Size()
			maxW := p.MaxWidth
			if maxW == 0 || maxW > winW-32 {
				maxW = winW - 32
			}
			if maxW < minW {
				maxW = minW
			}
			if *w > maxW {
				*w = maxW
			}
			if *w < minW {
				*w = minW
			}
			if p.MaxHeight > 0 && *h > p.MaxHeight {
				*h = p.MaxHeight
			}
			if *h < minH {
				*h = minH
			}
			pos := ui.Box(c).Absolute().Width(*w).Height(*h)
			switch p.Placement {
			case FloatBottomLeft:
				pos.Left(offset).Bottom(offset)
			case FloatTopRight:
				pos.Right(offset).Top(offset)
			case FloatTopLeft:
				pos.Left(offset).Top(offset)
			default:
				pos.Right(offset).Bottom(offset)
			}
			panel := pos.Background(t.Surface).Border(1, t.Border).
				Radius(Radius(t, RadiusXl)).
				Shadow(0, 8, 30, -8, t.Text.Alpha(0.18)).
				Clip().Transition(StateTransition())
			OverlayEnter(c, panel, "floating panel enter")
			if !p.NoKeyboard && panel.OverlayShortcut(0, ui.KeyEscape) {
				*open = false
			}
			panel.Children(func() {
				// Header: title, actions and a close affordance.
				head := ui.Row(c).FillWidth().AlignItems(ui.Center).
					Gap(t.Space(2)).Padding(t.Space(4)).
					BorderWidth(0, 0, 1, 0).BorderColor(t.Border)
				head.Children(func() {
					ui.Text(c, p.Title).FontSize(TextSize(t, TextMd)).Bold().Grow(1)
					if p.Actions != nil {
						p.Actions()
					}
					close := controls.NewIconAction(c, controls.IconActionOptions{
						SVG: icons.X(), Label: Message(c, "close"),
					})
					if close.Clicked() {
						*open = false
					}
				})
				body := ui.Box(c).Grow(1).MinHeight(0)
				body.Children(func() {
					if build != nil {
						build()
					}
				})
				if p.Footer != nil {
					foot := ui.Row(c).FillWidth().AlignItems(ui.Center).
						Gap(t.Space(2)).Padding(t.Space(3), t.Space(4)).
						BorderWidth(1, 0, 0, 0).BorderColor(t.Border)
					foot.Children(p.Footer)
				}
				// Resize handles on the two edges facing away from the
				// corner. Only Dragged writes the remembered size.
				f.resizeHandle(c, p, w, h, minW, minH, true)
				f.resizeHandle(c, p, w, h, minW, minH, false)
			})
		})
	})
	return f
}

// resizeHandle builds one edge handle. horizontal picks the edge that
// changes the width (true) or the height (false). The edge facing the
// corner is not resizable, so a drag always grows the panel inward.
func (f *FloatingPanel) resizeHandle(c *ui.Context, p FloatingPanelProps, w, h *float32, minW, minH float32, horizontal bool) {
	t := c.Theme()
	grip := t.Space(1.5)
	var hnd ui.Element
	if horizontal {
		hnd = ui.Box(c).Absolute().Top(0).Bottom(0).Width(grip).Cursor(ui.CursorResizeEW)
		if p.Placement == FloatBottomRight || p.Placement == FloatTopRight {
			hnd = hnd.Left(0)
		} else {
			hnd = hnd.Right(0)
		}
		dx, _, held := hnd.Dragged()
		if held && dx != 0 {
			sign := float32(1)
			if p.Placement == FloatBottomRight || p.Placement == FloatTopRight {
				sign = -1
			}
			v := *w + sign*dx
			if v < minW {
				v = minW
			}
			if p.MaxWidth > 0 && v > p.MaxWidth {
				v = p.MaxWidth
			}
			*w = v
		}
		return
	}
	hnd = ui.Box(c).Absolute().Start(0).End(0).Height(grip).Cursor(ui.CursorResizeNS)
	if p.Placement == FloatBottomRight || p.Placement == FloatBottomLeft {
		hnd = hnd.Top(0)
	} else {
		hnd = hnd.Bottom(0)
	}
	_, dy, held := hnd.Dragged()
	if held && dy != 0 {
		sign := float32(1)
		if p.Placement == FloatBottomRight || p.Placement == FloatBottomLeft {
			sign = -1
		}
		v := *h + sign*dy
		if v < minH {
			v = minH
		}
		if p.MaxHeight > 0 && v > p.MaxHeight {
			v = p.MaxHeight
		}
		*h = v
	}
}

// Open reports whether the panel shows.
func (f *FloatingPanel) Open() bool { return f.open != nil && *f.open }
