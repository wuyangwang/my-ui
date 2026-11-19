package bind

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	displaypkg "github.com/wuyangwang/my-ui/display"
	formspkg "github.com/wuyangwang/my-ui/forms"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
)

// UI is the handle to my-ui in one frame. It holds the context of the
// frame, and whatever the library holds beside it later.
//
// Bind it once in a view, and the components take no context of their
// own:
//
//	func view(c *ui.Context) {
//		h := bind.Bind(c)
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
// The handle is sugar over the constructors: controls.NewButton(c, "Save")
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

// ctrlpkg.Button creates a button.
func (u UI) Button(label string) *ctrlpkg.Button { return ctrlpkg.NewButton(u.c, label) }

// formspkg.Input creates a text input editing *value.
func (u UI) Input(value *string) *formspkg.Input { return formspkg.NewInput(u.c, value) }

// formspkg.TextArea creates a multi-line text field editing *value.
func (u UI) TextArea(value *string) *formspkg.TextArea { return formspkg.NewTextArea(u.c, value) }

// formspkg.Select creates a drop-down choosing one of options into *value.
func (u UI) Select(value *string, options []formspkg.SelectOption) *formspkg.Select {
	return formspkg.NewSelect(u.c, value, options)
}

// formspkg.ComboBox creates a drop-down whose trigger is a text field: what is
// typed narrows the options down.
func (u UI) ComboBox(text *string, options []formspkg.SelectOption) *formspkg.ComboBox {
	return formspkg.NewComboBox(u.c, formspkg.ComboOptions{
		Value:   text,
		Options: options,
	})
}

// formspkg.MultiSelect creates a field of tags: the chosen options show as chips.
func (u UI) MultiSelect(tokens *[]string, options []formspkg.SelectOption) *formspkg.MultiSelect {
	return formspkg.NewMultiSelect(u.c, tokens, options)
}

// ctrlpkg.Checkbox creates a check box editing *checked, with a label.
func (u UI) Checkbox(checked *bool, label string) *ctrlpkg.Checkbox {
	return ctrlpkg.NewCheckbox(u.c, checked, label)
}

// ctrlpkg.Switch creates a switch editing *on.
func (u UI) Switch(on *bool) *ctrlpkg.Switch { return ctrlpkg.NewSwitch(u.c, on) }

// ctrlpkg.Slider creates a slider setting *value between lo and hi.
func (u UI) Slider(value *float64, lo, hi float64) *ctrlpkg.Slider {
	return ctrlpkg.NewSlider(u.c, value, lo, hi)
}

// overlaypkg.Tooltip shows a tip with a text over the anchor.
func (u UI) Tooltip(anchor ui.Element, text string) *overlaypkg.Tooltip {
	return overlaypkg.NewTooltip(u.c, anchor, text)
}

// displaypkg.Badge creates a badge.
func (u UI) Badge(text string) *displaypkg.Badge { return displaypkg.NewBadge(u.c, text) }

// displaypkg.Card creates a card with a title.
func (u UI) Card(title string) *displaypkg.Card { return displaypkg.NewCard(u.c, title) }

// displaypkg.Tabs creates tabs, one of them selected.
func (u UI) Tabs(selected *int, tabs ...displaypkg.Tab) *displaypkg.Tabs {
	return displaypkg.NewTabs(u.c, selected, tabs...)
}

// displaypkg.Progress creates a progress bar at value.
func (u UI) Progress(value float64) *displaypkg.Progress { return displaypkg.NewProgress(u.c, value) }

// displaypkg.Kbd shows the keys of a shortcut.
func (u UI) Kbd(keys ...string) *displaypkg.Kbd { return displaypkg.NewKbd(u.c, keys...) }

// ctrlpkg.Icon creates an icon drawing svg.
func (u UI) Icon(svg *ui.SVG) *ctrlpkg.Icon { return ctrlpkg.NewIcon(u.c, svg) }
