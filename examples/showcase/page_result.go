package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/display"
)

// resultPage demonstrates the Result component on its own page.
func (s *showcase) resultPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "状态", "成功、警告、信息、错误，各自用主题的语义色", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Start).Children(func() {
				for _, st := range []display.ResultStatus{
					display.ResultSuccess, display.ResultWarning,
					display.ResultInfo, display.ResultError,
				} {
					display.NewResult(c, display.ResultOptions{
						Status: st, Title: "提交完成", Description: "这一页展示结果状态",
					})
				}
			})
		})
	})
}
