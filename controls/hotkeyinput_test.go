package controls

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// TestHotkeyInputRecordsAndRefuses checks that a modifier plus a key is
// recorded, a bare key is refused, and a conflicting shortcut does not
// replace the value.
func TestHotkeyInputRecordsAndRefuses(t *testing.T) {
	value := ""
	conflicts := ""
	focused := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		h := NewHotkeyInput(c, &value).Reset("mod+k").AllowClear(true).
			Conflicts("mod+s").OnConflict(func(keys string) { conflicts = keys })
		focused = h.Control().Focused()
	}
	tt := ui.NewTester(view, 480, 240)
	tt.Frame()
	if err := tt.Click("Press keys to record shortcut"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if !focused {
		t.Fatal("the recorder did not take the focus")
	}
	tt.Key(ui.Ctrl, ui.KeyK)
	tt.Frame()
	if value != "mod+k" {
		t.Errorf("the recorder holds %q, want mod+k", value)
	}
	// A shortcut taken elsewhere is refused and reported.
	tt.Key(ui.Ctrl, ui.KeyS)
	tt.Frame()
	if value != "mod+k" || conflicts != "mod+s" {
		t.Errorf("conflict: value=%q, reported=%q", value, conflicts)
	}
	// A key with no modifier is refused too.
	tt.Key(0, ui.KeyA)
	tt.Frame()
	if value != "mod+k" {
		t.Errorf("a bare key overwrote the value: %q", value)
	}
	if _, ok := tt.Access("A shortcut needs a modifier and a key"); !ok {
		t.Error("the invalid message is not on screen")
	}
}

// TestHotkeyInputResetAndClear checks the two icon affordances beside a
// recorded shortcut. Both reuse the shared icon action.
func TestHotkeyInputResetAndClear(t *testing.T) {
	value := "mod+k"
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewHotkeyInput(c, &value).Reset("mod+j").AllowClear(true)
	}
	tt := ui.NewTester(view, 480, 240)
	tt.Frame()
	for _, name := range []string{"Reset shortcut", "Clear shortcut"} {
		if _, ok := tt.Access(name); !ok {
			t.Errorf("%s is not on screen", name)
		}
	}
	if err := tt.Click("Clear shortcut"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if value != "" {
		t.Errorf("Clear left the value at %q", value)
	}
}

// TestHotkeyInputFluentConfig checks that options set after the
// constructor are read by the next frame's build.
func TestHotkeyInputFluentConfig(t *testing.T) {
	value := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewHotkeyInput(c, &value).Placeholder("按下组合键").Messages("需要修饰键", "已被占用")
	}
	tt := ui.NewTester(view, 480, 240)
	tt.Frame()
	tt.Frame()
	if _, ok := tt.Access("按下组合键"); !ok {
		t.Error("the localized placeholder is not on screen")
	}
}

// TestHotkeyName checks the string a pressed key is written as.
func TestHotkeyName(t *testing.T) {
	for _, tc := range []struct {
		mods ui.Modifiers
		key  ui.Key
		want string
	}{
		{ui.Ctrl, ui.KeyK, "mod+k"},
		{ui.Ctrl | ui.Shift, ui.KeyP, "mod+shift+p"},
		{ui.Alt | ui.Super, ui.Key1, "meta+alt+1"},
	} {
		got, ok := hotkeyName(tc.mods, tc.key)
		if !ok || got != tc.want {
			t.Errorf("hotkeyName(%v, %v) = %q, %v; want %q", tc.mods, tc.key, got, ok, tc.want)
		}
	}
	if _, ok := hotkeyName(0, ui.Key(9999)); ok {
		t.Error("an unnamed key was accepted")
	}
}
