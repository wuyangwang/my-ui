package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestStepsShowsTitlesAndCurrent(t *testing.T) {
	current := 1
	view := func(c *ui.Context) {
		NewSteps(c, StepsProps{Current: current, Items: []StepItem{
			{Title: "账号"}, {Title: "资料"}, {Title: "完成"},
		}})
	}
	tt := ui.NewTester(view, 480, 120)
	tt.Advance(time.Second)
	for _, title := range []string{"账号", "资料", "完成"} {
		if _, ok := tt.Find(title); !ok {
			t.Fatalf("the step %q is missing", title)
		}
	}
	info, ok := tt.Access("资料")
	if !ok {
		t.Fatal("the current step is not on screen")
	}
	if info.Bounds.W == 0 {
		t.Error("the current step has no visible box")
	}
	_ = current
}
