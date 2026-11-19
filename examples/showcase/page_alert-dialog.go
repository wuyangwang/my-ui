package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
)

// alertDialogDemo is the page's own state, kept on a stable element so
// the window remembers the open flag and the last result across frames.
type alertDialogDemo struct {
	open   bool
	result string
}

// alertDialogPage demonstrates an alert whose buttons are myui presets:
// the last one is destructive and focused, and the result comes back
// through OnResult.
func (s *showcase) alertDialogPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("alert-dialog-demo")).Hidden(true)
	st := ui.Local(holder, "alert-dialog-demo", func() alertDialogDemo {
		return alertDialogDemo{result: "尚未选择"}
	})
	if st == nil {
		return
	}
	s.sectionNote(c, t, "AlertDialog", "包 mygo 的警告框契约；按钮走 myui 预设，最后一个为 Destructive，索引由 OnResult 交回", func() {
		s.card(c, t, func() {
			trigger := ctrlpkg.NewButton(c, "删除项目").Destructive()
			if trigger.Clicked() {
				st.open = true
			}
			buttons := []string{"取消", "确认删除"}
			overlaypkg.NewAlertDialog(c, &st.open, overlaypkg.AlertDialogOptions{
				Title:   "删除这个项目？",
				Message: "删除后无法恢复，请确认继续。",
				Buttons: buttons,
				OnResult: func(i int) {
					if i == ui.AlertDismissed {
						st.result = "已取消（Escape）"
						return
					}
					if i >= 0 && i < len(buttons) {
						st.result = fmt.Sprintf("按钮 %d：%s", i, buttons[i])
					}
				},
			})
			ui.Row(c).AlignItems(ui.Center).Children(func() {
				ui.Text(c, "结果："+st.result).FontSize(myui.TextSize(t, myui.TextSm)).
					TextColor(t.TextMuted)
			})
		})
	})
}
