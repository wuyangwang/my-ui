package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestBlockVariantsAndClick(t *testing.T) {
	clicks := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewBlock(c, BlockOptions{Label: "filled"}).Children(func() {
				ui.Text(c, "Filled")
			})
			NewBlock(c, BlockOptions{Label: "outlined"}).Outline().Children(func() {
				ui.Text(c, "Outlined")
			})
			NewBlock(c, BlockOptions{Label: "clickable", Clickable: true,
				OnChanged: func() { clicks++ }}).Children(func() {
				ui.Text(c, "Clickable")
			})
		})
	}
	tt := ui.NewTester(view, 420, 160)
	tt.Advance(time.Second)
	for _, text := range []string{"Filled", "Outlined", "Clickable"} {
		if !tt.HasText(text) {
			t.Fatalf("%s missing", text)
		}
	}
	if err := tt.Click("clickable"); err != nil {
		t.Fatalf("click: %v", err)
	}
	tt.Advance(time.Second)
	if clicks != 1 {
		t.Fatalf("clicks = %d, want 1", clicks)
	}
}

func TestBlockDisabled(t *testing.T) {
	clicks := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewBlock(c, BlockOptions{Label: "off", Clickable: true, Disabled: true,
			OnChanged: func() { clicks++ }}).Children(func() { ui.Text(c, "Off") })
	}
	tt := ui.NewTester(view, 200, 120)
	tt.Advance(time.Second)
	_ = tt.Click("off")
	tt.Advance(time.Second)
	if clicks != 0 {
		t.Fatal("disabled block must not fire OnChanged")
	}
}
