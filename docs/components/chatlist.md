# ChatList 规格书

> 位置：组件在 [`ai/`](../../ai) 子包；主题/tokens 与文案表在根包 `myui`。
>
> 状态：**已实现**（`ai/chatlist.go`）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| shadcn AI Elements `Chat` / `Message` / `Message Scroller` | 消息流是「会话 + 单条消息」两层；流式回复逐 token 到来，滚动要贴底但允许用户上滚离开 |
| Vercel AI Elements `Conversation` / `Reasoning` | 助手消息带头像与状态，用户消息靠右，系统消息弱化居中 |
| mygo `ui.List` + `ui.ListState` | 锚点式虚拟化：行高动态、可增删；`FollowEnd`/`FollowAnimation` 原生贴底跟随，`Key` 稳定标识，`EndReached` 加载更早消息 |
| my-ui `markdown/` | 流式 `NewStream` / 完稿 `NewMarkdown`，主题每帧注入 |

## 2. 能否实现

- 虚拟化与贴底：`ui.List` + `ListState.FollowEnd`/`FollowAnimation` 已覆盖「新消息贴底、上滚即停、回到底部再跟随」，无需手写 `ScrollState`/`MaxFloat32` 弹簧 ✓
- 稳定行键：`ListState.Key` 传消息 `ID`，前面插入历史消息时位置与选择不丢 ✓
- 流式正文：`markdown.NewStream` 每帧读取 app 的源文本；完稿换 `markdown.NewMarkdown` ✓
- 头像：`display.NewAgentAvatar`（助手）与 `display.NewAvatar`（用户），颜色取主题槽位 ✓
- 输入条：本组件只画消息流；底部输入框由 app 用 `forms`/`layout` 组合，见「已知限制」 ✓

## 3. 实现到何种程度

- `ChatMessage`：app 持有的消息结构（`ID`、`Role`、`Text`、`Stream`、`Time`、`Pending`）。
- `ChatListOptions`：`NoAvatar`（默认画头像）、`AvatarSize`、`Time`、`NoFollow`（默认跟随）、`OnEndReached`、`EmptyText`、`UserName`、`AgentName`。
- `NewChatList(c, messages, state, o)`：配置 `state.Key`（消息 ID）、`state.FollowEnd`、`state.FollowAnimation`（`ui.Smooth`）、`state.EndReached`，返回 `ui.List` 包一层主题容器；`Justify(End)` 让少量消息贴在底部。`state` 为 nil 时用元素的 `ui.Local` 保存，跨帧不丢位置。`NewChatListStatic` 是固定数据的极简入口。
- 行渲染：用户消息靠右、助手消息靠左、系统消息居中弱化；正文按 `Stream` 是否为 nil 选流式或完稿 Markdown；`Pending` 显示由框架 `Loop` 驱动的 Spin（无计时器）。
- `EmptyText`：空会话显示一行提示。`ScrollToEnd()`、`AtEnd()`、`HoldPlace()` 透传给 `ListState`。
- `OnEndReached` 是 mygo 的 `ListState.EndReached`，在**接近列表末尾**时触发（信息流加载更多），不是顶部加载历史。
- 颜色、间距、圆角、字号、动效时长每帧从 `c.Theme()` 与根包 tokens 取；`StateTransition()` 让气泡换色渐变。无包级颜色缓存、无 goroutine。

## 4. 项目结构

- `ai/` 子包，与 `markdown/`、`chart/` 同层；dot-import 根包取词表，具名导入 `display/`、`markdown/`；根包不导入它。
- 只依赖 mygo 的 `ui` 与 `ui/markdown` 公开 API，不碰 internal，不带资产。

## 5. API 契约

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

## 6. 状态

### 已知限制

- `state` 为 nil 时组件把 `ListState` 存进元素的 `ui.Local`，跨帧保持滚动位置与跟随状态；app 只是读不到 `Selected` 等字段，需要时传自己的指针。
- `ChatMessage.ID` 决定稳定行键；留空退化为行号，前面插入历史消息时滚动位置可能跳。固定数据不受影响，接口数据请自己分配唯一 ID。
- 顶部加载历史消息：mygo 只有「接近末尾」的 `EndReached`，没有顶部回调。app 可在构建列表后读 `ListState.Visible()` 判断首行是否为 0，再配合一个 `ui.Local` 标志防重复来触发加载。
- `Stream` 返回的必须是**累积**源文本（只允许追加），这是上游 `markdown.Stream` 的约定；回复结束后要写回 `Text` 并把 `Stream` 置 nil。

- [x] 调研与规格书（按 mygo v0.10.0 `ListState` 实际 API 对齐）
- [x] 实现（`chatlist.go`）
- [x] 测试（`chatlist_test.go`）
- [x] showcase「ChatList」页
