package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// layoutPage shows the docs shell: header, sidebar, article and its table
// of contents. The active item is kept in a ui.Local of a hidden holder.
func (s *showcase) layoutPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("layout-demo")).Hidden(true)
	active := ui.Local(holder, "layout-demo", func() string { return "简介" })
	if active == nil {
		return
	}
	s.sectionNote(c, t, "Layout", "文档版式：头栏 + 侧栏 + 正文 + 目录，内容在头栏下滚动", func() {
		s.card(c, t, func() {
			frame := ui.Column(c).FillWidth().Height(t.Space(100)).
				Border(1, t.Border).Radius(myui.Radius(t, myui.RadiusLg)).Clip()
			frame.Children(func() {
				layoutpkg.NewLayout(c, layoutpkg.LayoutProps{
					Header: func() {
						ui.Row(c).FillWidth().Height(t.Space(10)).Padding(0, t.Space(4)).
							AlignItems(ui.Center).Background(t.Surface).
							BorderWidth(0, 0, 1, 0).BorderColor(t.Border).Children(func() {
							ui.Text(c, "文档").Bold()
						})
					},
					Sidebar: func() {
						ui.Column(c).Padding(t.Space(3)).Gap(t.Space(1)).Children(func() {
							for _, p := range []string{"简介", "安装", "主题", "常见问题"} {
								p := p
								b := ctrlpkg.NewButton(c, p).Ghost().Sm()
								if *active == p {
									b.Primary()
								}
								if b.Clicked() {
									*active = p
								}
							}
						})
					},
					Toc: func() {
						ui.Column(c).Padding(t.Space(4)).Gap(t.Space(2)).Children(func() {
							ui.Text(c, "本页目录").FontSize(myui.TextSize(t, myui.TextXs)).Bold().
								TextColor(t.TextMuted)
							for _, h := range []string{"概览", "快速开始", "组件"} {
								ui.Text(c, h).TextColor(t.TextMuted)
							}
						})
					},
					Footer: func() {
						ui.Text(c, "版权所有 © 2026").TextColor(t.TextMuted)
					},
				}, func() {
					ui.Column(c).Padding(t.Space(6)).Gap(t.Space(3)).Children(func() {
						ui.Text(c, "页面："+*active).FontSize(myui.TextSize(t, myui.TextXl)).Bold()
						ui.Text(c, "正文区域在头栏下方独立滚动，侧栏和目录保持不动。").
							TextColor(t.TextMuted)
					})
				})
			})
		})
	})
}
