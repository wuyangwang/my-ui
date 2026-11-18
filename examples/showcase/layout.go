package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// sectionNote is one section of a page: a heading, a muted line of
// context, then the demo. Every page uses this shape, so the right
// column reads the same way whichever component is open.
func (s *showcase) sectionNote(c *ui.Context, t *ui.Theme, block, note string, build func()) {
	ui.Column(c).Gap(t.Space(4)).Children(func() {
		ui.Column(c).Gap(t.Space(1)).Children(func() {
			ui.Text(c, block).FontSize(myui.TextSize(t, myui.TextLg)).Bold()
			ui.Text(c, note).FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
		})
		build()
	})
}

// card is the one demo surface every page uses: the page's background,
// a hairline border, shadcn's xl radius, one soft shadow and the same
// inset. A Row above grows it across the content column.
func (s *showcase) card(c *ui.Context, t *ui.Theme, build func()) {
	ui.Row(c).Children(func() {
		ui.Column(c).Grow(1).Background(t.Background).Border(1, t.Border).
			Radius(myui.Radius(t, myui.RadiusXl)).
			Padding(t.Space(5), t.Space(6)).Gap(t.Space(4)).
			Shadow(0, 1, 3, 0, t.Text.Alpha(0.05)).Children(build)
	})
}

// demoCard is a uniform tile for variant grids: a label, a short note
// and the control itself, all on the same card as the other demos.
func (s *showcase) demoCard(c *ui.Context, t *ui.Theme, name, cn, note string, build func()) {
	s.card(c, t, func() {
		ui.Row(c).FillWidth().AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
			ui.Column(c).Grow(1).Gap(t.Space(1)).Children(func() {
				ui.Row(c).AlignItems(ui.Baseline).Gap(t.Space(1.5)).Children(func() {
					ui.Text(c, cn).FontSize(myui.TextSize(t, myui.TextSm)).Bold()
					ui.Text(c, name).FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
				})
				ui.Text(c, note).FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			})
			build()
		})
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
func pickVariant(b *ctrlpkg.Button, active bool) *ctrlpkg.Button {
	if active {
		return b.Primary()
	}
	return b.Outline()
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
		Grow(1).Background(sub.Background).Padding(t.Space(4)).Gap(t.Space(2)).
		Radius(myui.Radius(t, myui.RadiusLg)).Children(func() {
		st := c.Theme()
		ui.Column(c).Grow(1).Background(st.Surface).Border(1, st.Border).
			Radius(st.Radius).Padding(t.Space(4)).Gap(t.Space(3)).Children(func() {
			st2 := c.Theme()
			ui.Text(c, label).FontSize(myui.TextSize(st2, myui.TextMd)).Bold()
			ctrlpkg.NewButton(c.Key("preview "+label), "在卡片里按下").Primary()
			ui.Text(c, "子树主题独立，提示与弹层也继承这一层。").
				FontSize(myui.TextSize(st2, myui.TextXs)).TextColor(st2.TextMuted)
		})
	})
}
