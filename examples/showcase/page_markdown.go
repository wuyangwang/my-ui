package main

import (
	"github.com/egoist/mygo/ui"
	md "github.com/wuyangwang/my-ui/markdown"
)

// markdownDemo is the document both demos show: one source, read whole or
// a little more of it each frame.
const markdownDemo = "# 流式 Markdown\n\n这是一段 **加粗** 与 `行内代码`，还有一个 [链接](https://go.dev)。\n\n" +
	"- 第一项\n- 第二项\n\n> 引用一行。\n\n```go\nfunc main() {}\n```\n"

// markdownStreamDemo holds how much of the document the page has revealed.
// It lives on the demo element, so leaving the page drops it, and no
// package-level state or timer is involved.
type markdownStreamDemo struct {
	frame int
}

// markdownPage demonstrates a finished document and a reply streaming in.
func (s *showcase) markdownPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "静态文档", "标题、强调、列表、引用与代码，颜色随 my-ui 主题", func() {
		s.card(c, t, func() {
			md.NewMarkdown(c.Key("markdown-static"), markdownDemo)
		})
	})

	s.sectionNote(c, t, "流式输出", "每帧读取 app 持有的源文本，框架负责增量解析与淡入", func() {
		s.card(c, t, func() {
			box := ui.Box(c.Key("markdown-stream-demo"))
			st := ui.Local(box, "markdown", func() markdownStreamDemo { return markdownStreamDemo{} })
			st.frame++
			rs := []rune(markdownDemo)
			n := min(len(rs), 2+st.frame/2)
			source := string(rs[:n])
			box.Grow(1).Children(func() {
				md.NewStream(c.Key("markdown-stream"), func() string { return source })
			})
		})
	})
}
