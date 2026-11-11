package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestDrawerOpensOverTheEdgeOfTheWindow(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewDrawer(c, &open, DrawerRight, func() ui.Element {
			return NewButton(c, "设置").Element
		}, func(panel ui.Element) {
			ui.Text(c, "面板里的内容")
		})
	}
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
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewDrawer(c, &open, DrawerRight, func() ui.Element {
			return NewButton(c, "设置").Element
		}, func(panel ui.Element) {
			ui.Text(c, "面板里的内容")
		})
	}
	tt := ui.NewTester(view, 320, 240)
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
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewDrawer(c, &open, DrawerRight, func() ui.Element {
			return NewButton(c, "设置").Element
		}, func(panel ui.Element) {
			ui.Text(c, "面板里的内容")
		})
	}
	tt := ui.NewTester(view, 320, 240)
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
