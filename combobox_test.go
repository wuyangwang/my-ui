package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestComboBoxFiltersItsOptions(t *testing.T) {
	var text string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, &text, []SelectOption{
			{Label: "Alpha", Value: "a"},
			{Label: "Beta", Value: "b"},
		}).Outline().Label("选择")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)

	// Typing in the field narrows the options down.
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Type("B")
	tt.Advance(time.Second)
	if _, ok := tt.Find("Beta"); !ok {
		t.Error("the option that holds the text is not on screen")
	}
	if _, ok := tt.Find("Alpha"); ok {
		t.Error("the option that does not hold the text is still on screen")
	}
}

func TestComboBoxChoosesAnOption(t *testing.T) {
	var text string
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, &text, []SelectOption{
			{Label: "Alpha", Value: "a"},
			{Label: "Beta", Value: "b"},
		}).Outline().Label("选择").OnChanged(func(v string) { chosen = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("Beta"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if chosen != "b" {
		t.Errorf("the choice reported %q, want %q", chosen, "b")
	}
}

func TestComboBoxGroupsItsOptions(t *testing.T) {
	var text string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, &text, []SelectOption{
			{Group: "最近", Label: "Alpha", Value: "a"},
			{Group: "更多", Label: "Beta", Value: "b"},
		}).Outline().Label("选择")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("最近"); !ok {
		t.Error("the group is not on screen")
	}
	if _, ok := tt.Find("更多"); !ok {
		t.Error("the second group is not on screen")
	}
}
