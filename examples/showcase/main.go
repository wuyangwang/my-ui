// Showcase lays my-ui's widgets out in one native window: one page per
// component, plus a page for the theming switches and subtree previews.
//
//	go run ./examples/showcase
//
// Every page lives in a file of its own and registers itself in
// pages.go; the window knows nothing about them.
package main

import (
	"log"
	"strconv"

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
	openFold   string // the fold the sidebar keeps open
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
}

// view builds one frame. Apply runs first, every frame, so the theme
// follows the state; everything below it takes what it sets.
func (s *showcase) view(c *ui.Context) {
	t := myui.Apply(c, myui.Options{Appearance: s.appearance, Palette: s.palette})

	// The frame's own background: behind every element MyGo paints the
	// window theme, which the forced appearance disagrees with, so the
	// root paints its own. A column lets the app bar own the top strip.
	ui.Column(c).Fill().Background(t.Background).Children(func() {
		s.appBar(c, t)
		ui.Row(c).Grow(1).AlignItems(ui.Stretch).Children(func() {
			s.sidebar(c, t)
			s.main(c, t)
		})
	})
	// The toasts of the window: one place for them, whatever page is
	// open.
	myui.NewToastViewport(c)
}

// appBar is the docs-style strip across the top of the window: the
// library's name, the component search, and the tools the showcase
// really has — a version and the appearance switch. Nothing in it
// pretends to be a product account.
func (s *showcase) appBar(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Children(func() {
		ui.Row(c).FillWidth().Height(t.Space(13)).PaddingX(t.Space(6)).
			Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
			// The name, with the same mark the sidebar shows.
			mark := ui.Box(c).Size(t.Space(6), t.Space(6)).Radius(t.Space(2)).
				Background(t.Accent)
			mark.Children(func() {
				ui.Icon(c, icons.Braces()).Size(t.Space(4), t.Space(4)).
					TextColor(t.AccentText)
			})
			ui.Text(c, "my-ui").FontSize(myui.TextSize(t, myui.TextLg)).Bold()
			ui.Text(c, "组件库").FontSize(myui.TextSize(t, myui.TextSm)).
				TextColor(t.TextMuted)

			// The component search, the one control a docs bar needs.
			ui.Box(c).Width(t.Space(6))
			myui.NewInput(c.Key("nav search"), &s.navQuery).Outline().Sm().
				Width(t.Space(64)).PrefixIcon(icons.Search()).
				Placeholder("搜索组件…").Label("搜索组件")

			ui.Box(c).Grow(1)

			// A version pill, then the appearance switch.
			version := ui.Box(c).Padding(t.Space(1), t.Space(2.5)).
				Radius(t.Space(2)).Background(t.Surface)
			version.Children(func() {
				ui.Text(c, "mygo v0.8.3").FontSize(myui.TextSize(t, myui.TextXs)).
					TextColor(t.TextMuted)
			})
			icon := icons.Moon()
			if t.Dark {
				icon = icons.Sun()
			}
			// A myui button sizes itself to content with AlignSelf(Start),
			// so an icon-only one in the centered bar needs that undone.
			toggle := myui.NewButton(c.Key("app theme"), "").
				Icon(icon).Label("切换明暗").Ghost().Sm()
			toggle.AlignSelf(ui.Center)
			if toggle.Clicked() {
				if t.Dark {
					s.appearance = myui.AppearanceLight
				} else {
					s.appearance = myui.AppearanceDark
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
	ui.Column(c).Width(248).Background(t.Surface).Padding(t.Space(6), t.Space(4)).
		Gap(t.Space(4)).Children(func() {
		// A label over the folds, so the sidebar reads as a tree of
		// components, with the count the filter leaves.
		ui.Text(c, "组件 · "+strconv.Itoa(len(pages))).
			FontSize(myui.TextSize(t, myui.TextSm)).Bold().
			TextColor(t.Text).Padding(t.Space(1), t.Space(2), 0, t.Space(2))

		ui.Scroll(c).Grow(1).Children(func() {
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

// fold is the title of a fold of pages: a button that opens the fold and
// closes the others.
func (s *showcase) fold(c *ui.Context, t *ui.Theme, name string, n int) {
	label := name + " (" + strconv.Itoa(n) + ")"
	b := myui.NewButton(c.Key("fold "+name), label).Ghost().Block()
	b.FontSize(myui.TextSize(t, myui.TextSm)).Bold()
	b.Justify(ui.Start)
	b.Padding(t.Space(3), t.Space(2)).Radius(t.Space(2))
	b.TextColor(t.TextMuted)
	if s.foldOpen(name) {
		b.Icon(icons.ChevronDown())
		b.TextColor(t.Text)
	} else {
		b.Icon(icons.ChevronRight())
	}
	if b.Clicked() {
		if s.openFold == name {
			s.openFold = ""
		} else {
			s.openFold = name
		}
	}
}

// foldOpen tells whether the pages of the fold g show: a search opens
// them all, and else the fold of the page that is open shows.
func (s *showcase) foldOpen(g string) bool {
	if s.navQuery != "" {
		return true
	}
	if s.openFold != "" {
		return s.openFold == g
	}
	return pageOf(s.page).group == g
}

// nav is one sidebar entry: the icon of the page and its name. The open
// page sits on a white pill of its own, with a bar of accent ink on its
// left edge.
func (s *showcase) nav(c *ui.Context, t *ui.Theme, p page) {
	active := s.page == p.key
	ui.Row(c).AlignItems(ui.Center).Gap(t.Space(0)).Children(func() {
		if active {
			ui.Box(c).Width(3).Height(t.Space(6)).Radius(2).Background(t.Accent)
		} else {
			ui.Box(c).Width(3).Height(t.Space(6))
		}
		b := myui.NewButton(c.Key("nav "+p.key), p.label).Ghost().Block()
		b.FontSize(myui.TextSize(t, myui.TextMd))
		b.Justify(ui.Start)
		b.Padding(t.Space(2.5), t.Space(2.5)).Radius(t.Space(2))
		if p.icon != nil {
			b.Icon(p.icon)
		}
		if active {
			b.Background(t.Background).TextColor(t.Text).Bold().
				Shadow(0, 1, 2, 0, t.Text.Alpha(0.10))
		} else {
			b.TextColor(t.TextMuted)
		}
		if b.Clicked() {
			s.page = p.key
		}
	})
}

// main is the page column beside the sidebar: the page title, then the
// page on a scroll.
func (s *showcase) main(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Grow(1).Padding(t.Space(7), t.Space(8)).Gap(t.Space(6)).
		Children(func() {
		// The name of the page, the group it is in, and a short line.
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

		ui.Scroll(c).Grow(1).Gap(t.Space(6)).Children(func() {
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
	s := &showcase{page: pages[0].key}
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
