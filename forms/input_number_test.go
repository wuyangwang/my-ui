package forms

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestInputNumberPresetsAndRange(t *testing.T) {
	v := 2.5
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Width(300).Children(func() {
			NewInputNumber(c, &v, 0, 10, 0.5).Variant(Outline).Size(SizeSm).Width(180)
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(1e9)
	if v != 2.5 {
		t.Errorf("the value changed to %v without an edit", v)
	}
}
