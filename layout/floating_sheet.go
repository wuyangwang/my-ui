package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// FloatingSheetProps configures NewFloatingSheet. Heights up to 1 are
// fractions of the container; larger values are DIPs.
type FloatingSheetProps struct {
	// Title is the heading of the sheet's drag handle.
	Title string
	// SnapPoints are the heights the sheet rests at after a drag, as
	// fractions of the container height (for example 0.3, 0.6, 0.9).
	SnapPoints []float32
	// MinHeight and RestingHeight bound the drag and set where it starts.
	// Zero means the theme's Space(30) and Space(50) respectively.
	MinHeight, RestingHeight float32
	// MaxHeight is the largest height; zero means the container height.
	MaxHeight float32
	// CloseThreshold, as a fraction of the lowest snap point, closes the
	// sheet when the drag goes below it. Zero disables drag-to-close.
	CloseThreshold float32
	// HeaderActions builds controls at the right of the drag header.
	HeaderActions func()
	// Inline keeps room for the sheet in the flow; dragged taller, it
	// covers the content above. Without it the sheet is absolutely
	// positioned at the bottom of its container.
	Inline bool
	// Embedded draws a border instead of a shadow.
	Embedded bool
	// NoDismiss leaves dragging down from closing the sheet.
	NoDismiss bool
}

// FloatingSheet is a sheet along the bottom of its container while the
// app's *bool is true. Drag the header to resize; the height snaps to the
// nearest SnapPoint when the pointer is let go. The window's height is
// used as the container height for the snap fractions: mygo's public API
// has no per-container size event, so the fraction is an approximation
// when the sheet sits in a smaller box.
type FloatingSheet struct {
	ui.Element
	open *bool
}

// NewFloatingSheet builds the sheet. build fills its scrollable content
// and is called inside the sheet's own Children. The open flag is the
// app's; a drag below the close threshold or the header's close button
// writes false to it.
func NewFloatingSheet(c *ui.Context, open *bool, p FloatingSheetProps, build func()) *FloatingSheet {
	s := &FloatingSheet{open: open}
	if open == nil || !*open {
		return s
	}
	t := c.Theme()
	minH := p.MinHeight
	if minH == 0 {
		minH = t.Space(30)
	}
	rest := p.RestingHeight
	if rest == 0 {
		rest = t.Space(50)
	}
	_, containerH := c.Size()
	maxH := p.MaxHeight
	if maxH == 0 || maxH > containerH {
		maxH = containerH
	}
	if rest < minH {
		rest = minH
	}
	if rest > maxH {
		rest = maxH
	}
	root := ui.Box(c).FillWidth()
	if p.Inline {
		root = root.Height(rest)
	} else {
		root = root.Height(rest).Absolute().Bottom(0).Start(0).End(0)
	}
	s.Element = root
	root.Children(func() {
		h := ui.Local(root, "floating sheet height", func() float32 { return rest })
		// The header is the drag handle. Dragged is the delta since the
		// last frame, so a downward drag shortens the sheet.
		head := ui.Row(c).FillWidth().AlignItems(ui.Center).
			Gap(t.Space(2)).Padding(t.Space(2), t.Space(4)).
			Cursor(ui.CursorResizeNS).
			Background(t.Surface).
			BorderWidth(0, 0, 1, 0).BorderColor(t.Border).
			Transition(StateTransition())
		dx, dy, held := head.Dragged()
		_ = dx
		if held && dy != 0 {
			v := *h - dy
			if v < minH {
				v = minH
			}
			if v > maxH {
				v = maxH
			}
			*h = v
		}
		if _, _, ok := head.DragEnded(); ok {
			if !p.NoDismiss && p.CloseThreshold > 0 && len(p.SnapPoints) > 0 &&
				*h < p.SnapPoints[0]*p.CloseThreshold*containerH {
				*open = false
			} else if len(p.SnapPoints) > 0 {
				*h = nearestSnap(*h, p.SnapPoints, containerH, minH, maxH)
			}
		}
		head.Children(func() {
			// The grabber tells the user the header is a handle.
			ui.Box(c).Width(t.Space(8)).Height(t.Space(0.5)).
				Radius(t.Space(1)).Background(t.Border).
				MarginEnd(t.Space(2))
			ui.Text(c, p.Title).FontSize(TextSize(t, TextSm)).Bold().Grow(1)
			if p.HeaderActions != nil {
				p.HeaderActions()
			} else {
				close := controls.NewIconAction(c, controls.IconActionOptions{
					SVG: icons.X(), Label: Message(c, "close"),
				})
				if close.Clicked() {
					*open = false
				}
			}
		})

		body := ui.Box(c).Grow(1).MinHeight(0)
		body.Children(func() {
			if build != nil {
				build()
			}
		})
		// Apply the remembered height to the root at the end of the
		// build: the children have been attached, so the size is the
		// one this frame's drag chose.
		s.Element = root.Height(*h)
		if p.Inline {
			s.Element = s.Element.Height(*h)
		} else {
			s.Element = s.Element.Height(*h).Absolute().Bottom(0).Start(0).End(0)
		}
		if p.Embedded {
			s.Element = s.Element.Border(1, t.Border)
		} else {
			s.Element = s.Element.
				Border(1, t.Border).
				Radius(Radius(t, RadiusXl)).
				Shadow(0, -8, 24, -8, t.Text.Alpha(0.18))
		}
		s.Element = s.Element.Background(t.Surface).Transition(StateTransition())
		OverlayEnter(c, s.Element, "floating sheet enter")
	})
	return s
}

// nearestSnap picks the snap point closest to h, in DIPs, clamped to the
// allowed range.
func nearestSnap(h float32, points []float32, containerH, minH, maxH float32) float32 {
	best, dist := h, float32(1e20)
	for _, p := range points {
		v := p * containerH
		if p > 1 {
			v = p
		}
		if v < minH {
			v = minH
		}
		if v > maxH {
			v = maxH
		}
		d := v - h
		if d < 0 {
			d = -d
		}
		if d < dist {
			best, dist = v, d
		}
	}
	return best
}

// Open reports whether the sheet shows.
func (s *FloatingSheet) Open() bool { return s.open != nil && *s.open }
