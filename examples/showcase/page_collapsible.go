package main

import (
	"github.com/egoist/mygo/ui"
	display "github.com/wuyangwang/my-ui/display"
)

// collapsiblePage demonstrates the Collapsible component.
func (s *showcase) collapsiblePage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "折叠", "单段折叠面板，包 mygo Collapsible", func() {
		s.card(c, t, func() {
			display.NewCollapsible(c, "更多选项", &s.collapsibleOpen, func() { ui.Text(c, "折叠里的内容") })
		})
	})
}
