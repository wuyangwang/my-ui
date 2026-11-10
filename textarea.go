package myui

import "github.com/egoist/mygo/ui"

// TextArea is a multi-line text field with variant, size and events:
//
//	myui.NewTextArea(c, &draft).Outline().Placeholder("Notes")
//	myui.NewTextArea(c, &draft).Lines(3, 8).OnChanged(func(v string) { ... })
//
// The text area is a field, as an Input is: a frame with the text in
// it, and the prefix, the suffix and the clear button beside it. The
// frame carries the look of the preset, so the text itself has no
// border of its own.
//
// The chain applies the presets from the current theme every frame, so
// it follows palette and appearance changes. Call it every frame, ahead
// of anything else the app styles on the field, as the view runs.
type TextArea struct {
	field[*TextArea]
}

// NewTextArea creates a multi-line text field editing *value with the
// native preset.
func NewTextArea(c *ui.Context, value *string) *TextArea {
	ta := &TextArea{}
	ta.field = *newField[*TextArea](c, value, func(c *ui.Context, value *string, host ui.Element) ui.Element {
		return ui.TextArea(c, value)
	})
	ta.self = ta
	return ta
}

// Lines makes the text area as high as its text, wrapped at its width,
// from min lines up to max, past which it scrolls, as a message field
// grows with what is typed. It returns the text area, for chaining.
func (ta *TextArea) Lines(min, max int) *TextArea {
	ta.inner = ta.inner.Lines(min, max)
	return ta
}
