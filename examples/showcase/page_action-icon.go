package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// actionIconPage demonstrates the ActionIcon component.
func (s *showcase) actionIconPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "ActionIcon", "纯图标按钮，Borderless / Filled / Outlined 三种面", func() {
		s.card(c, t, func() {
			ui.Row(c).AlignItems(ui.Center).Gap(t.Space(4)).Children(func() {
				ctrlpkg.NewActionIcon(c, icons.Copy(), "复制").Borderless()
				ctrlpkg.NewActionIcon(c, icons.Settings(), "设置").Filled()
				ctrlpkg.NewActionIcon(c, icons.Trash(), "删除").Outlined().Danger(true)
				pin := ctrlpkg.NewActionIcon(c, icons.Pin(), "固定").Filled().Active(s.toggleOn)
				if pin.Clicked() {
					s.toggleOn = !s.toggleOn
				}
				ctrlpkg.NewActionIcon(c, icons.Bell(), "通知").Xs()
				ctrlpkg.NewActionIcon(c, icons.Shapes(), "组件").Lg()
			})
		})
	})
}
