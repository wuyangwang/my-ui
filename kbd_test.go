package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestKbdShowsItsKeys(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewKbd(c, "Ctrl", "K")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	for _, key := range []string{"Ctrl", "K"} {
		if _, ok := tt.Find(key); !ok {
			t.Errorf("the key %q is not on screen", key)
		}
	}
}
