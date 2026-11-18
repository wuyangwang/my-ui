package overlays

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestContextMenuChoosesAnItem(t *testing.T) {
	open := false
	chosen := ""
	view := func(c *ui.Context) {
		target := ui.Row(c).Height(40).Children(func() {
			ui.Text(c, "目标")
		})
		NewContextMenu(c, &open, target, []MenuItem{
			{Label: "重命名", OnClick: func() { chosen = "重命名" }},
			{Label: "移除", OnClick: func() { chosen = "移除" }},
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.RightClick("目标"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("重命名"); !ok {
		t.Fatal("the context menu did not show the items")
	}
	if err := tt.Click("重命名"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if chosen != "重命名" {
		t.Errorf("the chosen item is %q, want 重命名", chosen)
	}
}
