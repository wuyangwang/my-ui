package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestDividerForms(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(3)).Children(func() {
			NewDivider(c)
			NewDivider(c, DividerOptions{Dashed: true})
			NewDivider(c, DividerOptions{Label: "OR"})
			ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
				ui.Text(c, "left")
				NewDivider(c, DividerOptions{Vertical: true})
				ui.Text(c, "right")
			})
		})
	}
	tt := ui.NewTester(view, 320, 200)
	tt.Advance(time.Second)
	if !tt.HasText("OR") || !tt.HasText("left") || !tt.HasText("right") {
		t.Fatal("divider content missing")
	}
}

func TestDividerDarkTheme(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewDivider(c, DividerOptions{Label: "或"}).Width(200)
	}
	tt := ui.NewTester(view, 260, 100)
	tt.Advance(time.Second)
	if !tt.HasText("或") {
		t.Fatal("labelled divider missing")
	}
}
