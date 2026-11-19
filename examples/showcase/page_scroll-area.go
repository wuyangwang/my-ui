package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// scrollAreaPage shows a thin scrollbar and the soft fades at the edges
// that still have content.
func (s *showcase) scrollAreaPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("scroll-area-demo")).Hidden(true)
	state := ui.Local(holder, "scroll-area-demo", func() ui.ScrollState { return ui.ScrollState{} })
	if state == nil {
		return
	}
	s.sectionNote(c, t, "ScrollArea", "带细滚动条与边缘淡出的滚动区；淡出按滚动状态即时计算", func() {
		s.card(c, t, func() {
			layoutpkg.NewScrollArea(c, state, func() {
				ui.Column(c).Gap(t.Space(1)).Padding(t.Space(2)).Children(func() {
					for i := 0; i < 20; i++ {
						ui.Row(c.Key(i)).Padding(t.Space(2), t.Space(3)).
							Radius(myui.Radius(t, myui.RadiusSm)).Background(t.Surface).Children(func() {
							ui.Text(c, fmt.Sprintf("消息 %d", i+1))
						})
					}
				})
			}).Width(360).Height(240)
		})
	})
}
