package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// snippetPage shows one-line commands with prefixes and variants.
func (s *showcase) snippetPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "命令", "前缀只显示，不进入剪贴板", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).AlignItems(ui.Start).Children(func() {
				displaypkg.NewSnippet(c, "go get github.com/wuyangwang/my-ui", displaypkg.SnippetOptions{
					Language: "sh", Prefix: "$", Variant: displaypkg.CodeFilled,
				})
				displaypkg.NewSnippet(c, "bun run dev", displaypkg.SnippetOptions{
					Language: "sh", Prefix: "$", Variant: displaypkg.CodeOutlined,
				})
				displaypkg.NewSnippet(c, "npx wrangler deploy", displaypkg.SnippetOptions{
					Language: "sh", Variant: displaypkg.CodeBorderless,
				})
			})
		})
	})
	s.sectionNote(c, t, "复制回调", "复制后写回 lastAction", func() {
		s.card(c, t, func() {
			displaypkg.NewSnippet(c, "go test ./...", displaypkg.SnippetOptions{
				Prefix: "$", OnCopy: func(string) { s.lastAction = "Copied command" },
			})
			ui.Text(c, s.lastAction).FontSize(12).TextColor(t.TextMuted)
		})
	})
}
