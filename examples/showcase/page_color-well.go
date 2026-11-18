package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// colorWellPage demonstrates the ColorWell component.
func (s *showcase) colorWellPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "颜色井", "色块由 mygo 原生面板编辑，库给边框和圆角", func() {
		s.card(c, t, func() { ctrlpkg.NewColorWell(c, &s.pickedColor) })
	})
}
