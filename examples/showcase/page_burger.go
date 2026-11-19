package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// burgerPage shows the header menu button in its dropped and fullscreen
// forms. The open flag and the active key live in a ui.Local.
func (s *showcase) burgerPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("burger-demo")).Hidden(true)
	st := ui.Local(holder, "burger-demo", func() struct {
		open, fullscreen bool
		active           string
	} {
		return struct {
			open, fullscreen bool
			active           string
		}{active: "chat"}
	})
	if st == nil {
		return
	}
	items := []layoutpkg.BurgerItem{
		{Key: "chat", Label: "聊天", Icon: icons.MessageSquare()},
		{Key: "discover", Label: "发现", Icon: icons.Compass()},
		{Key: "files", Label: "文件", Icon: icons.Folder()},
		{Divider: true},
		{Key: "settings", Label: "设置", Icon: icons.Settings()},
	}
	s.sectionNote(c, t, "Burger", "移动端风格的菜单按钮：下拉面板或全屏菜单；桌面实现见限制说明", func() {
		s.card(c, t, func() {
			bar := ui.Row(c).FillWidth().Height(t.Space(12)).
				AlignItems(ui.Center).Padding(0, t.Space(4)).
				BorderWidth(0, 0, 1, 0).BorderColor(t.Border)
			bar.Children(func() {
				ui.Text(c, "LobeHub").Bold().Grow(1)
				layoutpkg.NewBurger(c, &st.open, layoutpkg.BurgerProps{
					Items: items, Selected: &st.active, Fullscreen: st.fullscreen,
					Footer: func() { ctrlpkg.NewButton(c, "登录").Primary().Sm() },
				})
			})
			ui.Text(c, "当前页面："+st.active).TextColor(t.TextMuted)
			toggle := ctrlpkg.NewButton(c, "切换全屏菜单").Ghost().Sm()
			if toggle.Clicked() {
				st.fullscreen = !st.fullscreen
			}
		})
	})
}
