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

// tabsVariant is the look of the strip: pills in a tray, an underline, or
// a dot under the chosen tab.
type tabsVariant uint8

const (
	tabsPills tabsVariant = iota
	tabsLine
	tabsPoint
)

// tabsMatch is the shared identity of the selected indicator. It rides on
// the app's *int, which is stable across frames, so the framework matches
// the one indicator built this frame with the one built last frame — even
// in another tab — and slides it across.
type tabsMatch struct{ selected *int }

// Tabs are the tabs of a page: one of them is selected, and the app sees
// which. The selected indicator is its own element with an
// ElementTransition that slides it from tab to tab; hover changes the
// label's ink. Only the selected tab's content is built.
//
//	display.NewTabs(c, &selected, display.Tab{Label: "概览", Content: func() {
//		ui.Text(c, "今天的情况")
//	}}).Line().Size(myui.SizeSm)
type Tabs struct {
	ui.Element
	c         *ui.Context
	selected  *int
	tabs      []Tab
	variant   tabsVariant
	size      Size
	vertical  bool
	root      ui.Element
	panel     ui.Element
	list      ui.Element
	triggers  []ui.Element
	labels    []ui.Element
	indicator ui.Element
}

// NewTabs creates tabs, one of them selected. Only the content of the
// selected tab is built, as the others are not on screen.
func NewTabs(c *ui.Context, selected *int, tabs ...Tab) *Tabs {
	tb := &Tabs{c: c, selected: selected, tabs: tabs}
	root := ui.Column(c).AlignSelf(ui.Stretch)
	tb.root = root
	root.Children(func() {
		tb.buildList(c, selected)
		// The content of the tab that is selected.
		if *selected >= 0 && *selected < len(tabs) {
			if content := tabs[*selected].Content; content != nil {
				tb.panel = ui.Column(c).Gap(c.Theme().Space(3))
				tb.panel.Children(content)
			}
		}
	})
	tb.Element = root
	tb.apply()
	return tb
}

// buildList builds the strip and its triggers into the root.
func (tb *Tabs) buildList(c *ui.Context, selected *int) {
	parts := ui.TabsBase(c, selected, len(tb.tabs))
	list := parts.List.Gap(c.Theme().Space(1))
	tb.list = list
	list.Children(func() {
		tb.triggers = make([]ui.Element, len(tb.tabs))
		tb.labels = make([]ui.Element, len(tb.tabs))
		slide := ui.ElementTransition{
			Duration: MotionNormal,
			Ease:     ui.EaseOut,
			Match:    tabsMatch{selected},
		}
		for i := range tb.tabs {
			tab := parts.Tab(i).Transition(StateTransition())
			tb.triggers[i] = tab
			tab.Children(func() {
				if i == *selected {
					// One indicator, built in whichever tab is selected.
					// Its Match links it across frames so it slides.
					ind := ui.Box(c).Absolute().ZIndex(-1).IgnoreInput(true).
						Transition(slide)
					tb.indicator = ind
				}
				if tb.tabs[i].Icon != nil {
					ui.Icon(c, tb.tabs[i].Icon).Size(IconSizeSm, IconSizeSm)
				}
				tb.labels[i] = ui.Text(c, tb.tabs[i].Label).SingleLine()
			})
			if tb.tabs[i].Disabled {
				tab.Disabled(true)
			}
		}
	})
}

