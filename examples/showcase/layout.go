package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// sectionNote is a titled block of a page: heading, note, then cards.
func (s *showcase) sectionNote(c *ui.Context, t *ui.Theme, block, note string, build func()) {
	ui.Column(c).Gap(t.Space(4)).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2.5)).Children(func() {
			ui.Box(c).Width(3).Height(t.Space(5)).Radius(2).Background(t.Accent)
			ui.Text(c, block).FontSize(myui.TextSize(t, myui.TextLg)).Bold()
			ui.Text(c, note).FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
		})
		build()
	})
}

// card is a plain surface panel a page lays content into. On shadcn's
// theme a card sits on the same white as the page, so a hairline border
// and two soft shadows do the separating. A Row above grows it across
// the page width.
func (s *showcase) card(c *ui.Context, t *ui.Theme, build func()) {
	ui.Row(c).Children(func() {
		ui.Column(c).Grow(1).Background(t.Background).Border(1, t.Border).
			Radius(myui.Radius(t, myui.RadiusLg)).
			Padding(t.Space(4), t.Space(5)).Gap(t.Space(3)).
			Shadow(0, 1, 2, 0, t.Text.Alpha(0.06)).
			Shadow(0, 8, 24, -8, t.Text.Alpha(0.10)).Children(build)
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
			myui.NewButton(c.Key("preview "+label), "在卡片里按下").Primary()
			ui.Text(c, "子树主题独立，提示与弹层也继承这一层。").
				FontSize(myui.TextSize(st2, myui.TextXs)).TextColor(st2.TextMuted)
		})
	})
}
