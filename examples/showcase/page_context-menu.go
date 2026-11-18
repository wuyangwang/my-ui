package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/overlays"
)

// contextMenuPage demonstrates the ContextMenu component on its own page.
func (s *showcase) contextMenuPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "右键菜单", "给任意目标元素附上一组操作", func() {
		s.card(c, t, func() {
			target := ui.Box(c).Padding(t.Space(2), t.Space(4)).Radius(t.Radius).
				Background(t.Surface).Children(func() { ui.Text(c, "右键我") })
			overlays.NewContextMenu(c, &s.contextOpen, target, []overlays.MenuItem{
				{Label: "重命名", OnClick: func() { s.lastAction = "重命名" }},
			})
			ui.Text(c, "最近："+s.lastAction).TextColor(t.TextMuted)
		})
	})
}
