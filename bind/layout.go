package bind

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// Burger builds an open/closed menu button editing *open, configured by p.
func (u UI) Burger(open *bool, p layoutpkg.BurgerProps) *layoutpkg.Burger {
	return layoutpkg.NewBurger(u.c, open, p)
}

// FloatingSheet builds a sheet that floats over the window while *open is
// true, configured by p; build makes its content.
func (u UI) FloatingSheet(open *bool, p layoutpkg.FloatingSheetProps, build func()) *layoutpkg.FloatingSheet {
	return layoutpkg.NewFloatingSheet(u.c, open, p, build)
}

// DraggablePanel builds a panel whose *size and *expanded the user drags,
// configured by p; build makes its content.
func (u UI) DraggablePanel(size *float32, expanded *bool, p layoutpkg.DraggablePanelProps, build func()) *layoutpkg.DraggablePanel {
	return layoutpkg.NewDraggablePanel(u.c, size, expanded, p, build)
}

// DraggableSideNav builds a side navigation whose *width and *expanded the
// user drags, configured by p; body makes each item.
func (u UI) DraggableSideNav(width *float32, expanded *bool, p layoutpkg.DraggableSideNavProps, body func(nav layoutpkg.SideNavState)) *layoutpkg.DraggableSideNav {
	return layoutpkg.NewDraggableSideNav(u.c, width, expanded, p, body)
}

// FloatingPanel builds a panel that floats over the window while *open is
// true, configured by p; build makes its content.
func (u UI) FloatingPanel(open *bool, p layoutpkg.FloatingPanelProps, build func()) *layoutpkg.FloatingPanel {
	return layoutpkg.NewFloatingPanel(u.c, open, p, build)
}

// Grid builds a themed grid container whose children fill cells row by row.
func (u UI) Grid() ui.Element {
	return layoutpkg.NewGrid(u.c)
}

// GridPanel builds a themed grid with a card face chosen by variant.
func (u UI) GridPanel(variant myui.Variant) *layoutpkg.Grid {
	return layoutpkg.NewGridPanel(u.c, variant)
}

// Header builds a themed header bar configured by p.
func (u UI) Header(p layoutpkg.HeaderProps) *layoutpkg.Header {
	return layoutpkg.NewHeader(u.c, p)
}

// Layout builds a themed layout configured by p; build makes its content.
func (u UI) Layout(p layoutpkg.LayoutProps, build func()) *layoutpkg.Layout {
	return layoutpkg.NewLayout(u.c, p, build)
}

// Scroll builds a themed vertical scroll container.
func (u UI) Scroll() ui.Element {
	return layoutpkg.NewScroll(u.c)
}

// ScrollHorizontal builds a themed horizontal scroll container.
func (u UI) ScrollHorizontal() ui.Element {
	return layoutpkg.NewScrollHorizontal(u.c)
}

// ScrollPanel builds a themed scroll container over state.
func (u UI) ScrollPanel(state *ui.ScrollState) *layoutpkg.Scroll {
	return layoutpkg.NewScrollPanel(u.c, state)
}

// ScrollArea builds a themed scroll area over state; build makes its content.
func (u UI) ScrollArea(state *ui.ScrollState, build func()) *layoutpkg.ScrollArea {
	return layoutpkg.NewScrollArea(u.c, state, build)
}

// Sidebar builds a themed sidebar selecting one of its items into *selected.
func (u UI) Sidebar(selected *string, fn func()) ui.Element {
	return layoutpkg.NewSidebar(u.c, selected, fn)
}

// Split builds a resizable split with *size between first and second.
func (u UI) Split(size *float32, first, second func()) ui.Element {
	return layoutpkg.NewSplit(u.c, size, first, second)
}

// Steps builds a themed step indicator configured by p.
func (u UI) Steps(p layoutpkg.StepsProps) *layoutpkg.Steps {
	return layoutpkg.NewSteps(u.c, p)
}

// SideNav builds a themed side navigation configured by p.
func (u UI) SideNav(p layoutpkg.SideNavProps) *layoutpkg.SideNav {
	return layoutpkg.NewSideNav(u.c, p)
}
