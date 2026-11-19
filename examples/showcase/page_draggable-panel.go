package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// draggablePanelPage shows an inspector that resizes by dragging its seam.
func (s *showcase) draggablePanelPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("draggable-panel-demo")).Hidden(true)
	st := ui.Local(holder, "draggable-panel-demo", func() struct {
		size     float32
		expanded bool
	} {
		return struct {
			size     float32
			expanded bool
		}{size: 220, expanded: true}
	})
	if st == nil {
		return
	}
	s.sectionNote(c, t, "DraggablePanel", "拖接缝改宽度，折叠后留下可点开的窄条；尺寸与状态由 app 持有", func() {
		s.card(c, t, func() {
			frame := ui.Row(c).FillWidth().Height(t.Space(60)).
				AlignItems(ui.Stretch).Clip().
				Border(1, t.Border).Radius(myui.Radius(t, myui.RadiusLg))
			frame.Children(func() {
				ui.Column(c).Grow(1).Padding(t.Space(4)).Children(func() {
					ui.Text(c, "会话").Bold()
					ui.Text(c, "拖动右侧接缝改变面板宽度。").TextColor(t.TextMuted)
				})
				layoutpkg.NewDraggablePanel(c, &st.size, &st.expanded,
					layoutpkg.DraggablePanelProps{MinSize: 120, MaxSize: 320}, func() {
						ui.Column(c).Padding(t.Space(4)).Gap(t.Space(2)).Children(func() {
							ui.Text(c, "详情").Bold()
							ui.Text(c, "参与者、文件与置顶消息。").TextColor(t.TextMuted)
						})
					})
			})
		})
	})
}
