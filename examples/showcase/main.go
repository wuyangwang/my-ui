// Showcase lays my-ui's widgets out in one native window: one page per
// component, plus a page for the theming switches and subtree previews.
//
//	go run ./examples/showcase
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// showcase is the state the window shows.
type showcase struct {
	page       int
	appearance myui.Appearance
	palette    myui.PaletteID
}

// pages is one sidebar entry: the page key and its label.
var pages = []struct {
	key   string
	label string
}{
	{"button", "Button"},
	{"theme", "外观"},
}

// swatches is one palette switch entry: the id and its label.
var swatches = []struct {
	id     myui.PaletteID
	button string
}{
	{myui.PaletteNone, "系统"},
	{myui.PaletteBlue, "蓝"},
	{myui.PaletteViolet, "紫"},
	{myui.PaletteRose, "玫红"},
	{myui.PaletteOrange, "橙"},
	{myui.PaletteEmerald, "翠绿"},
	{myui.PaletteTeal, "青"},
	{myui.PaletteSky, "天蓝"},
	{myui.PalettePink, "粉"},
}

// view builds one frame. Apply runs first, every frame, so the theme
// follows the state; everything below it takes what it sets.
func (s *showcase) view(c *ui.Context) {
	t := myui.Apply(c, myui.Options{Appearance: s.appearance, Palette: s.palette})
	if t == nil {
		t = c.Theme()
	}

	// Draw the frame's own background with the applied theme: behind
	// every element MyGo paints the window theme, which the forced
	// appearance disagrees with, so the root paints its own.
	ui.Row(c).Fill().Background(t.Background).AlignItems(ui.Stretch).Children(func() {
		s.sidebar(c, t)
		s.main(c, t)
	})
}

// sidebar carries the navigation: the brand and the pages, one group
// per kind, on the surface tone so the structure reads.
func (s *showcase) sidebar(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(190).Background(t.Surface).Padding(t.Space(4)).
		Gap(t.Space(4)).Children(func() {
		ui.Text(c, "my-ui").FontSize(myui.TextSize(t, myui.TextLg)).Bold()
		ui.Text(c, "MyGo 原生组件库").
			FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)

		ui.Column(c).Gap(t.Space(2)).Children(func() {
			ui.Text(c, "组件").FontSize(myui.TextSize(t, myui.TextXs)).Bold().TextColor(t.TextMuted)
			s.nav(c, t, 0)
		})
		ui.Column(c).Gap(t.Space(2)).Children(func() {
			ui.Text(c, "主题").FontSize(myui.TextSize(t, myui.TextXs)).Bold().TextColor(t.TextMuted)
			s.nav(c, t, 1)
		})
	})
}

// nav is one sidebar entry; the open page paints the accent ink.
func (s *showcase) nav(c *ui.Context, t *ui.Theme, i int) {
	item := pages[i]
	b := myui.NewButton(c.Key("nav "+item.key), item.label).Ghost()
	if s.page == i {
		a := t.Accent
		b.Background(a.Alpha(0.10)).TextColor(a)
	}
	if b.Clicked() {
		s.page = i
	}
}

// main is the page column beside the sidebar: the page title, then the
// page on a scroll.
func (s *showcase) main(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Grow(1).Padding(t.Space(4)).Gap(t.Space(4)).Children(func() {
		ui.Text(c, pages[s.page].label).
			FontSize(myui.TextSize(t, myui.TextXl)).Bold()

		ui.Scroll(c).Grow(1).Gap(t.Space(4)).Children(s.pages(c, t))
	})
}

// pages hands back the builder of the open page.
func (s *showcase) pages(c *ui.Context, t *ui.Theme) func() {
	switch s.page {
	case 1:
		return func() { s.themePage(c, t) }
	default:
		return func() { s.buttonPage(c, t) }
	}
}

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

