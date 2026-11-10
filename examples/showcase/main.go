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

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// showcase is the state the window shows.
type showcase struct {
	page       string // the key of the open page
	appearance myui.Appearance
	palette    myui.PaletteID

	name        string
	email       string
	query       string
	draft       string
	lang        string
	langs       []string
	size        string
	password    string
	popoverOpen bool
	modalOpen   bool
	disabled    string
	readonly    string
	lastEvent   string
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
}

// sidebar carries the navigation: the brand and the pages, one group
// per kind, on the surface tone so the structure reads.
func (s *showcase) sidebar(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(190).Background(t.Surface).Padding(t.Space(4)).
		Gap(t.Space(4)).Children(func() {
		ui.Text(c, "my-ui").FontSize(myui.TextSize(t, myui.TextLg)).Bold()
		ui.Text(c, "MyGo 原生组件库").
			FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)

		for _, g := range groups() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, g).FontSize(myui.TextSize(t, myui.TextXs)).Bold().TextColor(t.TextMuted)
				for _, p := range pages {
					if p.group == g {
						s.nav(c, t, p)
					}
				}
			})
		}
	})
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
