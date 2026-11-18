// Showcase lays my-ui's widgets out in one native window: one page per
// component, plus a page for the theming switches and subtree previews.
//
//	go run ./examples/showcase
//
// Every page lives in a file of its own and registers itself in
// pages.go; the window knows nothing about them.
package main

import (
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	formspkg "github.com/wuyangwang/my-ui/forms"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
	"log"
	"strconv"
	"time"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// showcase is the state the window shows.
type showcase struct {
	page       string // the key of the open page
	appearance myui.Appearance
	palette    myui.PaletteID
	navQuery   string // what the sidebar's search holds
	openFold   string // the fold the user opened explicitly
	collapsed  string // the fold the user closed, overriding the page's group
	iconQuery  string // what the icon page's search holds
	drawerOpen bool

	name          string
	email         string
	query         string
	draft         string
	lang          string
	langs         []string
	size          string
	password      string
	mail          bool
	calendar      bool
	notifications bool
	sync          bool
	switchOff     bool
	volume        float64
	step          float64
	tab           int
	popoverOpen   bool
	modalOpen     bool
	disabled      string
	readonly      string
	lastEvent     string

	titleBarEvent   string // the last thing the title bar demo's buttons did
	pageNum         int    // the status page's pager
	crumb           int    // the more page's breadcrumb
	amount          float64
	ratio           float64
	searchQuery     string
	hoverOpen       bool
	menuOpen        bool
	contextOpen     bool
	lastAction      string
	tableState      ui.ListState
	autoValue       string
	pickedDate      time.Time
	slide           int
	tagClosed       bool
	cmdOpen         bool
	cmdQuery        string
	scroll          ui.ScrollState // the content column's scroll position
	lastPage        string         // the page the scroll position belongs to
	otpValue        string
	listState       ui.ListState
	todo            []string
	low, high       float64
	stepValue       float64
	rating          int
	pickedColor     ui.Color
	pickedTime      time.Time
	editableText    string
	accordionOpen   bool
	collapsibleOpen bool
	decimalValue    string
	splitAction     string
	toggleOn        bool
}

// theme returns the options the window's frame applies, every frame:
// the appearance and palette the theme page and the app bar switch.
func (s *showcase) theme() myui.Options {
	return myui.Options{Appearance: s.appearance, Palette: s.palette}
}

// setAppearance switches the view's theme and the native window
// appearance with it: myui.Apply alone leaves the system drawing a
// light frame (a white edge) around a dark view.
func (s *showcase) setAppearance(a myui.Appearance) {
	s.appearance = a
	switch a {
	case myui.AppearanceLight:
		mygo.Theme.SetSource(mygo.ThemeLight)
	case myui.AppearanceDark:
		mygo.Theme.SetSource(mygo.ThemeDark)
	default:
		mygo.Theme.SetSource(mygo.ThemeSystem)
	}
}

// view builds the content under the window's title bar: the app bar,
// then the sidebar beside the open page. Apply runs first, every
// frame, so the theme follows the state — window.New has applied the
// same options for the bar already, and the parts take what it sets.
func (s *showcase) view(c *ui.Context) {
	t := myui.Apply(c, s.theme())

	// The frame's own background: behind every element MyGo paints the
	// window theme, which the forced appearance disagrees with, so the
	// root paints its own. A column lets the app bar own the top strip.
	ui.Column(c).Fill().Background(t.Background).Children(func() {
		s.appBar(c, t)
		ui.Row(c).Grow(1).AlignItems(ui.Stretch).Children(func() {
			s.sidebar(c, t)
			// A hairline separates the nav from the page, as in
			// shadcn's inset sidebar layout.
			ui.Box(c).Width(1).Background(t.Border)
			s.main(c, t)
		})
	})
	// The toasts of the window: one place for them, whatever page is
	// open.
	overlaypkg.NewToastViewport(c)
}

// appBar is the docs-style strip across the top of the window: the
// library's name, the component search, and the tools the showcase
// really has — a version and the appearance switch. Nothing in it
// pretends to be a product account.
func (s *showcase) appBar(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Children(func() {
		ui.Row(c).FillWidth().Height(t.Space(14)).AlignItems(ui.Center).Children(func() {
			// The brand takes exactly the sidebar's width, so the
			// top bar and the nav below line up on the same edge.
			brand := ui.Row(c).Width(256).PaddingX(t.Space(5)).
				AlignItems(ui.Center).Gap(t.Space(2.5))
			brand.Children(func() {
				mark := ui.Row(c).AlignItems(ui.Center).Justify(ui.Center).
					Size(t.Space(6), t.Space(6)).Radius(myui.Radius(t, myui.RadiusMd)).
					Background(t.Accent)
				mark.Children(func() {
					ui.Icon(c, icons.Component()).Size(t.Space(4), t.Space(4)).
						TextColor(t.AccentText)
				})
				ui.Text(c, "my-ui").FontSize(myui.TextSize(t, myui.TextMd)).Bold()
				ui.Text(c, "组件库").FontSize(myui.TextSize(t, myui.TextXs)).
					TextColor(t.TextMuted)
			})

			// The component search, the one control a docs bar needs.
			formspkg.NewInput(c.Key("nav search"), &s.navQuery).Outline().Sm().
				Width(t.Space(70)).PrefixIcon(icons.Search()).
				Placeholder("搜索组件…").Label("搜索组件")

			ui.Box(c).Grow(1)

			// A version pill: the library release and the pinned framework.
			version := ui.Box(c).Padding(t.Space(1), t.Space(2.5)).
				Radius(t.Space(2)).Background(t.Surface)
			version.Children(func() {
				ui.Text(c, "my-ui v"+myui.Version+" · mygo v0.10.0").
					FontSize(myui.TextSize(t, myui.TextXs)).
					TextColor(t.TextMuted)
			})
			icon := icons.Moon()
			if t.Dark {
				icon = icons.Sun()
			}
			// A myui button sizes itself to content with AlignSelf(Start),
			// so an icon-only one in the centered bar needs that undone.
			toggle := ctrlpkg.NewButton(c.Key("app theme"), "").
				Icon(icon).Label("切换明暗").Ghost().Sm()
			toggle.AlignSelf(ui.Center)
			if toggle.Clicked() {
				if t.Dark {
					s.setAppearance(myui.AppearanceLight)
				} else {
					s.setAppearance(myui.AppearanceDark)
				}
			}
		})
		ui.Box(c).FillWidth().Height(1).Background(t.Border)
	})
}

// sidebar carries the navigation: one fold per kind, then the pages of
// the open fold. The filter that narrows them lives in the app bar.
//
// The folds are closed one at a time: the fold of the page that is open
// shows, and clicking a fold's title opens that fold and closes the
// other. A search opens them all and leaves only what it finds, so the
// sidebar holds a hundred pages as well as a dozen.
func (s *showcase) sidebar(c *ui.Context, t *ui.Theme) {
	// 16rem, the width shadcn's docs sidebar uses; the page gets the
	// rest. A quiet label heads the scroll region, like a SidebarGroup
	// label, instead of a row that looks like another nav entry.
	ui.Column(c).Width(256).Background(t.Background).
		Padding(t.Space(4), t.Space(3)).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "组件").FontSize(myui.TextSize(t, myui.TextXs)).Bold().
			TextColor(t.TextMuted).Padding(t.Space(1), t.Space(3))

		ui.Scroll(c).Grow(1).Gap(t.Space(2)).Children(func() {
			found := 0
			for _, g := range groups() {
				shown := filterPages(pagesOf(g), s.navQuery)
				if len(shown) == 0 {
					continue
				}
				found += len(shown)
				s.fold(c, t, g, len(shown))
				if s.foldOpen(g) {
					for _, p := range shown {
						s.nav(c, t, p)
					}
				}
			}
			if found == 0 {
				ui.Text(c, "没有匹配的组件").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted).
					Padding(t.Space(2))
			}
		})
	})
}

