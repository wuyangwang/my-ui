package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// TestTheOverlaysPageOpensItsOverlays takes the overlays page for a
// drive: the tip shows over its anchor, the popover opens on a click,
// and the dialog opens when its trigger is pressed.
func TestTheOverlaysPageOpensItsOverlays(t *testing.T) {
	s := &showcase{page: "overlays"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)

	// The tip shows when the pointer rests on the anchor.
	anchor, ok := tt.Find("帮助")
	if !ok {
		t.Fatal("the tooltip's anchor is missing")
	}
	tt.Move(anchor.X+anchor.W/2, anchor.Y+anchor.H/2)
	tt.Advance(2 * time.Second)
	if _, ok := tt.Find("这是帮助文字"); !ok {
		t.Fatal("the tip did not show")
	}

	// The popover opens on a click on its anchor.
	if err := tt.Click("更多"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板里的内容"); !ok {
		t.Fatal("the popover did not open")
	}

	// The dialog opens on a click on its trigger.
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("确定删除这个项目吗？"); !ok {
		t.Fatal("the dialog did not open")
	}
	if !s.modalOpen {
		t.Error("the dialog's flag is not true")
	}
}
