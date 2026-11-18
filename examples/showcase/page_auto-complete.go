package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/forms"
)

// autoCompletePage demonstrates the AutoComplete component on its own page.
func (s *showcase) autoCompletePage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "自动补全", "过滤与选择由 mygo 负责，外观跟随主题", func() {
		s.card(c, t, func() {
			forms.NewAutoComplete(c, &s.autoValue, []string{"苹果", "香蕉", "橘子"}).
				Variant(myui.Outline).Placeholder("输入或选择水果").Width(280)
		})
	})
}
