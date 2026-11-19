// Package ai holds application components for chat interfaces: a virtualized
// message list, a searchable model picker, and a reasoning-settings panel.
//
// These components compose my-ui's styled widgets with mygo's public API and
// its Markdown view. They hold no application state: messages, selections and
// parameter values are pointers the app owns, and colors and metrics are read
// from the frame's theme, so a theme change lands on the next frame.
//
//	func view(c *ui.Context) {
//		h := bind.Bind(c)
//		h.ChatList(app.messages, nil, ai.ChatListOptions{})
//		h.ModelPicker(&app.query, app.models).Outline()
//		h.Reasoning(ai.ReasoningOptions{Effort: &app.effort})
//	}
//
// The package imports no assets and never reaches into mygo's internal
// packages.
package ai
