package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// inputPage is the Input's page: variants, sizes and states.
func (s *showcase) inputPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体", "边框与底色随 Variant 预设变化", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				myui.NewInput(c.Key("in-native"), &s.name).Placeholder("Native / 默认").Label("Native")
				myui.NewInput(c.Key("in-outline"), &s.email).Outline().Placeholder("Outline").Label("Outline")
				myui.NewInput(c.Key("in-ghost"), &s.name).Ghost().Placeholder("Ghost").Label("Ghost")
				myui.NewInput(c.Key("in-primary"), &s.name).Primary().Placeholder("Primary（强调边框）").Label("Primary")
				myui.NewInput(c.Key("in-danger"), &s.name).Destructive().Placeholder("Destructive").Label("Destructive")
			})
		})
	})

	s.sectionNote(c, t, "尺寸", "内衬与字号沿用 Size 档，与变体正交", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewInput(c.Key("in-sm"), &s.name).Sm().Outline().Placeholder("小尺寸").Label("小")
				myui.NewInput(c.Key("in-md"), &s.name).Md().Outline().Placeholder("中尺寸（默认）").Label("中")
				myui.NewInput(c.Key("in-lg"), &s.name).Lg().Outline().Placeholder("大尺寸").Label("大")
			})
		})
	})

	s.sectionNote(c, t, "状态", "禁用与只读遵循 ui 输入约定", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewInput(c.Key("in-disabled"), &s.disabled).Outline().Placeholder("不可编辑").Label("禁用").Disabled(true)
				s.readonly = "只读内容"
				myui.NewInput(c.Key("in-readonly"), &s.readonly).Outline().Label("只读").ReadOnly(true)
			})
		})
	})

	s.sectionNote(c, t, "组合", "前缀图标（如搜索框的放大镜）、后缀、清除按钮与错误提示都由库排版", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				myui.NewInput(c.Key("in-search"), &s.query).Outline().Label("搜索").
					PrefixIcon(icons.Search()).Placeholder("搜索…").Clearable()
				myui.NewInput(c.Key("in-clear"), &s.name).Outline().Label("可清除").Placeholder("输入后可清除").
					Clearable().OnCleared(func() { s.lastEvent = "cleared" })
				myui.NewInput(c.Key("in-affix"), &s.email).Outline().Label("带前后缀").
					Prefix(func() { ui.Text(c, "https://") }).
					Suffix(func() { ui.Text(c, ".com") })
				myui.NewInput(c.Key("in-error"), &s.email).Outline().Label("错误态").
					Invalid(true).ErrorText("邮箱格式不对")
			})
		})
	})

	s.sectionNote(c, t, "事件", "值变化、失焦、清除与回车都有回调", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewInput(c.Key("in-events"), &s.name).Outline().Label("带事件").
					OnChanged(func(v string) { s.lastEvent = "changed: " + v }).
					OnBlurred(func() { s.lastEvent = "blurred" }).
					OnSubmit(func() { s.lastEvent = "submitted" })
				ui.Text(c, "最近事件："+s.lastEvent).
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			})
		})
	})
}
