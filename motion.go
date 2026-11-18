package myui

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// StateTransition is the transition the presets use: the colors of a
// state change fade, and everything else happens at once.
//
// The duration comes from the motion tokens, so the whole library moves
// at one pace. The desktop's Reduce Motion preference is the framework's:
// it shows the change at once.
func StateTransition() ui.ElementTransition {
	return ui.ElementTransition{
		Duration: MotionFast,
		Colors:   true,
		Position: false,
		Size:     false,
		Shadows:  false,
	}
}

// overlayFade is how long an overlay panel takes to come in.
const overlayFade = 160 * time.Millisecond

// OverlayEnter fades a panel in when it first appears. It uses the
// element's own animation, so the desktop's Reduce Motion takes the value
// straight to its target. The Local flag makes the first build start at 0
// and ask for another frame; the second and later builds animate to 1.
func OverlayEnter(c *ui.Context, panel ui.Element, key any) {
	entered := ui.Local(panel, key, func() bool { return false })
	if !*entered {
		panel.Jump("overlay-in", 0)
		*entered = true
		c.Invalidate()
	}
	op := panel.AnimateTo("overlay-in", 1, ui.Animation{Duration: overlayFade, NonMotion: true})
	panel.Opacity(op)
}
