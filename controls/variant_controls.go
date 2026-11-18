package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// The widgets in this file wrap MyGo's own controls with my-ui's look.
// MyGo keeps the value pointer, the keyboard and accessibility; the
// library adds the shared radius, colors and color transitions.

// RangeSlider picks a low and a high value on one track.
type RangeSlider struct {
	ui.Element
}

// NewRangeSlider builds a two-ended slider.
func NewRangeSlider(c *ui.Context, low, high *float64, lo, hi, step float64) *RangeSlider {
	t := c.Theme()
	return &RangeSlider{Element: ui.RangeSlider(c, low, high, lo, hi, step).
		Padding(t.Space(1)).Background(t.Surface).
		Radius(Radius(t, RadiusMd)).Transition(StateTransition())}
}

// StepSlider picks a value on a track with visible steps.
type StepSlider struct {
	ui.Element
}

// NewStepSlider builds a stepped slider.
func NewStepSlider(c *ui.Context, value *float64, lo, hi, step float64) *StepSlider {
	t := c.Theme()
	return &StepSlider{Element: ui.StepSlider(c, value, lo, hi, step).
		Padding(t.Space(1)).Background(t.Surface).
		Radius(Radius(t, RadiusMd)).Transition(StateTransition())}
}

// Stepper steps a value with up/down arrows.
type Stepper struct {
	ui.Element
	c       *ui.Context
	variant Variant
	size    Size
	width   float32
}

// NewStepper builds a stepper editing *value.
func NewStepper(c *ui.Context, value *float64, lo, hi, step float64) *Stepper {
	s := &Stepper{c: c}
	s.Element = ui.Stepper(c, value, lo, hi, step).Transition(StateTransition())
	s.apply()
	return s
}

// Variant applies the shared look of v.
func (s *Stepper) Variant(v Variant) *Stepper {
	s.variant = v
	s.apply()
	return s
}

// Size picks the text scale.
func (s *Stepper) Size(sz Size) *Stepper {
	s.size = sz
	s.apply()
	return s
}

// Width pins the stepper's width in DIPs; zero fills the parent.
func (s *Stepper) Width(w float32) *Stepper {
	s.width = w
	s.apply()
	return s
}

func (s *Stepper) apply() {
	t := s.c.Theme()
	s.Element = s.Element.Background(t.Background).Border(1, t.Border).
		Radius(Radius(t, RadiusMd)).Padding(t.Space(1), t.Space(2)).
		FontSize(stepperFont(t, s.size))
	if s.width > 0 {
		s.Element = s.Element.Width(s.width)
	} else {
		s.Element = s.Element.FillWidth()
	}
}

func stepperFont(t *ui.Theme, size Size) float32 {
	switch size {
	case SizeXs:
		return TextSize(t, TextXs)
	case SizeLg:
		return TextSize(t, TextLg)
	default:
		return TextSize(t, TextSm)
	}
}

// Rating picks a number of stars.
type Rating struct {
	ui.Element
}

// NewRating builds a star rating bound to *value.
func NewRating(c *ui.Context, value *int, max int) *Rating {
	return &Rating{Element: ui.Rating(c, value, max).Transition(StateTransition())}
}

// Meter shows a value in a range, colored by level.
type Meter struct {
	ui.Element
}

// NewMeter builds a meter; levels may be nil to use mygo's defaults.
func NewMeter(c *ui.Context, value, lo, hi float64, levels *ui.MeterLevels) *Meter {
	return &Meter{Element: ui.Meter(c, value, lo, hi, levels).Transition(StateTransition())}
}

// ColorPicker is a swatch that opens MyGo's color picker.
type ColorPicker struct {
	ui.Element
}

// NewColorPicker builds a swatch for *color, with my-ui's border and
// radius, on the theme's surface.
func NewColorPicker(c *ui.Context, color *ui.Color) *ColorPicker {
	t := c.Theme()
	return &ColorPicker{Element: ui.ColorPicker(c, color).
		Border(1, t.Border).Radius(Radius(t, RadiusMd)).
		Background(t.Surface).Padding(t.Space(1)).
		Transition(StateTransition())}
}
