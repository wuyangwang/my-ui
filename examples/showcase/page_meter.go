package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// meterPage demonstrates the Meter component.
func (s *showcase) meterPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "量度", "值在区间内的位置由 mygo 着色", func() {
		s.card(c, t, func() { ctrlpkg.NewMeter(c, 65, 0, 100, nil) })
	})
}
