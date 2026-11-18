package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/display"
	"github.com/wuyangwang/my-ui/icons"
)

// statusPage demonstrates the four display components added in this
// batch: loading, empty, result and pagination.
func (s *showcase) statusPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "Skeleton", "数据到达前的占位条，颜色取主题 Surface", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(4)).Children(func() {
				display.NewSkeleton(c, 3).Width(360)
				ui.Row(c).Gap(t.Space(4)).Children(func() {
					display.NewSkeleton(c, 1).Circle()
					display.NewSkeleton(c, 2).Width(220).Size(myui.SizeSm)
				})
			})
		})
	})
	s.sectionNote(c, t, "EmptyState", "没有内容时的图标、说明和下一步", func() {
		s.card(c, t, func() {
			display.NewEmptyState(c, display.EmptyStateOptions{
				Icon:        icons.Inbox(),
				Title:       "还没有任何记录",
				Description: "建一条记录后，它会出现在这里",
				Action: func() {
					ctrlpkg.NewButton(c.Key("empty action"), "新建").Primary().Sm()
				},
			})
		})
	})
	s.sectionNote(c, t, "Result", "四种结果状态，图标与语义色随主题", func() {
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
	s.sectionNote(c, t, "Pagination", "页码按钮写回当前页指针，首尾与省略号自动处理", func() {
		s.card(c, t, func() {
			if s.pageNum < 1 {
				s.pageNum = 1
			}
			display.NewPagination(c, &s.pageNum, 12).OnChanged(func(int) {})
		})
	})
}
