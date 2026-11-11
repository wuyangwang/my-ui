package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

func TestAlertShowsItsTitleAndDescription(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewAlert(c, "注意").Description("这条信息需要留意")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("注意"); !ok {
		t.Fatal("the title of the alert is not on screen")
	}
	if _, ok := tt.Find("这条信息需要留意"); !ok {
		t.Fatal("the description of the alert is not on screen")
	}
}

func TestAlertIsAnAlertToAssistiveTechnology(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewAlert(c, "出错了").Destructive()
		NewAlert(c, "好了").Success()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	// The alert is a live region: a screen reader announces it as it
	// comes.
	_, ok := tt.AccessByRole(ui.RoleAlert, "")
	if !ok {
		t.Fatal("no alert is in the tree of the assistive technology")
	}
	// The title of an alert reads as its text.
	for _, name := range []string{"出错了", "好了"} {
		if _, ok := tt.Access(name); !ok {
			t.Errorf("the title %q of the alert is not on screen", name)
		}
	}
}

func TestAlertHasItsIconAtTheLeft(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewAlert(c, "注意").Icon(icons.TriangleAlert()).Description("这条信息需要留意")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	title, ok := tt.Find("注意")
	if !ok {
		t.Fatal("the title of the alert is not on screen")
	}
	// The icon takes the room at the left of the title: the title does
	// not sit at the edge of the window.
	if title.X < 8 {
		t.Errorf("the title of the alert is at %v, with no room for an icon", title.X)
	}
}
