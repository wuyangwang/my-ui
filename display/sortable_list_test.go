package display

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestSortableListDrawsItsRows(t *testing.T) {
	var state ui.ListState
	state.Reorder = func(rows []int, to int) {}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSortableList(c, &state, 3, func(i int) {
			ui.Text(c, []string{"苹果", "香蕉", "梨"}[i])
		})
	}
	tt := ui.NewTester(view, 320, 240)
	for _, s := range []string{"苹果", "香蕉", "梨"} {
		if _, ok := tt.Find(s); !ok {
			t.Fatalf("the row %q is missing", s)
		}
	}
}
