package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// dividerPage shows the horizontal, labeled, dashed and vertical lines.
func (s *showcase) dividerPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "横线", "普通、虚线与带文字", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(4)).Width(360).Children(func() {
				ui.Text(c, "Sign in with your work email.").TextColor(t.TextMuted)
				displaypkg.NewDivider(c)
				ui.Text(c, "Or use a passkey on this device.").TextColor(t.TextMuted)
				displaypkg.NewDivider(c, displaypkg.DividerOptions{Dashed: true})
				displaypkg.NewDivider(c, displaypkg.DividerOptions{Label: "OR"})
			})
		})
	})
	s.sectionNote(c, t, "竖线", "行内分隔", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				ui.Text(c, "Overview")
				displaypkg.NewDivider(c, displaypkg.DividerOptions{Vertical: true})
				ui.Text(c, "Docs")
				displaypkg.NewDivider(c, displaypkg.DividerOptions{Vertical: true, Dashed: true})
				ui.Text(c, "Examples")
			})
		})
	})
}
