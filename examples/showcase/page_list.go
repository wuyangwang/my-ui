package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/display"
)

// listDemo holds a list state with a stable key for each row, so the
// choice follows its item when the rows change.
type listDemo struct {
	state    ui.ListState
	selected int
	items    []string
}

// listPage demonstrates mygo's virtual list inside the library's
// rounded, bordered container. Selection stays with ListState.
func (s *showcase) listPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("list-demo")).Hidden(true)
	st := ui.Local(holder, "list-demo", func() listDemo {
		return listDemo{
			selected: -1,
			items:    []string{"收件箱", "已发送", "草稿箱", "归档", "垃圾箱"},
		}
	})
	if st == nil {
		return
	}
	// Point the state at the value ui.Local keeps, not at the temporary
	// the initializer returned: ui.Local copies that value into its slot.
	st.state.Selected = &st.selected
	st.state.Key = func(i int) any { return st.items[i] }
	st.state.Label = func(i int) string { return st.items[i] }
	s.sectionNote(c, t, "List", "包 mygo 的虚拟列表，只加圆角、描边与底色容器；位置与选择仍由 ListState 管", func() {
		s.card(c, t, func() {
			display.NewList(c, &st.state, len(st.items), func(i int) {
				ui.Text(c, st.items[i]).Padding(t.Space(2), t.Space(3))
			}).Height(t.Space(64))
			ui.Text(c, fmtChoice(st.selected, st.items)).
				FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
		})
	})
}

// fmtChoice is the line under the demo list: the name of the chosen row.
func fmtChoice(selected int, items []string) string {
	if selected < 0 || selected >= len(items) {
		return "当前未选择任何行。"
	}
	return "当前选择：" + items[selected]
}
