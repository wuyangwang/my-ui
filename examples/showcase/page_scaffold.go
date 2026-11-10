package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// scaffoldPage demonstrates the Scaffold layout skeleton.
func (s *showcase) scaffoldPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "用途", "Scaffold 用于统一页面结构：可选标题栏安全区/顶部栏、可滚动内容区、底部栏，并为底部栏预留滚动空间（避免最后内容被遮挡）", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				ui.Text(c, "示例预览（卡片内模拟一个页面）").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
				ui.Column(c).Height(300).Background(t.Background).Border(1, t.Border).Radius(t.Radius).Children(func() {
					myui.Scaffold(c, myui.ScaffoldOptions{
						Top: func() {
							ui.Row(c).Fill().Padding(t.Space(2), t.Space(3)).Background(t.Surface).AlignItems(ui.Center).Children(func() {
								ui.Text(c, "页面标题 / 顶部栏").FontSize(myui.TextSize(t, myui.TextSm)).Bold()
							})
						},
						Bottom: func() {
							ui.Row(c).Fill().Padding(t.Space(2), t.Space(3)).Background(t.Surface).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
								ui.Row(c).Grow(1).Children(func() {})
								myui.NewButton(c.Key("scaffold-cancel"), "取消").Outline().Sm()
								myui.NewButton(c.Key("scaffold-ok"), "保存").Primary().Sm()
							})
						},
						BottomInset: 56,
						Content: func() {
							ui.Column(c).Padding(t.Space(3)).Gap(t.Space(2)).Children(func() {
								ui.Text(c, "这里是可滚动内容区").Bold()
								for i := 1; i <= 10; i++ {
									ui.Text(c, "列表项 / 内容行")
								}
								ui.Text(c, "最后一行（应在底部栏上方可见）").TextColor(t.TextMuted)
							})
						},
					})
				})
			})
		})
	})
	s.sectionNote(c, t, "标题栏安全区", "隐藏系统标题栏时可用 UseTitleBar 让顶部区域避开窗口控件并可拖拽", func() {
		s.card(c, t, func() {
			ui.Text(c, "myui.Scaffold(c, myui.ScaffoldOptions{UseTitleBar: true, Top: ..., Bottom: ..., Content: ...})").
				FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
		})
	})
}
