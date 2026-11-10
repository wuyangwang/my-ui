package myui

import "github.com/egoist/mygo/ui"

// UI is the handle to my-ui in one frame. It holds the context of the
// frame, and whatever the library holds beside it later.
//
// Bind it once in a view, and the components take no context of their
// own:
//
//	func view(c *ui.Context) {
//		h := myui.Bind(c)
//		h.Apply(myui.Options{Appearance: myui.AppearanceDark})
//		h.Button("Save").Primary()
//		h.Input(&app.name).Outline().Placeholder("Name")
//	}
//
// The handle holds the context of one frame only: bind it again in the
// next frame, as the view runs. It keeps no state between frames, so
// the theme, the palette and the desktop's own changes reach the
// components on the next frame, as they should.
//
// The handle is sugar over the constructors: myui.NewButton(c, "Save")
// says the same thing, and both are there to choose from. The handle is
// a value, so binding it costs nothing.
type UI struct {
	c *ui.Context
}

// Bind binds a frame's context to the components of my-ui.
func Bind(c *ui.Context) UI {
	return UI{c: c}
}

// Context returns the context of the frame the handle is bound to.
func (u UI) Context() *ui.Context { return u.c }

// Theme returns the theme of the frame.
func (u UI) Theme() *ui.Theme { return u.c.Theme() }

// Apply applies the options to the theme of the frame and returns the
// theme that is in place: nil when the frame follows the desktop.
func (u UI) Apply(o Options) *ui.Theme { return Apply(u.c, o) }

// Button creates a button.
func (u UI) Button(label string) *Button { return NewButton(u.c, label) }

// Input creates a text input editing *value.
func (u UI) Input(value *string) *Input { return NewInput(u.c, value) }

// TextArea creates a multi-line text field editing *value.
func (u UI) TextArea(value *string) *TextArea { return NewTextArea(u.c, value) }

// Select creates a drop-down choosing one of options into *value.
func (u UI) Select(value *string, options []SelectOption) *Select {
	return NewSelect(u.c, value, options)
}

// ComboBox creates a drop-down whose trigger is a text field: what is
// typed narrows the options down.
func (u UI) ComboBox(text *string, options []SelectOption) *ComboBox {
	return NewComboBox(u.c, text, options)
}

// MultiSelect creates a field of tags: the chosen options show as chips.
func (u UI) MultiSelect(tokens *[]string, options []SelectOption) *MultiSelect {
	return NewMultiSelect(u.c, tokens, options)
}

// Tooltip shows a tip with a text over the anchor.
func (u UI) Tooltip(anchor ui.Element, text string) *Tooltip {
	return NewTooltip(u.c, anchor, text)
}

// Popover shows a panel over its anchor while open is true.
func (u UI) Popover(open *bool, anchor func() ui.Element, build func(panel ui.Element)) *Popover {
	return NewPopover(u.c, open, anchor, build)
}

// Modal shows a dialog over the window while open is true.
func (u UI) Modal(open *bool, trigger func() ui.Element, build func(panel ui.Element)) *Modal {
	return NewModal(u.c, open, trigger, build)
}

// Badge creates a badge.
func (u UI) Badge(text string) *Badge { return NewBadge(u.c, text) }

// Icon creates an icon drawing svg.
func (u UI) Icon(svg *ui.SVG) *Icon { return NewIcon(u.c, svg) }

// Scaffold lays out a page: the title bar's safe area, a header, the
// scrollable content and a bottom bar.
func (u UI) Scaffold(o ScaffoldOptions) { Scaffold(u.c, o) }
