package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// tabsPage shows the tabs: one of them is open, and it shows its own
// content.
func (s *showcase) tabsPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "标签页", "Tabs：一次看一块，只建开着的那块", func() {
		s.card(c, t, func() {
			displaypkg.NewTabs(c, &s.tab,
				displaypkg.Tab{Label: "概览", Content: func() {
					ui.Text(c, "概览的内容").TextColor(t.TextMuted)
				}},
				displaypkg.Tab{Label: "数据", Content: func() {
					ui.Text(c, "数据的内容").TextColor(t.TextMuted)
				}},
				displaypkg.Tab{Label: "设置", Content: func() {
					ui.Text(c, "设置的内容").TextColor(t.TextMuted)
				}},
			)
		})
	})
}
