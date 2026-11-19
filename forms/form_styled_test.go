package forms

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestNewFormShowsItsFields(t *testing.T) {
	view := func(c *ui.Context) {
		NewForm(c, func() {
			ui.Field(c, "邮箱", func() { ui.TextInput(c, new(string)) })
			ui.Field(c, "城市", func() { ui.TextInput(c, new(string)) })
		})
	}
	tt := ui.NewTester(view, 400, 300)
	tt.Advance(time.Second)
	if !tt.HasText("邮箱") || !tt.HasText("城市") {
		t.Fatal("the styled form did not lay out its fields")
	}
}
