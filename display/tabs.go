package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// Tab is one tab of a Tabs: its label, an optional icon, whether it can be
// chosen, and what it shows.
type Tab struct {
	Label    string
	Icon     *ui.SVG
	Disabled bool
	Content  func()
}

// tabsVariant is the look of the strip: pills in a tray, or an
// underline.
type tabsVariant uint8

const (
	tabsPills tabsVariant = iota
	tabsLine
)

// Tabs are the tabs of a page: one of them is selected, and the app
// sees which.
//
//	display.NewTabs(c, &selected, display.Tab{Label: "概览", Content: func() {
//		ui.Text(c, "今天的情况")
//	}})
//
// The tabs are styled by the library, as shadcn's: a row of pills, the
// selected one on a surface of its own. Line() switches to shadcn's
// underline style.
type Tabs struct {
	ui.Element
	c        *ui.Context
	selected *int
	tabs     []Tab
	variant  tabsVariant
	list     ui.Element
	triggers []ui.Element
}

// NewTabs creates tabs, one of them selected. Only the content of the
// selected tab is built, as the others are not on screen.
func NewTabs(c *ui.Context, selected *int, tabs ...Tab) *Tabs {
	tb := &Tabs{c: c, selected: selected, tabs: tabs}
	parts := ui.TabsBase(c, selected, len(tabs))
	tb.list = parts.List.Gap(c.Theme().Space(1))
	tb.list.Children(func() {
		tb.triggers = make([]ui.Element, len(tabs))
		for i := range tabs {
			tb.triggers[i] = parts.Tab(i).Transition(StateTransition())
			tab := tabs[i]
			tb.triggers[i].Children(func() {
				if tab.Icon != nil {
					ui.Icon(c, tab.Icon).Size(c.Theme().Space(4), c.Theme().Space(4)).
						Margin(0, c.Theme().Space(1.5), 0, 0)
				}
				ui.Text(c, tab.Label)
			})
			if tab.Disabled {
				tb.triggers[i].Disabled(true)
			}
		}
	})
	tb.Element = tb.list
	tb.apply()
	// The content of the tab that is selected.
	if *selected >= 0 && *selected < len(tabs) {
		if content := tabs[*selected].Content; content != nil {
			ui.Column(c).Gap(c.Theme().Space(3)).Children(content)
		}
	}
	return tb
}

// apply paints the strip and its triggers for the current variant.
func (tb *Tabs) apply() {
	t := tb.c.Theme()
	if tb.variant == tabsLine {
		tb.list.Background(ui.Transparent).Border(0, ui.Transparent).Padding(t.Space(1), 0)
	} else {
		tb.list.Background(t.SurfaceHover).Border(0, ui.Transparent).Radius(Radius(t, RadiusMd)).Padding(t.Space(1))
	}
	for i, tr := range tb.triggers {
		if tb.variant == tabsLine {
			tr.Padding(t.Space(2), t.Space(3)).Radius(0).BorderWidth(0, 0, 2, 0)
		} else {
			tr.Padding(t.Space(1.5), t.Space(3)).Radius(Radius(t, RadiusSm)).BorderWidth(0, 0, 0, 0)
		}
		chosen := i == *tb.selected
		disabled := i < len(tb.tabs) && tb.tabs[i].Disabled
		if chosen {
			tr.TextColor(t.Text)
			if tb.variant == tabsLine {
				tr.BorderColor(t.Accent).Background(ui.Transparent).Shadow(0, 0, 0, 0, ui.Transparent)
			} else {
				tr.BorderColor(ui.Transparent).Background(t.Surface).
					Shadow(0, 1, 2, 0, t.TextMuted.Alpha(0.15))
			}
		} else {
			tr.Background(ui.Transparent).Shadow(0, 0, 0, 0, ui.Transparent)
			if disabled {
				tr.TextColor(t.TextMuted.Alpha(0.5))
			} else {
				tr.TextColor(t.TextMuted).BorderColor(ui.Transparent)
			}
		}
	}
}

// Line switches the tabs to shadcn's underline style and returns them,
// for chaining.
func (tb *Tabs) Line() *Tabs {
	tb.variant = tabsLine
	tb.apply()
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