// apply paints the strip, the triggers and the indicator for the current
// variant, size and direction. The hovered tab's label takes the full ink.
func (tb *Tabs) apply() {
	t := tb.c.Theme()
	if tb.vertical {
		tb.root = tb.root.Row().AlignItems(ui.Start)
		tb.list = tb.list.Column().AlignItems(ui.Stretch)
		if tb.panel.Valid() {
			tb.panel = tb.panel.Grow(1).MinWidth(0).Padding(0, 0, 0, t.Space(3))
		}
	} else {
		tb.root = tb.root.Column()
		tb.list = tb.list.Row().AlignItems(ui.Center).AlignSelf(ui.Start)
		if tb.panel.Valid() {
			tb.panel = tb.panel.Grow(1).MinWidth(0).Padding(t.Space(3), 0, 0, 0)
		}
	}

	height, padX, font := tb.metrics(t)
	switch tb.variant {
	case tabsLine:
		tb.list.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0, 0)
		if tb.vertical {
			tb.list = tb.list.Gap(t.Space(1))
		} else {
			tb.list = tb.list.Gap(t.Space(4))
		}
	default: // tabsPills and tabsPoint share the tray
		tb.list.Background(t.SurfaceHover).Border(0, ui.Transparent).
			Radius(Radius(t, RadiusMd)).Padding(t.Space(1))
	}

	for i, tr := range tb.triggers {
		disabled := i < len(tb.tabs) && tb.tabs[i].Disabled
		chosen := i == *tb.selected
		switch tb.variant {
		case tabsLine:
			tr.PaddingY(t.Space(2)).PaddingX(padX).Radius(0).BorderWidth(0, 0, 0, 0).
				Gap(t.Space(1.5)).FontSize(font)
		case tabsPoint:
			tr.Padding(t.Space(2), padX, t.Space(3.5), padX).Radius(Radius(t, RadiusSm)).
				BorderWidth(0, 0, 0, 0).Gap(t.Space(1.5)).FontSize(font)
		default:
			tr.Height(height).PaddingX(padX).Radius(Radius(t, RadiusSm)).
				BorderWidth(0, 0, 0, 0).Gap(t.Space(1.5)).FontSize(font)
		}
		// Hover changes the label's ink; the frame rebuild reads the
		// current hover state, and StateTransition fades the change.
		ink := t.TextMuted
		switch {
		case disabled:
			ink = t.TextMuted.Alpha(0.5)
		case chosen:
			ink = t.Text
		case tr.Hovered():
			ink = t.Text
		}
		tr.TextColor(ink)
		if tb.labels[i].Valid() {
			tb.labels[i] = tb.labels[i].TextColor(ink)
		}
	}

	if !tb.indicator.Valid() {
		return
	}
	ind := tb.indicator
	switch tb.variant {
	case tabsLine:
		ind.Background(t.Accent)
		if tb.vertical {
			ind = ind.Top(0).Bottom(0).Right(0).Width(t.Space(0.5)).Radius(0)
		} else {
			ind = ind.Left(0).Right(0).Bottom(0).Height(t.Space(0.5)).Radius(0)
		}
	case tabsPoint:
		d := t.Space(1.25)
		ind = ind.Size(d, d).Radius(d/2).Background(t.Accent).
			LeftPercent(50).Margin(0, 0, 0, -d/2).Bottom(t.Space(1))
	default:
		ind = ind.Top(0).Left(0).Right(0).Bottom(0).
			Radius(Radius(t, RadiusSm)).Background(t.Surface).
			Shadow(0, 1, 2, 0, t.TextMuted.Alpha(0.15))
	}
	tb.indicator = ind
}

func (tb *Tabs) metrics(t *ui.Theme) (height, padX, font float32) {
	height, padX, font = 32, t.Space(3), TextSize(t, TextSm)
	switch tb.size {
	case SizeXs:
		height, padX, font = 26, t.Space(2), TextSize(t, TextXs)
	case SizeLg:
		height, padX, font = 36, t.Space(4), TextSize(t, TextMd)
	}
	return
}

// Line switches the tabs to an underline style and returns them, for
// chaining.
func (tb *Tabs) Line() *Tabs {
	tb.variant = tabsLine
	tb.apply()
	return tb
}

// Point switches the tabs to a dot under the chosen tab and returns them.
func (tb *Tabs) Point() *Tabs {
	tb.variant = tabsPoint
	tb.apply()
	return tb
}

// Vertical stacks the tabs in a column beside their panel and returns
// them, for chaining.
func (tb *Tabs) Vertical() *Tabs {
	tb.vertical = true
	tb.apply()
	return tb
}

// Size sets the tab height, padding and font size.
func (tb *Tabs) Size(s Size) *Tabs {
	tb.size = s
	tb.apply()
	return tb
}

// OnChanged runs fn when another tab is selected in this frame, and
// returns the tabs, for chaining.
func (tb *Tabs) OnChanged(fn func(index int)) *Tabs {
	if fn != nil && tb.list.Valid() && tb.list.Changed() {
		fn(*tb.selected)
	}
	return tb
}
