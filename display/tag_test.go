package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestTagCloses(t *testing.T) {
	closed := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTag(c, TagOptions{Label: "新功能", Closed: &closed}).Variant(Success)
	}
	tt := ui.NewTester(view, 320, 120)
	if _, ok := tt.Find("新功能"); !ok {
		t.Fatal("the tag label is missing")
	}
	close, ok := tt.Access("Close")
	if !ok {
		t.Fatal("the close icon is missing")
	}
	x, y := close.Bounds.X+close.Bounds.W/2, close.Bounds.Y+close.Bounds.H/2
	tt.Press(x, y)
	tt.Advance(time.Second)
	tt.Release(x, y)
	tt.Advance(time.Second)
	if !closed {
		t.Error("the close icon did not write the app's flag")
	}
}
