package overlays

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestDropdownMenuChoosesAnItem(t *testing.T) {
	open := false
	chosen := ""
	view := func(c *ui.Context) {
		NewDropdownMenu(c, &open, func() ui.Element {
			return ui.Text(c, "更多")
		}, []MenuItem{
			{Label: "复制", OnClick: func() { chosen = "复制" }},
			{Separator: true},
			{Label: "删除", OnClick: func() { chosen = "删除" }},
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("更多"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	t.Logf("open after click: %v", open)
	if _, ok := tt.Find("删除"); !ok {
		t.Fatal("the menu panel did not show the items")
	}
	if err := tt.Click("删除"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if chosen != "删除" {
		t.Errorf("the chosen item is %q, want 删除", chosen)
	}
	if open {
		t.Error("choosing an item did not close the menu")
	}
}
