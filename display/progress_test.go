package display

import (
	. "github.com/wuyangwang/my-ui"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestProgressIsOnScreen(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewProgress(c, 40).Label("上传进度").ShowInfo(true)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Access("上传进度"); !ok {
		t.Error("the bar is not on screen")
	}
	if _, ok := tt.AccessByRole(ui.RoleProgress, "上传进度"); !ok {
		t.Error("the bar has no progressbar role")
	}
}

func TestProgressReportsPercent(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewProgress(c, 40).Label("forty")
			NewProgress(c, 0.4).Label("fraction")
			NewProgress(c, 2).Value(1, 4).Label("quarter")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	for name, want := range map[string]float64{"forty": 40, "fraction": 40, "quarter": 25} {
		node, ok := tt.AccessByRole(ui.RoleProgress, name)
		if !ok {
			t.Fatalf("progress %q missing", name)
		}
		if node.Now != want {
			t.Errorf("progress %q reports %v, want %v", name, node.Now, want)
		}
	}
}

func TestProgressFormsAndStatuses(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewProgress(c, 40).Variant(ProgressSegments).Label("segments")
			NewProgress(c, 60).Variant(ProgressInset).Label("inset")
			NewProgress(c, 80).Type(ProgressCircle).ShowInfo(true).Label("circle")
			NewProgress(c, 100).Status(ProgressSuccess).ShowInfo(true).Label("done")
			NewProgress(c, 50).Status(ProgressException).ShowInfo(true).Label("failed")
			NewProgress(c, -1).Label("indeterminate")
		})
	}
	tt := ui.NewTester(view, 360, 420)
	tt.Advance(time.Second)
	for _, name := range []string{"segments", "inset", "circle", "done", "failed", "indeterminate"} {
		if _, ok := tt.AccessByRole(ui.RoleProgress, name); !ok {
			t.Errorf("progress %q missing", name)
		}
	}
}

func TestProgressFormat(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewProgress(c, 30).ShowInfo(true).Format(func(p float32) string { return "custom" })
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if _, ok := tt.Find("custom"); !ok {
		t.Error("the formatted percent is not on screen")
	}
}
