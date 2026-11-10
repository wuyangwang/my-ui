package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// iconPage shows the icons of the set and the presets of the Icon.
func (s *showcase) iconPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "常用图标", "来自 icons 包（Lucide, ISC），未被调用的图标不进二进制", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).Wrap().AlignItems(ui.Center).Children(func() {
				for _, ic := range []struct {
					svg  *ui.SVG
					name string
				}{
					{icons.Save(), "save"},
					{icons.Search(), "search"},
					{icons.Settings(), "settings"},
					{icons.Trash(), "trash-2"},
					{icons.Download(), "download"},
					{icons.Upload(), "upload"},
					{icons.Plus(), "plus"},
					{icons.Close(), "x"},
					{icons.Check(), "check"},
					{icons.Warning(), "triangle-alert"},
					{icons.Info(), "info"},
					{icons.Loading(), "loader-circle"},
					{icons.User(), "user"},
					{icons.Mail(), "mail"},
					{icons.Calendar(), "calendar"},
					{icons.Folder(), "folder"},
				} {
					myui.NewIcon(c.Key("icon "+ic.name), ic.svg).Md().Label(ic.name)
				}
			})
		})
	})

	s.sectionNote(c, t, "尺寸与语义色", "尺寸档为 14 / 20 / 24 DIP；变体给图标上语义色", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				myui.NewIcon(c.Key("icon-sm"), icons.Star()).Sm().Label("小")
				myui.NewIcon(c.Key("icon-md"), icons.Star()).Md().Label("中")
				myui.NewIcon(c.Key("icon-lg"), icons.Star()).Lg().Label("大")
				myui.NewIcon(c.Key("icon-accent"), icons.Star()).Primary().Md().Label("强调")
				myui.NewIcon(c.Key("icon-danger"), icons.Star()).Destructive().Md().Label("危险")
			})
		})
	})
}
