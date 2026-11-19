package overlays

import (
	"runtime"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

func TestTooltipShowsItsHotkey(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Children(func() {
			// The anchor sits low so the tip has room above it: a tip
			// that lands in the corner shows the mistake this test checks.
			ui.Box(c).Height(140)
			NewTooltip(c, ctrlpkg.NewButton(c, "帮助").Element, "这是帮助文字").Hotkey("mod+k")
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
	mod := "Ctrl"
	if runtime.GOOS == "darwin" {
		mod = "Cmd"
	}
	for _, key := range []string{mod, "K"} {
		if _, ok := tt.Find(key); !ok {
			t.Errorf("the key %q of the hotkey is not on screen", key)
		}
	}
	// The key caps are built inside the tip, next to the text, not in the
	// window's corner where overlay content lands without tip.Children.
	cap, ok := tt.Find(mod)
	if !ok {
		t.Fatalf("the key %q is not on screen", mod)
	}
	if cap.X == 0 && cap.Y == 0 {
		t.Error("the hotkey is in the window's corner, not in the tip")
	}
	if cap.X < text.X || cap.X > text.X+text.W+40 {
		t.Errorf("the hotkey is at %v, not beside the text at %v", cap, text)
	}
	if cap.Y < text.Y-4 || cap.Y > text.Y+text.H+4 {
		t.Errorf("the hotkey is at %v, not on the text's row at %v", cap, text)
	}
}

// TestTooltipOptionsHotkey covers setting the same value at construction
// time, before the tip ever builds.
func TestTooltipOptionsHotkey(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Children(func() {
			ui.Box(c).Height(140)
			NewTooltip(c, ctrlpkg.NewButton(c, "帮助").Element, "这是帮助文字",
				TooltipOptions{Hotkey: "mod+k"})
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
	mod := "Ctrl"
	if runtime.GOOS == "darwin" {
		mod = "Cmd"
	}
	for _, key := range []string{mod, "K"} {
		if _, ok := tt.Find(key); !ok {
			t.Errorf("the key %q of the hotkey is not on screen", key)
		}
	}
}
