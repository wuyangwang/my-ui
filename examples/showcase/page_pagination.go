package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/display"
)

// paginationPage demonstrates the Pagination component on its own page.
func (s *showcase) paginationPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "页码", "点击写回 *int，超出七页时显示省略号", func() {
		s.card(c, t, func() {
			if s.pageNum < 1 {
				s.pageNum = 1
			}
			display.NewPagination(c, &s.pageNum, 12).OnChanged(func(int) {})
		})
	})
}
