package display

import (
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// SkeletonAnimation is how a placeholder shows it is loading.
type SkeletonAnimation uint8

const (
	// SkeletonFade pulses the bars' opacity (the zero value).
	SkeletonFade SkeletonAnimation = iota
	// SkeletonSweep moves a highlight across the bars.
	SkeletonSweep
	// SkeletonStatic keeps the bars still.
	SkeletonStatic
)

// Skeleton is a loading placeholder: the bars a page will fill in when
// its data arrives. It is a column of themed surfaces, so nothing flashes
// when the content swaps in. The last of several lines is 66% of the
// container's width, following the container instead of a fixed DIP
// length. Fade and Sweep use Element.Loop, so the desktop's Reduce Motion
// comes free — there are no timers in the component.
type Skeleton struct {
	ui.Element
	c         *ui.Context
	lines     int
	size      Size
	shape     Shape
	width     float32
	animation SkeletonAnimation
}

// NewSkeleton builds lines placeholder bars; a number below 1 means
// three lines.
func NewSkeleton(c *ui.Context, lines int) *Skeleton {
	if lines < 1 {
		lines = 3
	}
	s := &Skeleton{c: c, lines: lines}
	s.build()
	return s
}

// Circle turns the placeholder into one round avatar block instead of a
// column of lines.
func (s *Skeleton) Circle() *Skeleton {
	s.shape = ShapeCircle
	s.build()
	return s
}

// Width pins the placeholder's width in DIPs; zero fills its parent.
func (s *Skeleton) Width(w float32) *Skeleton {
	s.width = w
	s.build()
	return s
}

// Size picks the height of one bar.
func (s *Skeleton) Size(sz Size) *Skeleton {
	s.size = sz
	s.build()
	return s
}

// Animation picks how the placeholder shows loading: Fade pulses its
// opacity, Sweep moves a highlight across it, Static keeps it still.
func (s *Skeleton) Animation(a SkeletonAnimation) *Skeleton {
	s.animation = a
	s.build()
	return s
}

// Fade sets the pulse animation and returns the skeleton.
func (s *Skeleton) Fade() *Skeleton { return s.Animation(SkeletonFade) }

// Sweep sets the highlight animation and returns the skeleton.
func (s *Skeleton) Sweep() *Skeleton { return s.Animation(SkeletonSweep) }

// Static keeps the placeholder still and returns the skeleton.
func (s *Skeleton) Static() *Skeleton { return s.Animation(SkeletonStatic) }

func (s *Skeleton) build() {
	t := s.c.Theme()
	h := s.barHeight()
	col := ui.Column(s.c).Gap(t.Space(2))
	if s.width > 0 {
		col = col.Width(s.width)
	} else {
		col = col.FillWidth()
	}
	if s.shape == ShapeCircle {
		d := h * 3
		col.Children(func() {
			s.animate(ui.Box(s.c).Size(d, d).Radius(d/2).Background(t.Surface), t)
		})
		s.Element = col
		return
	}
	col.Children(func() {
		for i := 0; i < s.lines; i++ {
			bar := ui.Box(s.c).Height(h).Radius(h / 2).Background(t.Surface)
			if i == s.lines-1 && s.lines > 1 {
				// The short last line follows the container: 66% of
				// its width, not a fixed number of DIPs.
				bar = bar.WidthPercent(66)
			} else {
				bar = bar.FillWidth()
			}
			s.animate(bar, t)
		}
	})
	s.Element = col
}

// animate applies the chosen animation to one bar. It builds the bar's
// children only when needed, so a Static placeholder stays plain.
func (s *Skeleton) animate(bar ui.Element, t *ui.Theme) {
	switch s.animation {
	case SkeletonFade:
		// A bounce easing makes the opacity go down and back up in one
		// period; Loop keeps frames coming without a timer.
		bar.Opacity(1 - 0.5*bar.Loop("skeleton-fade", 1600*time.Millisecond, ui.Bounce(ui.EaseInOut)))
	case SkeletonSweep:
		at := bar.Loop("skeleton-sweep", 1400*time.Millisecond, ui.EaseInOut)
		bar = bar.Clip()
		bar.Children(func() {
			ui.Box(s.c).Absolute().Fill().TranslatePercent(200*at-100, 0).
				LinearGradient(ui.LinearGradient{
					Angle: 90,
					Stops: []ui.GradientStop{
						{Color: ui.Transparent, At: 0},
						{Color: t.SurfaceHover, At: 0.5},
						{Color: ui.Transparent, At: 1},
					},
				})
		})
	}
}

func (s *Skeleton) barHeight() float32 {
	switch s.size {
	case SizeXs:
		return 10
	case SizeSm:
		return 12
	case SizeLg:
		return 20
	default:
		return 16
	}
}
