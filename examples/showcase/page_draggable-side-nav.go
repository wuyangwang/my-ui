package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// draggableSideNavPage shows a workspace rail that collapses to a slim
// strip. The builders receive the SideNavState and fade their rows with
// its Collapse value.
func (s *showcase) draggableSideNavPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("draggable-side-nav-demo")).Hidden(true)
	st := ui.Local(holder, "draggable-side-nav-demo", func() struct {
		width    float32
		expanded bool
	} {
		return struct {
			width    float32
			expanded bool
		}{width: 260, expanded: true}
	})
	if st == nil {
		return
	}
	s.sectionNote(c, t, "DraggableSideNav", "拖边缘改宽度，拖到三分之一以下收成轨道；头部与正文都收到 SideNavState", func() {
		s.card(c, t, func() {
			frame := ui.Row(c).FillWidth().Height(t.Space(80)).
				Clip().Border(1, t.Border).Radius(myui.Radius(t, myui.RadiusLg))
			frame.Children(func() {
				layoutpkg.NewDraggableSideNav(c, &st.width, &st.expanded,
					layoutpkg.DraggableSideNavProps{
						MinWidth: 96, MaxWidth: 320,
						Header: func(nav layoutpkg.SideNavState) {
							ui.Row(c).Padding(t.Space(3)).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
								ui.Icon(c, icons.Users()).Size(myui.IconSizeMd, myui.IconSizeMd).
									TextColor(t.Accent)
								if nav.Expanded {
									ui.Text(c, "工作区").Bold()
								}
							})
						},
					}, func(nav layoutpkg.SideNavState) {
						ui.Column(c).Gap(t.Space(1)).Padding(t.Space(2)).Children(func() {
							for _, name := range []string{"首页", "集成", "社区"} {
								ui.Row(c).Padding(t.Space(2)).Children(func() {
									ui.Text(c, name).Opacity(1 - nav.Collapse)
								})
							}
						})
					})
				ui.Column(c).Grow(1).Padding(t.Space(6)).Children(func() {
					label := "当前：收起"
					if st.expanded {
						label = "当前：展开"
					}
					ui.Text(c, "主内容区").Bold()
					ui.Text(c, label).TextColor(t.TextMuted)
					toggle := ctrlpkg.NewButton(c, "切换侧栏").Ghost().Sm()
					if toggle.Clicked() {
						st.expanded = !st.expanded
					}
				})
			})
		})
	})
}
