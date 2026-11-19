package tests

import (
	. "github.com/wuyangwang/my-ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// TestSelectInACardChoosesAnOption puts the select in a card in a
// column, as the showcase does, and chooses an option: the click lands
// on an option of the popup.
func TestSelectInACardChoosesAnOption(t *testing.T) {
	var value string
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(4)).Children(func() {
			formspkg.NewSelect(c.Key("sel"), &value, []formspkg.SelectOption{
				{Label: "简体中文", Value: "zh-Hans"},
				{Label: "English", Value: "en"},
			}).Outline().Label("语言").OnChanged(func(v string) { chosen = v })
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("语言"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("English"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if value != "en" || chosen != "en" {
		t.Errorf("the choice is %q / %q, want %q", value, chosen, "en")
	}
}

// TestComboBoxInACardChoosesAfterSearching puts the combo box in a card
// in a column, as the showcase does: the user searches, and the click
// lands on an option.
func TestComboBoxInACardChoosesAfterSearching(t *testing.T) {
	var text string
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(4)).Children(func() {
			formspkg.NewComboBox(c.Key("cb"), formspkg.ComboOptions{
				Text: &text,
				Options: []formspkg.SelectOption{
					{Label: "简体中文", Value: "zh-Hans"},
					{Label: "English", Value: "en"},
					{Label: "日本語", Value: "ja"},
				},
			}).Outline().Label("语言").OnChanged(func(v string) { chosen = v })
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("语言"); err != nil {
		t.Fatal(err)
	}
	tt.Type("E")
	tt.Advance(time.Second)
	if err := tt.Click("English"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if chosen != "en" {
		t.Errorf("the choice reported %q, want %q", chosen, "en")
	}
}
