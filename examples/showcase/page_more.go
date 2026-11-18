package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/display"
	"github.com/wuyangwang/my-ui/forms"
	"github.com/wuyangwang/my-ui/overlays"
)

// morePage demonstrates the next batch: breadcrumb, numeric input,
// search bar and hover card.
func (s *showcase) morePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "Breadcrumb", "当前项强调墨色，其余可点并写回索引", func() {
		s.card(c, t, func() {
			display.NewBreadcrumb(c, []string{"首页", "组件", "更多组件"}, &s.crumb)
		})
	})
	s.sectionNote(c, t, "InputNumber", "步进、范围与步长走 mygo 底座，外观用共享字段词表", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				forms.NewInputNumber(c, &s.amount, 0, 100, 1).Variant(myui.Outline).Width(180)
				forms.NewInputNumber(c, &s.ratio, 0, 1, 0.05).Size(myui.SizeSm).Width(180)
			})
		})
	})
	s.sectionNote(c, t, "SearchBar", "放大镜 + 清除，默认占满、Variant/Size 与 Input 一致", func() {
		s.card(c, t, func() {
			forms.NewSearchBar(c, &s.searchQuery).Placeholder("搜索仓库…").Width(280)
		})
	})
	s.sectionNote(c, t, "HoverCard", "指针停在触发项上展开，移到面板上保持", func() {
		s.card(c, t, func() {
			overlays.NewHoverCard(c, &s.hoverOpen, func() ui.Element {
				return ctrlpkg.NewButton(c, "悬停查看").Outline().Sm().Element
			}, func(panel ui.Element) {
				panel.Children(func() {
					ui.Text(c, "这里是浮层的内容").Bold()
					ui.Text(c, "内容必须挂在 panel 上，才会落在浮层里").TextColor(t.TextMuted)
				})
			})
		})
	})
}
