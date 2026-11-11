package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// sliderPage shows the sliders: the knobs that pick a number.
func (s *showcase) sliderPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "滑杆", "Slider：在两个数之间取值，方向键与 Home/End 也能动", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(4)).Children(func() {
				ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
					myui.NewSlider(c.Key("sl volume"), &s.volume, 0, 100).Label("音量")
					ui.Text(c, fmt.Sprintf("%d", int(s.volume))).TextColor(t.TextMuted)
				})
				ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
					myui.NewSlider(c.Key("sl step"), &s.step, 0, 10).Step(1).Label("档位")
					ui.Text(c, fmt.Sprintf("%d", int(s.step))).TextColor(t.TextMuted)
				})
			})
		})
	})
}
