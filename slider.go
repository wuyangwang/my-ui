package myui

import (
	"strconv"

	"github.com/egoist/mygo/ui"
)

// Slider is a knob on a track, that picks a number between two.
//
//	myui.NewSlider(c, &app.volume, 0, 100)
//
// The track and the knob are styled by the library, as shadcn's: a
// thin line with the accent up to the knob, and a knob on it. The
// arrows, Home and End move it while it has the focus, and the pointer
// drags it. Vertical() turns it upright; ShowValue() prints the value
// beside it.
type Slider struct {
	ui.Element
	c         *ui.Context
	value     *float64
	lo        float64
	hi        float64
	invalid   bool
	vertical  bool
	showValue bool
	format    func(float64) string
}

// NewSlider creates a slider setting *value between lo and hi.
func NewSlider(c *ui.Context, value *float64, lo, hi float64) *Slider {
	s := &Slider{c: c, value: value, lo: lo, hi: hi}
	s.build()
	return s
}

// build makes the base and its painter. A preset builds again so the
// slider follows the theme and the chosen orientation.
func (s *Slider) build() {
	t := s.c.Theme()
	knob := t.Space(4)
	// A min width leaves room for the track to draw: without one the
	// base collapses to the padding and only the knob shows. This is
	// what MyGo's own slider does with MinWidth + PaddingX(kw/2).
	sl := ui.SliderBase(s.c, s.value, s.lo, s.hi).
		Height(t.Space(6)).MinWidth(t.Space(40)).PaddingX(knob / 2)
	if s.vertical {
		sl = sl.Vertical().Width(t.Space(6)).MinWidth(0).
			Height(t.Space(40)).PaddingX(0).PaddingY(knob / 2)
	}
	sl.Draw(func(p *ui.Painter, r ui.Rect) {
		s.draw(p, r)
	})
	s.Element = sl
}

// draw paints the track of the slider and its knob, along the main axis
// the orientation says.
func (s *Slider) draw(p *ui.Painter, r ui.Rect) {
	t := s.c.Theme()
	knob := t.Space(4)
	track := t.Space(0.5)
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
	if s.vertical {
		// Bottom to top: the filled part grows upward.
		x := r.X + (r.W-knob)/2
		cy := r.Y + knob/2 + (1-at)*(r.H-knob)
		trackRect := ui.Rect{X: x + knob/2 - track/2, Y: r.Y + knob/2, W: track, H: r.H - knob}
		p.Fill(trackRect, t.Border, track/2)
		p.Fill(ui.Rect{X: trackRect.X, Y: cy, W: track, H: at * (r.H - knob)}, onColor, track/2)
		k := ui.Rect{X: x, Y: cy - knob/2, W: knob, H: knob}
		p.Shadow(k, knob/2, 0, 1, 3, 0, ui.RGBA(0, 0, 0, 0.2))
		p.Fill(k, t.Background, knob/2)
		p.Fill(k, onColor, knob/2)
		return
	}
	y := r.Y + (r.H-knob)/2
	x := r.X + knob/2 + at*(r.W-knob)
	p.Fill(ui.Rect{X: r.X + knob/2, Y: y + knob/2 - track/2, W: r.W - knob, H: track}, t.Border, track/2)
	p.Fill(ui.Rect{X: r.X + knob/2, Y: y + knob/2 - track/2, W: at * (r.W - knob), H: track}, onColor, track/2)
	k := ui.Rect{X: x - knob/2, Y: y, W: knob, H: knob}
	p.Shadow(k, knob/2, 0, 1, 3, 0, ui.RGBA(0, 0, 0, 0.2))
	p.Fill(k, t.Background, knob/2)
	p.Fill(k, onColor, knob/2)
}

// Vertical turns the slider upright and returns it, for chaining.
func (s *Slider) Vertical() *Slider {
	s.vertical = true
	s.build()
	return s
}

// Step sets the step of the slider: the arrows and the drag move it by
// that much, and it returns the slider, for chaining. A step of zero
// lets it take any value.
func (s *Slider) Step(step float64) *Slider {
	s.Element = s.Element.Step(step)
	return s
}

// ShowValue prints the current value beside the slider and returns it,
// for chaining. Format controls how the number reads.
func (s *Slider) ShowValue(on bool) *Slider {
	s.showValue = on
	s.Element = s.Element.Value(s.valueText())
	return s
}

// Format sets how ShowValue renders the number, and returns the slider,
// for chaining. The default is a plain decimal.
func (s *Slider) Format(fn func(float64) string) *Slider {
	s.format = fn
	if s.showValue {
		s.Element = s.Element.Value(s.valueText())
	}
	return s
}

// valueText is the number as the slider reports it to assistive
// technology and shows it when ShowValue is on.
func (s *Slider) valueText() string {
	if s.value == nil {
		return ""
	}
	if s.format != nil {
		return s.format(*s.value)
	}
	return strconv.FormatFloat(*s.value, 'f', -1, 64)
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
	s.build()
	s.Element = s.Element.Invalid(on)
	return s
}

// Control returns the element a field describes: the slider itself.
func (s *Slider) Control() ui.Element { return s.Element }
