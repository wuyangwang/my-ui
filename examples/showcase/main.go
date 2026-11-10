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
	"github.com/wuyangwang/my-ui/icons"
)

// showcase is the state the window shows.
type showcase struct {
	page       int
	appearance myui.Appearance
	palette    myui.PaletteID

	name      string
	email     string
	disabled  string
	readonly  string
	lastEvent string
}

// pages is one sidebar entry: the page key and its label.
var pages = []struct {
	key   string
	label string
}{
	{"button", "Button"},
	{"input", "Input"},
	{"badge", "Badge"},
	{"icon", "Icon"},
	{"scaffold", "Scaffold"},
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
			s.nav(c, t, 1)
			s.nav(c, t, 2)
			s.nav(c, t, 3)
			s.nav(c, t, 4)
		})
		ui.Column(c).Gap(t.Space(2)).Children(func() {
			ui.Text(c, "主题").FontSize(myui.TextSize(t, myui.TextXs)).Bold().TextColor(t.TextMuted)
			s.nav(c, t, 5)
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
		return func() { s.inputPage(c, t) }
	case 2:
		return func() { s.badgePage(c, t) }
	case 3:
		return func() { s.iconPage(c, t) }
	case 4:
		return func() { s.scaffoldPage(c, t) }
	case 5:
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

// inputPage is the Input's page: variants, sizes and states.
func (s *showcase) inputPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体", "边框与底色随 Variant 预设变化", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				myui.NewInput(c.Key("in-native"), &s.name).Placeholder("Native / 默认").Label("Native")
				myui.NewInput(c.Key("in-outline"), &s.email).Outline().Placeholder("Outline").Label("Outline")
				myui.NewInput(c.Key("in-ghost"), &s.name).Ghost().Placeholder("Ghost").Label("Ghost")
				myui.NewInput(c.Key("in-primary"), &s.name).Primary().Placeholder("Primary（强调边框）").Label("Primary")
				myui.NewInput(c.Key("in-danger"), &s.name).Destructive().Placeholder("Destructive").Label("Destructive")
			})
		})
	})

	s.sectionNote(c, t, "尺寸", "内衬与字号沿用 Size 档，与变体正交", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewInput(c.Key("in-sm"), &s.name).Sm().Outline().Placeholder("小尺寸").Label("小")
				myui.NewInput(c.Key("in-md"), &s.name).Md().Outline().Placeholder("中尺寸（默认）").Label("中")
				myui.NewInput(c.Key("in-lg"), &s.name).Lg().Outline().Placeholder("大尺寸").Label("大")
			})
		})
	})

	s.sectionNote(c, t, "状态", "禁用与只读遵循 ui 输入约定", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewInput(c.Key("in-disabled"), &s.disabled).Outline().Placeholder("不可编辑").Label("禁用").Disabled(true)
				s.readonly = "只读内容"
				myui.NewInput(c.Key("in-readonly"), &s.readonly).Outline().Label("只读").ReadOnly(true)
			})
		})
	})

	s.sectionNote(c, t, "组合", "清除按钮、前后缀与错误提示都由库排版", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				myui.NewInput(c.Key("in-clear"), &s.name).Outline().Label("可清除").Placeholder("输入后可清除").
					Clearable().OnCleared(func() { s.lastEvent = "cleared" })
				myui.NewInput(c.Key("in-affix"), &s.email).Outline().Label("带前后缀").
					Prefix(func() { ui.Text(c, "https://") }).
					Suffix(func() { ui.Text(c, ".com") })
				myui.NewInput(c.Key("in-error"), &s.email).Outline().Label("错误态").
					Invalid(true).ErrorText("邮箱格式不对")
			})
		})
	})

	s.sectionNote(c, t, "事件", "值变化、失焦、清除与回车都有回调", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				myui.NewInput(c.Key("in-events"), &s.name).Outline().Label("带事件").
					OnChanged(func(v string) { s.lastEvent = "changed: " + v }).
					OnBlurred(func() { s.lastEvent = "blurred" }).
					OnSubmit(func() { s.lastEvent = "submitted" })
				ui.Text(c, "最近事件："+s.lastEvent).
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			})
		})
	})
}

