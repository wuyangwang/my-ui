package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestHeaderSlots(t *testing.T) {
	view := func(c *ui.Context) {
		NewHeader(c, HeaderProps{
			Logo:    func() { ui.Text(c, "品牌") },
			Nav:     func() { ui.Text(c, "导航") },
			Actions: func() { ui.Text(c, "操作") },
		})
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	for _, s := range []string{"品牌", "导航", "操作"} {
		if !tt.HasText(s) {
			t.Errorf("the header slot %q is not on screen", s)
		}
	}
}

func TestMobileHeaderPutsNavFirst(t *testing.T) {
	view := func(c *ui.Context) {
		NewHeader(c, HeaderProps{
			Mobile: true,
			Nav:    func() { ui.Text(c, "菜单") },
			Logo:   func() { ui.Text(c, "标题") },
		})
	}
	tt := ui.NewTester(view, 360, 240)
	tt.Advance(time.Second)
	nav, ok1 := tt.Find("菜单")
	logo, ok2 := tt.Find("标题")
	if !ok1 || !ok2 {
		t.Fatal("the mobile header did not draw both slots")
	}
	if nav.X >= logo.X {
		t.Errorf("the menu must lead the brand: menu x=%v, brand x=%v", nav.X, logo.X)
	}
}
