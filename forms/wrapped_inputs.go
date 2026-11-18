package forms

import (
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// Calendar is MyGo's month view on my-ui's card-like surface: a
// hairline border, the theme's xl radius and a soft inset. The chosen
// date stays the app's pointer, and mygo still marks today.
type Calendar struct {
	ui.Element
}

// NewCalendar builds a calendar editing *date.
func NewCalendar(c *ui.Context, date *time.Time) *Calendar {
	t := c.Theme()
	return &Calendar{Element: ui.Calendar(c, date).
		Background(t.Background).Border(1, t.Border).
		Radius(Radius(t, RadiusXl)).Padding(t.Space(2)).
		Transition(StateTransition())}
}

// TimeInput edits a time with the same field look as DatePicker.
type TimeInput struct {
	ui.Element
	c       *ui.Context
	variant Variant
	size    Size
	width   float32
}

// NewTimeInput builds a time field editing *tm.
func NewTimeInput(c *ui.Context, tm *time.Time) *TimeInput {
	ti := &TimeInput{c: c}
	ti.Element = ui.TimeInput(c, tm).Transition(StateTransition())
	ti.apply()
	return ti
}

// Variant applies the shared look of v.
func (t *TimeInput) Variant(v Variant) *TimeInput {
	t.variant = v
	t.apply()
	return t
}

// Size picks the text scale.
func (t *TimeInput) Size(sz Size) *TimeInput {
	t.size = sz
	t.apply()
	return t
}

// Width pins the field's width in DIPs; zero fills the parent.
func (t *TimeInput) Width(w float32) *TimeInput {
	t.width = w
	t.apply()
	return t
}

func (t *TimeInput) apply() {
	th := t.c.Theme()
	look := FrameLookOf(t.variant, th)
	if !look.Native {
		t.Element = t.Element.Background(look.Bg).Border(look.Width, look.Border)
		if look.HasText {
			t.Element = t.Element.TextColor(look.Text)
		}
	}
	v, h := FramePadding(th, t.size)
	t.Element = t.Element.Radius(Radius(th, RadiusMd)).
		Padding(v, h).FontSize(FrameFontSize(th, t.size))
	if t.width > 0 {
		t.Element = t.Element.Width(t.width)
	} else {
		t.Element = t.Element.FillWidth()
	}
}

// EditableText is text renamed in place, on a subtle field face.
type EditableText struct {
	ui.Element
}

// NewEditableText builds an in-place editor for *value.
func NewEditableText(c *ui.Context, value *string) *EditableText {
	t := c.Theme()
	return &EditableText{Element: ui.EditableText(c, value).
		Background(t.Background).Border(1, t.Border).
		Radius(Radius(t, RadiusMd)).Padding(t.Space(1), t.Space(2)).
		FontWeight(500).Transition(StateTransition())}
}
