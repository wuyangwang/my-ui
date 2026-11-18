package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestAccordionAndCollapsible(t *testing.T) {
	open, section := false, false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(3)).Children(func() {
			NewAccordion(c, func() {
				NewAccordionItem(c, "详情", &open, func() { ui.Text(c, "详情内容") })
			})
			NewCollapsible(c, "更多", &section, func() { ui.Text(c, "更多内容") })
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	for _, s := range []string{"详情", "更多"} {
		if _, ok := tt.Find(s); !ok {
			t.Fatalf("%q is missing", s)
		}
	}
	time.Sleep(0)
}
