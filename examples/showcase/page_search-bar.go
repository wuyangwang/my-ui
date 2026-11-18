package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/forms"
)

// searchBarPage demonstrates the SearchBar component on its own page.
func (s *showcase) searchBarPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "搜索", "输入有值时才显示清除图标，尺寸与 Input 同档", func() {
		s.card(c, t, func() {
			forms.NewSearchBar(c, &s.searchQuery).Placeholder("搜索仓库…").Width(280)
		})
	})
}
