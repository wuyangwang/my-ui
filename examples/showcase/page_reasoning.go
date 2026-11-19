package main

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	aipkg "github.com/wuyangwang/my-ui/ai"
)

type reasoningDemo struct {
	effort int
	temp   float64
	topP   float64
}

func (s *showcase) reasoningPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "Reasoning", "思考强度低/中/高，温度与 top_p 三个语义预设，值由 app 持有", func() {
		s.card(c, t, func() {
			box := ui.Box(c.Key("reasoning-demo"))
			st := ui.Local(box, "reasoning", func() reasoningDemo {
				return reasoningDemo{effort: 1, temp: 0.7, topP: 0.95}
			})
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				aipkg.NewReasoning(c, aipkg.ReasoningOptions{
					Effort: &st.effort, Temperature: &st.temp, TopP: &st.topP,
				}).Label("推理设置")
				ui.Text(c, "档位 "+strconv.Itoa(st.effort)+" · 温度 "+
					strconv.FormatFloat(st.temp, 'f', 2, 64)+" · top_p "+
					strconv.FormatFloat(st.topP, 'f', 2, 64)).
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			})
		})
	})
}
