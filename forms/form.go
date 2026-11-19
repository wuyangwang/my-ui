package forms

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// Field is a form field: a label, the control it names, a description
// under it and an error message. The type parameter is the type of the
// value the control edits; rules and the touched/dirty state are kept in
// the library, but the value itself stays in the app's pointer.
//
//	myui.NewField(c, "邮箱", &app.email, func() ui.Element {
//		return forms.NewInput(c.Key("email"), &app.email).Outline()
//	}, myui.FieldOptions[string]{
//		Description: "我们只用它发通知",
//		Rules:       []myui.Rule{myui.Required("请填写邮箱"), myui.Email("邮箱格式不对")},
//		Required:    true,
//	})
//
// The builder returns the widget itself — the field finds the element it
// describes by asking for its Control(), so a caller never has to
// remember the trailing .Control(). A raw ui.Element is used as-is.
//
// Clicking the label focuses the control. The error is announced as it
// appears and the control is marked invalid. Whether an error is shown
// is decided by When; the value is checked every frame, so Valid() is
// meaningful before the field has ever been left.
type Field[T any] struct {
	ui.Element
	c           *ui.Context
	control     ui.Element
	hasCtrl     bool
	value       *T
	rules       []Rule
	when        ValidateOn
	label       ui.Element
	labelText   ui.Element
	desc        ui.Element
	err         ui.Element
	touched     *bool
	wasFocused  *bool
	initial     *string
	descRef     *string
	manualErr   *string
	errorsRef   *[]string
	dirty       bool
	invalid     bool
	gainedFocus bool
	required    bool
}

// ValidateOn is when a field shows the message of a rule that fails:
// after the field is left (the default), on a change, on gaining focus,
// or always. The value itself is checked on every frame at every timing,
// so Valid() never depends on Touched().
type ValidateOn uint8

const (
	OnBlur ValidateOn = iota
	OnChange
	OnFocus
	Always
)

// FieldOptions configures a Field at construction. Every field is set up
// in one call: the value pointer, the rules, the timing, the description
// and the required mark travel together, so validation cannot be forgotten
// and the pointer is never passed twice.
type FieldOptions[T any] struct {
	// Value is the app-owned pointer to the value the control edits. It
	// overrides the value argument of NewField when both are set.
	Value *T
	// Rules are checked in order; the first failing rule's message shows.
	Rules []Rule
	// When decides when a failing message is shown. Zero is OnBlur.
	When ValidateOn
	// Description is the muted help text under the control.
	Description string
	// Required puts a danger-colored star after the label.
	Required bool
}

// NewField creates a labelled field around the control built by build.
// The value is read from o.Value, or from value when o.Value is nil; it
// is never read a second time through a separate Validate call.
//
// build may return a widget of this library directly — the field asks it
// for Control(), so the trailing .Control() is not needed — or a raw
// ui.Element, which is used as-is. The widget type is inferred from build;
// callers do not write it.
func NewField[T any, B any](c *ui.Context, label string, value *T, build func() B, o FieldOptions[T]) *Field[T] {
	f := &Field[T]{c: c, rules: o.Rules, when: o.When, required: o.Required}
	f.value = o.Value
	if f.value == nil {
		f.value = value
	}

	col := ui.Column(c).Gap(c.Theme().Space(1.5))
	col.Children(func() {
		// A face-less button: clicking the label focuses the control.
		lbl := ui.ButtonBase(c).TabStop(false)
		lbl.Children(func() {
			f.labelText = ui.Text(c, label)
		})
		f.label = lbl

		if build != nil {
			var ctrl ui.Element
			switch built := any(build()).(type) {
			case ui.Element:
				// A raw element is already the control.
				ctrl = built
			case interface{ Control() ui.Element }:
				// A widget of this library: use the element it
				// describes, so callers never need .Control().
				ctrl = built.Control()
			}
			f.control = ctrl.Label(label)
			f.hasCtrl = true
		}

		f.desc = ui.Column(c)
		f.err = ui.Column(c)

		if lbl.Clicked() && f.hasCtrl {
			f.control.Focus()
		}
	})
	f.Element = col

	// The state lives on the control, so it survives from frame to frame
	// even though the field struct is rebuilt each build pass.
	anchor := col
	if f.hasCtrl {
		anchor = f.control
	}
	f.touched = ui.Local(anchor, "field-touched", func() bool { return false })
	f.wasFocused = ui.Local(anchor, "field-focused", func() bool { return false })
	f.initial = ui.Local(anchor, "field-initial", f.text)
	f.descRef = ui.Local(anchor, "field-desc", func() string { return o.Description })
	f.manualErr = ui.Local(anchor, "field-error", func() string { return "" })
	f.errorsRef = ui.Local(anchor, "field-errors", func() []string { return nil })

	if f.required {
		c := f.c
		f.label.Children(func() {
			ui.Text(c, "*").TextColor(c.Theme().Danger)
		})
	}

	if f.hasCtrl {
		if f.control.FocusLeft() {
			*f.touched = true
		}
		focused := f.control.Focused()
		f.gainedFocus = focused && !*f.wasFocused
		*f.wasFocused = focused
	}
	f.dirty = f.text() != *f.initial

	// The description can be set once and still show on every frame.
	f.desc.Children(func() {
		if *f.descRef == "" {
			return
		}
		text := ui.Text(f.c, *f.descRef).TextColor(f.c.Theme().TextMuted)
		if f.hasCtrl {
			f.control = f.control.DescribedBy(text)
		}
	})
	f.check()
	return f
}