// iconPage shows the icons of the set and the presets of the Icon.
func (s *showcase) iconPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "常用图标", "来自 icons 包（Lucide, ISC），未被调用的图标不进二进制", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).Wrap().AlignItems(ui.Center).Children(func() {
				for _, ic := range []struct {
					svg  *ui.SVG
					name string
				}{
					{icons.Save(), "save"},
					{icons.Search(), "search"},
					{icons.Settings(), "settings"},
					{icons.Trash(), "trash-2"},
					{icons.Download(), "download"},
					{icons.Upload(), "upload"},
					{icons.Plus(), "plus"},
					{icons.Close(), "x"},
					{icons.Check(), "check"},
					{icons.Warning(), "triangle-alert"},
					{icons.Info(), "info"},
					{icons.Loading(), "loader-circle"},
					{icons.User(), "user"},
					{icons.Mail(), "mail"},
					{icons.Calendar(), "calendar"},
					{icons.Folder(), "folder"},
				} {
					myui.NewIcon(c.Key("icon "+ic.name), ic.svg).Md().Label(ic.name)
				}
			})
		})
	})

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
}

// badgePage is the Badge's page: variants and sizes.
func (s *showcase) badgePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "变体", "标签/徽标的语义预设", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				myui.NewBadge(c.Key("badge-native"), "Native").Native()
				myui.NewBadge(c.Key("badge-primary"), "Primary").Primary()
				myui.NewBadge(c.Key("badge-outline"), "Outline").Outline()
				myui.NewBadge(c.Key("badge-ghost"), "Ghost").Ghost()
				myui.NewBadge(c.Key("badge-destructive"), "Danger").Destructive()
				myui.NewBadge(c.Key("badge-link"), "Link").Link()
			})
		})
	})

	s.sectionNote(c, t, "尺寸", "内衬与字号随 Size 变化", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
				myui.NewBadge(c.Key("badge-sm"), "Sm").Primary().Sm()
				myui.NewBadge(c.Key("badge-md"), "Md").Primary().Md()
				myui.NewBadge(c.Key("badge-lg"), "Lg").Primary().Lg()
			})
		})
	})

	s.sectionNote(c, t, "组合示例", "与按钮等组件并排展示", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
				myui.NewButton(c.Key("badge-btn"), "消息").Outline()
				myui.NewBadge(c.Key("badge-count"), "12").Primary().Sm()
			})
		})
	})
}

// scaffoldPage demonstrates the Scaffold layout skeleton.
func (s *showcase) scaffoldPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "用途", "Scaffold 用于统一页面结构：可选标题栏安全区/顶部栏、可滚动内容区、底部栏，并为底部栏预留滚动空间（避免最后内容被遮挡）", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				ui.Text(c, "示例预览（卡片内模拟一个页面）").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
				ui.Column(c).Height(300).Background(t.Background).Border(1, t.Border).Radius(t.Radius).Children(func() {
					myui.Scaffold(c, myui.ScaffoldOptions{
						Top: func() {
							ui.Row(c).Fill().Padding(t.Space(2), t.Space(3)).Background(t.Surface).AlignItems(ui.Center).Children(func() {
								ui.Text(c, "页面标题 / 顶部栏").FontSize(myui.TextSize(t, myui.TextSm)).Bold()
							})
						},
						Bottom: func() {
							ui.Row(c).Fill().Padding(t.Space(2), t.Space(3)).Background(t.Surface).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
								ui.Row(c).Grow(1).Children(func() {})
								myui.NewButton(c.Key("scaffold-cancel"), "取消").Outline().Sm()
								myui.NewButton(c.Key("scaffold-ok"), "保存").Primary().Sm()
							})
						},
						BottomInset: 56,
						Content: func() {
							ui.Column(c).Padding(t.Space(3)).Gap(t.Space(2)).Children(func() {
								ui.Text(c, "这里是可滚动内容区").Bold()
								for i := 1; i <= 10; i++ {
									ui.Text(c, "列表项 / 内容行")
								}
								ui.Text(c, "最后一行（应在底部栏上方可见）").TextColor(t.TextMuted)
							})
						},
					})
				})
			})
		})
	})
	s.sectionNote(c, t, "标题栏安全区", "隐藏系统标题栏时可用 UseTitleBar 让顶部区域避开窗口控件并可拖拽", func() {
		s.card(c, t, func() {
			ui.Text(c, "myui.Scaffold(c, myui.ScaffoldOptions{UseTitleBar: true, Top: ..., Bottom: ..., Content: ...})").
				FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
		})
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
