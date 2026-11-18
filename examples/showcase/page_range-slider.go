package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// rangeSliderPage demonstrates the RangeSlider component.
func (s *showcase) rangeSliderPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "范围", "两个旋钮分别写回 low 与 high", func() {
		s.card(c, t, func() { ctrlpkg.NewRangeSlider(c, &s.low, &s.high, 0, 100, 1) })
	})
}
