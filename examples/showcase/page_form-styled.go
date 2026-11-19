package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// formStyledPage shows the my-ui styled form container. It keeps mygo's
// ui.Form signature, so the fields keep mygo's aligned-label layout while
// the frame and spacing come from the library.
func (s *showcase) formStyledPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("form-styled-demo")).Hidden(true)
	values := ui.Local(holder, "form-styled-demo", func() struct{ name, email string } {
		return struct{ name, email string }{}
	})
	if values == nil {
		return
	}
	s.sectionNote(c, t, "Form", "带 my-ui 外观的表单容器：标签仍由 mygo 对齐，表面与间距取自主题", func() {
		s.card(c, t, func() {
			formspkg.NewFormWith(c, myui.Outline, func() {
				ui.Field(c, "姓名", func() { ui.TextInput(c, &values.name) })
				ui.Field(c, "邮箱", func() { ui.TextInput(c, &values.email) })
			})
		})
	})
}
