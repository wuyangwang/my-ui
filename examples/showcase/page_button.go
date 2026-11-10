package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// buttonPage is the Button's page: its variants first, then the sizes
// and the states, each in a card of its own.
func (s *showcase) buttonPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体", "同一基底，把鼠标放进卡里比色阶", func() {
		variants := []struct {
			name, cn, tip, label string
			preset               func(*myui.Button) *myui.Button
		}{
			{"Primary", "主要", "悬停渐深，压下更深", "保存更改", (*myui.Button).Primary},
			{"Secondary", "次要", "内置原样，零值即得", "取消", (*myui.Button).Secondary},
			{"Outline", "描边", "白底描边，悬停转淡", "浏览更多", (*myui.Button).Outline},
			{"Ghost", "幽灵", "裸底，悬停才现形", "轻触一下", (*myui.Button).Ghost},
			{"Destructive", "危险", "删除类动作", "删除项目", (*myui.Button).Destructive},
			{"Link", "链接", "文字态，悬停浮起", "阅读文档", (*myui.Button).Link},
		}
		for i := 0; i < len(variants); i += 2 {
			pair := variants[i : i+2]
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Start).Children(func() {
				for _, v := range pair {
					v := v
					ui.Column(c).Grow(1).Children(func() {
						s.presetCard(c, t, v.name, v.cn, v.tip, v.label, v.preset)
					})
				}
			})
		}
	})

	s.sectionNote(c, t, "尺寸", "内衬与字号只是 Size 档，与变体正交", func() {
		s.card(c, t, func() {
			sizes := []struct {
				label string
				pick  func(*myui.Button) *myui.Button
			}{{"小", (*myui.Button).Sm}, {"中", (*myui.Button).Md}, {"大", (*myui.Button).Lg}}
			for _, sz := range sizes {
				b := myui.NewButton(c.Key("size "+sz.label), sz.label).Primary()
				sz.pick(b)
			}
		})
	})

	s.sectionNote(c, t, "图标", "图标在文字之前或之后；只有图标时给出无障碍名；忙碌态转圈且不响应点击", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				myui.NewButton(c.Key("btn-icon"), "保存").Icon(icons.Save()).Primary()
				myui.NewButton(c.Key("btn-endicon"), "导出").EndIcon(icons.Download()).Outline()
				myui.NewButton(c.Key("btn-icon-only"), "").Icon(icons.Settings()).Label("设置")
				myui.NewButton(c.Key("btn-loading"), "保存").Loading()
			})
		})
	})

	s.sectionNote(c, t, "状态", "不便操作的样子随主题的暗淡走", func() {
		s.card(c, t, func() {
			s.cardHead(c, t, "禁用", "暗淡，点击不报")
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				myui.NewButton(c.Key("disabled"), "保存更改").Primary().Disabled(true)
				myui.NewButton(c.Key("disabled-outline"), "浏览更多").Outline().Disabled(true)
				myui.NewButton(c.Key("disabled-ghost"), "轻触一下").Ghost().Disabled(true)
			})
		})
		s.card(c, t, func() {
			s.cardHead(c, t, "链接", "文字态预设，点击逻辑照收")
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				myui.NewButton(c.Key("link"), "回复这条评论").Link()
				myui.NewButton(c.Key("link-sm"), "取消设置").Link().Sm()
			})
		})
	})
}

// presetCard is one variant's cell: the preset names, a tip, and the
// widget itself under them.
func (s *showcase) presetCard(c *ui.Context, t *ui.Theme, name, cn, tip, label string, preset func(*myui.Button) *myui.Button) {
	s.card(c, t, func() {
		s.cardHead(c, t, cn, name+" · "+tip)
		b := myui.NewButton(c.Key("button "+name), label)
		preset(b)
	})
}
