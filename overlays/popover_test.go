package overlays

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

func TestPopoverSidePlacesThePanel(t *testing.T) {
	for _, tc := range []struct {
		name string
		side PopoverSide
		want string // "above" or "below" the anchor
	}{
		{"top", PopoverTop, "above"},
		{"bottom", PopoverBottom, "below"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			open := false
			side := tc.side
			view := func(c *ui.Context) {
				ui.Column(c).Children(func() {
					// Room on both sides of the anchor, so the framework
					// need not flip the panel to the opposite side.
					ui.Box(c).Height(170)
					NewPopover(c, &open, func() ui.Element {
						return ctrlpkg.NewButton(c, "更多").Element
					}, func(panel ui.Element) {
						ui.Text(c, "面板内容")
					}).Side(side)
				})
			}
			tt := ui.NewTester(view, 400, 420)
			tt.Advance(time.Second)
			anchor, ok := tt.Find("更多")
			if !ok {
				t.Fatal("the anchor is not on screen")
			}
			if err := tt.Click("更多"); err != nil {
				t.Fatal(err)
			}
			tt.Advance(time.Second)
			content, ok := tt.Find("面板内容")
			if !ok {
				t.Fatal("the panel is not on screen after the anchor is pressed")
			}
			if tc.want == "above" && content.Y >= anchor.Y {
				t.Errorf("Side(Top): the panel is at %v, not above the anchor at %v", content, anchor)
			}
			if tc.want == "below" && content.Y <= anchor.Y+anchor.H {
				t.Errorf("Side(Bottom): the panel is at %v, not below the anchor at %v", content, anchor)
			}
		})
	}
}

func TestPopoverOptionsSidePlacesThePanel(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		ui.Column(c).Children(func() {
			ui.Box(c).Height(170)
			NewPopover(c, &open, func() ui.Element {
				return ctrlpkg.NewButton(c, "更多").Element
			}, func(panel ui.Element) {
				ui.Text(c, "面板内容")
			}, PopoverOptions{Side: PopoverTop})
		})
	}
	tt := ui.NewTester(view, 400, 420)
	tt.Advance(time.Second)
	anchor, ok := tt.Find("更多")
	if !ok {
		t.Fatal("the anchor is not on screen")
	}
	if err := tt.Click("更多"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	content, ok := tt.Find("面板内容")
	if !ok {
		t.Fatal("the panel is not on screen after the anchor is pressed")
	}
	if content.Y >= anchor.Y {
		t.Errorf("PopoverOptions{Side: PopoverTop}: the panel is at %v, not above the anchor at %v", content, anchor)
	}
}
