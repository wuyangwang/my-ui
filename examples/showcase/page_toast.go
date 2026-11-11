package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// toastPage shows the toasts: the buttons show them, and they come out
// at the corner of the window.
func (s *showcase) toastPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "通知", "Toast：操作的结果，停在窗口的右下角，几秒后自己走", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				plain := myui.NewButton(c, "普通通知")
				if plain.Clicked() {
					myui.NewToast(c, "已保存").Description("改动已经写入。").Show()
				}
				ok := myui.NewButton(c, "成功")
				if ok.Clicked() {
					myui.NewToast(c, "操作成功").Description("改动已经写入。").Success().Show()
				}
				warn := myui.NewButton(c, "警告")
				if warn.Clicked() {
					myui.NewToast(c, "磁盘快满了").Description("清理一些旧文件吧。").Warning().Show()
				}
				bad := myui.NewButton(c, "失败")
				if bad.Clicked() {
					myui.NewToast(c, "出错了").Description("网络没有连上，请稍后再试。").Destructive().Show()
				}
				undo := myui.NewButton(c, "带撤销")
				if undo.Clicked() {
					myui.NewToast(c, "已删除").Action("撤销", func() {
						s.lastEvent = "撤销了删除"
					}).Show()
				}
			})
		})
	})
}
