package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// decimalInputPage demonstrates the DecimalInput component.
func (s *showcase) decimalInputPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "精确小数", "字符串与步长交给 mygo，不做浮点舍入", func() {
		s.card(c, t, func() {
			formspkg.NewDecimalInput(c, &s.decimalValue, "0", "100", "0.01").Variant(myui.Outline).Width(180)
		})
	})
}
