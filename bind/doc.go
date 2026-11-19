// Package bind is the sugar over my-ui's component constructors. [Bind]
// takes one frame's context and returns a [UI] handle: a method per
// component with the constructor's arguments minus the *ui.Context.
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
// Two constructors are generic and cannot become methods, since Go methods
// have no type parameters: controls.NewRadio[T] and forms.NewFrame[T].
// UI.Radio offers the common string instance of the former; for any other
// value type, call controls.NewRadio[T] directly. Use forms.NewFrame when
// building a custom field wrapper.
package bind
