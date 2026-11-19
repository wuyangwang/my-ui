package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
	layoutpkg "github.com/wuyangwang/my-ui/layout"
)

// sideNavPage shows an app's icon rail. The active tab is kept in a
// ui.Local of a hidden holder.
func (s *showcase) sideNavPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("side-nav-demo")).Hidden(true)
	active := ui.Local(holder, "side-nav-demo", func() string { return "chat" })
	if active == nil {
		return
	}
	tabs := []struct {
		key  string
		icon func() *ui.SVG
	}{
		{"chat", icons.MessageSquare},
		{"discover", icons.Compass},
		{"files", icons.Folder},
	}
	s.sectionNote(c, t, "SideNav", "应用左侧的图标轨道：头像在上，主操作居中，帮助与设置沉底", func() {
		s.card(c, t, func() {
			frame := ui.Row(c).FillWidth().Height(t.Space(100)).
				Clip().Border(1, t.Border).Radius(myui.Radius(t, myui.RadiusLg))
			frame.Children(func() {
				layoutpkg.NewSideNav(c, layoutpkg.SideNavProps{
					Avatar: func() {
						ui.Row(c).Size(t.Space(8), t.Space(8)).
							Radius(t.Space(8) / 2).Background(t.Accent).
							AlignItems(ui.Center).Justify(ui.Center).Children(func() {
							ui.Text(c, "A").TextColor(t.AccentText)
						})
					},
					TopActions: func() {
						for _, tab := range tabs {
							tab := tab
							b := ctrlpkg.NewButton(c, "").Icon(tab.icon()).Label(tab.key).Ghost()
							if *active == tab.key {
								b.Background(t.Accent.Alpha(0.12)).TextColor(t.Accent)
							}
							if b.Clicked() {
								*active = tab.key
							}
						}
					},
					BottomActions: func() {
						ctrlpkg.NewButton(c, "").Icon(icons.CircleHelp()).Label("帮助").Ghost()
						ctrlpkg.NewButton(c, "").Icon(icons.Settings()).Label("设置").Ghost()
					},
				})
				ui.Column(c).Grow(1).Center().Children(func() {
					ui.Text(c, "当前标签："+*active).TextColor(t.TextMuted)
				})
			})
		})
	})
}
