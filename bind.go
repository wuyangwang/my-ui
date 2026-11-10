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

// Badge creates a badge.
func (u UI) Badge(text string) *Badge { return NewBadge(u.c, text) }

// Icon creates an icon drawing svg.
func (u UI) Icon(svg *ui.SVG) *Icon { return NewIcon(u.c, svg) }

// Scaffold lays out a page: the title bar's safe area, a header, the
// scrollable content and a bottom bar.
func (u UI) Scaffold(o ScaffoldOptions) { Scaffold(u.c, o) }
