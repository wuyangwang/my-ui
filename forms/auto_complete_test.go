package forms

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestAutoCompletePresetsBuild(t *testing.T) {
	value := "苹"
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Width(300).Children(func() {
			NewAutoComplete(c, &value, []string{"苹果", "香蕉"}).
				Variant(Outline).Size(SizeSm).Width(180)
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	// The editor and its suggestions are mygo's; this checks that the
	// style presets wrap them without eating their input.
	if value != "苹" {
		t.Errorf("the value changed to %q without an edit", value)
	}
}
