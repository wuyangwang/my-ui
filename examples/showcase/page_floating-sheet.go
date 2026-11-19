package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// floatingSheetPage shows a bottom sheet that snaps to a few heights.
func (s *showcase) floatingSheetPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("floating-sheet-demo")).Hidden(true)
	st := ui.Local(holder, "floating-sheet-demo", func() struct{ open bool } {
		return struct{ open bool }{open: true}
	})
	if st == nil {
		return
	}
	s.sectionNote(c, t, "FloatingSheet", "容器底部的浮层：拖头部改高度，松开吸附到最近的档位", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, "拖住下面的标题栏，把便签拉高或收低。").
					FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
			})
			layoutpkg.NewFloatingSheet(c, &st.open, layoutpkg.FloatingSheetProps{
				Title:      "便签",
				SnapPoints: []float32{0.3, 0.6, 0.9},
			}, func() {
				ui.Scroll(c).Grow(1).Children(func() {
					ui.Column(c).Gap(t.Space(1)).Padding(t.Space(3)).Children(func() {
						for i := 0; i < 10; i++ {
							ui.Text(c.Key(i), fmt.Sprintf("第 %d 条记录", i+1)).
								TextColor(t.TextMuted)
						}
					})
				})
			})
		})
	})
}
