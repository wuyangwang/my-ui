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
	if t == nil {
		t = c.Theme()
	}

	// Draw the frame's own background with the applied theme: behind
	// every element MyGo paints the window theme, which the forced
	// appearance disagrees with, so the root paints its own.
	ui.Row(c).Fill().Background(t.Background).AlignItems(ui.Stretch).Children(func() {
		s.sidebar(c, t)
		s.main(c, t)
	})
	// The toasts of the window: one place for them, whatever page is
	// open.
	myui.NewToastViewport(c)
}

// sidebar carries the navigation: the brand, a search and the pages, one
// fold per kind, on the surface tone so the structure reads.
//
// The folds are closed one at a time: the fold of the page that is open
// shows, and clicking a fold's title opens that fold and closes the
// other. A search opens them all and leaves only what it finds, so the
// sidebar holds a hundred pages as well as a dozen.
func (s *showcase) sidebar(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(190).Background(t.Surface).Padding(t.Space(4)).
		Gap(t.Space(3)).Children(func() {
		ui.Text(c, "my-ui").FontSize(myui.TextSize(t, myui.TextLg)).Bold()
		ui.Text(c, "MyGo 原生组件库").
			FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)

		myui.NewInput(c.Key("nav search"), &s.navQuery).Outline().Sm().
			PrefixIcon(icons.Search()).
			Placeholder("搜索组件…").
			Label("搜索组件")

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
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			}
		})
	})
}

// fold is the title of a fold of pages: a button that opens the fold and
// closes the others.
func (s *showcase) fold(c *ui.Context, t *ui.Theme, name string, n int) {
	label := name + " (" + strconv.Itoa(n) + ")"
	b := myui.NewButton(c.Key("fold "+name), label).Ghost().Block()
	if s.foldOpen(name) {
		b.Icon(icons.ChevronDown())
		b.Background(t.Accent.Alpha(0.10)).TextColor(t.Accent)
	} else {
		b.Icon(icons.ChevronRight())
		b.TextColor(t.TextMuted)
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

// nav is one sidebar entry; the open page paints the accent ink.
func (s *showcase) nav(c *ui.Context, t *ui.Theme, p page) {
	b := myui.NewButton(c.Key("nav "+p.key), p.label).Ghost()
	if s.page == p.key {
		a := t.Accent
		b.Background(a.Alpha(0.10)).TextColor(a)
	}
	if b.Clicked() {
		s.page = p.key
	}
}

// main is the page column beside the sidebar: the page title, then the
// page on a scroll.
func (s *showcase) main(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Grow(1).Padding(t.Space(4)).Gap(t.Space(4)).Children(func() {
		ui.Text(c, pageLabel(s.page)).
			FontSize(myui.TextSize(t, myui.TextXl)).Bold()

		ui.Scroll(c).Grow(1).Gap(t.Space(4)).Children(func() {
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
			Width:   960,
			Height:  620,
			Content: ui.View(s.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
