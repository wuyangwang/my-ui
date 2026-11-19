package controls

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// TestSegmentedChoosesIntoTheIndex checks that clicking a segment
// writes its index and that Changed reports the choice in the frame of
// the click, as mygo's own segmented control does.
func TestSegmentedChoosesIntoTheIndex(t *testing.T) {
	selected := 0
	commits := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		current := selected
		seg := NewSegmented(c, &current, "List", "Grid", "Table").Label("View")
		if seg.Changed() {
			selected = current
			commits++
		}
	}
	tt := ui.NewTester(view, 480, 200)
	tt.Frame()
	var found bool
	for _, s := range tt.Texts() {
		if s == "Grid" {
			found = true
		}
	}
	if !found {
		t.Fatal("the segment labels are not on screen")
	}
	if err := tt.Click("Grid"); err != nil {
		t.Fatal(err)
	}
	if selected != 1 || commits != 1 {
		t.Errorf("clicking Grid gave selected=%d, commits=%d", selected, commits)
	}
}

// TestSegmentedVariants checks that the fluent presets keep the control
// usable and expose it to assistive technology as a radio group.
func TestSegmentedVariants(t *testing.T) {
	selected := 1
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewSegmented(c, &selected, "Day", "Week", "Month").Outlined().Sm().Block(true)
	}
	tt := ui.NewTester(view, 480, 200)
	tt.Frame()
	info, ok := tt.AccessByRole(ui.RoleRadioGroup, "")
	if !ok {
		t.Fatal("the segmented control is not a radio group")
	}
	if info.Bounds.W <= 0 {
		t.Error("the segmented control has no width")
	}
}
