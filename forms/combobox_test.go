package forms

import (
	. "github.com/wuyangwang/my-ui"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestComboBoxFiltersItsOptions(t *testing.T) {
	var text string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Text: &text,
			Options: []SelectOption{
				{Label: "Alpha", Value: "a"},
				{Label: "Beta", Value: "b"},
			},
			Label: "选择",
		}).Outline()
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
	var value string
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Value: &value,
			Options: []SelectOption{
				{Label: "Alpha", Value: "a"},
				{Label: "Beta", Value: "b"},
			},
			Label: "选择",
		}).Outline().OnChanged(func(v string) { chosen = v })
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
	// The selected value is the option's value, not its label.
	if value != "b" {
		t.Errorf("the selected value is %q, want %q", value, "b")
	}
}

func TestComboBoxChosenAndSelectedReadAcrossFrames(t *testing.T) {
	var value string
	var combo *ComboBox
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		combo = NewComboBox(c, ComboOptions{
			Value: &value,
			Options: []SelectOption{
				{Label: "Alpha", Value: "a"},
				{Label: "Beta", Value: "b"},
			},
			Label: "选择",
		}).Outline()
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
	tt.Advance(time.Second) // a frame with no choice must not expire them
	if combo.Chosen() != "b" {
		t.Errorf("Chosen() is %q after the frame, want %q", combo.Chosen(), "b")
	}
	if combo.Selected() != "b" {
		t.Errorf("Selected() is %q after the frame, want %q", combo.Selected(), "b")
	}
}

func TestComboBoxChoosesAnOptionAfterSearching(t *testing.T) {
	var value string
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Value: &value,
			Options: []SelectOption{
				{Label: "Alpha", Value: "a"},
				{Label: "Beta", Value: "b"},
				{Label: "Gamma", Value: "g"},
			},
			Label: "选择",
		}).Outline().OnChanged(func(v string) { chosen = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)

	// Search first, then choose: the options that do not hold the text
	// are left out of the popup.
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Type("B")
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
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Options: []SelectOption{
				{Group: "最近", Label: "Alpha", Value: "a"},
				{Group: "更多", Label: "Beta", Value: "b"},
			},
			Label: "选择",
		}).Outline()
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

func TestComboBoxClearRestoresAnEmptyValue(t *testing.T) {
	value := "b"
	text := "Beta"
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Value:     &value,
			Text:      &text,
			Options:   []SelectOption{{Label: "Alpha", Value: "a"}, {Label: "Beta", Value: "b"}},
			Clearable: true,
			Label:     "选择",
		}).Outline()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Access("Clear"); !ok {
		t.Fatal("the clear button is not on screen while a value is selected")
	}
	if err := tt.Click("Clear"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if value != "" {
		t.Errorf("the selected value is %q after clearing, want empty", value)
	}
	if text != "" {
		t.Errorf("the query is %q after clearing, want empty", text)
	}
}

func TestComboBoxShowsEmptyTextWhenNothingMatches(t *testing.T) {
	var text string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Text:      &text,
			Options:   []SelectOption{{Label: "Alpha", Value: "a"}},
			EmptyText: "没有匹配项",
			Label:     "选择",
		}).Outline()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Type("zzz")
	tt.Advance(time.Second)
	if _, ok := tt.Find("没有匹配项"); !ok {
		t.Error("the empty text does not show when nothing matches")
	}
}

func TestComboBoxHasADefaultEmptyText(t *testing.T) {
	var text string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Text:    &text,
			Options: []SelectOption{{Label: "Alpha", Value: "a"}},
			Label:   "选择",
		}).Outline()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Type("zzz")
	tt.Advance(time.Second)
	if _, ok := tt.Find("No results found"); !ok {
		t.Error("the default empty text does not show when nothing matches")
	}
}

func TestComboBoxCallbackAttachedAfterConstructionFires(t *testing.T) {
	var value string
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewComboBox(c, ComboOptions{
			Value:   &value,
			Options: []SelectOption{{Label: "Alpha", Value: "a"}, {Label: "Beta", Value: "b"}},
			Label:   "选择",
		}).Outline().OnChanged(func(v string) { chosen = v })
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
		t.Errorf("the callback attached after construction reported %q, want %q", chosen, "b")
	}
}

func TestComboBoxCallbackAttachedAfterTheChoiceFires(t *testing.T) {
	var value string
	chosen := ""
	attach := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		combo := NewComboBox(c, ComboOptions{
			Value:   &value,
			Options: []SelectOption{{Label: "Alpha", Value: "a"}, {Label: "Beta", Value: "b"}},
			Label:   "选择",
		}).Outline()
		if attach {
			combo.OnChanged(func(v string) { chosen = v })
		}
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
	if chosen != "" {
		t.Fatalf("the callback fired before it was attached: %q", chosen)
	}
	// Attach on a later frame: the choice is kept until a callback is
	// present, so late attachment still receives it.
	attach = true
	tt.Advance(time.Second)
	if chosen != "b" {
		t.Errorf("the callback attached after the choice reported %q, want %q", chosen, "b")
	}
}
