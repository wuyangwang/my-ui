package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// badgePage is the Badge's page: variants and sizes.
func (s *showcase) badgePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体", "标签/徽标的语义预设", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				myui.NewBadge(c.Key("badge-native"), "Native").Native()
				myui.NewBadge(c.Key("badge-primary"), "Primary").Primary()
				myui.NewBadge(c.Key("badge-outline"), "Outline").Outline()
				myui.NewBadge(c.Key("badge-ghost"), "Ghost").Ghost()
				myui.NewBadge(c.Key("badge-destructive"), "Danger").Destructive()
				myui.NewBadge(c.Key("badge-link"), "Link").Link()
			})
		})
	})

	s.sectionNote(c, t, "尺寸", "内衬与字号随 Size 变化", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
				myui.NewBadge(c.Key("badge-sm"), "Sm").Primary().Sm()
				myui.NewBadge(c.Key("badge-md"), "Md").Primary().Md()
				myui.NewBadge(c.Key("badge-lg"), "Lg").Primary().Lg()
			})
		})
	})

	s.sectionNote(c, t, "组合示例", "与按钮等组件并排展示", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
				myui.NewButton(c.Key("badge-btn"), "消息").Outline()
				myui.NewBadge(c.Key("badge-count"), "12").Primary().Sm()
			})
		})
	})
}
