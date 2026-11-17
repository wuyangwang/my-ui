package myui

import (
	"github.com/egoist/mygo/ui"
)

// Slider is a knob on a track, that picks a number between two.
//
//	myui.NewSlider(c, &app.volume, 0, 100)
//
// The track and the knob are styled by the library, as shadcn's: a
// thin line with the accent up to the knob, and a knob on it. The
// arrows, Home and End move it while it has the focus, and the pointer
// drags it.
type Slider struct {
	ui.Element
	c       *ui.Context
	value   *float64
	lo      float64
	hi      float64
	step    float64
	invalid bool
}

// NewSlider creates a slider setting *value between lo and hi.
func NewSlider(c *ui.Context, value *float64, lo, hi float64) *Slider {
	s := &Slider{c: c, value: value, lo: lo, hi: hi}
	t := c.Theme()
	sl := ui.SliderBase(c, value, lo, hi).Height(t.Space(6)).PaddingX(t.Space(3))
	sl.Draw(func(p *ui.Painter, r ui.Rect) {
		s.draw(p, r)
	})
	s.Element = sl
	return s
}

// draw paints the track of the slider and its knob.
func (s *Slider) draw(p *ui.Painter, r ui.Rect) {
	t := s.c.Theme()
	// The knob rides on the track, and the track is padded by half the
	// knob, as the framework's slider says.
	knob := t.Space(4)
	track := t.Space(0.5)
	y := r.Y + (r.H-knob)/2
	onColor := t.Accent
	if s.invalid {
		onColor = t.Danger
	}
	at := float32(0)
	if s.hi > s.lo && s.value != nil {
		at = float32((*s.value - s.lo) / (s.hi - s.lo))
	}
	if at < 0 {
		at = 0
	}
	if at > 1 {
		at = 1
	}
	x := r.X + knob/2 + at*(r.W-knob)
	// The track, and the part of it up to the knob in the accent.
	p.Fill(ui.Rect{X: r.X + knob/2, Y: y + knob/2 - track/2, W: r.W - knob, H: track}, t.Border, track/2)
	p.Fill(ui.Rect{X: r.X + knob/2, Y: y + knob/2 - track/2, W: at * (r.W - knob), H: track}, onColor, track/2)
	// The knob itself.
	k := ui.Rect{X: x - knob/2, Y: y, W: knob, H: knob}
	p.Shadow(k, knob/2, 0, 1, 3, 0, ui.RGBA(0, 0, 0, 0.2))
	p.Fill(k, t.Background, knob/2)
	p.Fill(k, onColor, knob/2)
}

// Step sets the step of the slider: the arrows and the drag move it by
// that much, and it returns the slider, for chaining. A step of zero
// lets it take any value.
func (s *Slider) Step(step float64) *Slider {
	s.step = step
	s.Element = s.Element.Step(step)
	return s
}

// OnChanged runs fn when the value of the slider changes in this frame,
// and returns the slider, for chaining.
func (s *Slider) OnChanged(fn func(value float64)) *Slider {
	if fn != nil && s.Changed() {
		fn(*s.value)
	}
	return s
}

// Label names the slider for assistive technology and returns it, for
// chaining: a slider says its value, and a name says what it is for.
func (s *Slider) Label(label string) *Slider {
	s.Element = s.Element.Label(label)
	return s
}

// Disabled dims the slider and takes its clicks away, and returns the
// slider, for chaining.
func (s *Slider) Disabled(disabled bool) *Slider {
	s.Element = s.Element.Disabled(disabled)
	return s
}

// Invalid marks the slider as not holding a good value and returns it,
// for chaining. The filled part and the knob take the danger color.
func (s *Slider) Invalid(on bool) *Slider {
	s.invalid = on
	s.Element = s.Element.Invalid(on)
	return s
}

// Control returns the element a field describes: the slider itself.
func (s *Slider) Control() ui.Element { return s.Element }
