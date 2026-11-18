package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// buttonPage is the Button's page: its variants first, then the sizes
// and the states, each in a card of its own.
func (s *showcase) buttonPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体", "同一基底，把鼠标放进卡里比色阶", func() {
		variants := []struct {
			name, cn, tip, label string
			preset               func(*ctrlpkg.Button) *ctrlpkg.Button
		}{
			{"Default", "默认", "内置原样，零值即得", "取消", (*ctrlpkg.Button).Default},
			{"Secondary", "次要", "弱填充面，无描边", "草稿", (*ctrlpkg.Button).Secondary},
			{"Primary", "主要", "悬停渐深，压下更深", "保存更改", (*ctrlpkg.Button).Primary},
			{"Outline", "描边", "白底描边，悬停转淡", "浏览更多", (*ctrlpkg.Button).Outline},
			{"Ghost", "幽灵", "裸底，悬停才现形", "轻触一下", (*ctrlpkg.Button).Ghost},
			{"Destructive", "危险", "删除类动作", "删除项目", (*ctrlpkg.Button).Destructive},
			{"Link", "链接", "文字态，悬停浮起", "阅读文档", (*ctrlpkg.Button).Link},
		}
		for i := 0; i < len(variants); i += 3 {
			row := variants[i:min(i+3, len(variants))]
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Stretch).Children(func() {
				for _, v := range row {
					v := v
					ui.Column(c).Grow(1).Children(func() {
						s.presetCard(c, t, v.name, v.cn, v.tip, v.label, v.preset)
					})
				}
			})
		}
	})

	s.sectionNote(c, t, "语义色", "成功/信息/警告各有自己的颜色，明暗各一套", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).Children(func() {
				ctrlpkg.NewButton(c.Key("btn-success"), "成功").Success()
				ctrlpkg.NewButton(c.Key("btn-info"), "信息").Info()
				ctrlpkg.NewButton(c.Key("btn-warning"), "警告").Warning()
			})
		})
	})

	s.sectionNote(c, t, "尺寸", "内衬与字号只是 Size 档，与变体正交", func() {
		s.card(c, t, func() {
			sizes := []struct {
				label string
				pick  func(*ctrlpkg.Button) *ctrlpkg.Button
			}{
				{"最小", (*ctrlpkg.Button).Xs},
				{"小", (*ctrlpkg.Button).Sm},
				{"中", (*ctrlpkg.Button).Md},
				{"大", (*ctrlpkg.Button).Lg},
			}
			for _, sz := range sizes {
				b := ctrlpkg.NewButton(c.Key("size "+sz.label), sz.label).Primary()
				sz.pick(b)
			}
		})
	})

	s.sectionNote(c, t, "形态", "Shape 决定圆角：胶囊、圆形", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				ctrlpkg.NewButton(c.Key("btn-rect"), "直角").Primary()
				ctrlpkg.NewButton(c.Key("btn-round"), "胶囊").Primary().Shape(myui.ShapeRound)
				ctrlpkg.NewButton(c.Key("btn-circle"), "圆形").Primary().Shape(myui.ShapeCircle)
			})
		})
	})

	s.sectionNote(c, t, "图标", "图标在文字之前或之后；只有图标时给出无障碍名；忙碌态转圈且不响应点击", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				ctrlpkg.NewButton(c.Key("btn-icon"), "保存").Icon(icons.Save()).Primary()
				ctrlpkg.NewButton(c.Key("btn-endicon"), "导出").EndIcon(icons.Download()).Outline()
				ctrlpkg.NewButton(c.Key("btn-icon-only"), "").Icon(icons.Settings()).Label("设置")
				ctrlpkg.NewButton(c.Key("btn-loading"), "保存").Loading()
			})
		})
	})

	s.sectionNote(c, t, "状态", "不便操作的样子随主题的暗淡走", func() {
		s.card(c, t, func() {
			s.cardHead(c, t, "禁用", "暗淡，点击不报")
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				ctrlpkg.NewButton(c.Key("disabled"), "保存更改").Primary().Disabled(true)
				ctrlpkg.NewButton(c.Key("disabled-outline"), "浏览更多").Outline().Disabled(true)
				ctrlpkg.NewButton(c.Key("disabled-ghost"), "轻触一下").Ghost().Disabled(true)
			})
		})
		s.card(c, t, func() {
			s.cardHead(c, t, "链接", "文字态预设，点击逻辑照收")
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				ctrlpkg.NewButton(c.Key("link"), "回复这条评论").Link()
				ctrlpkg.NewButton(c.Key("link-sm"), "取消设置").Link().Sm()
			})
		})
	})
}

// presetCard is one variant's cell: the names on the left, the widget
// itself on the right, so a gallery of them reads as a table of examples
// instead of a grid of near-empty boxes.
func (s *showcase) presetCard(c *ui.Context, t *ui.Theme, name, cn, tip, label string, preset func(*ctrlpkg.Button) *ctrlpkg.Button) {
	s.demoCard(c, t, name, cn, tip, func() {
		b := ctrlpkg.NewButton(c.Key("button "+name), label)
		preset(b)
	})
}
