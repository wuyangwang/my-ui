package overlays

import (
	"testing"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

func alertDialogFixture(open *bool, result *string, appearance Appearance) func(*ui.Context) {
	return func(c *ui.Context) {
		Apply(c, Options{Appearance: appearance})
		NewAlertDialog(c, AlertDialogOptions{
			Open:    open,
			Title:   "删除项目？",
			Message: "此操作不能撤销。",
			Actions: []DialogAction{
				{Label: "取消", OnClick: func() { *result = "cancelled" }},
				{Label: "确认删除", Kind: Destructive, OnClick: func() { *result = "deleted" }},
			},
			OnDismiss: func() { *result = "dismissed" },
		})
	}
}

func TestAlertDialogActionWritesThroughItsOnClick(t *testing.T) {
	open, result := false, ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		b := ctrlpkg.NewButton(c, "打开警告")
		if b.Clicked() {
			open = true
		}
		NewAlertDialog(c, AlertDialogOptions{
			Open:    &open,
			Title:   "删除项目？",
			Message: "此操作不能撤销。",
			Actions: []DialogAction{
				{Label: "取消", OnClick: func() { result = "cancelled" }},
				{Label: "确认删除", Kind: Destructive, OnClick: func() { result = "deleted" }},
			},
		})
	}
	tt := ui.NewTester(view, 500, 400)
	if err := tt.Click("打开警告"); err != nil {
		t.Fatal(err)
	}
	if _, ok := tt.Find("删除项目？"); !ok {
		t.Fatal("the alert's title is not on screen")
	}
	if result != "" {
		t.Fatalf("an idle alert reported %q", result)
	}
	if err := tt.Click("确认删除"); err != nil {
		t.Fatal(err)
	}
	if open {
		t.Error("the alert stayed open after its button")
	}
	if result != "deleted" {
		t.Errorf("the result is %q, want %q", result, "deleted")
	}
}

func TestAlertDialogEscapeCallsOnDismiss(t *testing.T) {
	open, result := true, ""
	tt := ui.NewTester(alertDialogFixture(&open, &result, AppearanceLight), 500, 400)
	tt.Key(0, ui.KeyEscape)
	if open {
		t.Error("Escape left the alert open")
	}
	if result != "dismissed" {
		t.Errorf("the result is %q, want %q", result, "dismissed")
	}
}

func TestAlertDialogLateOnDismissFires(t *testing.T) {
	open, result := true, ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		a := NewAlertDialog(c, AlertDialogOptions{
			Open:    &open,
			Title:   "删除项目？",
			Actions: []DialogAction{{Label: "确认", OnClick: func() { result = "ok" }}},
		})
		a.OnDismiss(func() { result = "dismissed" })
	}
	tt := ui.NewTester(view, 500, 400)
	tt.Key(0, ui.KeyEscape)
	if result != "dismissed" {
		t.Errorf("late OnDismiss got %q, want %q", result, "dismissed")
	}
}

func TestAlertDialogFollowsTheTheme(t *testing.T) {
	for _, a := range []Appearance{AppearanceLight, AppearanceDark} {
		open, result := true, ""
		tt := ui.NewTester(alertDialogFixture(&open, &result, a), 500, 400)
		for _, s := range []string{"删除项目？", "此操作不能撤销。", "取消", "确认删除"} {
			if _, ok := tt.Find(s); !ok {
				t.Errorf("%v: %q is not on screen", a, s)
			}
		}
	}
}
