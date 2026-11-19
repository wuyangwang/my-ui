package overlays

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

func TestModalTriggerOpensAndBodyLandsInThePanel(t *testing.T) {
	open := false
	var changes []bool
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		m := NewModal(c, ModalOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "打开对话框").Element
			},
			Title:       "删除项目",
			Description: "此操作不能撤销。",
			Body:        func() { ui.Text(c, "确定删除这个项目吗？") },
		})
		m.OnOpenChange(func(v bool) { changes = append(changes, v) })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("确定删除这个项目吗？"); ok {
		t.Fatal("the dialog shows before the trigger is pressed")
	}
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	if !open {
		t.Fatal("one click on the trigger did not open the dialog")
	}
	tt.Advance(time.Second)
	for _, s := range []string{"删除项目", "此操作不能撤销。", "确定删除这个项目吗？"} {
		if _, ok := tt.Find(s); !ok {
			t.Errorf("%q did not land in the panel", s)
		}
	}
	if len(changes) == 0 || !changes[len(changes)-1] {
		t.Errorf("late OnOpenChange got %v, want a final true", changes)
	}
}

func TestModalCloseButtonReportsOpenChange(t *testing.T) {
	open := false
	var changes []bool
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		m := NewModal(c, ModalOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "打开对话框").Element
			},
			Body: func() { ui.Text(c, "内容") },
		})
		m.OnOpenChange(func(v bool) { changes = append(changes, v) })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	close, ok := tt.Access("Close")
	if !ok {
		t.Fatal("the close button is not on screen")
	}
	b := close.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Advance(time.Second)
	if open {
		t.Error("the close button did not close the dialog")
	}
	if len(changes) < 2 || changes[len(changes)-1] {
		t.Errorf("OnOpenChange got %v, want a final false", changes)
	}
}
