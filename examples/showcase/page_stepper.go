package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// stepperPage demonstrates the Stepper component.
func (s *showcase) stepperPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "步进", "Variant/Size/Width 与 Input 同一套词表", func() {
		s.card(c, t, func() { ctrlpkg.NewStepper(c, &s.stepValue, 0, 10, 1).Variant(myui.Outline).Width(160) })
	})
}
