package ai

import (
	"strconv"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestChatListDrawsRolesAndKeysByMessageID(t *testing.T) {
	msgs := []ChatMessage{
		{ID: "s", Role: System, Text: "系统提示"},
		{ID: "u", Role: User, Text: "你好"},
		{ID: "a", Role: Assistant, Text: "收到，这是回复。"},
	}
	var state ui.ListState
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewChatList(c, msgs, &state, ChatListOptions{}).Height(240)
	}
	tt := ui.NewTester(view, 420, 300)
	for _, s := range []string{"系统提示", "你好", "收到，这是回复。"} {
		if _, ok := tt.Find(s); !ok {
			t.Errorf("the message %q is missing", s)
		}
	}
	if !state.FollowEnd {
		t.Error("a chat follows its end by default")
	}
	if state.Key == nil {
		t.Fatal("the list has no stable key")
	}
	for i, want := range []string{"s", "u", "a"} {
		if got := state.Key(i); got != want {
			t.Errorf("key %d is %v, want %q", i, got, want)
		}
	}
}

func TestChatListNoFollowAndEmptyText(t *testing.T) {
	var state ui.ListState
	msgs := []ChatMessage{}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewChatList(c, msgs, &state, ChatListOptions{
			NoFollow: true, EmptyText: "暂无消息",
		}).Height(240)
	}
	tt := ui.NewTester(view, 420, 300)
	if state.FollowEnd {
		t.Error("NoFollow leaves FollowEnd off")
	}
	if _, ok := tt.Find("暂无消息"); !ok {
		t.Error("the empty conversation shows no hint")
	}
}

func TestChatListStaticKeepsStateAndKeyAcrossFrames(t *testing.T) {
	var first *ui.ListState
	frame := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		frame++
		msgs := []ChatMessage{{ID: "id" + strconv.Itoa(frame), Role: User, Text: "固定数据"}}
		l := NewChatListStatic(c, msgs)
		if first == nil {
			first = l.state
		} else if first != l.state {
			t.Error("a nil state must be kept on the element across frames")
		}
		if want := "id" + strconv.Itoa(frame); l.state.Key(0) != want {
			t.Errorf("the key is %v, want %q", l.state.Key(0), want)
		}
	}
	tt := ui.NewTester(view, 420, 300)
	tt.Advance(100 * time.Millisecond)
	if first == nil {
		t.Fatal("the static list has no state")
	}
	if !first.FollowEnd {
		t.Error("the static list follows its end")
	}
	if _, ok := tt.Find("固定数据"); !ok {
		t.Error("the fixed message is missing")
	}
}

func TestChatListReadsTheStreamEachFrame(t *testing.T) {
	calls := 0
	msgs := []ChatMessage{{
		ID: "a", Role: Assistant,
		Stream: func() string { calls++; return "正在写入" },
	}}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewChatList(c, msgs, nil, ChatListOptions{}).Height(240)
	}
	tt := ui.NewTester(view, 420, 300)
	tt.Advance(50 * time.Millisecond)
	if calls == 0 {
		t.Error("the stream source was never read")
	}
	if _, ok := tt.Find("正在写入"); !ok {
		t.Error("the streamed text is missing")
	}
}
