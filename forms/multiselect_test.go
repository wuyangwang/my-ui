package forms

import (
	. "github.com/wuyangwang/my-ui"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestMultiSelectAddsAndRemovesTokens(t *testing.T) {
	var tokens []string
	var seen []string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewMultiSelect(c, &tokens, []SelectOption{
			{Label: "Alpha", Value: "a"},
			{Label: "Beta", Value: "b"},
		}).Outline().Label("选择").OnChanged(func(v []string) { seen = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)

	// Typing and pressing Enter adds a token.
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Type("Alpha")
	tt.Advance(time.Second)
	tt.KeyDown(ui.KeyEnter)
	tt.KeyUp(ui.KeyEnter)
	tt.Advance(time.Second)
	if len(tokens) != 1 || tokens[0] != "Alpha" {
		t.Fatalf("the field holds %v, want the one token", tokens)
	}
	if len(seen) != 1 {
		t.Errorf("the change reported %v, want the one token", seen)
	}
}

func TestMultiSelectShowsItsTokens(t *testing.T) {
	tokens := []string{"Alpha"}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewMultiSelect(c, &tokens, []SelectOption{
			{Label: "Alpha", Value: "a"},
			{Label: "Beta", Value: "b"},
		}).Outline().Label("选择")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("Alpha"); !ok {
		t.Error("the token is not on screen")
	}
}
