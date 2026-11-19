package display

import (
	"strings"
	"testing"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

var outlineTree = map[string][]string{
	"src":       {"src/ui", "src/main.go"},
	"src/ui":    {"src/ui/list.go"},
	"docs":      nil,
	"README.md": nil,
}

func outlineChildren(path string) []string {
	if path == "" {
		return []string{"src", "docs", "README.md"}
	}
	return outlineTree[path]
}

func outlineTitle(path string) string {
	if i := strings.LastIndexByte(path, '/'); i >= 0 {
		return path[i+1:]
	}
	return path
}

func TestOutlineShowsRootsAndOpensABranch(t *testing.T) {
	var state ui.OutlineState[string]
	selected := -1
	state.List.Selected = &selected
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewOutline(c, &state, outlineChildren, outlineTitle)
	}
	tt := ui.NewTester(view, 400, 400)
	if _, ok := tt.Find("src"); !ok {
		t.Fatal("the root src is missing")
	}
	if tt.HasText("main.go") {
		t.Fatal("a closed branch shows its children")
	}
	src, _ := tt.Find("src")
	tt.ClickAt(src.X-10, src.Y+src.H/2)
	if !state.Open.Has("src") || !tt.HasText("main.go") {
		t.Fatalf("the arrow did not open src; open %v", state.Open.Has("src"))
	}
	uiRow, _ := tt.Find("ui")
	tt.ClickAt(uiRow.X-10, uiRow.Y+uiRow.H/2)
	if !tt.HasText("list.go") {
		t.Fatal("opening src/ui did not show its child")
	}
	if err := tt.Click("list.go"); err != nil {
		t.Fatal(err)
	}
	if got := state.Item(selected); got != "src/ui/list.go" {
		t.Errorf("the choice is %q, want src/ui/list.go", got)
	}
}

func TestOutlineGenericKeysAndTitleFallback(t *testing.T) {
	var state ui.OutlineState[int]
	children := func(k int) []int {
		if k == 0 {
			return []int{1, 2}
		}
		return nil
	}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewOutline(c, &state, children, nil)
	}
	tt := ui.NewTester(view, 400, 400)
	for _, s := range []string{"1", "2"} {
		if _, ok := tt.Find(s); !ok {
			t.Errorf("the node %q is missing", s)
		}
	}
}

func TestOutlineFollowsTheTheme(t *testing.T) {
	for _, a := range []Appearance{AppearanceLight, AppearanceDark} {
		var state ui.OutlineState[string]
		view := func(c *ui.Context) {
			Apply(c, Options{Appearance: a})
			NewOutline(c, &state, outlineChildren, outlineTitle)
		}
		tt := ui.NewTester(view, 400, 400)
		if _, ok := tt.Find("README.md"); !ok {
			t.Errorf("%v: the root is missing", a)
		}
	}
}
