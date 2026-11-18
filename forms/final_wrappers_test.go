package forms

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestDecimalInputWrapper(t *testing.T) {
	value := "12.50"
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewDecimalInput(c, &value, "0", "100", "0.01").Variant(Outline).Width(160)
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if value != "12.50" {
		t.Errorf("the decimal value changed to %q", value)
	}
}
