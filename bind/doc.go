// Package bind is the sugar over my-ui's component constructors. [Bind]
// takes one frame's context and returns a [UI] handle: a method per
// frame-component constructor with the constructor's arguments minus the
// *ui.Context.
//
//	func view(c *ui.Context) {
//		h := bind.Bind(c)
//		h.Apply(myui.Options{Appearance: myui.AppearanceDark})
//		h.Button("Save").Primary()
//		h.Input(&app.name).Outline().Placeholder("Name")
//		h.ChatList(app.messages, nil, ai.ChatListOptions{})
//	}
//
// Bind the context again on every frame. The handle keeps no state of its
// own, so the theme and the desktop's changes reach the components on the
// next frame, as they do through the constructors.
//
// Coverage. The handle covers every non-generic frame-component constructor
// in the controls, forms, overlays, display, layout, chart, markdown and ai
// packages, plus the root NewTheme and chart.NewPalette helpers (whose only
// argument is taken from the frame's theme). Two families of constructors
// are not methods:
//
//   - Generic constructors, since Go methods cannot have type parameters:
//     controls.NewRadio[T], forms.NewFrame[T], forms.NewImageSelect[T] and
//     display.NewOutline[K]. UI.Radio offers the common string instance of
//     the first, with the same arguments minus the context; for any other
//     value type call controls.NewRadio[T] directly. Call the other three
//     directly with the frame's context when building a custom wrapper.
//   - The window package, whose New and Open create a real OS window, take
//     no frame context and return *mygo.Window. Only the window package
//     itself may import mygo's root package, so a handle method cannot name
//     that return type. Use window.New or window.Open directly for the
//     frameless shell.
package bind
