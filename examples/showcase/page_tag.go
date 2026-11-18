package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/display"
	"github.com/wuyangwang/my-ui/icons"
)

// tagPage demonstrates the Tag component on its own page.
func (s *showcase) tagPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "标签", "比 Badge 轻，常用在列表行的元信息旁", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
				display.NewTag(c, display.TagOptions{
					Label: "新功能", Icon: icons.Sparkles(), Closed: &s.tagClosed,
				}).Variant(myui.Success)
				display.NewTag(c, display.TagOptions{Label: "草稿"}).Variant(myui.Secondary)
			})
		})
	})
}
