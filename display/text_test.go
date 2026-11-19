package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestTextStyles(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(1)).Children(func() {
			NewText(c, "plain")
			NewText(c, "secondary", TextOptions{Type: TextSecondary})
			NewText(c, "success").Success()
			NewText(c, "danger", TextOptions{Type: TextDanger})
			NewText(c, "title").Heading(1).Strong()
			NewText(c, "code").Code()
		})
	}
	tt := ui.NewTester(view, 320, 280)
	tt.Advance(time.Second)
	for _, s := range []string{"plain", "secondary", "success", "danger", "title", "code"} {
		if !tt.HasText(s) {
			t.Fatalf("text %s missing", s)
		}
	}
}

func TestTextRowsAndDisabled(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewText(c, "a very long line of text that should be clipped", TextOptions{Rows: 1})
		NewText(c, "off", TextOptions{Disabled: true})
	}
	tt := ui.NewTester(view, 240, 120)
	tt.Advance(time.Second)
	if !tt.HasText("a very long line") || !tt.HasText("off") {
		t.Fatal("text missing")
	}
}
