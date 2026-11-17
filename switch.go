package myui

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// switchKnobTime is how long the knob of a switch takes to travel.
const switchKnobTime = 140 * time.Millisecond

// Switch is a knob that is on or off.
//
//	myui.NewSwitch(c, &app.dark)
//
// The switch is styled by the library, as shadcn's: a pill with a knob
// in it, the accent when it is on. The knob travels when the switch
// changes, as the desktop's does.
type Switch struct {
	ui.Element
	c       *ui.Context
	on      *bool
	invalid bool
}

// NewSwitch creates a switch editing *on.
func NewSwitch(c *ui.Context, on *bool) *Switch {
	s := &Switch{c: c, on: on}
	t := c.Theme()
	sw := ui.SwitchBase(c, on).Size(t.Space(9), t.Space(5)).Radius(t.Space(2.5))
	// The knob travels from one end to the other, as the desktop's
	// does.
	pos := sw.Animate("knob", switchPos(*on), switchKnobTime)
	sw.Draw(func(p *ui.Painter, r ui.Rect) {
		th := c.Theme()
		onColor := th.Accent
		if s.invalid {
			onColor = th.Danger
		}
		off := th.Border.Mix(th.Text, 0.15)
		p.Fill(r, off.Mix(onColor, pos), r.H/2)
		in := r.H / 8
		d := r.H - 2*in
		knob := ui.Rect{X: r.X + in + pos*(r.W-r.H), Y: r.Y + in, W: d, H: d}
		p.Shadow(knob, d/2, 0, 1, 3, 0, ui.RGBA(0, 0, 0, 0.25))
		p.Fill(knob, ui.RGB(255, 255, 255), d/2)
	})
	s.Element = sw
	return s
}

// switchPos is where the knob of a switch is: at the end that is on.
func switchPos(on bool) float32 {
	if on {
		return 1
	}
	return 0
}

// OnChanged runs fn when the switch changes in this frame, and returns
// the switch, for chaining.
func (s *Switch) OnChanged(fn func(on bool)) *Switch {
	if fn != nil && s.Changed() {
		fn(*s.on)
	}
	return s
}

// Label names the switch for assistive technology and returns it, for
// chaining: a switch that shows only a knob needs a name.
func (s *Switch) Label(label string) *Switch {
	s.Element = s.Element.Label(label)
	return s
}

// Disabled dims the switch and takes its clicks away, and returns the
// switch, for chaining.
func (s *Switch) Disabled(disabled bool) *Switch {
	s.Element = s.Element.Disabled(disabled)
	return s
}

// Invalid marks the switch as not holding a good value and returns it,
// for chaining. The on-color becomes the danger color.
func (s *Switch) Invalid(on bool) *Switch {
	s.invalid = on
	s.Element = s.Element.Invalid(on)
	return s
}

// Control returns the element a field describes: the switch itself.
func (s *Switch) Control() ui.Element { return s.Element }
