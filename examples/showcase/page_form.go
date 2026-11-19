package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// formPage shows the form field layer and the password field.
func (s *showcase) formPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "字段", "标签、控件、说明与错误由 Field 排成一体；点标签聚焦控件；值只传一次，规则进 options", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(4)).Children(func() {
				// The widget is returned as-is: the field finds its
				// Control(), so there is no trailing .Control() to forget.
				formspkg.NewField(c.Key("field-name"), "昵称", &s.name, func() *formspkg.Input {
					return formspkg.NewInput(c.Key("in-nickname"), &s.name).Outline()
				}, formspkg.FieldOptions[string]{
					Description: "显示在资料卡上",
				})

				// Rules and the value travel together in one call — no
				// second Validate that could be forgotten.
				formspkg.NewField(c.Key("field-email"), "邮箱", &s.email, func() *formspkg.Input {
					return formspkg.NewInput(c.Key("in-email"), &s.email).Outline()
				}, formspkg.FieldOptions[string]{
					Description: "我们只用它发通知",
					Required:    true,
					Rules: []myui.Rule{
						myui.Required("请填写邮箱"),
						myui.Email("邮箱格式不对"),
					},
				})

				formspkg.NewField(c.Key("field-size"), "尺寸", &s.size, func() *formspkg.Select {
					return formspkg.NewSelect(c.Key("sel-field"), &s.size, []formspkg.SelectOption{
						{Value: "小"}, {Value: "中"}, {Value: "大"},
					}).Outline()
				}, formspkg.FieldOptions[string]{
					Description: "下拉也是控件",
				})
			})
		})
	})

	s.sectionNote(c, t, "密码", "隐藏所输入的内容，眼睛按钮切换明暗", func() {
		s.card(c, t, func() {
			formspkg.NewPasswordInput(c.Key("in-password"), &s.password).Outline().Label("密码").
				Placeholder("输入密码")
		})
	})
}
