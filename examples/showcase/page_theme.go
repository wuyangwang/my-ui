package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// swatches is one palette switch entry: the id and its label.
var swatches = []struct {
	id     myui.PaletteID
	button string
}{
	{myui.PaletteNone, "系统"},
	{myui.PaletteBlue, "蓝"},
	{myui.PaletteViolet, "紫"},
	{myui.PaletteRose, "玫红"},
	{myui.PaletteOrange, "橙"},
	{myui.PaletteEmerald, "翠绿"},
	{myui.PaletteTeal, "青"},
	{myui.PaletteSky, "天蓝"},
	{myui.PalettePink, "粉"},
}

// themePage is the theming page: appearance, palette, and both
// appearances side by side as subtree previews.
func (s *showcase) themePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "外观与色板", "切换下一帧生效，整窗换装", func() {
		s.card(c, t, func() {
			s.cardHead(c, t, "明暗", "跟随系统，或强制一边")
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				choices := []struct {
					value  myui.Appearance
					button string
				}{
					{myui.AppearanceSystem, "跟随系统"},
					{myui.AppearanceLight, "浅色"},
					{myui.AppearanceDark, "深色"},
				}
				for _, ch := range choices {
					active := s.appearance == ch.value
					if pickVariant(myui.NewButton(c.Key("appearance "+ch.button), ch.button), active).
						Clicked() {
						s.appearance = ch.value
					}
				}
			})
		})

		s.card(c, t, func() {
			s.cardHead(c, t, "色板", "系统 · 8 个预设的强调色三元")
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				for _, p := range swatches {
					active := s.palette == p.id
					if pickVariant(myui.NewButton(c.Key("palette "+p.button), p.button), active).
						Clicked() {
						s.palette = p.id
					}
				}
			})
		})
	})

	s.sectionNote(c, t, "子树预览", "ui.Themed 里的另一外观，卡片独立换装", func() {
		ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Start).Children(func() {
			pane(c, t, s.palette, myui.AppearanceLight, "浅色子树")
			pane(c, t, s.palette, myui.AppearanceDark, "深色子树")
		})
	})
}
