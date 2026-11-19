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

// TestMultiSelectMaxDoesNotDoubleConsumeAClick builds the field with a
// chained Max in the same frame. The fluent setter must not build a second
// token field: one click on a chip's remove button must remove one token and
// report one change, never two.
func TestMultiSelectMaxDoesNotDoubleConsumeAClick(t *testing.T) {
	tokens := []string{"Alpha", "Beta"}
	changes := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewMultiSelect(c, &tokens, []SelectOption{
			{Label: "Alpha", Value: "a"},
			{Label: "Beta", Value: "b"},
		}).Outline().Label("选择").Max(2).
			OnChanged(func(v []string) { changes++ })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	tt.Advance(time.Second)

	if err := tt.Click("Remove Alpha"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	tt.Advance(time.Second)

	if len(tokens) != 1 || tokens[0] != "Beta" {
		t.Fatalf("one remove click left %v, want [Beta]", tokens)
	}
	if changes != 1 {
		t.Errorf("one remove click reported %d changes, want 1", changes)
	}
}

// TestMultiSelectMaxAddsOncePerClick is the same guarantee for adding: a
// click on a suggestion adds exactly one token when Max is chained in the
// frame.
func TestMultiSelectMaxAddsOncePerClick(t *testing.T) {
	var tokens []string
	changes := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewMultiSelect(c, &tokens, []SelectOption{
			{Label: "Alpha", Value: "a"},
			{Label: "Beta", Value: "b"},
		}).Outline().Label("选择").Max(2).
			OnChanged(func(v []string) { changes++ })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	tt.Advance(time.Second)

	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Type("Al")
	tt.Advance(time.Second)
	tt.Advance(time.Second)
	if err := tt.Click("Alpha"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	tt.Advance(time.Second)

	if len(tokens) != 1 || tokens[0] != "Alpha" {
		t.Fatalf("one add click left %v, want [Alpha]", tokens)
	}
	if changes != 1 {
		t.Errorf("one add click reported %d changes, want 1", changes)
	}
}
