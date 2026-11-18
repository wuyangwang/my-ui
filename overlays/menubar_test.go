package overlays

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestMenubarOpensTheChosenMenu(t *testing.T) {
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewMenubar(c, []MenubarMenu{
			{Label: "文件", Items: []MenuItem{
				{Label: "打开", OnClick: func() { chosen = "打开" }},
				{Label: "保存", OnClick: func() { chosen = "保存" }},
			}},
			{Label: "编辑", Items: []MenuItem{
				{Label: "撤销", OnClick: func() { chosen = "撤销" }},
			}},
		})
	}
	tt := ui.NewTester(view, 480, 240)
	tt.Advance(time.Second)
	if err := tt.Click("文件"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("保存"); !ok {
		t.Fatal("the file menu did not open")
	}
	if err := tt.Click("保存"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if chosen != "保存" {
		t.Errorf("the chosen item is %q, want 保存", chosen)
	}
}
