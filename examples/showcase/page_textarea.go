package main

import (
	"github.com/egoist/mygo/ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// textAreaPage is the TextArea's page: variants, sizes, the growing
// height and the states.
func (s *showcase) textAreaPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体与尺寸", "与 Input 同一套预设：变体管外观，尺寸管内衬与字号", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				formspkg.NewTextArea(c.Key("ta-outline"), &s.draft).Outline().Label("Outline").Placeholder("写点什么…")
				formspkg.NewTextArea(c.Key("ta-sm"), &s.draft).Sm().Outline().Label("小").Placeholder("小尺寸")
				formspkg.NewTextArea(c.Key("ta-ghost"), &s.draft).Ghost().Label("Ghost").Placeholder("无边框")
			})
		})
	})

	s.sectionNote(c, t, "随内容增高", "Lines(最少, 最多)：超出上限后滚动", func() {
		s.card(c, t, func() {
			formspkg.NewTextArea(c.Key("ta-lines"), &s.draft).Outline().Label("随内容增高").
				Lines(3, 6).Placeholder("输入几行试试")
		})
	})

	s.sectionNote(c, t, "状态与组合", "清除按钮、错误提示与事件", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				formspkg.NewTextArea(c.Key("ta-clear"), &s.draft).Outline().Label("可清除").
					Clearable().OnCleared(func() { s.lastEvent = "cleared" })
				formspkg.NewTextArea(c.Key("ta-error"), &s.draft).Outline().Label("错误态").
					Invalid(true).ErrorText("内容太长了")
			})
		})
	})
}
