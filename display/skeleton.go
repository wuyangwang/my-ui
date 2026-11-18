package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// Skeleton is a loading placeholder: the bars a page will fill in
// when its data arrives. It is a plain column of themed surfaces, so
// nothing flashes when the content swaps in.
type Skeleton struct {
	ui.Element
	c     *ui.Context
	lines int
	size  Size
	shape Shape
	width float32
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

// Circle turns the placeholder into one round avatar block instead of
// a column of lines.
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
			ui.Box(s.c).Size(d, d).Radius(d / 2).Background(t.Surface)
		})
		s.Element = col
		return
	}
	col.Children(func() {
		for i := 0; i < s.lines; i++ {
			bar := ui.Box(s.c).Height(h).Radius(h / 2).Background(t.Surface)
			if i == s.lines-1 && s.lines > 1 {
				bar = bar.Width(t.Space(24)) // the short last line
			} else {
				bar = bar.FillWidth()
			}
		}
	})
	s.Element = col
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
