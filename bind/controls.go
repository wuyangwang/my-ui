package bind

import (
	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// RangeSlider builds a two-ended slider editing *low and *high.
func (u UI) RangeSlider(low, high *float64, lo, hi, step float64) *ctrlpkg.RangeSlider {
	return ctrlpkg.NewRangeSlider(u.c, low, high, lo, hi, step)
}

// StepSlider builds a stepped slider editing *value.
func (u UI) StepSlider(value *float64, lo, hi, step float64) *ctrlpkg.StepSlider {
	return ctrlpkg.NewStepSlider(u.c, value, lo, hi, step)
}

// Stepper builds a stepper editing *value with up and down arrows.
func (u UI) Stepper(value *float64, lo, hi, step float64) *ctrlpkg.Stepper {
	return ctrlpkg.NewStepper(u.c, value, lo, hi, step)
}

// Rating builds a star rating editing *value, with at most max stars.
func (u UI) Rating(value *int, max int) *ctrlpkg.Rating {
	return ctrlpkg.NewRating(u.c, value, max)
}

// Meter builds a meter for *value between lo and hi; levels may be nil.
func (u UI) Meter(value, lo, hi float64, levels *ui.MeterLevels) *ctrlpkg.Meter {
	return ctrlpkg.NewMeter(u.c, value, lo, hi, levels)
}

// ColorPicker builds a swatch that opens the native color picker for *color.
func (u UI) ColorPicker(color *ui.Color) *ctrlpkg.ColorPicker {
	return ctrlpkg.NewColorPicker(u.c, color)
}

// Toggle builds a button that stays pressed, editing *on with a label.
func (u UI) Toggle(on *bool, label string) *ctrlpkg.Toggle {
	return ctrlpkg.NewToggle(u.c, on, label)
}

// ColorWell builds a swatch for *color that the native color panel edits.
func (u UI) ColorWell(color *ui.Color) *ctrlpkg.ColorWell {
	return ctrlpkg.NewColorWell(u.c, color)
}

// Radio builds a radio holding value while *selected equals it, with a
// label. It is controls.NewRadio instantiated at string, the common case.
// Go methods cannot have type parameters, so for any other comparable
// value type call controls.NewRadio[T] with the context directly.
func (u UI) Radio(selected *string, value string, label string) *ctrlpkg.Radio[string] {
	return ctrlpkg.NewRadio[string](u.c, selected, value, label)
}

// RadioGroup wraps radio buttons in a labeled group.
func (u UI) RadioGroup(build func()) ui.Element {
	return ctrlpkg.NewRadioGroup(u.c, build)
}

// IconAction builds an icon with a button role, configured by o.
func (u UI) IconAction(o ctrlpkg.IconActionOptions) ui.Element {
	return ctrlpkg.NewIconAction(u.c, o)
}

// CheckboxState builds a three-state box editing *state, with a label.
func (u UI) CheckboxState(state *ui.CheckState, label string) *ctrlpkg.CheckboxState {
	return ctrlpkg.NewCheckboxState(u.c, state, label)
}

// CheckboxGroup wraps check boxes in a group announced with label.
func (u UI) CheckboxGroup(label string, build func()) ui.Element {
	return ctrlpkg.NewCheckboxGroup(u.c, label, build)
}

// ButtonGroup wraps buttons in a group announced to assistive technology.
func (u UI) ButtonGroup(build func()) ui.Element {
	return ctrlpkg.NewButtonGroup(u.c, build)
}

// TitleBar builds a frameless window's title bar, configured by o.
func (u UI) TitleBar(o ctrlpkg.TitleBarOptions) *ctrlpkg.TitleBar {
	return ctrlpkg.NewTitleBar(u.c, o)
}

// ToggleGroup wraps toggle buttons in a group; build makes the toggles.
func (u UI) ToggleGroup(fn func()) *ctrlpkg.ToggleGroup {
	return ctrlpkg.NewToggleGroup(u.c, fn)
}

// HotkeyInput builds a field editing *value that records a key shortcut.
func (u UI) HotkeyInput(value *string) *ctrlpkg.HotkeyInput {
	return ctrlpkg.NewHotkeyInput(u.c, value)
}

// CopyButton builds a button that copies content to the clipboard.
func (u UI) CopyButton(content string) *ctrlpkg.CopyButton {
	return ctrlpkg.NewCopyButton(u.c, content)
}

// Segmented builds a segmented control selecting one of labels into *selected.
func (u UI) Segmented(selected *int, labels ...string) *ctrlpkg.Segmented {
	return ctrlpkg.NewSegmented(u.c, selected, labels...)
}

// ColorSwatches builds a row of swatches choosing one color into *value.
func (u UI) ColorSwatches(value *ui.Color, swatches ...ctrlpkg.ColorSwatch) *ctrlpkg.ColorSwatches {
	return ctrlpkg.NewColorSwatches(u.c, value, swatches...)
}

// ActionIconGroup builds a row of icon-only actions writing the clicked
// item's key into *chosen.
func (u UI) ActionIconGroup(chosen *string, items []ctrlpkg.ActionIconGroupItem) *ctrlpkg.ActionIconGroup {
	return ctrlpkg.NewActionIconGroup(u.c, chosen, items)
}

// ActionIcon builds an icon-only button with title for assistive technology.
func (u UI) ActionIcon(svg *ui.SVG, title string) *ctrlpkg.ActionIcon {
	return ctrlpkg.NewActionIcon(u.c, svg, title)
}
