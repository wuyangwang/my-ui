package overlays

import (
	"testing"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

func alertDialogFixture(open *bool, results *[]int, appearance Appearance) func(*ui.Context) {
	return func(c *ui.Context) {
		Apply(c, Options{Appearance: appearance})
		NewAlertDialog(c, open, AlertDialogOptions{
			Title:    "删除项目？",
			Message:  "此操作不能撤销。",
			Buttons:  []string{"取消", "确认删除"},
			OnResult: func(i int) { *results = append(*results, i) },
		})
	}
}

func TestAlertDialogReportsTheClickedButton(t *testing.T) {
	open, results := false, []int{}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		b := ctrlpkg.NewButton(c, "打开警告")
		if b.Clicked() {
			open = true
		}
		NewAlertDialog(c, &open, AlertDialogOptions{
			Title:    "删除项目？",
			Message:  "此操作不能撤销。",
			Buttons:  []string{"取消", "确认删除"},
			OnResult: func(i int) { results = append(results, i) },
		})
	}
	tt := ui.NewTester(view, 500, 400)
	if err := tt.Click("打开警告"); err != nil {
		t.Fatal(err)
	}
	if _, ok := tt.Find("删除项目？"); !ok {
		t.Fatal("the alert's title is not on screen")
	}
	if len(results) != 0 {
		t.Fatalf("an idle alert reported %v", results)
	}
	if err := tt.Click("确认删除"); err != nil {
		t.Fatal(err)
	}
	if open {
		t.Error("the alert stayed open after its button")
	}
	if len(results) != 1 || results[0] != 1 {
		t.Errorf("the result is %v, want [1]", results)
	}
}

func TestAlertDialogEscapeReportsDismissed(t *testing.T) {
	open, results := true, []int{}
	tt := ui.NewTester(alertDialogFixture(&open, &results, AppearanceLight), 500, 400)
	tt.Key(0, ui.KeyEscape)
	if open {
		t.Error("Escape left the alert open")
	}
	if len(results) != 1 || results[0] != ui.AlertDismissed {
		t.Errorf("the result is %v, want [%d]", results, ui.AlertDismissed)
	}
}

func TestAlertDialogFollowsTheTheme(t *testing.T) {
	for _, a := range []Appearance{AppearanceLight, AppearanceDark} {
		open, results := true, []int{}
		tt := ui.NewTester(alertDialogFixture(&open, &results, a), 500, 400)
		for _, s := range []string{"删除项目？", "此操作不能撤销。", "取消", "确认删除"} {
			if _, ok := tt.Find(s); !ok {
				t.Errorf("%v: %q is not on screen", a, s)
			}
		}
	}
}
