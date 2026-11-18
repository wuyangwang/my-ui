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
		NewHoverCard(c, &open, func() ui.Element {
			return ctrlpkg.NewButton(c, "悬停我").Element
		}, func(panel ui.Element) {
			// The panel's parent is the overlay layer, so its content
			// goes into the panel's children.
			panel.Children(func() { ui.Text(c, "卡片内容") })
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
