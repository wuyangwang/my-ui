package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// actionIconGroupPage demonstrates the ActionIconGroup component.
func (s *showcase) actionIconGroupPage(c *ui.Context, t *ui.Theme) {
	items := []ctrlpkg.ActionIconGroupItem{
		{Key: "复制", Icon: icons.Copy(), Label: "复制"},
		{Key: "分享", Icon: icons.Share(), Label: "分享"},
		{Key: "固定", Icon: icons.Pin(), Label: "固定"},
		{Key: "删除", Icon: icons.Trash(), Label: "删除", Danger: true},
		{Key: "禁用", Icon: icons.Settings(), Label: "禁用", Disabled: true},
	}
	s.sectionNote(c, t, "ActionIconGroup", "一排共用底面的图标动作，点击写入 chosen 指针", func() {
		s.demoCard(c, t, "Filled", "填充", "带弱填充的整组动作", func() {
			ctrlpkg.NewActionIconGroup(c, &s.lastAction, items).Filled()
		})
		s.demoCard(c, t, "Outlined", "描边", "整组共用一条边框", func() {
			ctrlpkg.NewActionIconGroup(c, &s.lastAction, items).Outlined().Sm()
		})
		s.card(c, t, func() {
			ui.Text(c, "最近点击："+s.lastAction).
				FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
		})
	})
}