// themePage is the theming page: appearance, palette, and both
// appearances side by side as subtree previews.
func (s *showcase) themePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "外观与色板", "切换下一帧生效，整窗换装", func() {
		s.card(c, t, func() {
			s.cardHead(c, t, "明暗", "跟随系统，或强制一边")
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				choices := []struct {
					value  myui.Appearance
					button string
				}{
					{myui.AppearanceSystem, "跟随系统"},
					{myui.AppearanceLight, "浅色"},
					{myui.AppearanceDark, "深色"},
				}
				for _, ch := range choices {
					active := s.appearance == ch.value
					if pickVariant(myui.NewButton(c.Key("appearance "+ch.button), ch.button), active).
						Clicked() {
						s.appearance = ch.value
					}
				}
			})
		})

		s.card(c, t, func() {
			s.cardHead(c, t, "色板", "系统 · 8 个预设的强调色三元")
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				for _, p := range swatches {
					active := s.palette == p.id
					if pickVariant(myui.NewButton(c.Key("palette "+p.button), p.button), active).
						Clicked() {
						s.palette = p.id
					}
				}
			})
		})
	})

	s.sectionNote(c, t, "子树预览", "ui.Themed 里的另一外观，卡片独立换装", func() {
		ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Start).Children(func() {
			pane(c, t, s.palette, myui.AppearanceLight, "浅色子树")
			pane(c, t, s.palette, myui.AppearanceDark, "深色子树")
		})
	})
}

// pane is one Themed preview card: a frame inside, where the subtree's
// own theme colors its parts. The subtree's theme comes from NewTheme,
// which builds without setting: only the frame's own Apply paints the
// window's ink, and a pane must not reach the texts outside it. The
// pane paints its own background too, since a subtree darker than its
// frame would show the window's.
func pane(c *ui.Context, t *ui.Theme, palette myui.PaletteID, a myui.Appearance, label string) {
	sub := myui.NewTheme(c, myui.Options{Appearance: a, Palette: palette})
	ui.Themed(c, sub).
		Grow(1).Background(sub.Background).Padding(t.Space(3)).Gap(t.Space(2)).Children(func() {
		st := c.Theme()
		ui.Column(c).Grow(1).Background(st.Surface).Border(1, st.Border).
			Radius(st.Radius).Padding(t.Space(3)).Gap(t.Space(2)).Children(func() {
			st2 := c.Theme()
			ui.Text(c, label).FontSize(myui.TextSize(st2, myui.TextMd)).Bold()
			myui.NewButton(c.Key("preview "+label), "在卡片里按下").Primary()
			ui.Text(c, "子树主题独立，提示与弹层也继承这一层。").
				FontSize(myui.TextSize(st2, myui.TextXs)).TextColor(st2.TextMuted)
		})
	})
}

// sectionNote is a titled block of a page: heading, note, then cards.
func (s *showcase) sectionNote(c *ui.Context, t *ui.Theme, block, note string, build func()) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Row(c).AlignItems(ui.Baseline).Gap(t.Space(2)).Children(func() {
			ui.Text(c, block).FontSize(myui.TextSize(t, myui.TextSm)).Bold()
			ui.Text(c, note).FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
		})
		build()
	})
}

// card is a plain surface panel a page lays content into; a Row above
// grows it across the page width.
func (s *showcase) card(c *ui.Context, t *ui.Theme, build func()) {
	ui.Row(c).Children(func() {
		ui.Column(c).Grow(1).Background(t.Surface).Border(1, t.Border).
			Radius(t.Radius).Padding(t.Space(3), t.Space(4)).Gap(t.Space(2)).Children(build)
	})
}

// cardHead is one card's title line; the note runs alongside, muted.
func (s *showcase) cardHead(c *ui.Context, t *ui.Theme, title, note string) {
	ui.Row(c).AlignItems(ui.Baseline).Gap(t.Space(2)).Children(func() {
		ui.Text(c, title).FontSize(myui.TextSize(t, myui.TextSm)).Bold()
		ui.Text(c, note).FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
	})
}

// pickVariant chooses the look of a switch's button: the chosen one
// fills with the accent, the others sit outlined.
func pickVariant(b *myui.Button, active bool) *myui.Button {
	if active {
		return b.Primary()
	}
	return b.Outline()
}

func main() {
	s := &showcase{}
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "my-ui 组件示例",
			Width:   960,
			Height:  620,
			Content: ui.View(s.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
