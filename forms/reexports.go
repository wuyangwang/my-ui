package forms

import "github.com/egoist/mygo/ui"

// MyGo's input widgets, re-exported unchanged. Use my-ui's Input and
// friends for the themed presets; use these when the framework's own
// widget is what the app needs.
var (
	// NewTextInput is ui.TextInput.
	NewTextInput = ui.TextInput
	// NewTextInputBase is ui.TextInputBase.
	NewTextInputBase = ui.TextInputBase
	// NewTextInputBuffer is ui.TextInputBuffer.
	NewTextInputBuffer = ui.TextInputBuffer
	// NewTextInputBufferBase is ui.TextInputBufferBase.
	NewTextInputBufferBase = ui.TextInputBufferBase
	// NewTextAreaBase is ui.TextAreaBase.
	NewTextAreaBase = ui.TextAreaBase
	// NewTextAreaBuffer is ui.TextAreaBuffer.
	NewTextAreaBuffer = ui.TextAreaBuffer
	// NewTextAreaBufferBase is ui.TextAreaBufferBase.
	NewTextAreaBufferBase = ui.TextAreaBufferBase
	// NewComboboxBase is ui.ComboboxBase.
	NewComboboxBase = ui.ComboboxBase
	// NewTokenField is ui.TokenField, the chips-and-suggestions widget.
	NewTokenField = ui.TokenField
	// NewNumberInput is ui.NumberInput.
	NewNumberInput = ui.NumberInput
	// NewDecimalInput is ui.DecimalInput.
	NewDecimalInput = ui.DecimalInput
	// NewDateInput is ui.DateInput.
	NewDateInput = ui.DateInput
	// NewTimeInput is ui.TimeInput.
	NewTimeInput = ui.TimeInput
	// NewSearchField is ui.SearchField.
	NewSearchField = ui.SearchField
	// NewAutocomplete is ui.Autocomplete.
	NewAutocomplete = ui.Autocomplete
	// NewEditableText is ui.EditableText.
	NewEditableText = ui.EditableText
	// NewForm is ui.Form.
	NewForm = ui.Form
	// NewFieldset is ui.Fieldset.
	NewFieldset = ui.Fieldset
	// NewFindBar is ui.FindBar.
	NewFindBar = ui.FindBar
)
