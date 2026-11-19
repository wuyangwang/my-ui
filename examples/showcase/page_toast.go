package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
)

// toastPage shows the toasts: the buttons show them through the window's
// Notifier, and they come out at the corner of the window.
func (s *showcase) toastPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "通知", "Toast：操作的结果，停在窗口的右下角，几秒后自己走", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				plain := ctrlpkg.NewButton(c, "普通通知")
				if plain.Clicked() {
					s.notifier.Show(overlaypkg.ToastOptions{
						Title: "已保存", Description: "改动已经写入。",
					})
				}
				ok := ctrlpkg.NewButton(c, "成功")
				if ok.Clicked() {
					s.notifier.Show(overlaypkg.ToastOptions{
						Title: "操作成功", Description: "改动已经写入。", Kind: myui.Success,
					})
				}
				warn := ctrlpkg.NewButton(c, "警告")
				if warn.Clicked() {
					s.notifier.Show(overlaypkg.ToastOptions{
						Title: "磁盘快满了", Description: "清理一些旧文件吧。", Kind: myui.Warning,
					})
				}
				bad := ctrlpkg.NewButton(c, "失败")
				if bad.Clicked() {
					s.notifier.Show(overlaypkg.ToastOptions{
						Title: "出错了", Description: "网络没有连上，请稍后再试。", Kind: myui.Destructive,
					})
				}
				undo := ctrlpkg.NewButton(c, "带撤销")
				if undo.Clicked() {
					s.notifier.Show(overlaypkg.ToastOptions{
						Title: "已删除", ActionLabel: "撤销",
						OnAction: func() { s.lastEvent = "撤销了删除" },
					})
				}
			})
		})
	})
}
