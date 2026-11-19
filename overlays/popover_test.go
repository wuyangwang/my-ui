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
					NewPopover(c, PopoverOptions{
						Open: &open,
						Trigger: func() ui.Element {
							return ctrlpkg.NewButton(c, "更多").Element
						},
						Body: func() { ui.Text(c, "面板内容") },
						Side: &side,
					})
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

func TestPopoverOpensWithOneClickAndBodyLandsInThePanel(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		NewPopover(c, PopoverOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "更多").Element
			},
			Body: func() { ui.Text(c, "面板内容") },
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板内容"); ok {
		t.Fatal("the panel shows before the anchor is pressed")
	}
	if err := tt.Click("更多"); err != nil {
		t.Fatal(err)
	}
	if !open {
		t.Fatal("one click on the anchor did not open the panel")
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板内容"); !ok {
		t.Fatal("the body did not land in the panel")
	}
}

func TestPopoverLateOnOpenChangeFires(t *testing.T) {
	open := false
	var changes []bool
	// The callback is attached after the constructor, so it cannot be a
	// one-frame callback tied to the constructor's frame.
	view := func(c *ui.Context) {
		p := NewPopover(c, PopoverOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "更多").Element
			},
			Body: func() { ui.Text(c, "面板内容") },
		})
		p.OnOpenChange(func(v bool) { changes = append(changes, v) })
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("更多"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if len(changes) == 0 || !changes[len(changes)-1] {
		t.Errorf("late OnOpenChange got %v, want a final true", changes)
	}
}
