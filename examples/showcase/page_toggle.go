package main

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// togglePage demonstrates a pressed-state button. The value beside it is
// read from the app's bool, so the effect of a click is visible at once.
func (s *showcase) togglePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "Toggle", "保持按下的按钮。点击或按空格切换 *bool；右边的文字读取它，变化走 Changed / OnChanged / Handle。", func() {
		s.card(c, t, func() {
			ui.Row(c).AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
				ctrlpkg.NewToggle(c, &s.toggleOn, "粗体").Primary().Icon(icons.Bold())
				ui.Text(c, "toggleOn = "+strconv.FormatBool(s.toggleOn)).
					FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
			})
		})
	})

	s.sectionNote(c, t, "状态与档位", "选中态、图标、禁用态，以及共享的 Xs/Sm/Md/Lg 尺寸。", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				ui.Row(c).AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
					ctrlpkg.NewToggle(c, &s.toggleOn, "开").Outline().Sm()
					ctrlpkg.NewToggle(c, &s.notifications, "通知").Secondary().Icon(icons.Bell())
					disabled := true
					ctrlpkg.NewToggle(c, &disabled, "禁用").Disabled(true)
				})
				ui.Row(c).AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
					ctrlpkg.NewToggle(c, &s.mail, "Xs").Xs()
					ctrlpkg.NewToggle(c, &s.toggleOn, "Sm").Sm()
					ctrlpkg.NewToggle(c, &s.notifications, "Md").Md()
					ctrlpkg.NewToggle(c, &s.mail, "Lg").Lg()
				})
			})
		})
	})
}
