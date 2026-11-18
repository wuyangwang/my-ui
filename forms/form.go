package forms

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// Field is a form field: its label, the control in it, a description
// under it and an error message.
//
//	myui.Field(c, "閭", func() ui.Element {
//		return h.Input(&app.email).Outline().Control()
//	}).Description("鎴戜滑鍙敤瀹冨彂閫氱煡").Required(true)
//
// The label is a button without a face: clicking it focuses the
// control, as a label does in a form. The error is announced as it
// appears, and the control is marked as not holding a good value.
//
// The control the builder returns is the one the label focuses and the
// one the description and the error describe: for a control of this
// library that is its Control().
type Field struct {
	ui.Element
	c         *ui.Context
	control   ui.Element
	hasCtrl   bool
	label     ui.Element
	labelText ui.Element
	desc      ui.Element
	err       ui.Element
	touched   *bool
	initial   *string
	dirty     bool
	required  bool
	invalid   bool
	when      ValidateOn
	validate  func() *Field
}

// ValidateOn is when a field runs its rules: on losing focus, on every
// change, on gaining focus, or always. OnBlur is the default, as a form
// does.
type ValidateOn uint8

const (
	OnBlur ValidateOn = iota
	OnChange
	OnFocus
	Always
)

// NewField creates a field of a label and a control. The builder makes
// the control, and returns the element the field describes.
func NewField(c *ui.Context, label string, build func() ui.Element) *Field {
	f := &Field{c: c}
	col := ui.Column(c).Gap(c.Theme().Space(1.5))
	col.Children(func() {
		// The label: a button without a face, so that a click on it
		// focuses the control.
		lbl := ui.ButtonBase(c).TabStop(false)
		lbl.Children(func() {
			f.labelText = ui.Text(c, label)
		})
		f.label = lbl

		// The control.
		if build != nil {
			f.control = build()
			f.hasCtrl = true
			// The control takes the name of the label, as a form's
			// control does.
			if f.hasCtrl {
				f.control = f.control.Label(label)
			}
		}

		// The slots of the description and of the error.
		f.desc = ui.Column(c)
		f.err = ui.Column(c)

		if lbl.Clicked() && f.hasCtrl {
			f.control.Focus()
		}
	})
	f.Element = col
	return f
}

// Description puts the text under the control and returns the field,
// for chaining.
func (f *Field) Description(s string) *Field {
	c := f.c
	var text ui.Element
	f.desc.Children(func() {
		text = ui.Text(c, s).TextColor(c.Theme().TextMuted)
	})
	if f.hasCtrl {
		f.control = f.control.DescribedBy(text)
	}
	return f
}

// Error puts the message of the error under the control and returns the
// field, for chaining. The message is announced as it appears, and the
// control is marked as not holding a good value.
//
// A field with a validation shows the message of the rules; the two
// write the same place, so the one written last is the one shown.
func (f *Field) Error(s string) *Field {
	c := f.c
	var text ui.Element
	f.err.Children(func() {
		text = ui.Text(c, s).TextColor(c.Theme().Danger).Role(ui.RoleAlert)
	})
	if f.hasCtrl {
		f.control = f.control.ErrorMessage(text).Invalid(true)
		f.invalid = true
	}
	return f
}

