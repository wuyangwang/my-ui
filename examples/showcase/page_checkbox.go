package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// checkboxPage shows the check boxes: their states, their sizes and
// their labels.
func (s *showcase) checkboxPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "勾选", "Checkbox：一个勾上的框，和它的标签", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				ctrlpkg.NewCheckbox(c.Key("cb mail"), &s.mail, "邮件通知")
				ctrlpkg.NewCheckbox(c.Key("cb calendar"), &s.calendar, "日历提醒")
				ctrlpkg.NewCheckbox(c.Key("cb off"), &s.switchOff, "已禁用").Disabled(true)
			})
		})
	})

	s.sectionNote(c, t, "尺寸", "尺寸档为 14 / 16 / 18 / 22 DIP", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				ctrlpkg.NewCheckbox(c.Key("cb xs"), &s.mail, "最小").Xs()
				ctrlpkg.NewCheckbox(c.Key("cb sm"), &s.mail, "小").Sm()
				ctrlpkg.NewCheckbox(c.Key("cb md"), &s.mail, "中").Md()
				ctrlpkg.NewCheckbox(c.Key("cb lg"), &s.mail, "大").Lg()
			})
		})
	})
}
