package main

import (
	"strings"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/display"
)

var showcaseOutlineTree = map[string][]string{
	"src":       {"src/ui", "src/main.go"},
	"src/ui":    {"src/ui/list.go", "src/ui/outline.go"},
	"docs":      {"docs/readme.md"},
	"README.md": nil,
}

// showcaseOutlineChildren returns the roots for the virtual root "" and
// a node's children otherwise; nil means a leaf.
func showcaseOutlineChildren(path string) []string {
	if path == "" {
		return []string{"src", "docs", "README.md"}
	}
	return showcaseOutlineTree[path]
}

func showcaseOutlineTitle(path string) string {
	if i := strings.LastIndexByte(path, '/'); i >= 0 {
		return path[i+1:]
	}
	return path
}

// outlineDemo keeps the tree state on a stable element: the open items,
// the choice and the generation all belong to the app.
type outlineDemo struct {
	state ui.OutlineState[string]
}

// outlinePage demonstrates the generic outline wrapper: mygo owns the
// tree, the library supplies the container and the row text.
func (s *showcase) outlinePage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("outline-demo")).Hidden(true)
	st := ui.Local(holder, "outline-demo", func() outlineDemo {
		return outlineDemo{}
	})
	if st == nil {
		return
	}
	s.sectionNote(c, t, "Outline", "泛型封装 mygo 的 ui.Outline[K]；K 零值作为虚拟根，开合与键盘仍归 mygo", func() {
		s.card(c, t, func() {
			display.NewOutline(c, &st.state, showcaseOutlineChildren, showcaseOutlineTitle).
				Height(t.Space(64))
			ui.Text(c, "点箭头展开节点；选中行可用左右方向键开合。").
				FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
		})
	})
}
