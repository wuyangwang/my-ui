package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
	myuiwin "github.com/wuyangwang/my-ui/window"
)

// windowPage demonstrates window.New: one button opens a custom window
// of its own, whose title bar follows the showcase's theme.
func (s *showcase) windowPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "用途", "Window：window.New 一行建窗——无框窗口（TitleBarHidden + HideWindowButtons）+ 内置 TitleBar + 窗口动作接线，标题栏随明暗主题换色。点下面的按钮开一个试试", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				open := ctrlpkg.NewButton(c.Key("open custom window"), "打开自定义窗口").Primary()
				if open.Clicked() {
					s.openCustomWindow()
				}
				ui.Text(c, "新窗口左侧是品牌、右侧是最小化/最大化/关闭；整条可拖拽，双击标题栏最大化，Windows 11 悬停最大化按钮出贴靠布局。").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			})
		})
	})
	s.sectionNote(c, t, "要点", "Theme 每帧取 showcase 的明暗/色板，新窗口与主窗口一起换色；New 返回 *mygo.Window，想继续操作窗口就留着它（这个演示在内容区放了关闭按钮）", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				for _, line := range []string{
					"window.New(window.Options{",
					"  Title: \"自定义窗口\", Name: \"my-ui\",",
					"  Width: 560, Height: 380,",
					"  Theme: s.theme, View: appView,",
					"})",
				} {
					ui.Text(c, line).FontSize(myui.TextSize(t, myui.TextXs))
				}
			})
		})
	})
}

// openCustomWindow builds one window of the shell's own: the brand on
// the bar, the showcase's theme lights it, and the content closes it.
func (s *showcase) openCustomWindow() {
	var win *mygo.Window
	win = myuiwin.New(myuiwin.Options{
		Title:  "自定义窗口",
		Name:   "my-ui",
		Icon:   icons.Component(), // 就一个图标 + 名字，其余全用默认
		Width:  560,
		Height: 380,
		Theme:  s.theme,
		View: func(c *ui.Context) {
			th := c.Theme()
			ui.Column(c).Fill().Background(th.Background).
				Padding(th.Space(6), th.Space(7)).Gap(th.Space(3)).Children(func() {
				ui.Text(c, "这是一个 window.New 建的自定义窗口。").
					FontSize(myui.TextSize(th, myui.TextLg)).Bold()
				ui.Text(c, "标题栏在主题自己的颜色里画——系统标题栏不会跟随应用的明暗切换，这个会。右上是窗口按钮，整条标题栏都可以拖拽。").
					FontSize(myui.TextSize(th, myui.TextSm)).TextColor(th.TextMuted)
				close := ctrlpkg.NewButton(c.Key("close custom window"), "关闭此窗口")
				if close.Clicked() && win != nil {
					win.Close()
				}
			})
		},
	})
	_ = win
}
