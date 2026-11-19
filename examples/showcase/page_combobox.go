package main

import (
	"github.com/egoist/mygo/ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// comboBoxDemo holds the state of the showcase's combos. Each combo owns
// its own value and query pointers: sharing one query was the old bug the
// options-based API fixes.
var comboBoxDemo struct {
	groupValue, groupText     string
	nativeValue, nativeText   string
	outlineValue, outlineText string
	smValue, smText           string
	lgValue, lgText           string
}

// comboBoxPage is the ComboBox's page: the searchable drop-down.
func (s *showcase) comboBoxPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "可搜索", "Value 是选中的值，Text 才是输入的查询；清除时两者一起清空", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				formspkg.NewComboBox(c.Key("cb-lang"), formspkg.ComboOptions{
					Value: &s.lang,
					Text:  &s.query,
					Options: []formspkg.SelectOption{
						{Label: "简体中文", Value: "zh-Hans"},
						{Label: "繁體中文", Value: "zh-Hant"},
						{Label: "English", Value: "en"},
						{Label: "日本語", Value: "ja"},
						{Label: "한국어", Value: "ko", Disabled: true},
					},
					Clearable:   true,
					Label:       "语言",
					Placeholder: "搜索语言…",
					EmptyText:   "没有匹配的语言",
					OnChanged:   func(v string) { s.lastEvent = "chosen: " + v },
				}).Outline()

				formspkg.NewComboBox(c.Key("cb-group"), formspkg.ComboOptions{
					Value: &comboBoxDemo.groupValue,
					Text:  &comboBoxDemo.groupText,
					Options: []formspkg.SelectOption{
						{Group: "最近", Label: "English", Value: "en"},
						{Group: "最近", Label: "日本語", Value: "ja"},
						{Group: "更多", Label: "한국어", Value: "ko"},
					},
					Label:       "分组",
					Placeholder: "搜索…",
				}).Outline()
			})
		})
	})

	s.sectionNote(c, t, "变体与尺寸", "与 Input 同一套预设：外观画在输入框外的框上", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				formspkg.NewComboBox(c.Key("cb-native"), formspkg.ComboOptions{
					Value: &comboBoxDemo.nativeValue, Text: &comboBoxDemo.nativeText,
					Options: []formspkg.SelectOption{{Value: "原生"}},
					Label:   "Native", Placeholder: "Native",
				})
				formspkg.NewComboBox(c.Key("cb-outline"), formspkg.ComboOptions{
					Value: &comboBoxDemo.outlineValue, Text: &comboBoxDemo.outlineText,
					Options: []formspkg.SelectOption{{Value: "描边"}},
					Label:   "Outline", Placeholder: "Outline",
				}).Outline()
				formspkg.NewComboBox(c.Key("cb-sm"), formspkg.ComboOptions{
					Value: &comboBoxDemo.smValue, Text: &comboBoxDemo.smText,
					Options: []formspkg.SelectOption{{Value: "小"}},
					Label:   "小", Placeholder: "小",
				}).Sm().Outline()
				formspkg.NewComboBox(c.Key("cb-lg"), formspkg.ComboOptions{
					Value: &comboBoxDemo.lgValue, Text: &comboBoxDemo.lgText,
					Options: []formspkg.SelectOption{{Value: "大"}},
					Label:   "大", Placeholder: "大",
				}).Lg().Outline()
			})
		})
	})
}
