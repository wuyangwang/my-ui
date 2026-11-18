package forms

import "github.com/egoist/mygo/ui"

// Only the MyGo inputs my-ui has no counterpart for are re-exported.
// Input, TextArea, PasswordInput, Select, ComboBox and MultiSelect have
// my-ui versions with the library's look and events, so their MyGo
// equivalents (including the Base and Buffer variants) are left out.
var (
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
