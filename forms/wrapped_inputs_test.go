package forms

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestWrappedMygoInputsBuild(t *testing.T) {
	date := time.Now()
	clock := time.Now()
	text := "文件名"
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(3)).Children(func() {
			NewCalendar(c, &date)
			NewTimeInput(c, &clock).Variant(Outline).Width(160)
			NewEditableText(c, &text)
		})
	}
	tt := ui.NewTester(view, 320, 480)
	tt.Advance(time.Second)
	if text != "文件名" {
		t.Errorf("the editable text changed to %q", text)
	}
}
