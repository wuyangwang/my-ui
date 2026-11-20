# ChatList 规格书

> 位置：组件在 [`ai/`](../../ai) 子包；主题/tokens 与文案表在根包 `myui`。
>

## API 契约

| API | 作用 |
|---|---|
| `type Role uint8` | `System`、`User`、`Assistant` 三种角色 |
| `type ChatMessage struct` | 一条消息：`ID`、`Role`、`Text`、`Stream func() string`、`Time`、`Pending` |
| `type ChatListOptions struct` | 头像、时间、跟随、加载更多、空文案等配置 |
| `func NewChatList(c *ui.Context, messages []ChatMessage, state *ui.ListState, o ChatListOptions) *ChatList` | 构建消息流；`messages` 是本帧快照，`state` 可为 nil |
| `func NewChatListStatic(c *ui.Context, messages []ChatMessage) *ChatList` | 固定数据的极简构造器：默认头像与贴底跟随 |
| `func (l *ChatList) ScrollToEnd()` | 平滑回到底部并恢复跟随 |
| `func (l *ChatList) AtEnd() bool` | 上次布局是否停在底部 |
| `func (l *ChatList) HoldPlace()` | 本帧改变行时保持位置（如展开末条消息的折叠） |
| `func (l *ChatList) EmptyText(s string) *ChatList` | 设置空会话提示 |
| `func (l *ChatList) Height(h float32) *ChatList` | 固定消息流高度 |

### 两种数据源

组件每帧重建，只读传入的消息快照，不写回，所以两种数据源用同一个 API：

**1. 本地固定数据**：直接传字面量，用 `NewChatListStatic`，连变量和 `ListState` 都不必声明。

```go
ai.NewChatListStatic(c, []ai.ChatMessage{
    {Role: ai.System, Text: "会话演示"},
    {Role: ai.User, Text: "帮我写个函数"},
    {Role: ai.Assistant, Text: "好。"},
}).Height(320)
```

**2. 接口返回的数据**：app 持有 slice，每帧传当前值；响应到达时 append，下一帧自动生效，没有刷新调用。

```go
type App struct {
    messages []ai.ChatMessage
    draft    string
}

// 每帧：
ai.NewChatList(c, a.messages, nil, ai.ChatListOptions{}).Height(360)

// 收到一条完整消息：
a.messages = append(a.messages, ai.ChatMessage{
    ID: strconv.Itoa(len(a.messages)), Role: ai.Assistant, Text: reply,
})

// 流式回复：闭包读 app 的累积文本，完成时按 ID 写回切片元素。
a.messages = append(a.messages, ai.ChatMessage{ID: "r1", Role: ai.Assistant,
    Stream: func() string { return a.draft }})
// ...a.draft 逐 token 追加...
for i := range a.messages {
    if a.messages[i].ID == "r1" {
        a.messages[i].Stream = nil
        a.messages[i].Text = a.draft
    }
}
```

等待接口时，在末尾放一条 `Pending: true` 的消息即可显示 typing 指示；出错时放一条 `Role: System` 的消息说明原因。

`bind.Bind(c)` 下省略 `c`：`h.ChatList(a.messages, nil, ai.ChatListOptions{})`，固定数据用 `h.ChatListStatic(msgs)`。

## 状态与限制

### 已知限制

- `state` 为 nil 时组件把 `ListState` 存进元素的 `ui.Local`，跨帧保持滚动位置与跟随状态；app 只是读不到 `Selected` 等字段，需要时传自己的指针。
- `ChatMessage.ID` 决定稳定行键；留空退化为行号，前面插入历史消息时滚动位置可能跳。固定数据不受影响，接口数据请自己分配唯一 ID。
- 顶部加载历史消息：mygo 只有「接近末尾」的 `EndReached`，没有顶部回调。app 可在构建列表后读 `ListState.Visible()` 判断首行是否为 0，再配合一个 `ui.Local` 标志防重复来触发加载。
- `Stream` 返回的必须是**累积**源文本（只允许追加），这是上游 `markdown.Stream` 的约定；回复结束后要写回 `Text` 并把 `Stream` 置 nil。
