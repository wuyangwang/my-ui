package overlays

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

func drawerFixture(open *bool, onOpenChange func(bool)) func(*ui.Context) {
	return func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewDrawer(c, DrawerOptions{
			Open:         open,
			Side:         DrawerRight,
			OnOpenChange: onOpenChange,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "设置").Element
			},
			Body: func() {
				ui.Column(c).Gap(c.Theme().Space(3)).Children(func() {
					ui.Text(c, "面板里的内容")
					ctrlpkg.NewButton(c, "面板里的按钮").Outline()
				})
			},
		})
	}
}

func TestDrawerOpensOverTheEdgeOfTheWindow(t *testing.T) {
	open := false
	view := drawerFixture(&open, nil)
	tt := ui.NewTester(view, 640, 400)
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板里的内容"); ok {
		t.Error("the panel shows before the trigger is pressed")
	}
	if err := tt.Click("设置"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	content, ok := tt.Find("面板里的内容")
	if !ok {
		t.Fatal("the panel is not on screen after the trigger is pressed")
	}
	// The panel is at the right edge of the window: the room left of it
	// shows the window behind.
	const panelMax = 360
	if content.X < 640-panelMax {
		t.Errorf("the content of the panel is at %v, not at the right edge", content.X)
	}
}

func TestDrawerClosesWithItsButton(t *testing.T) {
	open := false
	tt := ui.NewTester(drawerFixture(&open, nil), 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("设置"); err != nil {
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
	if _, ok := tt.Find("面板里的内容"); ok {
		t.Error("the panel is still on screen after the close button")
	}
	if open {
		t.Error("the flag is still true after the close button")
	}
}

func TestDrawerClosesWhenThePointerClicksOutside(t *testing.T) {
	open := false
	tt := ui.NewTester(drawerFixture(&open, nil), 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("设置"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板里的内容"); !ok {
		t.Fatal("the panel did not open")
	}
	// A click on the scrim closes the panel.
	tt.ClickAt(10, 10)
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板里的内容"); ok {
		t.Error("the panel is still on screen after a click outside it")
	}
}

// TestDrawerFluentTitleAndFooterShowWhenOpened is the regression for the
// old early return: Title and Footer set while closed must show up when
// the drawer opens, not silently do nothing.
func TestDrawerFluentTitleAndFooterShowWhenOpened(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewDrawer(c, DrawerOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "设置").Element
			},
			Body: func() { ui.Text(c, "面板里的内容") },
		}).Title("抽屉标题").Footer(func() { ui.Text(c, "抽屉页脚") })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("设置"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("抽屉标题"); !ok {
		t.Error("the fluent title did not show when the drawer opened")
	}
	if _, ok := tt.Find("抽屉页脚"); !ok {
		t.Error("the fluent footer did not show when the drawer opened")
	}
}

func TestDrawerReportsOpenChanges(t *testing.T) {
	open := false
	var changes []bool
	tt := ui.NewTester(drawerFixture(&open, func(v bool) { changes = append(changes, v) }), 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("设置"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if len(changes) == 0 || !changes[len(changes)-1] {
		t.Errorf("OnOpenChange got %v, want a final true", changes)
	}
}
