package tests

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// TestShadcnGeometry checks button heights against shadcn's scale:
// h-7/h-8/h-9/h-10. The field frame's inset uses the same spacing
// scale; its labeled element is the inner editor, not the frame.
func TestShadcnGeometry(t *testing.T) {
	view := func(c *ui.Context) {
		th := Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(th.Space(2)).Children(func() {
			ctrlpkg.NewButton(c, "xs").Xs().Label("btn-xs")
			ctrlpkg.NewButton(c, "sm").Sm().Label("btn-sm")
			ctrlpkg.NewButton(c, "md").Label("btn-md")
			ctrlpkg.NewButton(c, "lg").Lg().Label("btn-lg")
		})
	}
	tt := ui.NewTester(view, 320, 400)
	tt.Advance(time.Second)
	for name, want := range map[string]float32{
		"btn-xs": 28, "btn-sm": 32, "btn-md": 36, "btn-lg": 40,
	} {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		if absSize(info.Bounds.H-want) > 1 {
			t.Errorf("%s is %.1f high, want %.0f", name, info.Bounds.H, want)
		}
	}
}

func absSize(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
