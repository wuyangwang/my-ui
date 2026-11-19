package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestAgentAvatarStatesAndSeed(t *testing.T) {
	states := []AgentState{AgentIdle, AgentWorking, AgentWaiting, AgentBlocked, AgentThinking, AgentDone}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			for i, s := range states {
				NewAgentAvatar(c, AgentAvatarOptions{
					Seed: "agent", Name: string(rune('A' + i)), State: s,
				})
			}
		})
	}
	tt := ui.NewTester(view, 480, 120)
	tt.Advance(time.Second)
	for i := range states {
		name := string(rune('A' + i))
		if _, ok := tt.Access(name); !ok {
			t.Fatalf("agent %s missing", name)
		}
	}
	if a, ok := tt.Access("A"); ok {
		if b, ok2 := tt.Access("B"); ok2 && a.Bounds.W != b.Bounds.W {
			t.Fatal("same seed must give the same size")
		}
	}
}

func TestAgentAvatarShapes(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			for i := AgentFace; i <= AgentCat; i++ {
				NewAgentAvatar(c, AgentAvatarOptions{Shape: i, Color: AgentPurple, Size: 32})
			}
		})
	}
	tt := ui.NewTester(view, 560, 120)
	tt.Advance(time.Second)
	if tt.Image() == nil {
		t.Fatal("no frame drawn")
	}
}
