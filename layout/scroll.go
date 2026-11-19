package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// ShadcnScrollbar is the thin scrollbar used by the scroll components in
// this package. Its metrics are fixed DIPs; every color is left zero, so
// mygo fills it from the theme of the frame and it follows appearance,
// accent and high contrast. It is a function, not a package variable, so
// nothing about a theme is cached.
func ShadcnScrollbar(t *ui.Theme) ui.ScrollbarStyle {
	_ = t // the theme's own colors are used when the fields are left zero
	return ui.ScrollbarStyle{ThumbWidth: 6, TrackWidth: 12, Radius: 3}
}

// LobeScrollbar is the old name of ShadcnScrollbar.
//
// Deprecated: use ShadcnScrollbar. The scrollbar is shadcn-styled and
// theme-derived; the old name is kept only for source compatibility.
func LobeScrollbar(t *ui.Theme) ui.ScrollbarStyle { return ShadcnScrollbar(t) }

// NewScroll builds a vertical scroll container with my-ui's thin bar. It
// keeps mygo's ui.Scroll signature, so a call written for the old
// re-export compiles. Children are the content; the bar is painted by
// mygo and follows the theme.
func NewScroll(c *ui.Context) ui.Element {
	return ui.Scroll(c).ScrollbarStyle(ShadcnScrollbar(c.Theme()))
}

// NewScrollHorizontal is NewScroll with mygo's horizontal axis.
func NewScrollHorizontal(c *ui.Context) ui.Element {
	return ui.ScrollHorizontal(c).ScrollbarStyle(ShadcnScrollbar(c.Theme()))
}

// Scroll is a scroll container kept as a struct, for the fluent form.
type Scroll struct {
	ui.Element
	c                *ui.Context
	state            *ui.ScrollState
	horizontal, both bool
	width, height    float32
}

// NewScrollPanel creates a Scroll. state, when non-nil, keeps the scroll
// position in the app's own state; TrackScroll writes it every frame.
func NewScrollPanel(c *ui.Context, state *ui.ScrollState) *Scroll {
	return &Scroll{c: c, state: state}
}

// Horizontal switches the scroll to its horizontal axis.
func (s *Scroll) Horizontal() *Scroll {
	s.horizontal, s.both = true, false
	return s
}

// Both lets the content scroll on either axis.
func (s *Scroll) Both() *Scroll {
	s.horizontal, s.both = false, true
	return s
}

// Width pins the scroll's width; without it the scroll fills its parent.
func (s *Scroll) Width(w float32) *Scroll {
	s.width = w
	return s
}

// Height pins the scroll's height.
func (s *Scroll) Height(h float32) *Scroll {
	s.height = h
	return s
}

// Children fills the scroll with what build draws.
func (s *Scroll) Children(build func()) *Scroll {
	t := s.c.Theme()
	var el ui.Element
	switch {
	case s.both:
		el = ui.ScrollBoth(s.c)
	case s.horizontal:
		el = ui.ScrollHorizontal(s.c)
	default:
		el = ui.Scroll(s.c)
	}
	el = el.ScrollbarStyle(ShadcnScrollbar(t))
	if s.state != nil {
		el = el.TrackScroll(s.state)
	}
	if s.width > 0 {
		el = el.Width(s.width)
	}
	if s.height > 0 {
		el = el.Height(s.height)
	}
	el.Children(build)
	s.Element = el
	return s
}

// ScrollArea is a scroll region whose edges fade where there is more to
// scroll to, as shadcn's ScrollArea does. It is built on NewScroll, so
// the thumb colors come from the theme and no timer drives the fade: the
// state is read on the frame the container is built. The fade strips use
// t.Background on that frame, so they follow the theme like the bar does.
type ScrollArea struct {
	ui.Element
	c                *ui.Context
	state            *ui.ScrollState
	fade             bool
	horizontal, both bool
}

// NewScrollArea creates a ScrollArea. state is the app's scroll position;
// a nil state is allowed, in which case the fade is only an approximation
// (mygo still paints its own bar). build fills the region.
func NewScrollArea(c *ui.Context, state *ui.ScrollState, build func()) *ScrollArea {
	s := &ScrollArea{c: c, state: state, fade: true}
	t := c.Theme()
	root := ui.Box(c).Clip().Radius(Radius(t, RadiusMd))
	root.Children(func() {
		var el ui.Element
		switch {
		case s.both:
			el = ui.ScrollBoth(c)
		case s.horizontal:
			el = ui.ScrollHorizontal(c)
		default:
			el = ui.Scroll(c)
		}
		el = el.ScrollbarStyle(ShadcnScrollbar(t)).Grow(1)
		if s.state != nil {
			el = el.TrackScroll(s.state)
		}
		el.Children(build)
		if s.fade {
			fadeEdges(c, s, t)
		}
	})
	s.Element = root
	return s
}

// fadeEdges paints soft strips at the edges that still have content. The
// strips are built after the scroll so they sit on top, and ignore input
// so a press on them reaches the content below.
func fadeEdges(c *ui.Context, s *ScrollArea, t *ui.Theme) {
	fade := t.Space(10)
	if s.state == nil {
		return
	}
	if !s.horizontal {
		if s.state.Y > 0 {
			ui.Box(c).Absolute().Top(0).Start(0).End(0).Height(fade).
				IgnoreInput(true).
				LinearGradient(ui.LinearGradient{
					From: t.Background, To: ui.Transparent, Angle: 180,
				})
		}
		if s.state.Y < s.state.MaxY {
			ui.Box(c).Absolute().Bottom(0).Start(0).End(0).Height(fade).
				IgnoreInput(true).
				LinearGradient(ui.LinearGradient{
					From: ui.Transparent, To: t.Background, Angle: 180,
				})
		}
	}
	if s.horizontal || s.both {
		if s.state.X > 0 {
			ui.Box(c).Absolute().Start(0).Top(0).Bottom(0).Width(fade).
				IgnoreInput(true).
				LinearGradient(ui.LinearGradient{From: t.Background, To: ui.Transparent, Angle: 90})
		}
		if s.state.X < s.state.MaxX {
			ui.Box(c).Absolute().End(0).Top(0).Bottom(0).Width(fade).
				IgnoreInput(true).
				LinearGradient(ui.LinearGradient{From: ui.Transparent, To: t.Background, Angle: 90})
		}
	}
}

// Horizontal makes the ScrollArea scroll sideways.
func (s *ScrollArea) Horizontal() *ScrollArea {
	s.horizontal = true
	return s
}

// Both lets the ScrollArea scroll on either axis.
func (s *ScrollArea) Both() *ScrollArea {
	s.both = true
	return s
}

// Fade shows or hides the soft edges.
func (s *ScrollArea) Fade(on bool) *ScrollArea {
	s.fade = on
	return s
}

// Width pins the region's width; without it the region fills its parent.
func (s *ScrollArea) Width(w float32) *ScrollArea {
	s.Element = s.Element.Width(w)
	return s
}

// Height pins the region's height.
func (s *ScrollArea) Height(h float32) *ScrollArea {
	s.Element = s.Element.Height(h)
	return s
}
