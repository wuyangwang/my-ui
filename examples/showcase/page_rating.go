package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// ratingPage demonstrates the Rating component.
func (s *showcase) ratingPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "评分", "星数写回 *int，颜色随主题", func() {
		s.card(c, t, func() { ctrlpkg.NewRating(c, &s.rating, 5) })
	})
}
