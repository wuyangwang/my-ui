package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestGuideCardDismiss(t *testing.T) {
	closed := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewGuideCard(c, GuideCardOptions{
			Title: "Meet your agents", Desc: "They remember context.",
			Closed: &closed,
		}).Width(220)
	}
	tt := ui.NewTester(view, 300, 180)
	tt.Advance(time.Second)
	if !tt.HasText("Meet your agents") {
		t.Fatal("guide card missing")
	}
	if err := tt.Click("Close"); err != nil {
		t.Fatalf("close: %v", err)
	}
	tt.Advance(time.Second)
	if !closed {
		t.Fatal("close did not write the flag")
	}
	if tt.HasText("Meet your agents") {
		t.Fatal("closed guide card is still visible")
	}
}

func TestGuideCardVariants(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		ui.Row(c).Gap(c.Theme().Space(3)).Children(func() {
			NewGuideCard(c, GuideCardOptions{Title: "Filled", Variant: Default, Shadow: true, NoClose: true})
			NewGuideCard(c, GuideCardOptions{Title: "Outlined", Variant: Outline, NoClose: true})
			NewGuideCard(c, GuideCardOptions{Title: "Ghost", Variant: Ghost, NoClose: true})
		})
	}
	tt := ui.NewTester(view, 560, 180)
	tt.Advance(time.Second)
	for _, s := range []string{"Filled", "Outlined", "Ghost"} {
		if !tt.HasText(s) {
			t.Fatalf("variant %s missing", s)
		}
	}
}
