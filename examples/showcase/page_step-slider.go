package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// stepSliderPage demonstrates the StepSlider component.
func (s *showcase) stepSliderPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "步进滑杆", "步长由 app 给出，外观跟随主题", func() {
		s.card(c, t, func() { ctrlpkg.NewStepSlider(c, &s.stepValue, 0, 10, 1) })
	})
}
