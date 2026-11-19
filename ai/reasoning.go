package ai

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// ReasoningOptions configures NewReasoning. Every value is a pointer the
// app owns; a nil pointer leaves that control out of the panel.
type ReasoningOptions struct {
	// Effort is the index of the chosen level in Levels.
	Effort *int
	// Levels names the thinking levels. The frame's locale supplies
	// low/medium/high when this is empty.
	Levels []string
	// Temperature and TopP are the two sampling parameters.
	Temperature *float64
	TopP        *float64
}

// Reasoning is a panel for a model's thinking settings: a row of levels and,
// when the app binds them, the temperature and top_p sliders with three
// semantic presets. The controls are my-ui's Segmented and Slider, so their
// keyboard, pointer and assistive-technology behavior is the framework's.
//
// The panel holds no state: it writes through the pointers it was given,
// and every color and metric comes from the frame's theme.
type Reasoning struct {
	ui.Element
	c             *ui.Context
	o             ReasoningOptions
	effort        *ctrlpkg.Segmented
	temp, topP    *ctrlpkg.Slider
	effortRow     ui.Element
	tempRow       ui.Element
	topPRow       ui.Element
	presetRow     ui.Element
	presetChanged bool
}

// NewReasoning builds a settings panel over the app's values.
func NewReasoning(c *ui.Context, o ReasoningOptions) *Reasoning {
	r := &Reasoning{c: c, o: o}
	t := c.Theme()
	levels := o.Levels
	if len(levels) == 0 {
		levels = []string{
			Message(c, "effortLow"), Message(c, "effortMedium"), Message(c, "effortHigh"),
		}
	}
	if o.Effort != nil && (*o.Effort < 0 || *o.Effort >= len(levels)) {
		*o.Effort = 0
	}
	root := ui.Column(c).Gap(t.Space(3)).Padding(t.Space(2)).Background(t.Surface).
		Border(1, t.Border).Radius(Radius(t, RadiusLg)).Transition(StateTransition())
	root.Children(func() {
		if o.Effort != nil {
			r.effortRow = reasoningRow(c, t, Message(c, "effort"), func() {
				r.effort = ctrlpkg.NewSegmented(c, o.Effort, levels...).Sm()
			})
		}
		if o.Temperature != nil {
			r.tempRow = reasoningRow(c, t, Message(c, "temperature"), func() {
				r.temp = ctrlpkg.NewSlider(c, o.Temperature, 0, 2).
					ShowValue(true).Label(Message(c, "temperature"))
			})
		}
		if o.TopP != nil {
			r.topPRow = reasoningRow(c, t, Message(c, "topP"), func() {
				r.topP = ctrlpkg.NewSlider(c, o.TopP, 0, 1).
					ShowValue(true).Label(Message(c, "topP"))
			})
		}
		if o.Effort != nil || o.Temperature != nil || o.TopP != nil {
			r.presetRow = ui.Row(c).Gap(t.Space(2)).Children(func() {
				preset(c, 0, 0.2, 0.8, Message(c, "presetPrecise"), &r.presetChanged, o)
				preset(c, 1, 0.7, 0.95, Message(c, "presetBalanced"), &r.presetChanged, o)
				preset(c, 2, 1.2, 1, Message(c, "presetCreative"), &r.presetChanged, o)
			})
		}
	})
	r.Element = root
	return r
}

// reasoningRow puts a muted caption above a control.
func reasoningRow(c *ui.Context, t *ui.Theme, label string, build func()) ui.Element {
	row := ui.Column(c).Gap(t.Space(1))
	row.Children(func() {
		ui.Text(c, label).FontSize(TextSize(t, TextSm)).TextColor(t.TextMuted)
		build()
	})
	return row
}

// preset is a button that writes one semantic set of values into the app's
// pointers. Clicked is read during the build, as mygo expects.
func preset(c *ui.Context, effort int, temp, top float64, label string, changed *bool, o ReasoningOptions) {
	b := ctrlpkg.NewButton(c, label).Ghost().Sm()
	if b.Clicked() {
		*changed = true
		if o.Effort != nil {
			levels := len(o.Levels)
			if levels == 0 {
				levels = 3
			}
			*o.Effort = min(effort, levels-1)
		}
		if o.Temperature != nil {
			*o.Temperature = temp
		}
		if o.TopP != nil {
			*o.TopP = top
		}
	}
}

// Compact hides the sliders and presets, leaving only the levels row, and
// returns the panel, for chaining. It goes back when called with false.
func (r *Reasoning) Compact(on bool) *Reasoning {
	if r.o.Temperature != nil {
		r.tempRow.Hidden(on)
	}
	if r.o.TopP != nil {
		r.topPRow.Hidden(on)
	}
	if r.o.Effort != nil || r.o.Temperature != nil || r.o.TopP != nil {
		r.presetRow.Hidden(on)
	}
	return r
}

// Label names the levels group for assistive technology and returns the
// panel, for chaining.
func (r *Reasoning) Label(name string) *Reasoning {
	if r.effort != nil {
		r.effort.Label(name)
	}
	return r
}

// Changed reports whether a level, a slider or a preset changed this frame.
func (r *Reasoning) Changed() bool {
	if r.presetChanged {
		return true
	}
	if r.effort != nil && r.effort.Changed() {
		return true
	}
	if r.temp != nil && r.temp.Changed() {
		return true
	}
	return r.topP != nil && r.topP.Changed()
}

// OnChanged calls fn when a control in the panel changed this frame, and
// returns the panel, for chaining.
func (r *Reasoning) OnChanged(fn func()) *Reasoning {
	if fn != nil && r.Changed() {
		fn()
	}
	return r
}
