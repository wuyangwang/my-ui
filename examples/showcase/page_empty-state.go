package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/display"
	"github.com/wuyangwang/my-ui/icons"
)

// emptyStatePage demonstrates the EmptyState component on its own page.
func (s *showcase) emptyStatePage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "用途", "列表为空时告诉用户为什么，以及下一步做什么", func() {
		s.card(c, t, func() {
			display.NewEmptyState(c, display.EmptyStateOptions{
				Icon: icons.Inbox(), Title: "还没有任何记录",
				Description: "建一条记录后，它会出现在这里",
				Action:      func() { ctrlpkg.NewButton(c.Key("empty action"), "新建").Primary().Sm() },
			})
		})
	})
}
