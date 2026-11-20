# ai（二级包）

包含 `ChatList`、`ModelPicker`、`Reasoning`；文档见 `docs/components/`，示例见 `examples/showcase/` 的「AI」分组。

AI 应用常见组件。mygo 独有的底层能力：`ui/markdown` 专为 LLM 逐 token 流式回复构建（增量解析、半成品语法按终态显示、文本淡入），`ui.List` 为锚点式虚拟化列表（行高动态、行可增删、`FollowEnd` 贴底跟随）。本库把两者的组合包装成可直接用的应用组件。

## 组件

| 组件 | 作用 |
|---|---|
| `ChatList` | 虚拟化消息气泡：用户靠右、助手靠左、系统居中；流式或完稿 Markdown、头像、typing 指示、贴底跟随。`ListState.Key` 用消息 ID 稳定行 |
| `ModelPicker` | 可搜索的模型下拉：按厂商分组、厂商图标、上下文窗/速度/价格徽标、选中态跨帧可读 |
| `Reasoning` | 思考强度（低/中/高，档位可自定义）与 temperature/top_p 滑杆，附精确/均衡/创意三个预设 |

每个组件都有 `bind.Bind(c)` 对应的句柄方法。`ChatList` 同时支持本地固定数据（`NewChatListStatic(c, []ai.ChatMessage{...})`，不必声明变量与 `ListState`）和接口返回数据（app 持有 slice，每帧传当前值，append 后自动生效）。

## 约定

- 无 `go:embed` 资产；仅依赖 `myui` 根包与 mygo 公开 API，不改 mygo 行为。
- 消息、选中模型、强度与参数都由 app 持有指针；本包不持有状态。
- 文案走根包 i18n 文案表（新增 `effort*`、`temperature`、`topP`、`preset*`、`searchModels`、`noModels`、`user`、`assistant` 等键）。
- 颜色、间距、字号与动效时长每帧从 `c.Theme()` 与根包 tokens 取，无包级颜色状态。
