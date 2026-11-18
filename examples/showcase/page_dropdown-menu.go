package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
	"github.com/wuyangwang/my-ui/overlays"
)

// dropdownMenuPage demonstrates the DropdownMenu component on its own page.
func (s *showcase) dropdownMenuPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "下拉菜单", "触发项传的是菜单按钮的内容，不要再套一个按钮", func() {
		s.card(c, t, func() {
			overlays.NewDropdownMenu(c, &s.menuOpen, func() ui.Element {
				return ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
					ui.Text(c, "操作")
					ui.Icon(c, icons.ChevronDown()).Size(myui.IconSizeXs, myui.IconSizeXs)
				})
			}, []overlays.MenuItem{
				{Label: "复制", Icon: icons.Copy(), OnClick: func() { s.lastAction = "复制" }},
				{Label: "删除", Icon: icons.Trash2(), OnClick: func() { s.lastAction = "删除" }},
			})
			ui.Text(c, "最近："+s.lastAction).TextColor(t.TextMuted)
		})
	})
}
