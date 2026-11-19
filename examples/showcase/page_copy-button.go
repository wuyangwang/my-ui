package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// copyButtonPage demonstrates the CopyButton component.
func (s *showcase) copyButtonPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "CopyButton", "点击复制文本，随后短暂显示对勾", func() {
		s.demoCard(c, t, "Borderless", "无面", "纯图标的复制动作", func() {
			ctrlpkg.NewCopyButton(c, "go get github.com/wuyangwang/my-ui")
		})
		s.demoCard(c, t, "Filled", "填充", "带弱填充的复制动作", func() {
			ctrlpkg.NewCopyButton(c, "myui.NewCopyButton(c, text)").Filled()
		})
		s.demoCard(c, t, "Outlined", "描边", "带边框的复制动作", func() {
			ctrlpkg.NewCopyButton(c, "v0.10.0").Outlined().Sm()
		})
	})
}
