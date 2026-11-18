package main

import (
	"github.com/egoist/mygo/ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// comboBoxPage is the ComboBox's page: the searchable drop-down.
func (s *showcase) comboBoxPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "可搜索", "触发器就是输入框：输入即过滤选项，选中即取替文字", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				formspkg.NewComboBox(c.Key("cb-lang"), &s.query, []formspkg.SelectOption{
					{Label: "简体中文", Value: "zh-Hans"},
					{Label: "繁體中文", Value: "zh-Hant"},
					{Label: "English", Value: "en"},
					{Label: "日本語", Value: "ja"},
					{Label: "한국어", Value: "ko", Disabled: true},
				}).Outline().Label("语言").Placeholder("搜索语言…").
					OnChanged(func(v string) { s.lastEvent = "chosen: " + v })

				formspkg.NewComboBox(c.Key("cb-group"), &s.query, []formspkg.SelectOption{
					{Group: "最近", Label: "English", Value: "en"},
					{Group: "最近", Label: "日本語", Value: "ja"},
					{Group: "更多", Label: "한국어", Value: "ko"},
				}).Outline().Label("分组").Placeholder("搜索…")
			})
		})
	})

	s.sectionNote(c, t, "变体与尺寸", "与 Input 同一套预设：外观画在输入框上", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				formspkg.NewComboBox(c.Key("cb-native"), &s.query, []formspkg.SelectOption{{Value: "原生"}}).Label("Native").Placeholder("Native")
				formspkg.NewComboBox(c.Key("cb-outline"), &s.query, []formspkg.SelectOption{{Value: "描边"}}).Outline().Label("Outline").Placeholder("Outline")
				formspkg.NewComboBox(c.Key("cb-sm"), &s.query, []formspkg.SelectOption{{Value: "小"}}).Sm().Outline().Label("小").Placeholder("小")
				formspkg.NewComboBox(c.Key("cb-lg"), &s.query, []formspkg.SelectOption{{Value: "大"}}).Lg().Outline().Label("大").Placeholder("大")
			})
		})
	})
}
