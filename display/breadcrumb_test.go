package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestBreadcrumbWritesTheChosenCrumb(t *testing.T) {
	chosen := 2
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewBreadcrumb(c, []string{"首页", "列表", "详情"}, &chosen)
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	for _, s := range []string{"首页", "列表", "详情"} {
		if _, ok := tt.Find(s); !ok {
			t.Fatalf("the crumb %q is missing", s)
		}
	}
	if err := tt.Click("首页"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if chosen != 0 {
		t.Errorf("the chosen crumb is %d, want 0", chosen)
	}
}
