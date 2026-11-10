package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// selectPage is the Select's page: the options, the states and the
// placeholder.
func (s *showcase) selectPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "选项", "标签与值可以不同：绑定的是值，触发器显示的是标签", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				myui.NewSelect(c.Key("sel-lang"), &s.lang, []myui.SelectOption{
					{Label: "简体中文", Value: "zh-Hans"},
					{Label: "繁體中文", Value: "zh-Hant"},
					{Label: "English", Value: "en"},
					{Value: "日本語"},
					{Label: "한국어", Value: "ko", Disabled: true},
				}).Outline().Label("语言").Placeholder("选择语言").
					OnChanged(func(v string) { s.lastEvent = "chosen: " + v })
				myui.NewSelect(c.Key("sel-size"), &s.size, []myui.SelectOption{
					{Value: "小"}, {Value: "中"}, {Value: "大"},
				}).Outline().Label("尺寸")
				myui.NewSelect(c.Key("sel-disabled"), &s.size, []myui.SelectOption{
					{Value: "小"},
				}).Outline().Label("禁用").Disabled(true)
			})
		})
	})

	s.sectionNote(c, t, "变体与尺寸", "与 Input 同一套预设", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewSelect(c.Key("sel-native"), &s.size, []myui.SelectOption{{Value: "原生"}}).Label("Native")
				myui.NewSelect(c.Key("sel-ghost"), &s.size, []myui.SelectOption{{Value: "幽灵"}}).Ghost().Label("Ghost")
				myui.NewSelect(c.Key("sel-sm"), &s.size, []myui.SelectOption{{Value: "小"}}).Sm().Outline().Label("小")
				myui.NewSelect(c.Key("sel-lg"), &s.size, []myui.SelectOption{{Value: "大"}}).Lg().Outline().Label("大")
			})
		})
	})

	s.sectionNote(c, t, "事件", "选择变化时回调", func() {
		s.card(c, t, func() {
			ui.Text(c, "最近事件："+s.lastEvent).
				FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
		})
	})
}
