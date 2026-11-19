package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// toggleGroupPage demonstrates the ToggleGroup component.
func (s *showcase) toggleGroupPage(c *ui.Context, t *ui.Theme) {
	build := func() {
		ctrlpkg.NewToggle(c, &s.toggleOn, "粗体").Size(myui.SizeSm)
		ctrlpkg.NewToggle(c, &s.notifications, "斜体").Size(myui.SizeSm)
		ctrlpkg.NewToggle(c, &s.mail, "下划线").Size(myui.SizeSm)
	}
	s.sectionNote(c, t, "ToggleGroup", "把若干开关合成一个控件，方向键在组内移动", func() {
		s.demoCard(c, t, "Outlined", "描边", "共用一条边框和分隔", func() {
			ctrlpkg.NewToggleGroup(c, build).Label("文字样式").Outlined()
		})
		s.demoCard(c, t, "Borderless", "无面", "彼此分开的开关", func() {
			ctrlpkg.NewToggleGroup(c, build).Label("文字样式").Borderless()
		})
		s.demoCard(c, t, "Segmented", "轨道", "放在分段轨道上的开关", func() {
			ctrlpkg.NewToggleGroup(c, build).Label("文字样式").Segmented()
		})
	})
}
