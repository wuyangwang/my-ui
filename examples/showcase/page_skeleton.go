package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/display"
)

// skeletonPage demonstrates the Skeleton component on its own page.
func (s *showcase) skeletonPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "用途", "加载态先用占位条撑住布局，数据到了再换成真实内容", func() {
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
}
