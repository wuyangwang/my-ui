package ai

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestReasoningDrawsLevelsParametersAndPresets(t *testing.T) {
	effort, temp, top := 1, 0.7, 0.95
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewReasoning(c, ReasoningOptions{
			Effort: &effort, Levels: []string{"低", "中", "高"},
			Temperature: &temp, TopP: &top,
		}).Label("推理设置")
	}
	tt := ui.NewTester(view, 420, 320)
	for _, s := range []string{"低", "中", "高", "Temperature", "Top P", "Precise", "Balanced", "Creative"} {
		if _, ok := tt.Find(s); !ok {
			t.Errorf("the panel is missing %q", s)
		}
	}
}

func TestReasoningWritesTheChosenLevelAndPreset(t *testing.T) {
	effort, temp, top := 0, 0.7, 0.95
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewReasoning(c, ReasoningOptions{
			Effort: &effort, Levels: []string{"低", "中", "高"},
			Temperature: &temp, TopP: &top,
		})
	}
	tt := ui.NewTester(view, 420, 320)
	if err := tt.Click("高"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	if effort != 2 {
		t.Errorf("effort is %d, want 2", effort)
	}
	if err := tt.Click("Precise"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	if effort != 0 || temp != 0.2 || top != 0.8 {
		t.Errorf("the precise preset set (%d, %v, %v), want (0, 0.2, 0.8)", effort, temp, top)
	}
}

func TestReasoningCompactHidesTheSliders(t *testing.T) {
	effort, temp := 0, 0.7
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewReasoning(c, ReasoningOptions{
			Effort: &effort, Levels: []string{"低", "中", "高"}, Temperature: &temp,
		}).Compact(true)
	}
	tt := ui.NewTester(view, 420, 320)
	if _, ok := tt.Find("Temperature"); ok {
		t.Error("Compact hides the temperature slider")
	}
	if _, ok := tt.Find("中"); !ok {
		t.Error("Compact keeps the levels row")
	}
}
