package forms

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// NewForm creates a my-ui styled form. It keeps mygo's ui.Form signature,
// so a call written for the old re-export compiles: the fn builds the
// fields (NewField, ui.Field, ui.Fieldset) as it would in a plain form.
//
// Unlike ui.Form, the container paints a surface with a hairline, rounded
// corners and the theme's spacing. The labels are still laid out by mygo,
// aligned beside their controls, and every color follows c.Theme() on the
// frame the form is built.
//
//	forms.NewForm(c, func() {
//		forms.NewField(c, "邮箱", &app.email, func() *forms.Input {
//			return forms.NewInput(c.Key("email"), &app.email).Outline()
//		}, forms.FieldOptions[string]{})
//	})
func NewForm(c *ui.Context, fn func()) ui.Element {
	return NewFormWith(c, Native, fn)
}

// NewFormWith is NewForm with a shared Variant preset. Native (the zero
// value) leaves the form bare of a face; Outline gives it a bordered
// surface, Secondary a muted one, Ghost no surface at all.
func NewFormWith(c *ui.Context, variant Variant, fn func()) ui.Element {
	t := c.Theme()
	el := ui.Form(c, fn).Gap(t.Space(4)).Padding(t.Space(4)).
		Transition(StateTransition())
	switch variant {
	case Outline:
		el = el.Background(t.Surface).Border(1, t.Border).
			Radius(Radius(t, RadiusLg))
	case Secondary:
		el = el.Background(t.Surface).Radius(Radius(t, RadiusLg))
	}
	return el
}
