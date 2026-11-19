package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// blockPage shows the block surfaces and its clickable state.
func (s *showcase) blockPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "表面", "淡底、带框与无面", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).Children(func() {
				for _, v := range []struct {
					name    string
					variant int
				}{{"Filled", 0}, {"Outline", 1}, {"Ghost", 2}} {
					b := displaypkg.NewBlock(c.Key("block-"+v.name), displaypkg.BlockOptions{})
					switch v.name {
					case "Outline":
						b.Outline()
					case "Ghost":
						b.Ghost()
					}
					b.Width(140).Children(func() {
						ui.Text(c, v.name).FontWeight(600)
						ui.Text(c, "块面内容").FontSize(12).TextColor(t.TextMuted)
					})
				}
			})
		})
	})
	s.sectionNote(c, t, "可点击", "点击写回 lastAction，悬停只改底色", func() {
		s.card(c, t, func() {
			displaypkg.NewBlock(c, displaypkg.BlockOptions{
				Clickable: true, Label: "clickable block",
				OnChanged: func() { s.lastAction = "Block clicked" },
			}).Width(180).Children(func() {
				ui.Text(c, "Click me").FontWeight(600)
				ui.Text(c, s.lastAction).FontSize(12).TextColor(t.TextMuted)
			})
		})
	})
}
