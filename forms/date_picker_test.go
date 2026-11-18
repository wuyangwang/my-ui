package forms

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestDatePickerDefaultsToToday(t *testing.T) {
	var date time.Time
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewDatePicker(c, &date)
	}
	tt := ui.NewTester(view, 320, 160)
	tt.Advance(time.Second)
	now := time.Now()
	if date.IsZero() {
		t.Fatal("a zero date was not defaulted to today")
	}
	if date.Year() != now.Year() || date.Month() != now.Month() || date.Day() != now.Day() {
		t.Errorf("the default date is %v, want today %v", date, now)
	}
}

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
