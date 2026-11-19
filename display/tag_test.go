package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestTagClosesAndCallsOnClose(t *testing.T) {
	closed := false
	called := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTag(c, TagOptions{
			Label: "新功能", Closed: &closed,
			OnClose: func() { called++ },
		}).Variant(Success)
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
	if called == 0 {
		t.Error("the close icon did not call OnClose")
	}
	if _, ok := tt.Find("新功能"); ok {
		t.Error("a closed tag still shows")
	}
}

func TestTagVariantsColorAndRound(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewTag(c, TagOptions{Label: "filled"}).TagVariant(TagFilled).Color(ui.Hex("#7c3aed"))
			NewTag(c, TagOptions{Label: "outlined"}).TagVariant(TagOutlined).Color(ui.Hex("#7c3aed"))
			NewTag(c, TagOptions{Label: "borderless"}).TagVariant(TagBorderless).Color(ui.Hex("#7c3aed"))
			NewTag(c, TagOptions{Label: "solid", Round: true}).TagVariant(TagSolid).Color(ui.Hex("#7c3aed"))
		})
	}
	tt := ui.NewTester(view, 420, 120)
	tt.Advance(time.Second)
	for _, name := range []string{"filled", "outlined", "borderless", "solid"} {
		if _, ok := tt.Find(name); !ok {
			t.Errorf("tag %q missing", name)
		}
	}
}
