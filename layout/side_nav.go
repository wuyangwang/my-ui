package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// SideNavProps are the slots of a SideNav rail.
type SideNavProps struct {
	// Avatar builds the account mark at the top of the rail.
	Avatar func()
	// Top builds the primary tabs below the avatar.
	TopActions func()
	// Bottom builds the help and settings actions at the foot.
	BottomActions func()
}

// SideNav is an app's vertical rail of icon buttons, 58 DIPs wide by
// default, with the actions pinned apart. It takes the height of its row,
// at least 640 DIPs, as the Lobe component does.
type SideNav struct {
	ui.Element
	c     *ui.Context
	width float32
}

// NewSideNav builds the rail. The width defaults to the theme's
// Space(15) (58 DIPs); Width overrides it.
func NewSideNav(c *ui.Context, p SideNavProps) *SideNav {
	t := c.Theme()
	s := &SideNav{c: c, width: t.Space(15)}
	col := ui.Column(c).Width(s.width).MinHeight(t.Space(160)).
		AlignItems(ui.Center).
		Background(t.Surface).
		BorderWidth(0, 1, 0, 0).BorderColor(t.Border).
		Padding(t.Space(2), 0).Gap(t.Space(2)).
		Transition(StateTransition())
	col.Children(func() {
		if p.Avatar != nil {
			p.Avatar()
		}
		if p.TopActions != nil {
			top := ui.Column(c).Gap(t.Space(1)).AlignItems(ui.Center)
			top.Children(p.TopActions)
		}
		ui.Box(c).Grow(1)
		if p.BottomActions != nil {
			foot := ui.Column(c).Gap(t.Space(1)).AlignItems(ui.Center)
			foot.Children(p.BottomActions)
		}
	})
	s.Element = col
	return s
}

// Width sets the rail's width in DIPs and returns it, for chaining.
func (s *SideNav) Width(v float32) *SideNav {
	s.width = v
	s.Element = s.Element.Width(v)
	return s
}