// Validate checks the value against the rules and shows the message of
// the first rule that does not hold. The rules are checked once the
// field has been left, and after that whenever the value changes; a
// field that has not been left shows nothing, as a form does.
//
//	myui.NewField(c, "邮箱", func() ui.Element { ... }).
//		Validate(&app.email, myui.Required("请填写邮箱"), myui.Email("邮箱格式不对"))
//
// The field keeps the states a form keeps: whether it has been left
// (Touched) and whether the value differs from the one it started with
// (Dirty).
func (f *Field) Validate(value *string, rules ...Rule) *Field {
	if !f.hasCtrl {
		return f
	}
	// The states the field keeps, on its control, so that they survive
	// from frame to frame.
	touched := ui.Local(f.control, "field-touched", func() bool { return false })
	initial := ui.Local(f.control, "field-initial", func() string { return *value })
	wasFocused := ui.Local(f.control, "field-focused", func() bool { return false })
	if f.control.FocusLeft() {
		*touched = true
	}
	gained := f.control.Focused() && !*wasFocused
	*wasFocused = f.control.Focused()
	f.touched = touched
	f.initial = initial
	f.dirty = *value != *initial

	// Keep the rules so the chosen timing can re-run them.
	f.validate = func() *Field { return f.check(value, rules) }

	switch f.when {
	case OnChange:
		return f.check(value, rules)
	case OnFocus:
		if gained {
			return f.check(value, rules)
		}
		return f
	case Always:
		return f.check(value, rules)
	default: // OnBlur
		if !*touched {
			return f
		}
		return f.check(value, rules)
	}
}

// check runs the rules against value and shows the first message that
// does not hold. It is the shared body of every ValidateOn timing.
func (f *Field) check(value *string, rules []Rule) *Field {
	if !f.hasCtrl {
		return f
	}
	for _, r := range rules {
		if msg := r(*value); msg != "" {
			return f.Error(msg)
		}
	}
	// The value holds: the error goes away.
	f.err = f.err.Hidden(true)
	f.invalid = false
	return f
}

// Touched reports whether the field has been left, so that it has been
// in use.
func (f *Field) Touched() bool { return f.touched != nil && *f.touched }

// Dirty reports whether the value differs from the one the field
// started with.
func (f *Field) Dirty() bool { return f.dirty }

// ValidateOn checks value against rules at the chosen timing and returns
// the field, for chaining. OnBlur is the default; use OnChange, OnFocus
// or Always for the other common form behaviors.
func (f *Field) ValidateOn(when ValidateOn, value *string, rules ...Rule) *Field {
	f.when = when
	return f.Validate(value, rules...)
}

// OnFocused calls fn when the control gains the focus in this frame, and
// returns the field for chaining. MyGo has no focus-gained event; the
// field computes the edge from Focused() across frames.
func (f *Field) OnFocused(fn func()) *Field {
	if fn != nil {
		if was := ui.Local(f.control, "field-focused", func() bool { return false }); f.control.Focused() && !*was {
			fn()
		}
	}
	return f
}

// Errors shows several messages under the field as a list, as shadcn's
// FieldError does, and marks the control invalid. Each message is a live
// region so a screen reader reads them.
func (f *Field) Errors(msgs ...string) *Field {
	c := f.c
	if len(msgs) == 0 {
		return f
	}
	f.err.Children(func() {
		list := ui.Column(c).Gap(c.Theme().Space(0.5))
		list.Children(func() {
			for _, m := range msgs {
				ui.Text(c, m).TextColor(c.Theme().Danger).Role(ui.RoleAlert)
			}
		})
	})
	if f.hasCtrl {
		f.control = f.control.ErrorMessage(f.err).Invalid(true)
	}
	f.invalid = true
	return f
}

// Valid reports whether the field holds a good value: it has been left
// and the rules hold.
func (f *Field) Valid() bool { return f.Touched() && !f.invalid }

// Required marks the field as one that must be filled, and returns it,
// for chaining: the label takes a star.
func (f *Field) Required(on bool) *Field {
	f.required = on
	c := f.c
	if on {
		f.label.Children(func() {
			ui.Text(c, "*").TextColor(c.Theme().Danger)
		})
	}
	return f
}

// Invalid marks the control as not holding a good value and returns the
// field, for chaining. The label takes the danger color with it.
func (f *Field) Invalid(on bool) *Field {
	f.invalid = on
	if f.hasCtrl {
		f.control = f.control.Invalid(on)
	}
	t := f.c.Theme()
	if on {
		f.labelText = f.labelText.TextColor(t.Danger)
	} else {
		f.labelText = f.labelText.TextColor(t.Text)
	}
	return f
}

// Control returns the element the field describes: the control itself,
// or the text of the control, for a control of this library.
func (f *Field) Control() ui.Element { return f.control }
