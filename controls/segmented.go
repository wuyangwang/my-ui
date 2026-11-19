package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// SegmentedVariant picks the track a Segmented sits on.
type SegmentedVariant uint8

const (
	// SegmentedFilled puts the segments on a muted surface.
	SegmentedFilled SegmentedVariant = iota
	// SegmentedOutlined leaves the window's background showing.
	SegmentedOutlined
)

// Segmented chooses one of labels into a bound index. It is my-ui's
// look over mygo's SegmentedBase: the keyboard, the radio-group role and
// the Changed event are the framework's, the track and the sliding
// indicator are the library's. The indicator moves through a framework
// transition whose duration comes from the motion tokens.
type Segmented struct {
	ui.Element
	c        *ui.Context
	selected *int
	labels   []string
	variant  SegmentedVariant
	size     Size
	block    bool
}

// NewSegmented builds a segmented control bound to *selected. The
// signature matches mygo's Segmented, with my-ui's look; call Label on
// the returned element to name the group.
func NewSegmented(c *ui.Context, selected *int, labels ...string) *Segmented {
	t := c.Theme()
	s := &Segmented{c: c, selected: selected, labels: labels}
	parts := ui.SegmentedBase(c, selected, len(labels))
	track := parts.Track
	track.AlignItems(ui.Stretch).Padding(t.Space(0.75)).Gap(t.Space(0.5)).
		Radius(Radius(t, RadiusMd)).Transition(StateTransition())
	indicator := t.Surface
	if s.variant == SegmentedOutlined {
		track.Border(1, t.Border)
		indicator = t.SurfaceHover
	} else {
		track.Background(t.Surface).Border(1, t.Border)
	}
	if s.block {
		track.AlignSelf(ui.Stretch)
	}
	// The Local pointer gives this control's indicator an identity of its
	// own, so its slide never borrows a position from another control.
	id := ui.Local(track, "segmented-indicator", func() *int { return new(int) })
	track.Children(func() {
		for i, label := range labels {
			height, padding, font := segmentedMetrics(t, s.size)
			seg := parts.Segment(i).Height(height).PaddingX(padding).
				Radius(Radius(t, RadiusSm)).FontSize(font).FontWeight(500).
				Cursor(ui.CursorPointer).Unselectable().Transition(StateTransition())
			if s.block {
				seg.Grow(1).Basis(0)
			}
			ink := t.TextMuted
			if i == *selected || seg.Hovered() {
				ink = t.Text
			}
			seg.TextColor(ink)
			seg.Children(func() {
				if i == *selected {
					ui.Box(c).Absolute().Top(0).Left(0).Right(0).Bottom(0).
						Radius(Radius(t, RadiusSm)).Background(indicator).ZIndex(-1).
						Transition(ui.ElementTransition{
							Duration: MotionNormal, Position: true, Colors: true, Match: id,
						})
				}
				ui.Text(c, label).SingleLine()
			})
		}
	})
	s.Element = track
	return s
}

// Variant applies the track look of v and returns the control, for
// chaining.
func (s *Segmented) Variant(v SegmentedVariant) *Segmented {
	s.variant = v
	return s
}

// Filled is Variant(SegmentedFilled).
func (s *Segmented) Filled() *Segmented { return s.Variant(SegmentedFilled) }

// Outlined is Variant(SegmentedOutlined).
func (s *Segmented) Outlined() *Segmented { return s.Variant(SegmentedOutlined) }

// Size applies the shared size s and returns the control, for chaining.
func (s *Segmented) Size(sz Size) *Segmented {
	s.size = sz
	return s
}

// Block stretches the track across its parent and shares the width
// between the segments, and returns the control, for chaining.
func (s *Segmented) Block(on bool) *Segmented {
	if on {
		s.Element = s.Element.AlignSelf(ui.Stretch)
	}
	return s
}

// Xs sizes the segments down to the smallest shared size.
func (s *Segmented) Xs() *Segmented { return s.Size(SizeXs) }

// Sm sizes the segments down.
func (s *Segmented) Sm() *Segmented { return s.Size(SizeSm) }

// Md gives the built-in sizing.
func (s *Segmented) Md() *Segmented { return s.Size(SizeMd) }

// Lg sizes the segments up.
func (s *Segmented) Lg() *Segmented { return s.Size(SizeLg) }

// Label names the group for assistive technology and returns the
// control, for chaining.
func (s *Segmented) Label(name string) *Segmented {
	s.Element = s.Element.Label(name)
	return s
}

// Changed reports whether a segment was chosen in this frame.
func (s *Segmented) Changed() bool { return s.Element.Changed() }

// segmentedMetrics is the height, horizontal inset and font of a
// segment on a shared size.
func segmentedMetrics(t *ui.Theme, size Size) (height, padding, font float32) {
	switch size {
	case SizeXs:
		return 26, t.Space(2), TextSize(t, TextXs)
	case SizeSm:
		return 30, t.Space(2.5), TextSize(t, TextSm)
	case SizeLg:
		return 40, t.Space(4), TextSize(t, TextLg)
	default:
		return 36, t.Space(3), TextSize(t, TextMd)
	}
}
