package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
	"github.com/wuyangwang/my-ui/icons"
)

// alertPage shows the alerts: the kinds of message a page carries.
func (s *showcase) alertPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "提示", "Alert：页面里的一段消息，读屏会念出它", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				displaypkg.NewAlert(c, "这是一条提示").Description("说明文字在标题下面，弱一些。")
				displaypkg.NewAlert(c, "保存前看一眼").Description("这条改动会影响所有人。").Warning().
					Icon(icons.TriangleAlert())
				displaypkg.NewAlert(c, "操作成功").Description("改动已经写入。").Success().
					Icon(icons.CircleCheck())
				displaypkg.NewAlert(c, "出错了").Description("网络没有连上，请稍后再试。").Destructive().
					Icon(icons.CircleX())
				displaypkg.NewAlert(c, "一条信息").Description("带图标的信息。").Info().
					Icon(icons.Info())
			})
		})
	})
}
