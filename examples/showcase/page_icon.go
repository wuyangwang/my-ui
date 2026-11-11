package main

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
	"github.com/wuyangwang/my-ui/icons/all"
	brandall "github.com/wuyangwang/my-ui/icons/brand/all"
)

// iconPage shows the icons of the two sets: every one of them, its name
// on the pointer, and a click that copies the name.
func (s *showcase) iconPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "尺寸与语义色", "尺寸档为 14 / 20 / 24 DIP；变体给图标上语义色", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				myui.NewIcon(c.Key("icon-sm"), icons.Star()).Sm().Label("小")
				myui.NewIcon(c.Key("icon-md"), icons.Star()).Md().Label("中")
				myui.NewIcon(c.Key("icon-lg"), icons.Star()).Lg().Label("大")
				myui.NewIcon(c.Key("icon-accent"), icons.Star()).Primary().Md().Label("强调")
				myui.NewIcon(c.Key("icon-danger"), icons.Star()).Destructive().Md().Label("危险")
			})
		})
	})

	shown, brands := filterIcons(s.iconQuery)
	s.sectionNote(c, t, "全量图标", "Lucide（ISC）与品牌标（CC0）：指针停在上面看名字，点击复制名字", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				myui.NewInput(c.Key("icon search"), &s.iconQuery).Outline().Sm().
					PrefixIcon(icons.Search()).
					Placeholder("搜索图标…").
					Label("搜索图标")
				ui.Text(c, fmt.Sprintf("Lucide %d / %d · 品牌标 %d / %d",
					len(shown), len(all.Names), len(brands), len(brandall.Names))).
					TextColor(t.TextMuted)
			})
			ui.Text(c, "Lucide").FontSize(myui.TextSize(t, myui.TextSm)).Bold()
			ui.Grid(c).Columns(16).Gap(t.Space(1)).Children(func() {
				for _, name := range shown {
					s.icon(c, name, all.ByName(name))
				}
			})
			ui.Text(c, "品牌标").FontSize(myui.TextSize(t, myui.TextSm)).Bold()
			ui.Grid(c).Columns(16).Gap(t.Space(1)).Children(func() {
				for _, name := range brands {
					s.icon(c, name, brandall.ByName(name))
				}
			})
		})
	})
}

// icon is one icon of a set: the name on the pointer, and a click that
// copies the name.
func (s *showcase) icon(c *ui.Context, name string, svg *ui.SVG) {
	if svg == nil {
		return
	}
	b := myui.NewButton(c.Key("icon "+name), "").Ghost().Xs()
	b.Icon(svg).Label(name)
	myui.NewTooltip(c, b.Element, name)
	if b.Clicked() {
		s.copy(c, name)
	}
}

// copy puts the text on the clipboard and says so, where there is one:
// a tester has none.
func (s *showcase) copy(c *ui.Context, text string) {
	if !mygo.App.IsReady() {
		return
	}
	mygo.Clipboard.WriteText(text)
	myui.NewToast(c, "已复制 "+text).Show()
}

// filterIcons are the icons whose name holds the query: the ones of the
// Lucide set and the ones of the brands.
func filterIcons(query string) (icons, brands []string) {
	q := strings.ToLower(query)
	match := func(name string) bool {
		return q == "" || strings.Contains(strings.ToLower(name), q)
	}
	for _, name := range all.Names {
		if match(name) {
			icons = append(icons, name)
		}
	}
	for _, name := range brandall.Names {
		if match(name) {
			brands = append(brands, name)
		}
	}
	return icons, brands
}
