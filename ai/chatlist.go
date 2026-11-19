package ai

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
	mdpkg "github.com/wuyangwang/my-ui/markdown"
)

// Role says who wrote a chat message.
type Role uint8

const (
	// System marks a short instruction or note, drawn centered and muted.
	System Role = iota
	// User marks the person's turn, drawn on the right.
	User
	// Assistant marks the model's turn, drawn on the left with an agent avatar.
	Assistant
)

// ChatMessage is one turn of a conversation. The app owns the slice these
// live in: the list reads it each frame and never copies or edits it.
//
// Text is the finished reply. Stream, when set, is read once per frame for a
// reply still being written; app code swaps it for Text when the reply is
// done. Pending draws a spinner in place of an empty body.
type ChatMessage struct {
	ID      string
	Role    Role
	Text    string
	Stream  func() string
	Time    string
	Pending bool
}

// ChatListOptions configures NewChatList.
type ChatListOptions struct {
	// NoAvatar leaves out the avatars beside user and assistant turns.
	// Avatars are drawn by default; set this to omit them.
	NoAvatar bool
	// AvatarSize is the avatar's side in DIPs; 32 when zero.
	AvatarSize float32
	// Time shows each message's Time under its body.
	Time bool
	// NoFollow leaves the list where it is as messages arrive. The default
	// follows the newest message, stops when the user scrolls up and takes
	// hold again when they scroll back to the end.
	NoFollow bool
	// OnEndReached asks the app to load more rows. Mygo calls it once as
	// the user nears the list's end, not again while they stay there.
	// Chat histories usually load older messages at the top: mygo has no
	// top-reached callback, so that needs an app-side trigger (see the
	// spec's known limits).
	OnEndReached func(last int)
	// EmptyText is shown instead of the list when there are no messages.
	EmptyText string
	// UserName and AgentName name the avatars; the frame's locale supplies
	// them when empty.
	UserName  string
	AgentName string
}

// ChatList is a virtualized message stream built on mygo's List. Rows keep
// their place by message ID, and the framework's FollowEnd does the stick-
// to-bottom work: no timers, no ScrollState arithmetic and no goroutines.
//
// messages is read, not written: a fixed set can be passed as a literal,
// and data from an interface can be passed as the app's slice each frame.
// Either way the app owns the data. ListState is optional; a nil one is
// kept on the element, so scrolling survives across frames in immediate
// mode. Pass your own state to read Selected or drive the list yourself.
type ChatList struct {
	ui.Element
	c        *ui.Context
	messages []ChatMessage
	state    *ui.ListState
	o        ChatListOptions
	empty    *string
	msgsRef  *[]ChatMessage
}

// NewChatList builds a message stream over messages and state. messages is
// a snapshot of this frame: pass a literal for fixed data, or app.messages
// for data that grows as replies arrive. state may be nil.
func NewChatList(c *ui.Context, messages []ChatMessage, state *ui.ListState, o ChatListOptions) *ChatList {
	l := &ChatList{c: c, messages: messages, state: state, o: o}
	t := c.Theme()
	root := ui.Box(c)
	l.empty = ui.Local(root, "chat-empty", func() string { return o.EmptyText })
	// The Key closure outlives this frame; it reads the latest snapshot
	// from the element's slot, not the slice head captured on frame one.
	l.msgsRef = ui.Local(root, "chat-messages", func() []ChatMessage { return messages })
	*l.msgsRef = messages
	if state == nil {
		// Local keeps one ListState for the life of the element. A plain
		// new(ListState) here would reset the place on every frame.
		l.state = *ui.Local(root, "chat-list-state", func() *ui.ListState { return new(ui.ListState) })
	}

	s := l.state
	s.Key = func(i int) any {
		msgs := *l.msgsRef
		if i >= 0 && i < len(msgs) && msgs[i].ID != "" {
			return msgs[i].ID
		}
		return i
	}
	s.FollowEnd = !o.NoFollow
	if !o.NoFollow {
		s.FollowAnimation = &ui.Smooth
	}
	if o.OnEndReached != nil {
		s.EndReached = o.OnEndReached
	}

	root.Children(func() {
		if len(messages) == 0 {
			if *l.empty != "" {
				ui.Row(c).Grow(1).Center().Children(func() {
					ui.Text(c, *l.empty).FontSize(TextSize(t, TextSm)).TextColor(t.TextMuted)
				})
			}
			return
		}
		ui.List(c, s, len(messages), func(i int) { l.row(c, i) }).
			Grow(1).Justify(ui.End).Gap(t.Space(3)).Padding(t.Space(3)).
			Background(t.Background).Radius(Radius(t, RadiusLg)).Transition(StateTransition())
	})
	l.Element = root
	return l
}

