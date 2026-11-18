package main

import (
	"github.com/egoist/mygo/ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// calendarPage demonstrates the Calendar component.
func (s *showcase) calendarPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "日历", "今天有强调色边框，选中的日期为实底", func() {
		s.card(c, t, func() { formspkg.NewCalendar(c, &s.pickedDate) })
	})
}
