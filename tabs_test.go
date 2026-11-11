package myui

import (
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
	// Another tab shows its own content.
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
