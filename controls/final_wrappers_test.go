package controls

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestToggleAndColorWell(t *testing.T) {
	on := false
	color := ui.Hex("#ff0000")
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(4)).Children(func() {
			NewToggle(c, &on, "开关").Size(SizeSm)
			NewColorWell(c, &color)
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if err := tt.Click("开关"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !on {
		t.Error("the toggle did not report its state")
	}
}
