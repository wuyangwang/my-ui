package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
	"github.com/wuyangwang/my-ui/overlays"
)

// commandPalettePage demonstrates a searchable command dialog.
func (s *showcase) commandPalettePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "命令面板", "输入即过滤命令，点击执行后自动关闭", func() {
		s.card(c, t, func() {
			open := ctrlpkg.NewButton(c, "打开命令面板").Outline().Sm()
			if open.Clicked() {
				s.cmdOpen = true
			}
			overlays.NewCommandPalette(c, &s.cmdOpen, &s.cmdQuery, []overlays.Command{
				{Label: "新建文件", Icon: icons.FilePlus(), Shortcut: "⌘N", OnSelect: func() { s.lastAction = "新建文件" }},
				{Label: "打开文件", Icon: icons.FolderOpen(), Shortcut: "⌘O", OnSelect: func() { s.lastAction = "打开文件" }},
				{Label: "保存文件", Icon: icons.Save(), Shortcut: "⌘S", OnSelect: func() { s.lastAction = "保存文件" }},
			})
			ui.Text(c, "最近："+s.lastAction).TextColor(t.TextMuted)
		})
	})
}
