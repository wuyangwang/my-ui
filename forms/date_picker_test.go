package forms

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestDatePickerPresetsBuild(t *testing.T) {
	date := time.Date(2026, 11, 18, 0, 0, 0, 0, time.UTC)
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Width(300).Children(func() {
			NewDatePicker(c, &date).Variant(Outline).Size(SizeSm).Width(180)
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if !date.Equal(time.Date(2026, 11, 18, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("the date changed to %v without an edit", date)
	}
}
