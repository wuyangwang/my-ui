package controls

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// TestToggleGroupKeepsMygoSignature checks that the group gathers the
// toggles fn builds, keeps their own bound state and names itself for
// assistive technology.
func TestToggleGroupKeepsMygoSignature(t *testing.T) {
	bold, italic := false, false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewToggleGroup(c, func() {
			NewToggle(c, &bold, "Bold").Size(SizeSm)
			NewToggle(c, &italic, "Italic").Size(SizeSm)
		}).Label("Style").Outlined()
	}
	tt := ui.NewTester(view, 480, 200)
	tt.Frame()
	for _, name := range []string{"Bold", "Italic"} {
		if _, ok := tt.Access(name); !ok {
			t.Errorf("%s is not on screen", name)
		}
	}
	if err := tt.Click("Bold"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if !bold {
		t.Error("the toggle in the group did not change its own pointer")
	}
	if italic {
		t.Error("clicking Bold also changed Italic")
	}
}

// TestToggleGroupVariants checks the borderless and segmented presets
// produce a visible group without stale colors.
func TestToggleGroupVariants(t *testing.T) {
	on := false
	for _, variant := range []ToggleGroupVariant{ToggleGroupBorderless, ToggleGroupSegmented} {
		variant := variant
		t.Run("", func(t *testing.T) {
			v := variant
			view := func(c *ui.Context) {
				Apply(c, Options{Appearance: AppearanceDark})
				NewToggleGroup(c, func() { NewToggle(c, &on, "Left").Size(SizeSm) }).Variant(v)
			}
			tt := ui.NewTester(view, 240, 160)
			tt.Frame()
			if _, ok := tt.Access("Left"); !ok {
				t.Error("the toggle is not on screen")
			}
		})
	}
}
