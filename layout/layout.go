package layout

import "github.com/egoist/mygo/ui"

// LayoutProps are the slots of a documentation page. Every slot is a
// builder the app supplies; the layout owns only the frame and the scroll
// regions, never the content's state.
type LayoutProps struct {
	// Header builds the top bar, typically with NewHeader.
	Header func()
	// Sidebar builds the navigation column.
	Sidebar func()
	// Toc builds the table of contents held at the right.
	Toc func()
	// TocWidth is the width of the Toc column; zero uses the theme's
	// Space(50) (200 DIPs).
	TocWidth float32
	// Footer builds the row at the foot of the content column.
	Footer func()
	// Mobile drops the sidebar and moves Toc above the content.
	Mobile bool
}

// Layout is a docs page: a header over a sidebar, the content and its
// table of contents. The content, Toc and Footer scroll together under the
// header; the sidebar scrolls on its own. All surfaces and hairlines come
// from c.Theme() on every frame.
type Layout struct {
	ui.Element
	c *ui.Context
}

// NewLayout builds the page. build draws the article; it runs inside the
// content column so its elements land there. Callers keep their own state
// (the active page, for instance) in the app or a ui.Local.
func NewLayout(c *ui.Context, p LayoutProps, build func()) *Layout {
	t := c.Theme()
	l := &Layout{c: c}
	root := ui.Column(c).Fill().Background(t.Background)
	root.Children(func() {
		if p.Header != nil {
			p.Header()
		}
		row := ui.Row(c).Grow(1).MinHeight(0).AlignItems(ui.Stretch)
		row.Children(func() {
			if !p.Mobile {
				side := ui.Box(c).Width(t.Space(72)).Clip().
					Background(t.Surface).
					BorderWidth(0, 1, 0, 0).BorderColor(t.Border)
				side.Children(func() {
					sc := ui.Scroll(c).ScrollbarStyle(LobeScrollbar(t)).Grow(1)
					sc.Children(func() {
						if p.Sidebar != nil {
							p.Sidebar()
						}
					})
				})
			}
			content := ui.Column(c).Grow(1).MinWidth(0)
			content.Children(func() {
				if p.Mobile && p.Toc != nil {
					ui.Box(c).FillWidth().Padding(t.Space(2), t.Space(4)).Children(p.Toc)
				}
				sc := ui.Scroll(c).ScrollbarStyle(LobeScrollbar(t)).Grow(1)
				sc.Children(func() {
					if build != nil {
						build()
					}
					if p.Footer != nil {
						foot := ui.Row(c).FillWidth().
							Margin(t.Space(8), 0, 0, 0).
							Padding(t.Space(4), 0, 0, 0).
							BorderWidth(1, 0, 0, 0).BorderColor(t.Border).
							Justify(ui.Center)
						foot.Children(p.Footer)
					}
				})
			})
			if !p.Mobile && p.Toc != nil {
				w := p.TocWidth
				if w == 0 {
					w = t.Space(50)
				}
				toc := ui.Box(c).Width(w).Clip().
					BorderWidth(0, 0, 0, 1).BorderColor(t.Border)
				toc.Children(func() {
					sc := ui.Scroll(c).ScrollbarStyle(LobeScrollbar(t)).Grow(1)
					sc.Children(p.Toc)
				})
			}
		})
	})
	l.Element = root
	return l
}
