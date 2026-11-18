package display

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestTableDrawsHeadersAndCells(t *testing.T) {
	var state ui.ListState
	columns := []ui.TableColumn{{Title: "名称", Width: 120}, {Title: "数量", Width: 80}}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTable(c, &state, columns, 2, func(row, col int) {
			if col == 0 {
				ui.Text(c, []string{"苹果", "梨"}[row])
			} else {
				ui.Text(c, []string{"2", "5"}[row])
			}
		})
	}
	tt := ui.NewTester(view, 320, 240)
	if _, ok := tt.Find("名称"); !ok {
		t.Fatal("the header is missing")
	}
	if _, ok := tt.Find("苹果"); !ok {
		t.Fatal("the cell is missing")
	}
	if _, ok := tt.Find("5"); !ok {
		t.Fatal("the last cell is missing")
	}
}
