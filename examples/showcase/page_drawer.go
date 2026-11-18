package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
)

// drawerPage shows the drawer: a panel that comes from the edge of the
// window.
func (s *showcase) drawerPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "抽屉", "Drawer：从窗口的边上来的一块面板，点外面或按 Escape 收起", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).Children(func() {
				overlaypkg.NewDrawer(c, &s.drawerOpen, overlaypkg.DrawerRight, func() ui.Element {
					return ctrlpkg.NewButton(c, "打开抽屉").Outline().Element
				}, func(panel ui.Element) {
					ui.Column(c).Gap(t.Space(3)).Children(func() {
						ui.Text(c, "设置").Bold().FontSize(t.FontSize + 2)
						ui.Text(c, "面板里的内容").TextColor(t.TextMuted)
						ctrlpkg.NewButton(c, "面板里的按钮").Outline()
					})
				})
				ctrlpkg.NewButton(c, "普通按钮").Outline()
			})
		})
	})
}
