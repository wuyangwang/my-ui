package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/layout"
)

// stepsPage demonstrates the process stepper.
func (s *showcase) stepsPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "步骤", "当前步之前的连线为强调色，支持圆点与竖排", func() {
		s.card(c, t, func() {
			items := []layout.StepItem{
				{Title: "账号", Description: "登录"},
				{Title: "资料", Description: "填写姓名与头像"},
				{Title: "团队", Description: "邀请成员"},
				{Title: "完成"},
			}
			layout.NewSteps(c, layout.StepsProps{Current: 1, Items: items})
		})
	})
}
