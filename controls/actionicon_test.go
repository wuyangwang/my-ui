package controls

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"
	testutil "github.com/wuyangwang/my-ui/internal/testutil"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// TestActionIconNamesSizesAndClick checks that an icon-only action has
// an accessible name, the shared box sizes, and a click that reports.
func TestActionIconNamesSizesAndClick(t *testing.T) {
	clicked := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewActionIcon(c, icons.Copy(), "Copy").Xs()
			NewActionIcon(c, icons.Settings(), "Settings").Sm()
			del := NewActionIcon(c, icons.Trash(), "Delete").Outlined().Danger(true)
			if del.Clicked() {
				clicked = true
			}
		})
	}
	tt := ui.NewTester(view, 480, 200)
	tt.Advance(time.Second)
	for name, want := range map[string]float32{"Copy": 28, "Settings": 32, "Delete": 36} {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		if info.Bounds.W != want || info.Bounds.H != want {
			t.Errorf("%s is %vx%v, want %v", name, info.Bounds.W, info.Bounds.H, want)
		}
	}
	if err := tt.Click("Delete"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if !clicked {
		t.Error("the delete icon did not report its click")
	}
}

// TestActionIconFilledFollowsTheme checks the surface of a filled icon:
// the theme's surface at rest and its hover shade under the pointer.
func TestActionIconFilledFollowsTheme(t *testing.T) {
	var th *ui.Theme
	view := func(c *ui.Context) {
		th = Apply(c, Options{Appearance: AppearanceLight})
		NewActionIcon(c, icons.Settings(), "Settings").Filled()
	}
	tt := ui.NewTester(view, 200, 160)
	info, _ := tt.Access("Settings")
	x, y := info.Bounds.X+2, info.Bounds.Y+info.Bounds.H/2
	tt.Advance(time.Second)
	r, g, b := testutil.Pixel(tt, x, y)
	testutil.SamePixel(t, r, g, b, th.Surface, "the filled icon at rest")

	tt.Move(info.Bounds.X+info.Bounds.W/2, info.Bounds.Y+info.Bounds.H/2)
	tt.Advance(time.Second)
	r, g, b = testutil.Pixel(tt, x, y)
	testutil.SamePixel(t, r, g, b, th.SurfaceHover, "the filled icon hovered")
}
