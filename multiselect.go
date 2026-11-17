package myui

import "github.com/egoist/mygo/ui"

// MultiSelect is a field of tags: the chosen options show as chips,
// each with a button that takes it out, and the options below narrow
// down with what is typed.
//
//	myui.NewMultiSelect(c, &app.langs, []myui.SelectOption{
//		{Label: "简体中文", Value: "zh-Hans"},
//		{Label: "English", Value: "en"},
//	}).Outline().Label("语言")
//
// The chips show the label of an option and the app holds that label:
// MyGo's token field keeps the text of its chips, so the value of a
// multi-select is its label, not a value of its own.
type MultiSelect struct {
	ui.Element // the field of chips
	c          *ui.Context
	tokens     *[]string
	options    []SelectOption
	size       Size
	variant    Variant
	shape      Shape
	invalid    bool
}

// NewMultiSelect creates a multi-select field. The tokens are the
// choices the app holds.
func NewMultiSelect(c *ui.Context, tokens *[]string, options []SelectOption) *MultiSelect {
	m := &MultiSelect{c: c, tokens: tokens, options: options}
	labels := make([]string, len(options))
	for i, o := range options {
		labels[i] = o.label()
	}
	m.Element = ui.TokenField(c, tokens, labels)
	return m
}

// Variant applies the look of v and returns the field, for chaining.
// The zero variant is Native, which keeps the built-in look.
func (m *MultiSelect) Variant(v Variant) *MultiSelect {
	m.variant = v
	m.apply()
	return m
}

// Primary is Variant(myui.Primary).
func (m *MultiSelect) Primary() *MultiSelect { return m.Variant(Primary) }

// Secondary is Variant(myui.Secondary), a muted fill.
func (m *MultiSelect) Secondary() *MultiSelect { return m.Variant(Secondary) }

// Default is Variant(myui.Default), the built-in look.
func (m *MultiSelect) Default() *MultiSelect { return m.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (m *MultiSelect) Native() *MultiSelect { return m.Variant(Native) }

// Outline is Variant(myui.Outline).
func (m *MultiSelect) Outline() *MultiSelect { return m.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (m *MultiSelect) Ghost() *MultiSelect { return m.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (m *MultiSelect) Destructive() *MultiSelect { return m.Variant(Destructive) }

// Link is Variant(myui.Link).
func (m *MultiSelect) Link() *MultiSelect { return m.Variant(Link) }

// Success is Variant(myui.Success).
func (m *MultiSelect) Success() *MultiSelect { return m.Variant(Success) }

// Info is Variant(myui.Info).
func (m *MultiSelect) Info() *MultiSelect { return m.Variant(Info) }

// Warning is Variant(myui.Warning).
func (m *MultiSelect) Warning() *MultiSelect { return m.Variant(Warning) }

// Size applies the sizing of s and returns the field, for chaining.
func (m *MultiSelect) Size(s Size) *MultiSelect {
	m.size = s
	m.apply()
	return m
}

// Sm sizes the field down.
func (m *MultiSelect) Sm() *MultiSelect { return m.Size(SizeSm) }

// Md gives the built-in sizing.
func (m *MultiSelect) Md() *MultiSelect { return m.Size(SizeMd) }

// Lg sizes the field up.
func (m *MultiSelect) Lg() *MultiSelect { return m.Size(SizeLg) }

// Xs sizes the field to the smallest size.
func (m *MultiSelect) Xs() *MultiSelect { return m.Size(SizeXs) }

// Shape sets the shape of the field and returns it, for chaining.
func (m *MultiSelect) Shape(s Shape) *MultiSelect {
	m.shape = s
	m.apply()
	return m
}

// Width sets the width of the field in DIPs and returns it, for
// chaining.
func (m *MultiSelect) Width(w float32) *MultiSelect {
	m.Element = m.Element.Width(w)
	return m
}

// Placeholder shows s in the empty field and returns the field, for
// chaining.
func (m *MultiSelect) Placeholder(s string) *MultiSelect {
	m.Element = m.Element.Placeholder(s)
	return m
}

// Label names the field for assistive technology and returns it, for
// chaining.
func (m *MultiSelect) Label(s string) *MultiSelect {
	m.Element = m.Element.Label(s)
	return m
}

// OnChanged calls fn with the choices when they changed in this frame,
// and returns the field, for chaining.
func (m *MultiSelect) OnChanged(fn func(values []string)) *MultiSelect {
	if fn != nil && m.Changed() {
		fn(*m.tokens)
	}
	return m
}

// Values returns the choices the field holds.
func (m *MultiSelect) Values() []string { return *m.tokens }

// Control returns the element a field describes: the field of chips,
// which is the one that takes the focus.
func (m *MultiSelect) Control() ui.Element { return m.Element }

// Disabled takes the multi-select out of use and returns it, for
// chaining.
func (m *MultiSelect) Disabled(on bool) *MultiSelect {
	m.Element = m.Element.Disabled(on)
	return m
}

// Invalid marks the multi-select as not holding a good value and returns
// it, for chaining. The field takes the danger border.
func (m *MultiSelect) Invalid(on bool) *MultiSelect {
	m.invalid = on
	m.Element = m.Element.Invalid(on)
	m.apply()
	return m
}

// Changed reports whether the choices changed in this frame.
func (m *MultiSelect) Changed() bool { return m.Element.Changed() }

// apply paints the look of the variant and the size on the field.
func (m *MultiSelect) apply() {
	t := m.c.Theme()
	l := fieldLookOf(m.variant, t)
	if !l.native {
		v, h := fieldPadding(t, m.size)
		border := l.border
		if m.invalid {
			border = t.Danger
		}
		m.Element = m.Element.Background(l.bg).Border(l.width, border).Radius(t.Radius).Padding(v, h)
	}
	switch m.shape {
	case ShapeRound, ShapeCircle:
		m.Element = m.Element.Radius(1000)
	default:
		m.Element = m.Element.Radius(t.Radius)
	}
	m.Element = m.Element.FontSize(fieldFontSize(t, m.size))
}
