package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/display"
)

// breadcrumbPage demonstrates the Breadcrumb component on its own page.
func (s *showcase) breadcrumbPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "路径", "层级路径，点击前面的项回到那一层", func() {
		s.card(c, t, func() {
			display.NewBreadcrumb(c, []string{"首页", "组件", "Breadcrumb"}, &s.crumb)
		})
	})
}
