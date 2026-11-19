package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// gridPage shows both the plain grid and the responsive auto-fitting grid.
func (s *showcase) gridPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("grid-demo")).Hidden(true)
	state := ui.Local(holder, "grid-demo", func() ui.GridState { return ui.GridState{} })
	if state == nil {
		return
	}
	s.sectionNote(c, t, "Grid", "等宽网格：固定列数与按宽度自动增减列两种用法", func() {
		s.card(c, t, func() {
			layoutpkg.NewGrid(c).Children(func() {
				for i := 0; i < 3; i++ {
					tile(c, t, fmt.Sprintf("固定 %d", i+1))
				}
			})
		})
		s.card(c, t, func() {
			layoutpkg.AutoGrid(c, state, 6, 120, 96, func(i int) {
				tile(c, t, fmt.Sprintf("自适应 %d", i+1))
			})
		})
	})
}

func tile(c *ui.Context, t *ui.Theme, label string) {
	ui.Column(c).Gap(t.Space(1)).Padding(t.Space(3)).
		Radius(myui.Radius(t, myui.RadiusLg)).Background(t.Surface).Children(func() {
		ui.Text(c, label).Bold()
		ui.Text(c, "随容器宽度增减列数。").TextColor(t.TextMuted)
	})
}
