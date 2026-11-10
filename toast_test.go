package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestToastShowsItsTitleAndDescription(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewToastViewport(c)
		btn := NewButton(c, "保存")
		if btn.Clicked() {
			NewToast(c, "已保存").Description("改动已写入").Success().Show()
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("保存"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("已保存"); !ok {
		t.Fatal("the toast is not on screen")
	}
	if _, ok := tt.Find("改动已写入"); !ok {
		t.Error("the description of the toast is not on screen")
	}
	// The library's viewport takes the place of the theme's own look of
	// the toasts: the toast shows once.
	shown := 0
	for _, s := range tt.Texts() {
		if s == "已保存" {
			shown++
		}
	}
	if shown != 1 {
		t.Errorf("the title shows %d times, want once", shown)
	}
}

func TestToastClosesWithItsButton(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewToastViewport(c)
		if NewButton(c, "保存").Element.Clicked() {
			NewToast(c, "已保存").Timeout(-1).Show()
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("保存"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("已保存"); !ok {
		t.Fatal("the toast is not on screen")
	}
	// The close button closes the toast.
	close, ok := tt.Access("关闭")
	if !ok {
		t.Fatal("the close button is not on screen")
	}
	b := close.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Advance(time.Second)
	if _, ok := tt.Find("已保存"); ok {
		t.Error("the toast is still on screen after the close button")
	}
}

func TestToastRunsItsAction(t *testing.T) {
	undone := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewToastViewport(c)
		if NewButton(c, "删除").Element.Clicked() {
			NewToast(c, "已删除").Action("撤销", func() { undone = true }).Timeout(-1).Show()
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("删除"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("撤销"); !ok {
		t.Fatal("the action of the toast is not on screen")
	}
	if err := tt.Click("撤销"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !undone {
		t.Error("the action did not run")
	}
	if _, ok := tt.Find("已删除"); ok {
		t.Error("the toast is still on screen after its action")
	}
}

func TestToastGoesWhenItsTimeIsOver(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewToastViewport(c)
		if NewButton(c, "保存").Element.Clicked() {
			NewToast(c, "已保存").Timeout(2 * time.Second).Show()
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("保存"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("已保存"); !ok {
		t.Fatal("the toast is not on screen")
	}
	tt.Advance(2 * time.Second)
	if _, ok := tt.Find("已保存"); ok {
		t.Error("the toast is still on screen after its time")
	}
}
