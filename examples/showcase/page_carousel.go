package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/display"
)

// carouselPage demonstrates the Carousel component on its own page.
func (s *showcase) carouselPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "轮播", "点击箭头或圆点写回当前索引", func() {
		s.card(c, t, func() {
			display.NewCarousel(c, &s.slide, 3, func(i int) {
				ui.Row(c).Height(t.Space(16)).AlignItems(ui.Center).Children(func() {
					ui.Text(c, []string{"第一张", "第二张", "第三张"}[i]).Bold()
				})
			})
		})
	})
}
