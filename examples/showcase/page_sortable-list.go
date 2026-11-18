package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/display"
)

// sortableListPage demonstrates a list whose rows can be dragged into a
// new order; the app owns the order in its slice.
func (s *showcase) sortableListPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "拖拽排序", "ListState.Reorder 由 app 提供，库只在每行加一个手柄", func() {
		s.card(c, t, func() {
			display.NewSortableList(c, &s.listState, len(s.todo), func(i int) {
				ui.Text(c, s.todo[i])
			})
		})
	})
}
