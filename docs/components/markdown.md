# Markdown 规格书

> 位置：组件在 [`markdown/`](../../markdown) 子包；主题/tokens 仍在根包 `myui`。
>
。

## API 契约

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
