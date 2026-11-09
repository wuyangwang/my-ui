package myui

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// The text scale, in multiples of the theme's font size: pass one to
// Theme.Rem, as ui.Text(c, "…").FontSize(myui.TextSize(t, myui.TextSm)).
// The build's widgets take the values from the theme's font size, so a
// bigger text scale (the desktop's TextScale preference) scales them all.
const (
	TextXs float32 = 0.75
	TextSm float32 = 0.875
	TextMd float32 = 1
	TextLg float32 = 1.25
	TextXl float32 = 1.5
	Text2Xl float32 = 2
)

// TextSize returns n times the theme's font size, the height of text at
// scale n (TextXs through Text2Xl).
func TextSize(t *ui.Theme, n float32) float32 { return t.Rem(n) }

// The spacing scale, in units of the theme's spacing. Sizes, paddings
// and gaps take the theme's Space of one of these, so a looser or denser
// theme (Theme.Spacing) spaces them all in kind.
const (
	SpacingXs float32 = 1
	SpacingSm float32 = 2
	SpacingMd float32 = 3
	SpacingLg float32 = 4
	SpacingXl float32 = 6
	Spacing2Xl float32 = 8
)

// Space returns n units of the theme's spacing (SpacingXs through
// Spacing2Xl pass well), the room one scale step takes.
func Space(t *ui.Theme, n float32) float32 { return t.Space(n) }

// The radius scale, in multiples of the theme's radius (RadiusLg rounds
// the corners of one of the widgets at Theme.Radius — its base scale).
const (
	RadiusSm float32 = 0.6
	RadiusMd float32 = 1
	RadiusLg float32 = 1.5
	RadiusXl float32 = 2.2
)

// Radius r returns Theme.Radius at scale s (RadiusSm through RadiusXl),
// the rounding of a control's corners.
func Radius(t *ui.Theme, s float32) float32 { return t.Radius * s }

// Motion durations, for the UI's own animations (Element.AnimateTo and
// Animate pass them as Animation.Duration with a ui.Ease, or leave the
// duration at 0 for the ui.Snappy spring, which is what the presets and
// widgets here do). The desktop's Reduce Motion preference shortens
// nothing here: Animate itself goes to its target at once.
const (
	MotionFast   time.Duration = 120 * time.Millisecond
	MotionNormal time.Duration = 200 * time.Millisecond
	MotionSlow   time.Duration = 320 * time.Millisecond
)
