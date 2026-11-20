# Outline 规格书

泛型封装 mygo 的 `ui.Outline[K]`，套 shadcn 的树行外观；开合、键盘与无障碍由 mygo 管。

## API 契约

```go
func NewOutline[K comparable](
    c *ui.Context,
    state *ui.OutlineState[K],
    children func(K) []K,
    title func(K) string,
) ui.Element
```

| 属性 | 类型 | 作用 |
|---|---|---|
| `K` | `comparable` | 节点键类型；与 `ui.OutlineState[K]` 一致 |
| `state` | `*ui.OutlineState[K]` | 树状态指针：`List.Selected` / `Open` / `Indent` / `Generation` 由 app 持有 |
| `children` | `func(K) []K` | 给定节点返回其子节点；`K` 零值的返回值作为根节点；nil 表示叶子 |
| `title` | `func(K) string` | 行的文字；返回空串时显示节点的 `fmt.Sprint` |
| 返回值 | `ui.Element` | 带容器外观的树视图 |