// text is the value as the rules see it. Rules take a string; a field of
// another value type is rendered with fmt.Sprint before checking.
func (f *Field[T]) text() string {
	if f.value == nil {
		return ""
	}
	return fmt.Sprint(*f.value)
}

// check runs the rules against the current value and renders the first
// failing message. The value is checked on every frame; only the moment a
// message is shown depends on When. Keeping the two apart is what lets
// Valid() answer before the first blur.
func (f *Field[T]) check() {
	msg := *f.manualErr
	if msg == "" {
		for _, r := range f.rules {
			if m := r(f.text()); m != "" {
				msg = m
				break
			}
		}
	}
	f.invalid = msg != ""

	show := false
	switch f.when {
	case OnChange:
		// dirty compares the current value with the initial one, so it
		// does not depend on the control's one-frame Changed edge.
		show = f.dirty || *f.touched
	case OnFocus:
		show = f.gainedFocus || *f.touched
	case Always:
		show = true
	default: // OnBlur
		show = *f.touched
	}
	if !show {
		msg = ""
	}

	c := f.c
	f.err.Children(func() {
		if len(*f.errorsRef) > 0 && show {
			list := ui.Column(c).Gap(c.Theme().Space(0.5))
			list.Children(func() {
				for _, m := range *f.errorsRef {
					ui.Text(c, m).TextColor(c.Theme().Danger).Role(ui.RoleAlert)
				}
			})
		} else if msg != "" {
			text := ui.Text(c, msg).TextColor(c.Theme().Danger).Role(ui.RoleAlert)
			if f.hasCtrl {
				f.control = f.control.ErrorMessage(text).Invalid(true)
			}
		}
	})
	if msg == "" && len(*f.errorsRef) == 0 {
		f.err = f.err.Hidden(true)
	} else {
		f.err = f.err.Hidden(false)
	}
	if f.hasCtrl && !f.invalid {
		f.control = f.control.Invalid(false)
	}
}

// Description sets the muted help text under the control and returns the
// field, for chaining. It is kept on the control, so setting it once is
// enough; the field relates it to the control with DescribedBy.
func (f *Field[T]) Description(s string) *Field[T] {
	if f.descRef != nil {
		*f.descRef = s
	}
	if s != "" {
		f.desc.Children(func() {
			text := ui.Text(f.c, s).TextColor(f.c.Theme().TextMuted)
			if f.hasCtrl {
				f.control = f.control.DescribedBy(text)
			}
		})
	}
	return f
}

// ErrorText puts a message under the control and returns the field, for
// chaining. The message is announced as it appears, and the control is
// marked invalid. A field with rules shows the failing rule's message
// instead; both write the same slot, so the one written last wins.
func (f *Field[T]) ErrorText(s string) *Field[T] {
	if f.manualErr != nil {
		*f.manualErr = s
	}
	if s != "" {
		f.err.Children(func() {
			text := ui.Text(f.c, s).TextColor(f.c.Theme().Danger).Role(ui.RoleAlert)
			if f.hasCtrl {
				f.control = f.control.ErrorMessage(text).Invalid(true)
			}
		})
		f.err = f.err.Hidden(false)
		f.invalid = true
	}
	return f
}

// Errors shows several messages under the field as a list, as shadcn's
// FieldError does, and marks the control invalid. Each message is a live
// region so a screen reader reads it. The list is kept on the control, so
// a single call lasts across frames.
func (f *Field[T]) Errors(msgs ...string) *Field[T] {
	if f.errorsRef != nil {
		*f.errorsRef = msgs
	}
	if len(msgs) > 0 {
		c := f.c
		f.err.Children(func() {
			list := ui.Column(c).Gap(c.Theme().Space(0.5))
			list.Children(func() {
				for _, m := range msgs {
					ui.Text(c, m).TextColor(c.Theme().Danger).Role(ui.RoleAlert)
				}
			})
		})
		f.err = f.err.Hidden(false)
		f.invalid = true
	}
	return f
}

// Required marks the field as one that must be filled, and returns it,
// for chaining: the label takes a star.
func (f *Field[T]) Required(on bool) *Field[T] {
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
func (f *Field[T]) Invalid(on bool) *Field[T] {
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

// Touched reports whether the field has been left, so that it has been
// in use.
func (f *Field[T]) Touched() bool { return f.touched != nil && *f.touched }

// Dirty reports whether the value differs from the one the field
// started with.
func (f *Field[T]) Dirty() bool { return f.dirty }

// Valid reports whether the rules hold for the current value. Unlike the
// old behavior it does not wait for Touched: a field already holding a
// good value is valid before it has been left, and an empty required
// field is invalid from the first frame.
func (f *Field[T]) Valid() bool { return !f.invalid }

// OnFocused calls fn when the control gains focus in this frame, and
// returns the field for chaining. mygo has no focus-gained event, so the
// field computes the edge from Focused() across frames in its own slot.
func (f *Field[T]) OnFocused(fn func()) *Field[T] {
	if !f.hasCtrl {
		return f
	}
	was := ui.Local(f.control, "field-onfocused", func() bool { return false })
	now := f.control.Focused()
	if fn != nil && now && !*was {
		fn()
	}
	*was = now
	return f
}

// Control returns the element the field describes: the widget's own
// control when build returned a widget of this library, or the raw
// element otherwise.
func (f *Field[T]) Control() ui.Element { return f.control }
