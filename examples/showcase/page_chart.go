package main

import (
	"github.com/egoist/mygo/ui"
	chartpkg "github.com/wuyangwang/my-ui/chart"
)

// chartPage demonstrates the four themed chart presets. Values belong to
// the page only as demo data; the charts themselves hold no state.
func (s *showcase) chartPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "柱状图", "my-ui 主色填充，主题边框作网格，圆角随 Radius", func() {
		s.card(c, t, func() {
			chartpkg.NewBars(c.Key("chart-bars"), []float64{18, 32, 24, 41, 29})
		})
	})

	s.sectionNote(c, t, "折线与面积", "同一组数值，柔和填充与描边都取当前主题", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(4)).Children(func() {
				chartpkg.NewLine(c.Key("chart-line"), []float64{12, 19, 15, 28, 22, 34})
				chartpkg.NewArea(c.Key("chart-area"), []float64{12, 19, 15, 28, 22, 34})
			})
		})
	})

	s.sectionNote(c, t, "饼图", "按序列色分扇区，图例用色块与标签自绘", func() {
		s.card(c, t, func() {
			chartpkg.NewPie(c.Key("chart-pie"), []float64{42, 28, 18, 12})
		})
	})
}
