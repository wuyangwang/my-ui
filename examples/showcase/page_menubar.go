package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
	"github.com/wuyangwang/my-ui/overlays"
)

// menubarPage demonstrates the application menubar.
func (s *showcase) menubarPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "菜单栏", "横排多个下拉菜单，面板与条目沿用 DropdownMenu 的外观", func() {
		s.card(c, t, func() {
			overlays.NewMenubar(c, []overlays.MenubarMenu{
				{Label: "文件", Items: []overlays.MenuItem{
					{Label: "打开", Icon: icons.FolderOpen(), OnClick: func() { s.lastAction = "打开" }},
					{Label: "保存", Icon: icons.Save(), OnClick: func() { s.lastAction = "保存" }},
				}},
				{Label: "编辑", Items: []overlays.MenuItem{
					{Label: "撤销", Icon: icons.Undo2(), OnClick: func() { s.lastAction = "撤销" }},
				}},
			})
			ui.Text(c, "最近："+s.lastAction).TextColor(t.TextMuted)
		})
	})
}
