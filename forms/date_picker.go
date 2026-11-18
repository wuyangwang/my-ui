package forms

import (
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// DatePicker is mygo's DateInput with my-ui's shared field look: the
// date pointer is the app's, and the calendar popup is mygo's.
type DatePicker struct {
	ui.Element
	c       *ui.Context
	variant Variant
	size    Size
	width   float32
}

// NewDatePicker builds a date field editing *date. A nil or zero date
// defaults to today, so the calendar opens on the current month with
// today marked, as a date picker should.
func NewDatePicker(c *ui.Context, date *time.Time) *DatePicker {
	if date == nil {
		date = new(time.Time)
	}
	if date.IsZero() {
		*date = time.Now()
	}
	d := &DatePicker{c: c}
	d.Element = ui.DateInput(c, date).Transition(StateTransition())
	d.apply()
	return d
}

// Variant applies the shared look of v.
func (d *DatePicker) Variant(v Variant) *DatePicker {
	d.variant = v
	d.apply()
	return d
}

// Size picks the text and padding scale.
func (d *DatePicker) Size(sz Size) *DatePicker {
	d.size = sz
	d.apply()
	return d
}

// Width pins the field's width in DIPs; zero fills the parent.
func (d *DatePicker) Width(w float32) *DatePicker {
	d.width = w
	d.apply()
	return d
}

func (d *DatePicker) apply() {
	t := d.c.Theme()
	look := FrameLookOf(d.variant, t)
	if !look.Native {
		d.Element = d.Element.Background(look.Bg).Border(look.Width, look.Border)
		if look.HasText {
			d.Element = d.Element.TextColor(look.Text)
		}
	}
	v, h := FramePadding(t, d.size)
	d.Element = d.Element.Radius(Radius(t, RadiusMd)).Padding(v, h).FontSize(FrameFontSize(t, d.size))
	if d.width > 0 {
		d.Element = d.Element.Width(d.width)
	} else {
		d.Element = d.Element.FillWidth()
	}
}
