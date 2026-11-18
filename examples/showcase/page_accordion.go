package main

import (
	"github.com/egoist/mygo/ui"
	display "github.com/wuyangwang/my-ui/display"
)

// accordionPage demonstrates the Accordion component.
func (s *showcase) accordionPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "手风琴", "mygo 管开合与键盘，my-ui 给整组一张卡片外观", func() {
		s.card(c, t, func() {
			display.NewAccordion(c, func() {
				display.NewAccordionItem(c, "详情", &s.accordionOpen, func() { ui.Text(c, "这里是详情内容") })
			})
		})
	})
}
