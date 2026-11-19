package display

import (
	. "github.com/wuyangwang/my-ui"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestTabsShowTheContentOfTheSelectedTab(t *testing.T) {
	selected := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTabs(c, &selected,
			Tab{Label: "概览", Content: func() { ui.Text(c, "概览的内容") }},
			Tab{Label: "设置", Content: func() { ui.Text(c, "设置的内容") }},
		)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("概览的内容"); !ok {
		t.Fatal("the content of the selected tab is not on screen")
	}
	if _, ok := tt.Find("设置的内容"); ok {
		t.Error("the content of a tab that is not selected shows")
	}
	if err := tt.Click("设置"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("设置的内容"); !ok {
		t.Error("the content of the tab that is selected is not on screen")
	}
	if _, ok := tt.Find("概览的内容"); ok {
		t.Error("the content of the tab that is not selected still shows")
	}
}

func TestTabsVerticalPointAndSize(t *testing.T) {
	selected := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTabs(c, &selected,
			Tab{Label: "一", Content: func() { ui.Text(c, "第一块") }},
			Tab{Label: "二", Content: func() { ui.Text(c, "第二块") }},
		).Vertical().Point().Size(SizeSm)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("第一块"); !ok {
		t.Fatal("the content of the selected vertical tab is missing")
	}
	if err := tt.Click("二"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("第二块"); !ok {
		t.Error("the content of the chosen vertical tab is missing")
	}
}

func TestTabsOnChanged(t *testing.T) {
	selected := 0
	changed := -1
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTabs(c, &selected,
			Tab{Label: "甲", Content: func() { ui.Text(c, "甲内容") }},
			Tab{Label: "乙", Content: func() { ui.Text(c, "乙内容") }},
		).OnChanged(func(i int) { changed = i })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("乙"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if changed != 1 {
		t.Errorf("OnChanged reported %d, want 1", changed)
	}
}
