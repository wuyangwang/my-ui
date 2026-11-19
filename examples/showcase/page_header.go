package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// headerPage shows the desktop and mobile arrangements of the top bar.
func (s *showcase) headerPage(c *ui.Context, t *ui.Theme) {
	logo := func() {
		ui.Text(c, "LobeHub").FontSize(myui.TextSize(t, myui.TextMd)).Bold()
	}
	actions := func() {
		ctrlpkg.NewButton(c, "").Icon(icons.Search()).Label("搜索").Ghost().Sm()
		ctrlpkg.NewButton(c, "").Icon(icons.Moon()).Label("主题").Ghost().Sm()
	}
	s.sectionNote(c, t, "Header", "站点顶栏：桌面为品牌、导航、操作；移动端把菜单放到最前", func() {
		s.card(c, t, func() {
			layoutpkg.NewHeader(c, layoutpkg.HeaderProps{
				Logo: logo,
				Nav: func() {
					ui.Row(c).Gap(t.Space(6)).Children(func() {
						for _, s := range []string{"组件", "文档", "更新"} {
							ui.Text(c, s).TextColor(t.TextMuted)
						}
					})
				},
				Actions: actions,
			})
			ui.Box(c).Height(t.Space(4))
			layoutpkg.NewHeader(c, layoutpkg.HeaderProps{
				Mobile:  true,
				Logo:    func() { ui.Row(c).Grow(1).Justify(ui.Center).Children(logo) },
				Nav:     func() { ctrlpkg.NewButton(c, "").Icon(icons.Menu()).Label("菜单").Ghost().Sm() },
				Actions: actions,
			})
		})
	})
}
