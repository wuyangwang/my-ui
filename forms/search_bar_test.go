package forms

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestSearchBarTypesAndClears(t *testing.T) {
	query := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Width(300).Children(func() {
			NewSearchBar(c, &query).Placeholder("搜索").Label("搜索框")
		})
	}
	tt := ui.NewTester(view, 320, 120)
	field, ok := tt.Access("搜索框")
	if !ok {
		t.Fatal("the search field is missing")
	}
	tt.ClickAt(field.Bounds.X+field.Bounds.W/2, field.Bounds.Y+field.Bounds.H/2)
	tt.Type("abc")
	tt.Advance(time.Second)
	if query != "abc" {
		t.Fatalf("the query is %q, want abc", query)
	}
	clear, ok := tt.Access("Clear")
	if !ok {
		t.Fatal("the clear icon is missing")
	}
	x, y := clear.Bounds.X+clear.Bounds.W/2, clear.Bounds.Y+clear.Bounds.H/2
	tt.Press(x, y)
	tt.Advance(time.Second)
	tt.Release(x, y)
	tt.Advance(time.Second)
	if query != "" {
		t.Errorf("the clear icon left %q", query)
	}
}
