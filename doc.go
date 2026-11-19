// Package myui is a component library for MyGo native UI (package ui):
// design tokens, theme presets, styled variants of built-in widgets, and
// composite desktop components such as pages, overlays and AI chat views
// that MyGo does not ship. It is a third-party module of its own; it
// depends only on the public API of github.com/egoist/mygo.
//
// Package myui holds the tokens, the theme presets and the lighter,
// always-useful components. Weightier packages — icons, fonts and the AI
// application components — live in subpackages the app imports only when it
// uses them, so nothing an app does not import is linked into it.
//
// # Components
//
// Each component is a preset over what MyGo already draws, in the shadcn
// vocabulary: a variant says how it looks, a size says how roomy it is.
//
//	Button   a push button: Primary, Outline, Ghost, Destructive, Link
//	Input    a text field: a frame, a prefix and a suffix, a clear button,
//	         events, and an error state
//	TextArea a multi-line text field, as an Input is, that grows with
//	         its text
//	Select   a drop-down: the trigger looks like an input, the popup
//	         holds the options, the chosen one marked
//	ComboBox a drop-down whose trigger is a text field: what is typed
//	         narrows the options down
//	MultiSelect a field of tags: the chosen options show as chips
//	Field    a form field: its label, the control in it, a description
//	         and an error, laid out as one thing
//	Checkbox a box that is checked or not, with a label
//	Radio    a circle that is selected or not: one of a group's choices
//	Switch   a knob that is on or off
//	Slider   a knob on a track, that picks a number between two
//	Tooltip  a tip over an anchor, when the pointer rests on it
//	Popover  a panel over an anchor, opened by a click on it
//	Modal    a dialog over the window, opened by a trigger
//	Drawer   a panel over the edge of the window, opened by a trigger
//	Toast    a message at the corner of the window, for a few seconds
//	Alert    a message in the page: a card with a title and a description
//	Badge    a small label: the variants, in three sizes
//	Icon     an SVG icon in the sizes 14, 20 and 24 DIPs
//	Card     a panel of the page: a title, a description and a content
//	Tabs     the tabs of a page: one of them is selected
//	Progress a bar that shows how far along something is
//	Kbd      the key caps of a shortcut
//	Scaffold a page: the title bar's safe area, a header, the scrollable
//	         content and a bottom bar
//
// Every component takes the frame's context and returns itself, so a
// preset may be chained onto it:
//
//	controls.NewButton(c, "Save").Primary().Sm()
//	forms.NewInput(c, &app.email).Outline().Clearable()
//
// # The shared presets
//
// The variants and the sizes are one vocabulary for all the components,
// so a button and an input take the same words:
//
//	Variant   Default (the zero value, the built-in look; Native is its
//	          old name), Secondary (a muted fill), Primary, Outline,
//	          Ghost, Destructive, Link, Success, Info, Warning
//	Size      SizeMd (the zero value), SizeSm, SizeLg, SizeXs
//	Shape     ShapeRect (the zero value), ShapeRound, ShapeCircle
//
// # The theme
//
// The theme follows the desktop: the appearance, the accent, the high
// contrast and the text scale. Apply sets it for the frame, NewTheme
// builds one for a subtree:
//
//	myui.Apply(c, myui.Options{Appearance: myui.AppearanceDark, Palette: myui.PaletteBlue})
//	ui.Themed(c, myui.NewTheme(c, myui.Options{Appearance: myui.AppearanceLight}))
//
// # Two ways to write it
//
// The constructors take the context of the frame. The handle binds it
// once per window, so a view need not pass it again:
//
//	controls.NewButton(c, "Save").Primary()
//
//	h := bind.Bind(c)
//	h.Button("Save").Primary()
//
// # The subpackages
//
// The icons of Lucide live in the icons package, which an app imports
// only when it uses icons:
//
//	ui.Icon(c, icons.Save())
//
// The fonts package ships OFL fonts for ui.RegisterFont; the AI
// application components are planned. See the roadmap in the repository
// for what is next.
package myui