// fold is a group label that opens the group. It reads as a label,
// not as a big nav button: small muted text and a quiet chevron, as
// shadcn's SidebarGroupLabel and Element Plus's group title do.
func (s *showcase) fold(c *ui.Context, t *ui.Theme, name string, n int) {
	label := name + " (" + strconv.Itoa(n) + ")"
	open := s.foldOpen(name)
	b := ctrlpkg.NewButton(c.Key("fold "+name), label).Ghost().Block().Xs()
	b.FontSize(myui.TextSize(t, myui.TextXs))
	b.Justify(ui.Start)
	b.Padding(t.Space(1.5), t.Space(3)).Radius(myui.Radius(t, myui.RadiusSm))
	b.TextColor(t.TextMuted)
	if open {
		b.Icon(icons.ChevronDown())
	} else {
		b.Icon(icons.ChevronRight())
	}
	if b.Clicked() {
		if s.foldOpen(name) {
			// An explicit close must beat the fallback to the open
			// page's group, or the current fold can never collapse.
			s.openFold = ""
			s.collapsed = name
		} else {
			s.openFold = name
			s.collapsed = ""
		}
	}
}

// foldOpen tells whether the pages of the fold g show: a search opens
// them all, and else the fold of the page that is open shows.
func (s *showcase) foldOpen(g string) bool {
	if s.navQuery != "" {
		return true
	}
	if s.collapsed == g {
		return false
	}
	if s.openFold != "" {
		return s.openFold == g
	}
	return pageOf(s.page).group == g
}

