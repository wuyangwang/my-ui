package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestTooltipShowsOverItsAnchor(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Children(func() {
			// The anchor sits low in the window: a tip that lands in
			// the corner shows the mistake this test is about.
			ui.Box(c).Height(140)
			NewTooltip(c, NewButton(c, "帮助").Element, "这是帮助文字")
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)

	anchor, ok := tt.Find("帮助")
	if !ok {
		t.Fatal("the anchor is not on screen")
	}
	// The tip shows when the pointer rests on the anchor.
	tt.Move(anchor.X+anchor.W/2, anchor.Y+anchor.H/2)
	tt.Advance(2 * time.Second)
	tip, ok := tt.Find("这是帮助文字")
	if !ok {
		t.Fatal("the tip is not on screen")
	}
	if tip.X == 0 && tip.Y == 0 {
		t.Error("the tip is in the window's corner, not over the anchor")
	}
	if tip.Y >= anchor.Y {
		t.Errorf("the tip is at %v, not above the anchor at %v", tip, anchor)
	}
	if tip.X > anchor.X+anchor.W || tip.X+tip.W < anchor.X {
		t.Errorf("the tip is at %v, beside the anchor at %v", tip, anchor)
	}
}

func TestTooltipDescribesItsAnchor(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTooltip(c, NewButton(c, "帮助").Element, "这是帮助文字")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	// A screen reader reads the tip with the anchor.
	info, ok := tt.Access("帮助")
	if !ok {
		t.Fatal("the anchor is not on screen")
	}
	if info.Description != "这是帮助文字" {
		t.Errorf("the anchor describes %q, want %q", info.Description, "这是帮助文字")
	}
}

func TestPopoverOpensOverItsAnchor(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Children(func() {
			ui.Box(c).Height(60)
			NewPopover(c, &open, func() ui.Element {
				return NewButton(c, "更多").Element
			}, func(panel ui.Element) {
				ui.Text(c, "面板里的内容")
			})
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)

	if _, ok := tt.Find("面板里的内容"); ok {
		t.Error("the panel shows before the anchor is pressed")
	}
	if err := tt.Click("更多"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	content, ok := tt.Find("面板里的内容")
	if !ok {
		t.Fatal("the panel is not on screen after the anchor is pressed")
	}
	// The content is in the panel, over the anchor, not in the corner.
	if content.Y == 0 && content.X == 0 {
		t.Error("the content is in the window's corner, not in the panel")
	}
}

func TestModalOpensOverTheWindow(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewModal(c, &open, func() ui.Element {
			return NewButton(c, "打开对话框").Element
		}, func(panel ui.Element) {
			ui.Text(c, "确定删除这个项目吗？")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)

	if _, ok := tt.Find("确定删除这个项目吗？"); ok {
		t.Error("the dialog shows before the trigger is pressed")
	}
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	content, ok := tt.Find("确定删除这个项目吗？")
	if !ok {
		t.Fatal("the dialog is not on screen after the trigger is pressed")
	}
	// The content is in the panel, in the middle of the window.
	if content.Y < 40 {
		t.Errorf("the content is at %v, not in the middle of the window", content)
	}
}

func TestModalClosesWithItsButton(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewModal(c, &open, func() ui.Element {
			return NewButton(c, "打开对话框").Element
		}, func(panel ui.Element) {
			ui.Text(c, "确定删除这个项目吗？")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)

	// The close button sits in the corner of the panel.
	close, ok := tt.Access("关闭")
	if !ok {
		t.Fatal("the close button is not on screen")
	}
	b := close.Bounds
	tt.ClickAt(b.X+b.W/2, b.Y+b.H/2)
	tt.Advance(time.Second)
	if _, ok := tt.Find("确定删除这个项目吗？"); ok {
		t.Error("the dialog is still on screen after the close button")
	}
	if open {
		t.Error("the flag is still true after the close button")
	}
}

func TestModalClosesWithEscape(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewModal(c, &open, func() ui.Element {
			return NewButton(c, "打开对话框").Element
		}, func(panel ui.Element) {
			ui.Text(c, "确定删除这个项目吗？")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("确定删除这个项目吗？"); !ok {
		t.Fatal("the dialog did not open")
	}
	tt.KeyDown(ui.KeyEscape)
	tt.KeyUp(ui.KeyEscape)
	tt.Advance(time.Second)
	if _, ok := tt.Find("确定删除这个项目吗？"); ok {
		t.Error("Escape did not close the dialog")
	}
}

func TestPopoverClosesWhenThePointerClicksOutside(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Children(func() {
			ui.Box(c).Height(60)
			NewPopover(c, &open, func() ui.Element {
				return NewButton(c, "更多").Element
			}, func(panel ui.Element) {
				ui.Text(c, "面板里的内容")
			})
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("更多"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板里的内容"); !ok {
		t.Fatal("the panel did not open")
	}
	// A click outside the panel closes it.
	tt.ClickAt(10, 10)
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板里的内容"); ok {
		t.Error("the panel is still on screen after a click outside it")
	}
}
