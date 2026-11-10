package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSelectShowsTheLabelOfTheChoice(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{
			{Label: "简体中文", Value: "zh-Hans"},
			{Label: "繁體中文", Value: "zh-Hant"},
		}).Outline().Label("语言")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	// Nothing is chosen yet, so the trigger shows no label.
	if _, ok := tt.Find("简体中文"); ok {
		t.Error("the trigger shows a label before a choice is made")
	}
	if _, ok := tt.Access("语言"); !ok {
		t.Fatal("the select is not on screen")
	}
}

func TestSelectOpensAndChooses(t *testing.T) {
	var value string
	var seen string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{
			{Label: "简体中文", Value: "zh-Hans"},
			{Label: "繁體中文", Value: "zh-Hant"},
		}).Outline().Label("语言").OnChanged(func(v string) { seen = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)

	if err := tt.Click("语言"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("繁體中文"); !ok {
		t.Fatal("the popup did not open")
	}
	if err := tt.Click("繁體中文"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if value != "zh-Hant" {
		t.Errorf("the select holds %q, want %q", value, "zh-Hant")
	}
	if seen != "zh-Hant" {
		t.Errorf("the change reported %q, want %q", seen, "zh-Hant")
	}
	// The trigger shows the label of the choice, not its value.
	if _, ok := tt.Find("繁體中文"); !ok {
		t.Error("the trigger does not show the label of the choice")
	}
}

func TestSelectPlaceholderShowsWhileNothingIsChosen(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{{Value: "小"}}).Outline().
			Placeholder("选择尺寸").Label("尺寸")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("选择尺寸"); !ok {
		t.Error("the placeholder is not on screen")
	}
}

func TestSelectDisabledTakesNoClick(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewSelect(c, &value, []SelectOption{{Value: "小"}}).Outline().
			Label("尺寸").Disabled(true)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("尺寸")
	if !ok {
		t.Fatal("the select is not on screen")
	}
	if !info.Disabled {
		t.Error("the select does not read as disabled")
	}
	if err := tt.Click("尺寸"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("小"); ok {
		t.Error("a disabled select opened its popup")
	}
}

func TestSelectSizesGiveRoomierBounds(t *testing.T) {
	height := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds.H
	}
	var a, b string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewSelect(c, &a, []SelectOption{{Value: "小"}}).Sm().Outline().Label("small")
			NewSelect(c, &b, []SelectOption{{Value: "中"}}).Md().Outline().Label("medium")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if h := height(tt, "small"); h >= height(tt, "medium") {
		t.Errorf("the small select is not shorter than the middle's: %v against %v", h, height(tt, "medium"))
	}
}
