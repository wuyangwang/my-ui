package main

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	aipkg "github.com/wuyangwang/my-ui/ai"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// chatDemo lives on the demo element: leaving the page drops the messages
// and the list's place with it, and no package-level state is involved.
type chatDemo struct {
	messages []aipkg.ChatMessage
	state    ui.ListState
	frame    int
}

func (s *showcase) chatListPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "ChatList", "虚拟化消息流；新消息贴底，上滚即停，行键跟随消息 ID", func() {
		s.card(c, t, func() {
			box := ui.Box(c.Key("chat-list-demo"))
			st := ui.Local(box, "chat", func() chatDemo { return chatDemo{} })
			if len(st.messages) == 0 {
				st.messages = []aipkg.ChatMessage{
					{ID: "0", Role: aipkg.System, Text: "这是一段会话演示"},
					{ID: "1", Role: aipkg.User, Text: "帮我写一个 Go 函数"},
					{ID: "2", Role: aipkg.Assistant, Text: "当然。**告诉**我它需要做什么。"},
				}
			}
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				if ctrlpkg.NewButton(c, "追加回复").Outline().Sm().Clicked() {
					id := strconv.Itoa(len(st.messages))
					st.messages = append(st.messages, aipkg.ChatMessage{
						ID: id, Role: aipkg.Assistant, Text: "第 " + id + " 条回复。",
					})
				}
				if ctrlpkg.NewButton(c, "回到底部").Ghost().Sm().Clicked() {
					st.state.ScrollToEnd()
				}
			})
			box.Height(280).Children(func() {
				aipkg.NewChatList(c, st.messages, &st.state, aipkg.ChatListOptions{
					Time: true,
				}).Height(280)
			})
		})
	})

	s.sectionNote(c, t, "流式回复", "Stream 每帧读取 app 持有的源文本，mygo 负责增量解析与淡入", func() {
		s.card(c, t, func() {
			box := ui.Box(c.Key("chat-stream-demo"))
			st := ui.Local(box, "chat-stream", func() chatDemo { return chatDemo{} })
			st.frame++
			rs := []rune("这是一条正在写入的回复，逐字显示。")
			n := min(len(rs), 1+st.frame/3)
			source := string(rs[:n])
			if st.frame > 3*len(rs)+60 {
				st.frame = 0
			}
			box.Height(160).Children(func() {
				aipkg.NewChatList(c, []aipkg.ChatMessage{{
					ID: "stream", Role: aipkg.Assistant, Stream: func() string { return source },
				}}, nil, aipkg.ChatListOptions{Time: true}).Height(160)
			})
		})
	})
}
