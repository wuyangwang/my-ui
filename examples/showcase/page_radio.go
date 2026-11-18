package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// radioPage shows the radio buttons: the choices of a group, and their
// states.
func (s *showcase) radioPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "单选", "Radio：一组里只选一个，选中的那个拿着值", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				ui.Row(c).Gap(t.Space(4)).Children(func() {
					ctrlpkg.NewRadio(c.Key("rd small"), &s.size, "small", "小")
					ctrlpkg.NewRadio(c.Key("rd medium"), &s.size, "medium", "中")
					ctrlpkg.NewRadio(c.Key("rd large"), &s.size, "large", "大")
				})
				ui.Row(c).Gap(t.Space(4)).Children(func() {
					ctrlpkg.NewRadio(c.Key("rd off"), &s.size, "none", "不可选").Disabled(true)
				})
			})
		})
	})
}
