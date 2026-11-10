package myui

import "github.com/egoist/mygo/ui"

// stateTransition is the transition the presets use: the colors of a
// state change fade, and everything else happens at once.
//
// The duration comes from the motion tokens, so the whole library
// moves at one pace. The desktop's Reduce Motion preference is the
// framework's: it shows the change at once.
//
// The pose animation (AnimatePose) is not used: its fade leaves the
// element it poses invisible, so the elements that come and go (the
// clear button of an input, the message under it) show at once.
func stateTransition() ui.ElementTransition {
	return ui.ElementTransition{
		Duration: MotionFast,
		Colors:   true,
		Position: false,
		Size:     false,
		Shadows:  false,
	}
}
