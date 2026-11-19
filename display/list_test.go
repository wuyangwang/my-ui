package display

import (
	"testing"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestListDrawsItsRowsAndKeepsSelection(t *testing.T) {
	var state ui.ListState
	state.Key = func(i int) any { return i }
	selected := 1
	state.Selected = &selected
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewList(c, &state, 3, func(i int) {
			ui.Text(c, []string{"苹果", "香蕉", "梨"}[i])
		})
	}
	tt := ui.NewTester(view, 320, 240)
	for _, s := range []string{"苹果", "香蕉", "梨"} {
		if _, ok := tt.Find(s); !ok {
			t.Fatalf("the row %q is missing", s)
		}
	}
	if err := tt.Click("梨"); err != nil {
		t.Fatal(err)
	}
	if selected != 2 {
		t.Errorf("the choice is %d, want 2", selected)
	}
}

func TestListWithoutStateDrawsItsRows(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewList(c, nil, 2, func(i int) {
			ui.Text(c, []string{"一", "二"}[i])
		})
	}
	tt := ui.NewTester(view, 320, 240)
	for _, s := range []string{"一", "二"} {
		if _, ok := tt.Find(s); !ok {
			t.Errorf("the row %q is missing", s)
		}
	}
}

func TestListFollowsTheTheme(t *testing.T) {
	for _, a := range []Appearance{AppearanceLight, AppearanceDark} {
		view := func(c *ui.Context) {
			Apply(c, Options{Appearance: a})
			NewList(c, nil, 1, func(int) { ui.Text(c, "主题行") })
		}
		tt := ui.NewTester(view, 320, 240)
		if _, ok := tt.Find("主题行"); !ok {
			t.Errorf("%v: the row is missing", a)
		}
	}
}
