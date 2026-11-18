package display

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

func TestEmptyStateShowsItsParts(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewEmptyState(c, EmptyStateOptions{
			Icon:        icons.Inbox(),
			Title:       "还没有内容",
			Description: "新建一条记录",
			Action: func() {
				// the action slot is built below the description
			},
		})
	}
	tt := ui.NewTester(view, 320, 240)
	if _, ok := tt.Find("还没有内容"); !ok {
		t.Fatal("the title is missing")
	}
	if _, ok := tt.Find("新建一条记录"); !ok {
		t.Fatal("the description is missing")
	}
}