// NewChatListStatic builds a message stream from a fixed set of messages,
// with avatars and stick-to-bottom on and no state the app has to manage.
//
//	ai.NewChatListStatic(c, []ai.ChatMessage{
//	    {Role: ai.User, Text: "你好"},
//	    {Role: ai.Assistant, Text: "你好！"},
//	}).Height(320)
func NewChatListStatic(c *ui.Context, messages []ChatMessage) *ChatList {
	return NewChatList(c, messages, nil, ChatListOptions{})
}

// row builds one message. User turns sit on the right, assistant turns on
// the left, and system notes are centered between them.
func (l *ChatList) row(c *ui.Context, i int) {
	m := l.messages[i]
	t := c.Theme()
	switch m.Role {
	case System:
		ui.Row(c).FillWidth().Justify(ui.Center).Children(func() {
			ui.Text(c, m.Text).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
				Padding(t.Space(1), t.Space(2))
		})
	case User:
		ui.Row(c).AlignSelf(ui.End).AlignItems(ui.Start).Gap(t.Space(2)).Children(func() {
			l.body(c, m, t, true)
			if !l.o.NoAvatar {
				l.avatar(c, m)
			}
		})
	default:
		ui.Row(c).AlignSelf(ui.Start).AlignItems(ui.Start).Gap(t.Space(2)).Children(func() {
			if !l.o.NoAvatar {
				l.avatar(c, m)
			}
			l.body(c, m, t, false)
		})
	}
}

// body paints a message's bubble. A reply still being written goes through
// markdown.NewStream; a finished one through NewMarkdown, so half-written
// syntax is shown as its likely final form while it streams.
func (l *ChatList) body(c *ui.Context, m ChatMessage, t *ui.Theme, user bool) {
	bubble := ui.Column(c).Gap(t.Space(1)).Padding(t.Space(2), t.Space(3)).
		Radius(Radius(t, RadiusLg)).Transition(StateTransition())
	if user {
		bubble.Background(t.Accent.Alpha(0.12))
	} else {
		bubble.Background(t.Surface).Border(1, t.Border)
	}
	bubble.Children(func() {
		switch {
		case m.Pending && m.Text == "" && m.Stream == nil:
			displaypkg.NewSpin(c).Size(SizeSm)
		case m.Stream != nil:
			mdpkg.NewStream(c, m.Stream)
		default:
			mdpkg.NewMarkdown(c, m.Text)
		}
		if l.o.Time && m.Time != "" {
			ui.Text(c, m.Time).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted)
		}
	})
}

// avatar draws a deterministic agent for the assistant and a user Initials
// avatar for the person. Both take their colors from the frame's theme.
func (l *ChatList) avatar(c *ui.Context, m ChatMessage) {
	size := l.o.AvatarSize
	if size <= 0 {
		size = 32
	}
	if m.Role == User {
		name := l.o.UserName
		if name == "" {
			name = Message(c, "user")
		}
		displaypkg.NewAvatar(c, name, displaypkg.AvatarOptions{
			Size: size, Shape: ShapeCircle, Fallback: displaypkg.AvatarText,
		})
		return
	}
	name := l.o.AgentName
	if name == "" {
		name = Message(c, "assistant")
	}
	state := displaypkg.AgentDone
	if m.Pending {
		state = displaypkg.AgentThinking
	}
	displaypkg.NewAgentAvatar(c, displaypkg.AgentAvatarOptions{
		Seed: m.ID, Name: name, Size: size, State: state,
	})
}

// ScrollToEnd glides to the newest message and resumes following it.
func (l *ChatList) ScrollToEnd() {
	if l.state != nil {
		l.state.ScrollToEnd()
	}
}

// AtEnd reports whether the last layout showed the newest message.
func (l *ChatList) AtEnd() bool {
	return l.state != nil && l.state.AtEnd()
}

// HoldPlace keeps the list from following its end in the frame that changes
// the rows, as opening a fold in the last message does.
func (l *ChatList) HoldPlace() {
	if l.state != nil {
		l.state.HoldPlace()
	}
}

// EmptyText sets the message shown when the conversation is empty and
// returns the list, for chaining.
func (l *ChatList) EmptyText(s string) *ChatList {
	if l.empty != nil {
		*l.empty = s
	}
	return l
}

// Height pins the list to h DIPs and returns it, for chaining. Without it
// the list takes the height its parent gives it.
func (l *ChatList) Height(h float32) *ChatList {
	l.Element = l.Element.Height(h)
	return l
}
