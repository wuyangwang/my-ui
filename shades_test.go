package myui

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestShadesStepsTheLightTowardBlack(t *testing.T) {
	base := ui.Hex("#dc2626")
	black := ui.RGB(0, 0, 0)
	hover, pressed, ink := shades(base, false)
	if hover != base.MixOklab(black, 0.14) {
		t.Errorf("the light hover is %v, want 0.14 toward black", hover)
	}
	if pressed != base.MixOklab(black, 0.28) {
		t.Errorf("the light pressed is %v, want 0.28 toward black", pressed)
	}
	if ink != OnAccent(base) {
		t.Errorf("the light ink is %v, want what reads on the base", ink)
	}
}

func TestShadesLiftsTheDarkHoverTowardWhite(t *testing.T) {
	base := ui.Hex("#dc2626")
	black, white := ui.RGB(0, 0, 0), ui.RGB(255, 255, 255)
	hover, pressed, _ := shades(base, true)
	if hover != base.MixOklab(white, 0.2) {
		t.Errorf("the dark hover is %v, want 0.2 toward white", hover)
	}
	if pressed != base.MixOklab(black, 0.28) {
		t.Errorf("the dark pressed is %v, want 0.28 toward black", pressed)
	}
}

// TestApplyAsksShadesForTheSystemAccent checks the system-accent path
// runs through Shades: hover and pressed, and the ink, are its answer.
func TestApplyAsksShadesForTheSystemAccent(t *testing.T) {
	accent := ui.Hex("#dc2626")
	var got *ui.Theme
	view := func(c *ui.Context) {
		got = Apply(c, Options{Appearance: AppearanceDark})
	}
	ui.NewTesterWith(view, 200, 100, ui.TesterOptions{
		Preferences: ui.Preferences{Accent: accent},
	})
	hover, pressed, ink := shades(accent, true)
	if got.AccentHover != hover || got.AccentPressed != pressed || got.AccentText != ink {
		t.Errorf("the system accent took %v, %v, %v; want Shades' answer %v, %v, %v",
			got.AccentHover, got.AccentPressed, got.AccentText, hover, pressed, ink)
	}
}
