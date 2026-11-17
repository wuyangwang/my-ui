package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSemanticColorComesFromTheTheme(t *testing.T) {
	th := *shadcnLight()
	custom := ui.Hex("#123456")
	th.Success = custom
	if got := semanticColor(&th, semanticSuccess); got != custom {
		t.Errorf("success color %v, want the theme's %v", got, custom)
	}
	// Info has no theme slot, so it keeps the fixed table.
	if got := semanticColor(&th, semanticInfo); got != semantics[semanticInfo][0] {
		t.Errorf("info color %v, want the fixed one", got)
	}
}

func TestDarkSemanticColorComesFromTheTheme(t *testing.T) {
	th := *shadcnDark()
	if got, want := semanticColor(&th, semanticSuccess), th.Success; got != want {
		t.Errorf("dark success %v, want the theme's %v", got, want)
	}
}

func TestComboBoxChosenPersistsAcrossFrames(t *testing.T) {
	var text string
	var seen []string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		box := NewComboBox(c, &text, []SelectOption{
			{Label: "Alpha", Value: "a"}, {Label: "Beta", Value: "b"},
		}).Label("选择")
		// Read Chosen every frame, as an app does. It must not go blank
		// on a frame with no input after a choice was made.
		seen = append(seen, box.Chosen())
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("选择"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("Beta"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if got := seen[len(seen)-1]; got != "b" {
		t.Errorf("the last frame saw Chosen %q, want b", got)
	}
	// Several more frames with no input keep the same value.
	count := len(seen)
	tt.Advance(time.Second)
	for _, v := range seen[count:] {
		if v != "b" {
			t.Errorf("Chosen went blank after the choice: %q", v)
		}
	}
}

func TestControlsHaveInvalid(t *testing.T) {
	var s string
	var i int
	var b bool
	var f float64
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewSelect(c, &s, []SelectOption{{Value: "a"}}).Invalid(true)
		NewCheckbox(c, &b, "x").Invalid(true)
		NewRadio(c, &i, 1, "y").Invalid(true)
		NewSwitch(c, &b).Invalid(true)
		NewSlider(c, &f, 0, 1).Invalid(true)
	}
	// The compile-time presence of all these methods is the point; the
	// view must also build without panic with them applied.
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	_ = tt
}
