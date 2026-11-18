package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/display"
)

// spinPage demonstrates the Spin component on its own page.
func (s *showcase) spinPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "加载", "Reduce Motion 由桌面决定，组件不自己计时", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				display.NewSpin(c).Variant(myui.Primary)
				display.NewSpin(c).Size(myui.SizeSm).Variant(myui.Success)
				ui.Text(c, "加载中…").TextColor(t.TextMuted)
			})
		})
	})
}
