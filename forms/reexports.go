package forms

import "github.com/egoist/mygo/ui"

// Only MyGo inputs my-ui neither provides nor plans to provide are
// re-exported. InputNumber/SearchBar/AutoComplete/DatePicker are on the
// roadmap, so NumberInput, SearchField, Autocomplete and DateInput are
// deliberately left out; Input, TextArea, Select, ComboBox and
// MultiSelect have my-ui versions.
var (
	// NewTimeInput is ui.TimeInput.
	NewTimeInput = ui.TimeInput
	// NewEditableText is ui.EditableText.
	NewEditableText = ui.EditableText
	// NewForm is ui.Form.
	NewForm = ui.Form
	// NewFieldset is ui.Fieldset.
	NewFieldset = ui.Fieldset
	// NewFindBar is ui.FindBar.
	NewFindBar = ui.FindBar
)
