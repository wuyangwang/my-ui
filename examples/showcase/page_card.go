package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// cardPage shows the display pieces: a card, a progress bar and a key
// cap.
func (s *showcase) cardPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "卡片", "Card：页面里的一块面板，有标题、说明与内容", func() {
		s.card(c, t, func() {
			displaypkg.NewCard(c, "概览").Description("今天的情况").Content(func() {
				ui.Column(c).Gap(t.Space(2)).Children(func() {
					ui.Text(c, "卡片里的内容").TextColor(t.TextMuted)
					ui.Row(c).Gap(t.Space(2)).Children(func() {
						ctrlpkg.NewButton(c, "查看").Primary().Sm()
						ctrlpkg.NewButton(c, "忽略").Ghost().Sm()
					})
				})
			})
		})
	})

	s.sectionNote(c, t, "进度", "Progress：一条线，走到哪算哪", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				displaypkg.NewProgress(c, 0.25).Label("上传进度")
				displaypkg.NewProgress(c, 0.65).Label("下载进度")
				displaypkg.NewProgress(c, 1).Label("已完成")
			})
		})
	})

	s.sectionNote(c, t, "快捷键", "Kbd：快捷键要按哪几个键", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				displaypkg.NewKbd(c, "Ctrl", "K")
				displaypkg.NewKbd(c, "Cmd", "Shift", "P")
				ui.Text(c, "搜索").TextColor(t.TextMuted)
			})
		})
	})
}
