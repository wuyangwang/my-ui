package main

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// segmentedPage demonstrates the Segmented component.
func (s *showcase) segmentedPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "Segmented", "几选一的轨道，指示器滑向选中的那一段", func() {
		s.demoCard(c, t, "Filled", "填充", "带底的轨道", func() {
			ctrlpkg.NewSegmented(c, &s.tab, "列表", "网格", "表格").Label("视图")
		})
		s.demoCard(c, t, "Outlined", "描边", "只有边框的轨道", func() {
			ctrlpkg.NewSegmented(c, &s.tab, "日", "周", "月").Outlined().Sm().Label("粒度")
		})
		s.demoCard(c, t, "Block", "占满", "平分整行宽度", func() {
			ctrlpkg.NewSegmented(c, &s.tab, "左", "中", "右").Block(true).Label("对齐")
		})
	})
}
