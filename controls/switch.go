package controls

import (
	. "github.com/wuyangwang/my-ui"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// switchKnobTime is how long the knob of a switch takes to travel.
const switchKnobTime = 140 * time.Millisecond

// Switch is a knob that is on or off.
//
//	controls.NewSwitch(c, &app.dark)
//
// The switch is styled by the library, as shadcn's: a pill with a knob
// in it, the accent when it is on. The knob travels when the switch
// changes, as the desktop's does. Size picks one of the shared size
// steps; Loading shows a spinner in the knob while a change is pending.
type Switch struct {
	ui.Element
	c       *ui.Context
	on      *bool
	invalid bool
	size    Size
	loading bool
}

// NewSwitch creates a switch editing *on.
func NewSwitch(c *ui.Context, on *bool) *Switch {
	s := &Switch{c: c, on: on}
	s.build()
	return s
}

// switchDims is the width, height and knob inset of the pill on a size.
func (s *Switch) dims() (w, h float32) {
	t := s.c.Theme()
	switch s.size {
	case SizeXs:
		return t.Space(7), t.Space(4)
	case SizeSm:
		return t.Space(8), t.Space(4.5)
	case SizeLg:
		return t.Space(11), t.Space(6)
	default:
		return t.Space(9), t.Space(5)
	}
}

// build makes the base and its painter. The presets only restyle this
// element: rebuilding the base in a fluent setter would read the frame's
// click twice.
func (s *Switch) build() {
	c := s.c
	on := s.on
	w, h := s.dims()
	sw := ui.SwitchBase(c, on).Size(w, h).Radius(h / 2)
	pos := sw.Animate("knob", switchPos(on != nil && *on), switchKnobTime)
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
		if s.loading {
			// A spinner in the knob, turning as the desktop animates.
			ang := float32(p.Now().UnixMilli()%1000) / 1000 * 360
			p.Transform(ui.Rotation(ang), knob.X+knob.W/2, knob.Y+knob.H/2, func() {
				p.Icon(icons.LoaderCircle(), knob, th.Accent)
			})
		} else {
			p.Fill(knob, ui.RGB(255, 255, 255), d/2)
		}
	})
	if s.loading {
		sw = sw.Disabled(true)
	}
	s.Element = sw
}

// switchPos is where the knob of a switch is: at the end that is on.
func switchPos(on bool) float32 {
	if on {
		return 1
	}
	return 0
}

// Size applies the shared size s to the switch and returns it, for
// chaining. It restyles the element built this frame rather than making a
// new SwitchBase, so the frame's click is still read once.
func (s *Switch) Size(size Size) *Switch {
	s.size = size
	w, h := s.dims()
	s.Element = s.Element.Size(w, h).Radius(h / 2)
	return s
}

// Xs sizes the switch down to the smallest step.
func (s *Switch) Xs() *Switch { return s.Size(SizeXs) }

// Sm sizes the switch down.
func (s *Switch) Sm() *Switch { return s.Size(SizeSm) }

// Md gives the built-in sizing.
func (s *Switch) Md() *Switch { return s.Size(SizeMd) }

// Lg sizes the switch up.
func (s *Switch) Lg() *Switch { return s.Size(SizeLg) }

// Loading shows a spinner in the knob and stops clicks, for the time a
// change is being saved. It returns the switch for chaining. The painter
// reads the flag from the switch, so no rebuild is needed.
func (s *Switch) Loading(on bool) *Switch {
	s.loading = on
	s.Element = s.Element.Disabled(on)
	return s
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
