package forms

import "github.com/egoist/mygo/ui"

// Only MyGo inputs my-ui neither provides nor plans to provide are
// re-exported unchanged. Complex inputs that my-ui wraps (Form,
// NumberInput, SearchField, Autocomplete, DateInput, TokenField) are
// deliberately left out.
var (
	// NewFieldset is ui.Fieldset.
	NewFieldset = ui.Fieldset
	// NewFindBar is ui.FindBar.
	NewFindBar = ui.FindBar
)
