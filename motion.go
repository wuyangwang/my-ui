package myui

import "github.com/egoist/mygo/ui"

// stateTransition is the transition the presets use: the colors of a
// state change fade, and everything else happens at once.
//
// The duration comes from the motion tokens, so the whole library
// moves at one pace. The desktop's Reduce Motion preference is the
// framework's: it shows the change at once.
func stateTransition() ui.ElementTransition {
	return ui.ElementTransition{
		Duration: MotionFast,
		Colors:   true,
		Position: false,
		Size:     false,
		Shadows:  false,
	}
}

// fadeIn makes an element that has just appeared fade in.
//
// A new element starts out invisible (its pose's fade is 0), and the
// spring brings it to the one it should have. An element that was
// there before keeps the pose it had, so the fade happens once.
func fadeIn(e ui.Element) ui.Element {
	return e.AnimatePose(ui.Pose{Fade: 1}, ui.Snappy)
}
