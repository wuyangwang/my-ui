package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/forms"
)

// datePickerPage demonstrates the DatePicker component on its own page.
func (s *showcase) datePickerPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "日期", "日历弹层由 mygo 提供，字段外观由库统一", func() {
		s.card(c, t, func() {
			forms.NewDatePicker(c, &s.pickedDate).Variant(myui.Outline).Width(180)
		})
	})
}
