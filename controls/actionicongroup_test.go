package controls

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// TestActionIconGroupWritesTheClickedKey checks that a group writes the
// key of the icon pressed into the app's pointer and exposes the same
// key through Chosen in the frame of the click.
func TestActionIconGroupWritesTheClickedKey(t *testing.T) {
	chosen := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewActionIconGroup(c, &chosen, []ActionIconGroupItem{
			{Key: "copy", Icon: icons.Copy(), Label: "Copy"},
			{Key: "share", Icon: icons.Share(), Label: "Share"},
			{Key: "pin", Icon: icons.Pin(), Label: "Pin", Disabled: true},
		}).Filled()
	}
	tt := ui.NewTester(view, 480, 200)
	tt.Frame()
	for _, name := range []string{"Copy", "Share", "Pin"} {
		if _, ok := tt.Access(name); !ok {
			t.Errorf("%s is not on screen", name)
		}
	}
	if err := tt.Click("Share"); err != nil {
		t.Fatal(err)
	}
	if chosen != "share" {
		t.Errorf("clicked Share: pointer=%q", chosen)
	}
	// A disabled item keeps the last choice.
	if err := tt.Click("Pin"); err == nil {
		tt.Frame()
	}
	if chosen != "share" {
		t.Errorf("a disabled icon changed the choice to %q", chosen)
	}
}
