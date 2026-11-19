package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	aipkg "github.com/wuyangwang/my-ui/ai"
	"github.com/wuyangwang/my-ui/icons"
)

type modelPickerDemo struct {
	text string
}

var modelPickerModels = []aipkg.ModelOption{
	{ID: "gpt-5", Name: "GPT-5", Group: "OpenAI", Icon: icons.Sparkles(),
		Description: "适合复杂推理", Context: "256K", Speed: "快", Price: "$$"},
	{ID: "gpt-5-mini", Name: "GPT-5 mini", Group: "OpenAI", Icon: icons.Sparkles(),
		Description: "便宜的日常任务", Context: "128K", Speed: "很快", Price: "$"},
	{ID: "claude", Name: "Claude", Group: "Anthropic", Icon: icons.BrainCircuit(),
		Description: "长文写作与代码", Context: "200K", Speed: "中", Price: "$$"},
	{ID: "local", Name: "本地模型", Group: "本地", Icon: icons.CpuAi(),
		Description: "离线运行", Context: "32K", Speed: "取决于机器", Price: "免费"},
}

func (s *showcase) modelPickerPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "ModelPicker", "可搜索的模型下拉，按厂商分组，图标与能力徽标随主题", func() {
		s.card(c, t, func() {
			box := ui.Box(c.Key("model-picker-demo"))
			st := ui.Local(box, "model", func() modelPickerDemo { return modelPickerDemo{} })
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				p := aipkg.NewModelPicker(c, &st.text, modelPickerModels).
					Outline().Placeholder("搜索模型…").Label("模型")
				ui.Text(c, "已选："+p.Chosen()).FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
			})
		})
	})
}
