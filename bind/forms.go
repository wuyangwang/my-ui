package bind

import (
	"time"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// Calendar builds a month view editing *date.
func (u UI) Calendar(date *time.Time) *formspkg.Calendar {
	return formspkg.NewCalendar(u.c, date)
}

// TimeInput builds a time field editing *tm.
func (u UI) TimeInput(tm *time.Time) *formspkg.TimeInput {
	return formspkg.NewTimeInput(u.c, tm)
}

// EditableText builds a text that turns into a field when clicked.
func (u UI) EditableText(value *string) *formspkg.EditableText {
	return formspkg.NewEditableText(u.c, value)
}

// SelectClear builds a drop-down that can clear *selected, as o says.
func (u UI) SelectClear(selected *string, options []formspkg.SelectOption, clearable bool) *formspkg.Select {
	return formspkg.NewSelectClear(u.c, selected, options, clearable)
}

// Upload builds a file drop zone, configured by o.
func (u UI) Upload(o formspkg.UploadOptions) ui.Element {
	return formspkg.NewUpload(u.c, o)
}

// PasswordInput builds a field editing *value that hides its text.
func (u UI) PasswordInput(value *string) *formspkg.PasswordInput {
	return formspkg.NewPasswordInput(u.c, value)
}

// AutoComplete builds a field suggesting words from suggestions.
func (u UI) AutoComplete(value *string, suggestions []string) *formspkg.AutoComplete {
	return formspkg.NewAutoComplete(u.c, value, suggestions)
}

// Field builds a labelled field around the control made by build. It is
// the string-valued shim of forms.NewField[string]: a Go method cannot
// have type parameters, so this method only covers fields whose value is
// a *string. For a field of another value type call forms.NewField[T]
// directly. The value is given once — as value, or as o.Value — and the
// rules and timing travel in o; there is no second Validate call, and
// build may return the widget itself (the field finds its Control()).
func (u UI) Field(label string, value *string, build func() any, o formspkg.FieldOptions[string]) *formspkg.Field[string] {
	return formspkg.NewField[string](u.c, label, value, build, o)
}

// DecimalInput builds a number field editing *value with string bounds.
func (u UI) DecimalInput(value *string, lo, hi, step string) *formspkg.DecimalInput {
	return formspkg.NewDecimalInput(u.c, value, lo, hi, step)
}

// FieldGroup wraps fields in a group announced to assistive technology.
func (u UI) FieldGroup(build func()) ui.Element {
	return formspkg.NewFieldGroup(u.c, build)
}

// FieldSet builds a group with a legend around the controls build makes.
func (u UI) FieldSet(legend string, build func()) *formspkg.FieldSet {
	return formspkg.NewFieldSet(u.c, legend, build)
}

// InputOTP builds a one-time-code field of digits characters editing *value.
func (u UI) InputOTP(value *string, digits int) *formspkg.InputOTP {
	return formspkg.NewInputOTP(u.c, value, digits)
}

// InputNumber builds a number field editing *value between lo and hi.
func (u UI) InputNumber(value *float64, lo, hi, step float64) *formspkg.InputNumber {
	return formspkg.NewInputNumber(u.c, value, lo, hi, step)
}

// SearchBar builds a search field editing *query.
func (u UI) SearchBar(query *string) *formspkg.SearchBar {
	return formspkg.NewSearchBar(u.c, query)
}

// DatePicker builds a field that opens a calendar for *date.
func (u UI) DatePicker(date *time.Time) *formspkg.DatePicker {
	return formspkg.NewDatePicker(u.c, date)
}

// OTP builds a one-time-code field configured by o, editing *value.
func (u UI) OTP(value *string, o formspkg.OTPOptions) *formspkg.InputOTP {
	return formspkg.NewOTP(u.c, value, o)
}

// Form builds a themed form; fn builds its fields.
func (u UI) Form(fn func()) ui.Element {
	return formspkg.NewForm(u.c, fn)
}

// FormWith builds a themed form with a shared Variant preset; fn builds its
// fields.
func (u UI) FormWith(variant myui.Variant, fn func()) ui.Element {
	return formspkg.NewFormWith(u.c, variant, fn)
}

// Frame is not offered as a method: forms.NewFrame[T] is generic, and Go
// methods cannot have type parameters. The T is the widget whose fluent
// chain the frame returns, so a concrete instance would not be useful.
// Call forms.NewFrame[T] directly when building a custom field wrapper.
// The same goes for forms.NewImageSelect[T]; call it directly.
