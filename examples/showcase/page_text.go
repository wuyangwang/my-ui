package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// textPage shows the type styles, headings and line limits.
func (s *showcase) textPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "语义色", "成功、警告、危险、信息", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).Children(func() {
				displaypkg.NewText(c, "默认").Variant(0)
				displaypkg.NewText(c, "次要").Secondary()
				displaypkg.NewText(c, "成功").Success()
				displaypkg.NewText(c, "警告").Warning()
				displaypkg.NewText(c, "危险").Danger()
				displaypkg.NewText(c, "信息").Info()
			})
		})
	})
	s.sectionNote(c, t, "标题与字饰", "Heading 1–5、粗斜线与代码", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				displaypkg.NewText(c, "一级标题").Heading(1)
				displaypkg.NewText(c, "三级标题").Heading(3)
				displaypkg.NewText(c, "粗体与斜体").Strong().Italic()
				displaypkg.NewText(c, "删除线与下划线", displaypkg.TextOptions{
					Underline: true, Delete: true,
				})
				displaypkg.NewText(c, "const x = 1", displaypkg.TextOptions{Code: true})
				displaypkg.NewText(c, "高亮").Mark()
			})
		})
	})
	s.sectionNote(c, t, "行数与禁用", "单行省略与禁用墨色", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Width(240).Children(func() {
				displaypkg.NewText(c, "这是一段很长很长、需要被截断成单行的示例说明文字。",
					displaypkg.TextOptions{Rows: 1})
				displaypkg.NewText(c, "禁用文本", displaypkg.TextOptions{Disabled: true})
			})
		})
	})
}
