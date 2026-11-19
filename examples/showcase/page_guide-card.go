package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// guideCardPage shows the dismissible card's surfaces.
func (s *showcase) guideCardPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体", "柔和底、带框与无面", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Start).Children(func() {
				displaypkg.NewGuideCard(c, displaypkg.GuideCardOptions{
					Title: "认识你的智能体", Desc: "它们记得上下文，接着上次继续。",
					Variant: myui.Default, Shadow: true,
				}).Width(200)
				displaypkg.NewGuideCard(c, displaypkg.GuideCardOptions{
					Title: "团队空间", Desc: "把常用工具放进同一个工作区。",
					Variant: myui.Outline, Shadow: true,
				}).Width(200)
				displaypkg.NewGuideCard(c, displaypkg.GuideCardOptions{
					Title: "快捷操作", Desc: "没有边框的轻量提示。",
					Variant: myui.Ghost,
				}).Width(200)
			})
		})
	})
	s.sectionNote(c, t, "关闭", "关闭指针由 ui.Local 持有", func() {
		s.card(c, t, func() {
			holder := ui.Box(c)
			closed := ui.Local(holder, "showcase-guide-closed", func() bool { return false })
			displaypkg.NewGuideCard(c, displaypkg.GuideCardOptions{
				Title: "可关闭的提示", Desc: "点击右上角关闭，下一帧隐藏。",
				Closed: closed, OnClose: func() { s.lastAction = "Guide closed" },
			}).Width(260)
		})
	})
}
