package main

import (
	"github.com/egoist/mygo/ui"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
)

// splitButtonPage demonstrates the SplitButton component.
func (s *showcase) splitButtonPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "拆分按钮", "主按钮 + 下拉操作，菜单沿用 my-ui 的 shadcn 面板", func() {
		s.card(c, t, func() {
			overlaypkg.NewSplitButton(c, "保存", func() { s.splitAction = "保存" }, []overlaypkg.MenuItem{
				{Label: "另存为", OnClick: func() { s.splitAction = "另存为" }},
			})
			ui.Text(c, "最近："+s.splitAction).TextColor(t.TextMuted)
		})
	})
}
