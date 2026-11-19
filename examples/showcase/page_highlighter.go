package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

const showcaseGoCode = `package main

import "fmt"

func main() {
	// a comment
	name := "Lobe"
	fmt.Println("Hello", name, 42)
}
`

const showcaseTSCode = `import { useState } from "react";

export function Counter() {
  const [n, setN] = useState(0);
  return <button onClick={() => setN(n + 1)}>{n}</button>;
}
`

// highlighterPage shows the code block's variants and header.
func (s *showcase) highlighterPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "基础", "小型分词器，颜色取自当前主题", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				displaypkg.NewHighlighter(c, showcaseGoCode, displaypkg.HighlighterOptions{
					Language: "go", Variant: displaypkg.CodeFilled,
				})
				displaypkg.NewHighlighter(c, showcaseTSCode, displaypkg.HighlighterOptions{
					Language: "ts", Variant: displaypkg.CodeOutlined,
				})
			})
		})
	})
	s.sectionNote(c, t, "完整头部", "折叠指针由 ui.Local 持有", func() {
		s.card(c, t, func() {
			holder := ui.Box(c)
			open := ui.Local(holder, "showcase-code-open", func() bool { return true })
			displaypkg.NewHighlighter(c, showcaseGoCode, displaypkg.HighlighterOptions{
				Language: "go", FileName: "main.go",
				FullFeatured: true, Collapsed: open,
				Variant: displaypkg.CodeOutlined, Shadow: true,
			})
		})
	})
}
