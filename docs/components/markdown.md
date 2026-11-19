# Markdown 规格书

> 位置：组件在 [`markdown/`](../../markdown) 子包；主题/tokens 仍在根包 `myui`。
>
> 状态：**已实现**（`markdown.go`），是 `ai/ChatList` 的前置件。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo `ui/markdown`（v0.10.0，`ui/markdown/view.go`、`theme.go`） | 专为 LLM 逐 token 回复而造：`View` 渲染完稿、`Stream` 增量解析并要求下一帧、半成品语法按最可能的终态显示、文本随框架动效淡入；`Options.Theme` 可注入外观，`DefaultTheme(t)` 从窗口主题派生，`Revealed(el)` 判断本帧是否已全部显示 |
| shadcn/ui 的排版观感 | 次要文字取 `TextMuted`、链接取主色、引用/规则取边框色、代码与表格取弱底 `Surface` |
| GitHub Alerts | 五种提示（Note/Tip/Important/Warning/Caution）按语义色着色 |

## 2. 能否实现

- 完稿与流式视图：`markdown.View(c, src, o)` / `markdown.Stream(c, src, o)` ✓（公开 API）。
- 主题映射：`markdown.Theme` 是公开结构体，`Options.Theme` 接受注入；`markdown.Color` 就是 `ui.Color` ✓。
- 流式数据源：上游只收 `string`，本包每帧调用 app 传入的 `stream func() string` 取当前源 ✓（状态在 app，本包不持有）。
- 显示完成判定：上游 `Revealed(el)` ✓，但元素句柄只在当帧有效，必须在同一帧构造后立刻调用。
- 图片加载、链接菜单、代码高亮等扩展：上游 `Options` 都有；本包的两个构造器只注入主题，未暴露这些选项（见「已知限制」）。

## 3. 实现到何种程度

**已实现**：

- `NewMarkdown(c, source)`：包 `markdown.View`，显示完整 Markdown（CommonMark + 上游默认扩展），文本可选中、可复制为文本/HTML/Markdown。
- `NewStream(c, stream)`：每帧读取 `stream()` 得到累积源文本，包 `markdown.Stream`；nil 函数显示空文档。app 在回复完成时改用 `NewMarkdown`，解析得以延续。
- `Theme(t)`：把 my-ui 的主题映射成 `markdown.Theme`——正文/次要/链接/强调取主题的 `Text`/`TextMuted`/`Accent`/`Text`，引用、规则、表格线取 `Border`，代码块与表头用 my-ui 的弱底 `Surface`，圆角取 `Radius(t, RadiusLg)`，`<mark>` 取 Warning 的半透明色，五种 Alert 依次取 Info/Success/Accent/Warning/Danger。主题字段为空时不覆盖，由上游默认补齐；`CodeColors` 留空，仍取 mygo 编辑器的语法配色。
- 每帧重建 `Options`：上游视图会保留传入的 `Options`，本包每次构造都新建，颜色与尺寸始终跟随 `c.Theme()`、`SetTheme` 与系统外观/强调色/高对比/文本缩放，下一帧生效。无包级缓存、无计时器、无 goroutine。
- `Revealed`：原样再导出，供 ChatList 判断何时显示时间与操作项。

**后置**：`Options` 的链接/图片/代码/Math 等扩展项在 ChatList 与 ModelPicker 需要时再以新组件或新构造器提供；本包不在这里开洞。

## 4. 项目结构是否合理

- 位置：`markdown/` 是独立子包，与 `ai/`、`chart/` 一样按需引入；根包不导入它，它 dot-import 根包取共享词表（`SemanticColor`、`Radius`、`StateTransition`），不与任何组件包互相引用。
- 依赖：仅 `ui`、`ui/markdown` 与根包 `myui`；不碰 mygo 的 internal 包，不带字体/图标资产。
- 状态：源文本由 app 持有，主题由 `c.Theme()` 取得，本包无状态。

## 5. API 契约

| API | 作用 |
|---|---|
| `func NewMarkdown(c *ui.Context, source string) ui.Element` | 显示完整 Markdown；每帧按当前主题构造视图 |
| `func NewStream(c *ui.Context, stream func() string) ui.Element` | 显示流式 Markdown；`stream` 每帧调用一次取累积源文本，nil 显示空文档 |
| `func Theme(t *ui.Theme) *markdown.Theme` | 把 my-ui 主题映射成上游 Markdown 主题；未覆盖字段留零，由上游默认补齐 |
| `var Revealed func(ui.Element) bool` | 再导出上游 `markdown.Revealed`；必须在构造视图的同一帧、紧接着调用 |

用法：

```go
md.NewMarkdown(c, "# 标题\n\n正文与 `代码`。")

var src string // app 持有，逐 token 追加
el := md.NewStream(c, func() string { return src })
if done {
    el = md.NewMarkdown(c, src)
}
if md.Revealed(el) { /* 显示时间与操作项 */ }
```

## 6. 状态

- [x] 调研与规格书（按 mygo v0.10.0 实际 API 对齐）
- [x] 实现（`markdown.go`：主题映射 + 两个构造器 + Revealed）
- [x] 测试（`markdown_test.go`：完稿文本、流式每帧读取、nil 流、明暗主题映射、Revealed）
- [x] showcase「Markdown」页（`page_markdown.go`，函数 `markdownPage`，由上层注册）
