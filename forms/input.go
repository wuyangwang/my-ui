package forms

import (
	"github.com/egoist/mygo/ui"
)

// Input is a text input with variant, size and events:
//
//	forms.NewInput(c, &value).Outline().Placeholder("Email")
//	forms.NewInput(c, &value).Clearable().OnChanged(func(v string) { ... })
//
// The input is a field: a frame with the text in it, and the prefix,
// the suffix and the clear button beside it. The frame carries the look
// of the preset, so the text itself has no border of its own.
//
// The chain applies the presets from the current theme every frame, so
// it follows palette and appearance changes. Call it every frame, ahead
// of anything else the app styles on the field, as the view runs.
type Input struct {
	Frame[*Input]
}

// NewInput creates a text input editing *value with the native preset.
func NewInput(c *ui.Context, value *string) *Input {
	in := &Input{}
	in.Frame = *NewFrame[*Input](c, value, func(c *ui.Context, value *string, host ui.Element) ui.Element {
		return ui.TextInput(c, value)
	})
	in.Self = in
	return in
}

// Password hides what the input holds and returns it, for chaining.
func (in *Input) Password() *Input {
	in.Inner = in.Inner.Password()
	return in
}

// OnSubmit calls fn when the user presses Enter in the input, and
// returns the input, for chaining.
func (in *Input) OnSubmit(fn func()) *Input {
	if fn != nil {
		in.Inner = in.Inner.OnTextKey(func(k ui.TextKey) bool {
			if k.Key == ui.KeyEnter {
				fn()
			}
			return false
		})
	}
	return in
}
