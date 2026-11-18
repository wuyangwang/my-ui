package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// togglePage demonstrates the Toggle component.
func (s *showcase) togglePage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "切换按钮", "保持按下的按钮，Size 走共享档位", func() {
		s.card(c, t, func() { ctrlpkg.NewToggle(c, &s.toggleOn, "开关") })
	})
}
