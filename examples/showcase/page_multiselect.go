package main

import (
	"strings"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// multiSelectPage is the MultiSelect's page: the field of tags.
func (s *showcase) multiSelectPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "多选", "选中的项以芯片呈现，各带一个移除按钮；输入即过滤还没选的选项", func() {
		s.card(c, t, func() {
			myui.NewMultiSelect(c.Key("ms-lang"), &s.langs, []myui.SelectOption{
				{Label: "简体中文", Value: "zh-Hans"},
				{Label: "繁體中文", Value: "zh-Hant"},
				{Label: "English", Value: "en"},
				{Label: "日本語", Value: "ja"},
			}).Outline().Label("语言").Placeholder("输入以添加…").
				OnChanged(func(v []string) { s.lastEvent = "chosen: " + strings.Join(v, ", ") })
		})
	})

	s.sectionNote(c, t, "变体与尺寸", "与 Input 同一套预设", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewMultiSelect(c.Key("ms-native"), &s.langs, []myui.SelectOption{{Value: "原生"}}).Label("Native")
				myui.NewMultiSelect(c.Key("ms-outline"), &s.langs, []myui.SelectOption{{Value: "描边"}}).Outline().Label("Outline")
				myui.NewMultiSelect(c.Key("ms-sm"), &s.langs, []myui.SelectOption{{Value: "小"}}).Sm().Outline().Label("小")
			})
		})
	})
}
