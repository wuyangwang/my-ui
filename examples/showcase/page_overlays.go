package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// overlaysPage shows the overlays: a tip, a popover and a dialog.
func (s *showcase) overlaysPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "提示", "Tooltip：指针停在锚点上时出现，读屏也会念出提示", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				myui.NewTooltip(c, myui.NewButton(c, "帮助").Element, "这是帮助文字")
				myui.NewTooltip(c, myui.NewButton(c, "只有图标").Outline().Element, "图标按钮的提示")
				myui.NewButton(c, "普通按钮").Outline()
			})
		})
	})

	s.sectionNote(c, t, "浮层", "Popover：点击锚点打开，点击外面关闭", func() {
		s.card(c, t, func() {
			myui.NewPopover(c, &s.popoverOpen, func() ui.Element {
				return myui.NewButton(c, "更多").Outline().Element
			}, func(panel ui.Element) {
				ui.Column(c).Gap(t.Space(2)).Children(func() {
					ui.Text(c, "面板里的内容").Bold()
					ui.Text(c, "点击面板外即可关闭")
				})
			})
		})
	})

	s.sectionNote(c, t, "对话框", "Modal：模态，点背景罩、按 Escape 或点右上角关闭", func() {
		s.card(c, t, func() {
			myui.NewModal(c, &s.modalOpen, func() ui.Element {
				return myui.NewButton(c, "打开对话框").Primary().Element
			}, func(panel ui.Element) {
				// The panel is the dialog as assistive technology sees
				// it: name it, so a screen reader says what it is for.
				panel.Label("删除项目")
				ui.Column(c).Gap(t.Space(3)).Children(func() {
					ui.Text(c, "确定删除这个项目吗？").Bold().FontSize(t.FontSize + 2)
					ui.Text(c, "这个操作不能撤销。").TextColor(t.TextMuted)
					ui.Row(c).Gap(t.Space(2)).AlignItems(ui.End).Children(func() {
						cancel := myui.NewButton(c, "取消").Ghost()
						if cancel.Clicked() {
							s.modalOpen = false
						}
						myui.NewButton(c, "删除").Destructive()
					})
				})
			})
		})
	})
}
