package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// colorPickerPage demonstrates the ColorPicker component.
func (s *showcase) colorPickerPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "颜色", "点击色块打开 mygo 的颜色选择器", func() {
		s.card(c, t, func() { ctrlpkg.NewColorPicker(c, &s.pickedColor) })
	})
}
