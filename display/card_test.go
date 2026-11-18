package display

import (
	. "github.com/wuyangwang/my-ui"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestCardShowsItsTitleAndDescription(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewCard(c, "概览").Description("今天的情况").Content(func() {
			ui.Text(c, "卡片里的内容")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	for _, what := range []string{"概览", "今天的情况", "卡片里的内容"} {
		if _, ok := tt.Find(what); !ok {
			t.Errorf("the text %q is not on screen", what)
		}
	}
}
