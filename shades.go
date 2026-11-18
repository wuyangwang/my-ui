package myui

import (
	"github.com/egoist/mygo/ui"
)

// Shades derives a color trio's full set of states from one base color
// and an appearance: the fill a control takes while hovered and while
// pressed, and the ink text shows in on any of them. The palette's
// authored triples and the system's accent color are the two sources
// that ask it.
//
// The rules live here so a preset never derives on its own: in the
// light appearance the states step toward black, 0.14 and 0.28 of the
// way in Oklab, and in the dark one the hover lifts toward white, 0.2
// of the way, so each state stays legible beside the base on both
// backgrounds. Shades is a pure ask: it holds no state across frames,
// and built views ask again next frame with the theme of the build.
//
// The theme carries the accent's trio in slots of its own; the Danger,
// the Warning and the Success come as one slot each, which is why the
// build asks Shades for the states of those.
func Shades(base ui.Color, dark bool) (hover, pressed, ink ui.Color) {
	black, white := ui.RGB(0, 0, 0), ui.RGB(255, 255, 255)
	if dark {
		hover, pressed = base.MixOklab(white, 0.2), base.MixOklab(black, 0.28)
	} else {
		hover, pressed = base.MixOklab(black, 0.14), base.MixOklab(black, 0.28)
	}
	return hover, pressed, OnAccent(base)
}
