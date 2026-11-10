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
//	Tooltip  a tip over an anchor, when the pointer rests on it
//	Popover  a panel over an anchor, opened by a click on it
//	Modal    a dialog over the window, opened by a trigger
//	Badge    a small label: the variants, in three sizes
//	Icon     an SVG icon in the sizes 14, 20 and 24 DIPs
//	Scaffold a page: the title bar's safe area, a header, the scrollable
//	         content and a bottom bar
//
// Every component takes the frame's context and returns itself, so a
// preset may be chained onto it:
//
//	myui.NewButton(c, "Save").Primary().Sm()
//	myui.NewInput(c, &app.email).Outline().Clearable()
//
// # The shared presets
//
// The variants and the sizes are one vocabulary for all the components,
// so a button and an input take the same words:
//
//	Variant   Native (the zero value, the built-in look), Secondary (an
//	          alias of Native), Primary, Outline, Ghost, Destructive, Link
//	Size      SizeMd (the zero value), SizeSm, SizeLg
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
//	myui.NewButton(c, "Save").Primary()
//
//	h := myui.Bind(c)
//	h.Button("Save").Primary()
//
// # The subpackages
//
// The icons of Lucide live in the icons package, which an app imports
// only when it uses icons:
//
//	ui.Icon(c, icons.Save())
//
// The fonts and the AI application components are planned; see the
// roadmap in the repository for what is next.
package myui
