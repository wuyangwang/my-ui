package overlays

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

func TestHoverCardOpensOnHoverAndClosesAfterTheGrace(t *testing.T) {
	open := false
	view := func(c *ui.Context) {
		NewHoverCard(c, HoverCardOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "悬停我").Element
			},
			// The library wraps Body in panel.Children, so the content
			// lands in the overlay's panel without any help from us.
			Body: func() { ui.Text(c, "卡片内容") },
		})
	}
	tt := ui.NewTester(view, 320, 240)
	adv := func() { tt.Advance(time.Second); tt.Advance(time.Second) }
	tt.Advance(time.Second)
	if open {
		t.Fatal("the card opens before the pointer arrives")
	}
	anchor, ok := tt.Access("悬停我")
	if !ok {
		t.Fatal("the trigger is missing")
	}
	tt.Move(anchor.Bounds.X+anchor.Bounds.W/2, anchor.Bounds.Y+anchor.Bounds.H/2)
	adv()
	if !open {
		t.Fatal("hovering the trigger did not open the card")
	}
	if _, ok := tt.Find("卡片内容"); !ok {
		t.Error("the panel content is missing")
	}
	tt.Move(300, 220)
	for i := 0; i < 6; i++ {
		tt.Advance(time.Second)
	}
	if open {
		t.Error("the card stayed open after the pointer left")
	}
}

func TestHoverCardReportsOpenChanges(t *testing.T) {
	open := false
	var changes []bool
	view := func(c *ui.Context) {
		card := NewHoverCard(c, HoverCardOptions{
			Open: &open,
			Trigger: func() ui.Element {
				return ctrlpkg.NewButton(c, "悬停我").Element
			},
			Body: func() { ui.Text(c, "卡片内容") },
		})
		card.OnOpenChange(func(v bool) { changes = append(changes, v) })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	anchor, ok := tt.Access("悬停我")
	if !ok {
		t.Fatal("the trigger is missing")
	}
	tt.Move(anchor.Bounds.X+anchor.Bounds.W/2, anchor.Bounds.Y+anchor.Bounds.H/2)
	tt.Advance(time.Second)
	tt.Advance(time.Second)
	if len(changes) == 0 || !changes[len(changes)-1] {
		t.Errorf("OnOpenChange got %v, want a final true", changes)
	}
}
