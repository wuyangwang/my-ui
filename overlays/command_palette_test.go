package overlays

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestCommandPaletteFiltersAndRuns(t *testing.T) {
	open, query := true, ""
	called := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewCommandPalette(c, &open, &query, []Command{
			{Label: "新建文件", OnSelect: func() { called = "新建文件" }},
			{Label: "打开文件", OnSelect: func() { called = "打开文件" }},
		})
	}
	tt := ui.NewTester(view, 640, 480)
	tt.Advance(time.Second)
	if err := tt.Click("打开文件"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if called != "打开文件" {
		t.Errorf("the selected command is %q, want 打开文件", called)
	}
	if open {
		t.Error("choosing a command did not close the palette")
	}
}
