package bind

import (
	"github.com/egoist/mygo/ui"
	aipkg "github.com/wuyangwang/my-ui/ai"
)

// ChatList builds a virtualized message stream over messages and state.
func (u UI) ChatList(messages []aipkg.ChatMessage, state *ui.ListState, o aipkg.ChatListOptions) *aipkg.ChatList {
	return aipkg.NewChatList(u.c, messages, state, o)
}

// ChatListStatic builds a message stream from a fixed set of messages.
func (u UI) ChatListStatic(messages []aipkg.ChatMessage) *aipkg.ChatList {
	return aipkg.NewChatListStatic(u.c, messages)
}

// ModelPicker builds a searchable model drop-down over text and options.
func (u UI) ModelPicker(text *string, options []aipkg.ModelOption) *aipkg.ModelPicker {
	return aipkg.NewModelPicker(u.c, text, options)
}

// Reasoning builds a thinking-settings panel over the app's values.
func (u UI) Reasoning(o aipkg.ReasoningOptions) *aipkg.Reasoning {
	return aipkg.NewReasoning(u.c, o)
}
