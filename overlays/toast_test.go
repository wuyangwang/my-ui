package overlays

import (
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestToastShowsItsTitleAndDescription(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewToastViewport(c)
		btn := ctrlpkg.NewButton(c, "保存")
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
		if ctrlpkg.NewButton(c, "保存").Element.Clicked() {
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
	close, ok := tt.Access("Close")
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
		if ctrlpkg.NewButton(c, "删除").Element.Clicked() {
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
		if ctrlpkg.NewButton(c, "保存").Element.Clicked() {
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

// showFromHandler uses only a Notifier, never a *ui.Context: the handler
// captures the stable handle, so it can be stored on the app's state and
// called from any later frame.
func showFromHandler(n *Notifier) string {
	return n.Show(ToastOptions{
		Title:       "已保存",
		Description: "改动已经写入。",
		Kind:        Success,
		Duration:    -1,
	})
}

func TestNotifierShowAndClose(t *testing.T) {
	var n *Notifier
	var id string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		if n == nil {
			n = NewNotifier(c)
		}
		n.Build()
		if ctrlpkg.NewButton(c, "保存").Element.Clicked() {
			id = showFromHandler(n)
		}
		if ctrlpkg.NewButton(c, "关闭通知").Element.Clicked() && id != "" {
			n.Close(id)
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("保存"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("已保存"); !ok {
		t.Fatal("Notifier.Show did not put the toast on screen")
	}
	if _, ok := tt.Find("改动已经写入。"); !ok {
		t.Error("Notifier.Show did not pass the description through")
	}
	if id == "" {
		t.Fatal("Notifier.Show returned no id")
	}
	if err := tt.Click("关闭通知"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("已保存"); ok {
		t.Error("Notifier.Close did not close the toast")
	}
}

func TestNotifierShowRunsAction(t *testing.T) {
	undone := false
	var n *Notifier
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		if n == nil {
			n = NewNotifier(c)
		}
		n.Build()
		if ctrlpkg.NewButton(c, "删除").Element.Clicked() {
			n.Show(ToastOptions{
				Title:       "已删除",
				Kind:        Destructive,
				Duration:    -1,
				ActionLabel: "撤销",
				OnAction:    func() { undone = true },
			})
		}
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("删除"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("撤销"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !undone {
		t.Error("the notifier's action did not run")
	}
}
