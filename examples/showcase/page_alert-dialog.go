package main

import (
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
// each action carries its own label and behavior, so no index-to-label
// mapping is left in the app.
func (s *showcase) alertDialogPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("alert-dialog-demo")).Hidden(true)
	st := ui.Local(holder, "alert-dialog-demo", func() alertDialogDemo {
		return alertDialogDemo{result: "尚未选择"}
	})
	if st == nil {
		return
	}
	s.sectionNote(c, t, "AlertDialog", "包 mygo 的警告框契约；按钮走 myui 预设，每个动作自带标签与回调", func() {
		s.card(c, t, func() {
			trigger := ctrlpkg.NewButton(c, "删除项目").Destructive()
			if trigger.Clicked() {
				st.open = true
			}
			overlaypkg.NewAlertDialog(c, overlaypkg.AlertDialogOptions{
				Open:    &st.open,
				Title:   "删除这个项目？",
				Message: "删除后无法恢复，请确认继续。",
				Actions: []overlaypkg.DialogAction{
					{Label: "取消", OnClick: func() { st.result = "已取消" }},
					{Label: "确认删除", Kind: myui.Destructive, OnClick: func() { st.result = "已删除" }},
				},
				OnDismiss: func() { st.result = "已取消（Escape）" },
			})
			ui.Row(c).AlignItems(ui.Center).Children(func() {
				ui.Text(c, "结果："+st.result).FontSize(myui.TextSize(t, myui.TextSm)).
					TextColor(t.TextMuted)
			})
		})
	})
}
