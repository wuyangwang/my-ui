package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/display"
	"github.com/wuyangwang/my-ui/forms"
	"github.com/wuyangwang/my-ui/icons"
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
	s.sectionNote(c, t, "菜单", "mygo MenuBase / ContextMenuBase 加 shadcn 风格的面板与条目", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(6)).AlignItems(ui.Center).Children(func() {
				overlays.NewDropdownMenu(c, &s.menuOpen, func() ui.Element {
					return ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
						ui.Text(c, "操作")
						ui.Icon(c, icons.ChevronDown()).Size(myui.IconSizeXs, myui.IconSizeXs)
					})
				}, []overlays.MenuItem{
					{Label: "复制", Icon: icons.Copy(), OnClick: func() { s.lastAction = "复制" }},
					{Label: "删除", Icon: icons.Trash2(), OnClick: func() { s.lastAction = "删除" }},
				})
				target := ui.Box(c).Padding(t.Space(2), t.Space(4)).Radius(t.Radius).
					Background(t.Surface).Children(func() { ui.Text(c, "右键我") })
				overlays.NewContextMenu(c, &s.contextOpen, target, []overlays.MenuItem{
					{Label: "重命名", OnClick: func() { s.lastAction = "重命名" }},
				})
				ui.Text(c, "最近："+s.lastAction).TextColor(t.TextMuted)
			})
		})
	})
	s.sectionNote(c, t, "AutoComplete / DatePicker", "都包 mygo 的输入底座，套共享 Variant/Size/Width", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				forms.NewAutoComplete(c, &s.autoValue, []string{"苹果", "香蕉", "橘子"}).
					Variant(myui.Outline).Placeholder("输入或选择水果").Width(280)
				forms.NewDatePicker(c, &s.pickedDate).Variant(myui.Outline).Width(180)
			})
		})
	})
	s.sectionNote(c, t, "Carousel / Tag", "轮播索引由 app 持有；标签是轻量元信息标记", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				display.NewCarousel(c, &s.slide, 3, func(i int) {
					ui.Row(c).Height(t.Space(16)).AlignItems(ui.Center).Children(func() {
						ui.Text(c, []string{"第一张", "第二张", "第三张"}[i]).Bold()
					})
				})
				ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
					display.NewTag(c, display.TagOptions{
						Label: "新功能", Icon: icons.Sparkles(),
						Closed: &s.tagClosed,
					}).Variant(myui.Success)
					display.NewTag(c, display.TagOptions{Label: "草稿"}).Variant(myui.Secondary)
					ui.Text(c, "关闭后标签仍会保留，演示不移除数据").TextColor(t.TextMuted)
				})
			})
		})
	})
	s.sectionNote(c, t, "Spin / Table", "旋转图标走框架动效；表格由 mygo ListState 管排序与选择", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				display.NewSpin(c).Variant(myui.Primary)
				display.NewSpin(c).Size(myui.SizeSm).Variant(myui.Success)
				ui.Text(c, "加载中…").TextColor(t.TextMuted)
			})
			display.NewTable(c, &s.tableState, []ui.TableColumn{
				{Title: "名称", Width: 160}, {Title: "状态", Width: 120},
			}, 2, func(row, col int) {
				if col == 0 {
					ui.Text(c, []string{"组件 A", "组件 B"}[row])
				} else {
					ui.Text(c, []string{"正常", "草稿"}[row])
				}
			})
		})
	})
}
