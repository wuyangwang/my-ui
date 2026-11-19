package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// agentAvatarPage shows the seeded agents and their states.
func (s *showcase) agentAvatarPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "种子", "同一名字得到同一形状与颜色", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				for _, name := range []string{"Ada", "Grace", "Alan", "Lobe", "Claude", "GPT"} {
					displaypkg.NewAgentAvatar(c.Key(name), displaypkg.AgentAvatarOptions{
						Seed: name, Name: name, Size: 44,
					})
				}
			})
		})
	})
	s.sectionNote(c, t, "状态", "眼睛与角标随状态变化，无动画", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				states := []struct {
					name  string
					state displaypkg.AgentState
				}{
					{"Idle", displaypkg.AgentIdle}, {"Working", displaypkg.AgentWorking},
					{"Waiting", displaypkg.AgentWaiting}, {"Blocked", displaypkg.AgentBlocked},
					{"Thinking", displaypkg.AgentThinking}, {"Done", displaypkg.AgentDone},
				}
				for _, st := range states {
					ui.Column(c).Gap(t.Space(1)).AlignItems(ui.Center).Children(func() {
						displaypkg.NewAgentAvatar(c, displaypkg.AgentAvatarOptions{
							Seed: "demo", Name: st.name, Size: 40, State: st.state,
						})
						ui.Text(c, st.name).FontSize(11).TextColor(t.TextMuted)
					})
				}
			})
		})
	})
}