// nav is one sidebar entry: a 40 DIP-tall row with the page's icon and
// label. The active entry takes a tint of the accent and the accent's
// ink, as Ant Design's selected menu item and shadcn's sidebar button
// do — no white pill, no left bar, no shadow.
func (s *showcase) nav(c *ui.Context, t *ui.Theme, p page) {
	active := s.page == p.key
	ui.Row(c).AlignItems(ui.Center).Padding(0, 0, 0, t.Space(2)).Children(func() {
		b := ctrlpkg.NewButton(c.Key("nav "+p.key), p.label).Ghost().Block()
		b.FontSize(myui.TextSize(t, myui.TextSm))
		b.Justify(ui.Start)
		b.Height(t.Space(9)).Padding(0, t.Space(3)).
			Radius(myui.Radius(t, myui.RadiusMd))
		if p.icon != nil {
			b.Icon(p.icon)
		}
		if active {
			b.Background(t.Accent.Alpha(0.10)).TextColor(t.Accent).Bold()
		} else {
			b.TextColor(t.TextMuted)
		}
		if b.Clicked() {
			s.page = p.key
			s.collapsed = "" // navigating to a page opens its group again
		}
	})
}

// main is the page column beside the sidebar: the page title, then the
// page on a scroll.
func (s *showcase) main(c *ui.Context, t *ui.Theme) {
	// The scroll belongs to the open page. Reset it before building the
	// new page, so the next component starts at the top rather than
	// inheriting the previous page's offset.
	if s.lastPage != s.page {
		s.scroll.Y = 0
		s.lastPage = s.page
	}
	ui.Column(c).Grow(1).Padding(t.Space(7), t.Space(8)).Gap(t.Space(6)).
		Children(func() {
			// Page name, group pill, then one line of context.
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
					ui.Text(c, pageLabel(s.page)).
						FontSize(myui.TextSize(t, myui.Text2Xl)).Bold()
					badge := ui.Box(c).Padding(t.Space(0.5), t.Space(2)).
						Radius(t.Space(2)).Background(t.Surface)
					badge.Children(func() {
						ui.Text(c, pageOf(s.page).group).
							FontSize(myui.TextSize(t, myui.TextXs)).
							TextColor(t.TextMuted)
					})
				})
				ui.Text(c, pageOf(s.page).desc).
					FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
			})
			ui.Box(c).FillWidth().Height(1).Background(t.Border)

			// The scroll's identity follows the page: switching to
			// another component builds a new scroll, starting at the
			// top, instead of reusing the old page's offset.
			ui.Scroll(c).TrackScroll(&s.scroll).Grow(1).
				Gap(t.Space(6)).Children(func() {
				s.openPage(c, t)()
			})
		})
}

// openPage hands back the builder of the open page.
func (s *showcase) openPage(c *ui.Context, t *ui.Theme) func() {
	p := pageOf(s.page)
	return func() { p.build(s, c, t) }
}
func main() {
	s := &showcase{page: pages[0].key, pickedDate: time.Now(),
		todo: []string{"设计规格", "实现组件", "补测试", "写文档"}}
	s.listState.Reorder = func(rows []int, to int) {
		// mygo hands the rows and the insertion point; move them in the
		// app's slice, like any other bound state.
		ids := append([]int(nil), rows...)
		if len(ids) == 0 {
			return
		}
		picked := make([]string, 0, len(ids))
		for _, i := range ids {
			if i >= 0 && i < len(s.todo) {
				picked = append(picked, s.todo[i])
			}
		}
		rest := []string{}
		for i := range s.todo {
			keep := false
			for _, j := range ids {
				if i == j {
					keep = true
				}
			}
			if !keep {
				rest = append(rest, s.todo[i])
			}
		}
		if to > len(rest) {
			to = len(rest)
		}
		if to < 0 {
			to = 0
		}
		next := append([]string(nil), rest[:to]...)
		next = append(next, picked...)
		next = append(next, rest[to:]...)
		s.todo = next
	}
	// The showcase speaks Chinese: the strings the controls hold fixed
	// follow the language of the app.
	mygo.App.SetLocale("zh-Hans")
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "my-ui 组件示例",
			Width:   1100,
			Height:  700,
			Content: ui.View(s.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
