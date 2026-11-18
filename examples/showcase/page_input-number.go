package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/forms"
)

// inputNumberPage demonstrates the InputNumber component on its own page.
func (s *showcase) inputNumberPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "数字输入", "min/max/step 由 mygo 处理，Variant/Size/Width 由库统一", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				forms.NewInputNumber(c, &s.amount, 0, 100, 1).Variant(myui.Outline).Width(180)
				forms.NewInputNumber(c, &s.ratio, 0, 1, 0.05).Size(myui.SizeSm).Width(180)
			})
		})
	})
}
