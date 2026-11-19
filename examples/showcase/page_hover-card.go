package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/overlays"
)

// hoverCardPage demonstrates the HoverCard component on its own page.
func (s *showcase) hoverCardPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "悬停卡片", "面板内容挂在 Body 里，库会把它放进浮层的面板", func() {
		s.card(c, t, func() {
			overlays.NewHoverCard(c, overlays.HoverCardOptions{
				Open: &s.hoverOpen,
				Trigger: func() ui.Element {
					return ctrlpkg.NewButton(c, "悬停查看").Outline().Sm().Element
				},
				Body: func() {
					ui.Text(c, "这里是浮层的内容").Bold()
					ui.Text(c, "内容由库放进面板，才会落在浮层里").TextColor(t.TextMuted)
				},
			})
		})
	})
}
