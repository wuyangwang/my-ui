package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// floatingPanelPage shows a corner panel that leaves the page behind it
// usable. The open flag lives in a ui.Local of a hidden holder, so the
// showcase needs no new field in its own state.
func (s *showcase) floatingPanelPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("floating-panel-demo")).Hidden(true)
	st := ui.Local(holder, "floating-panel-demo", func() struct{ open bool } {
		return struct{ open bool }{open: true}
	})
	if st == nil {
		return
	}
	s.sectionNote(c, t, "FloatingPanel", "窗口角落的浮层面板：打开时页面仍可用，Escape 与关闭按钮收起", func() {
		s.card(c, t, func() {
			trigger := ctrlpkg.NewButton(c, "打开任务面板").Primary()
			if trigger.Clicked() {
				st.open = true
			}
			layoutpkg.NewFloatingPanel(c, &st.open, layoutpkg.FloatingPanelProps{
				Title:     "任务运行",
				Width:     320,
				Placement: layoutpkg.FloatBottomRight,
				Footer: func() {
					ctrlpkg.NewButton(c, "继续").Primary().Sm()
				},
			}, func() {
				ui.Column(c).Gap(t.Space(2)).Padding(t.Space(4)).Children(func() {
					for _, line := range []string{"读取任务", "搜索代码", "定位问题", "写入修复"} {
						ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
							ui.Icon(c, icons.CircleCheck()).Size(myui.IconSizeSm, myui.IconSizeSm).
								TextColor(t.Success)
							ui.Text(c, line)
						})
					}
				})
			})
		})
	})
}
