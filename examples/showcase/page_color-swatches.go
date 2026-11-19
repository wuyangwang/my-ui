package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// colorSwatchesPage demonstrates the ColorSwatches component.
func (s *showcase) colorSwatchesPage(c *ui.Context, t *ui.Theme) {
	swatches := []ctrlpkg.ColorSwatch{
		{Color: ui.Hex("#ef4444"), Label: "红"},
		{Color: ui.Hex("#f59e0b"), Label: "橙"},
		{Color: ui.Hex("#22c55e"), Label: "绿"},
		{Color: ui.Hex("#3b82f6"), Label: "蓝"},
		{Color: ui.Transparent, Label: "透明"},
	}
	s.sectionNote(c, t, "ColorSwatches", "一组色块选一个颜色，方向键在色块间移动", func() {
		s.demoCard(c, t, "Circle", "圆形", "默认圆点，选中带勾", func() {
			ctrlpkg.NewColorSwatches(c, &s.pickedColor, swatches...)
		})
		s.demoCard(c, t, "Square", "方形", "方角色块，尺寸更大", func() {
			ctrlpkg.NewColorSwatches(c, &s.pickedColor, swatches...).Square(true).Size(32)
		})
	})
}
