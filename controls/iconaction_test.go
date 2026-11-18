package controls

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// TestIconActionSharedSizeAndHover checks the icon affordance used by
// clear and close icons in every component: a 20 DIP hit area by
// default, a custom one when asked, and ink that changes on hover.
func TestIconActionSharedSizeAndHover(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(4)).Children(func() {
			NewIconAction(c, IconActionOptions{SVG: icons.X(), Label: "default"})
			NewIconAction(c, IconActionOptions{SVG: icons.X(), Label: "small", Size: 16, IconSize: IconSizeXs})
		})
	}
	tt := ui.NewTester(view, 240, 120)
	tt.Advance(time.Second)
	for name, want := range map[string]float32{"default": 20, "small": 16} {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		if info.Bounds.W != want || info.Bounds.H != want {
			t.Errorf("%s is %vx%v, want %v", name, info.Bounds.W, info.Bounds.H, want)
		}
	}
	info, _ := tt.Access("default")
	x, y := info.Bounds.X+info.Bounds.W/2, info.Bounds.Y+info.Bounds.H/2
	rest := ui.Hex("#71717a")
	r, g, b := pixelOf(tt, x, y)
	if absDiff(int(r), int(rest.R)) > 12 || absDiff(int(g), int(rest.G)) > 12 {
		t.Errorf("the icon at rest drew (%d, %d, %d), want muted ink", r, g, b)
	}
	tt.Move(x, y)
	tt.Advance(time.Second)
	r2, g2, b2 := pixelOf(tt, x, y)
	if r == r2 && g == g2 && b == b2 {
		t.Error("the icon keeps its muted ink on hover")
	}
}

func pixelOf(tt *ui.Tester, x, y float32) (uint8, uint8, uint8) {
	p := tt.Image().At(int(x), int(y))
	r, g, b, _ := p.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
