package tests

import (
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
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
			overlaypkg.NewTooltip(c, ctrlpkg.NewButton(c, "帮助").Element, "这是帮助文字")
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
		overlaypkg.NewTooltip(c, ctrlpkg.NewButton(c, "帮助").Element, "这是帮助文字")
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

// TestTooltipDrawsItsText looks at the pixels: the tip is the theme
// turned over, and the text of it is the ink on that ground.
func TestTooltipDrawsItsText(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Children(func() {
			ui.Box(c).Height(140)
			overlaypkg.NewTooltip(c, ctrlpkg.NewButton(c, "帮助").Element, "这是帮助文字")
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	anchor, ok := tt.Find("帮助")
	if !ok {
		t.Fatal("the anchor is not on screen")
	}
	tt.Move(anchor.X+anchor.W/2, anchor.Y+anchor.H/2)
	tt.Advance(2 * time.Second)
	text, ok := tt.Find("这是帮助文字")
	if !ok {
		t.Fatal("the tip is not on screen")
	}
	tip, ok := tt.AccessByRole(ui.RoleTooltip, "")
	if !ok {
		t.Fatal("the tip is not in the tree of the assistive technology")
	}
	// The tip has a ground of its own: the ink of the theme.
	b := tip.Bounds
	cr, _, _, _ := tt.Image().At(int(b.X+2), int(b.Y+b.H/2)).RGBA()
	if cr > 0x8000 {
		t.Error("the tip has no ground of its own")
	}
	// The text of it is drawn in the ink on that ground.
	light := 0
	for x := int(text.X); x < int(text.X+text.W); x++ {
		for y := int(text.Y); y < int(text.Y+text.H); y++ {
			r, g, bl, _ := tt.Image().At(x, y).RGBA()
			if r > 0xA000 && g > 0xA000 && bl > 0xA000 {
				light++
			}
		}
	}
	if light == 0 {
		t.Error("the text of the tip is not drawn")
	}
}

func TestPopoverOpensOverItsAnchor(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Children(func() {
			ui.Box(c).Height(60)
			overlaypkg.NewPopover(c, overlaypkg.PopoverOptions{
				Open: &open,
				Trigger: func() ui.Element {
					return ctrlpkg.NewButton(c, "更多").Element
				},
				Body: func() { ui.Text(c, "面板里的内容") },
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
		overlaypkg.NewModal(c, overlaypkg.ModalOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "打开对话框").Element
			},
			Body: func() { ui.Text(c, "确定删除这个项目吗？") },
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
		overlaypkg.NewModal(c, overlaypkg.ModalOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "打开对话框").Element
			},
			Body: func() { ui.Text(c, "确定删除这个项目吗？") },
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)

	// The close button sits in the corner of the panel.
	close, ok := tt.Access("Close")
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
		overlaypkg.NewModal(c, overlaypkg.ModalOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "打开对话框").Element
			},
			Body: func() { ui.Text(c, "确定删除这个项目吗？") },
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
			overlaypkg.NewPopover(c, overlaypkg.PopoverOptions{
				Open: &open,
				Trigger: func() ui.Element {
					return ctrlpkg.NewButton(c, "更多").Element
				},
				Body: func() { ui.Text(c, "面板里的内容") },
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

// TestModalCentersItsButtons checks the dialog's layout: the buttons
// of it sit in the middle of the panel, as the content around them does.
func TestModalCentersItsButtons(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		overlaypkg.NewModal(c, overlaypkg.ModalOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "打开").Element
			},
			Body: func() {
				ui.Column(c).AlignItems(ui.Center).Children(func() {
					ui.Text(c, "确定吗？")
					ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
						ctrlpkg.NewButton(c, "取消")
						ctrlpkg.NewButton(c, "删除")
					})
				})
			},
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("打开"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	cancel, ok := tt.Find("取消")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	remove, ok := tt.Find("删除")
	if !ok {
		t.Fatal("the button is not on screen")
	}
	dialog, ok := tt.AccessByRole(ui.RoleDialog, "")
	if !ok {
		t.Fatal("the dialog is not in the tree of the assistive technology")
	}
	b := dialog.Bounds
	left := cancel.X - b.X
	right := (b.X + b.W) - (remove.X + remove.W)
	if diff := left - right; diff > 2 || diff < -2 {
		t.Errorf("the buttons sit %v from the left and %v from the right of the dialog", left, right)
	}
}
