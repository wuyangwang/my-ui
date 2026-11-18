package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/display"
)

// tablePage demonstrates the Table component on its own page.
func (s *showcase) tablePage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "表格", "列宽、排序与选择仍是 mygo 的能力", func() {
		s.card(c, t, func() {
			display.NewTable(c, &s.tableState, []ui.TableColumn{
				{Title: "名称", Width: 160}, {Title: "状态", Width: 120},
			}, 2, func(row, col int) {
				if col == 0 {
					ui.Text(c, []string{"组件 A", "组件 B"}[row])
				} else {
					ui.Text(c, []string{"正常", "草稿"}[row])
				}
			})
		})
	})
}
