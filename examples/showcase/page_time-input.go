package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// timeInputPage demonstrates the TimeInput component.
func (s *showcase) timeInputPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "时间", "Variant/Size/Width 与 DatePicker 一致", func() {
		s.card(c, t, func() { formspkg.NewTimeInput(c, &s.pickedTime).Variant(myui.Outline).Width(160) })
	})
}
