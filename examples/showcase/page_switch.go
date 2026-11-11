package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// switchPage shows the switches: the knobs that are on or off.
func (s *showcase) switchPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "开关", "Switch：一个开或关的旋钮，变了会走一段", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
					myui.NewSwitch(c.Key("sw notifications"), &s.notifications).Label("通知")
					myui.NewSwitch(c.Key("sw sync"), &s.sync).Label("同步")
					myui.NewSwitch(c.Key("sw off"), &s.switchOff).Label("已禁用").Disabled(true)
				})
				ui.Text(c, "开关没有文字，名字是无障碍名：读屏读到的是它。").TextColor(t.TextMuted)
			})
		})
	})
}
