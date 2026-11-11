package myui

import (
	"github.com/egoist/mygo/ui"
)

// Tab is one tab of a Tabs: its label and what it shows.
type Tab struct {
	Label   string
	Content func()
}

// Tabs are the tabs of a page: one of them is selected, and the app
// sees which.
//
//	myui.NewTabs(c, &selected, myui.Tab{Label: "概览", Content: func() {
//		ui.Text(c, "今天的情况")
//	}})
//
// The tabs are styled by the library, as shadcn's: a row of pills, the
// selected one on a surface of its own.
type Tabs struct {
	ui.Element
	c        *ui.Context
	selected *int
	tabs     []Tab
}

// NewTabs creates tabs, one of them selected. Only the content of the
// selected tab is built, as the others are not on screen.
func NewTabs(c *ui.Context, selected *int, tabs ...Tab) *Tabs {
	tb := &Tabs{c: c, selected: selected, tabs: tabs}
	t := c.Theme()
	parts := ui.TabsBase(c, selected, len(tabs))
	list := parts.List.Gap(t.Space(1)).Background(t.SurfaceHover).Radius(t.Radius).
		Padding(t.Space(1))
	list.Children(func() {
		for i, tab := range tabs {
			tr := parts.Tab(i).Padding(t.Space(1.5), t.Space(3)).Radius(t.Radius / 2)
			tr.TextColor(t.TextMuted)
			if i == *selected {
				tr.TextColor(t.Text).Background(t.Surface).
					Shadow(0, 1, 2, 0, t.TextMuted.Alpha(0.15))
			}
			// The label belongs to the tab: a click on it reaches the
			// tab.
			tr.Children(func() {
				ui.Text(c, tab.Label)
			})
		}
	})
	tb.Element = list
	// The content of the tab that is selected.
	if *selected >= 0 && *selected < len(tabs) {
		if content := tabs[*selected].Content; content != nil {
			ui.Column(c).Gap(t.Space(3)).Children(content)
		}
	}
	return tb
}

// OnChanged runs fn when another tab is selected in this frame, and
// returns the tabs, for chaining.
func (tb *Tabs) OnChanged(fn func(index int)) *Tabs {
	if fn != nil && tb.Changed() {
		fn(*tb.selected)
	}
	return tb
}
