# List 规格书

把 mygo 的虚拟列表放进一张带圆角、描边和底色的容器里；滚动、选择、键盘与无障碍仍归 mygo。

## API 契约

```go
func NewList(c *ui.Context, state *ui.ListState, n int, row func(i int)) ui.Element
```

| 属性 | 类型 | 作用 |
|---|---|---|
| `c` | `*ui.Context` | 构建上下文；颜色每帧从 `c.Theme()` 取 |
| `state` | `*ui.ListState` | 列表状态指针，原样交给 `ui.List`；可为 nil |
| `n` | `int` | 行数，交给 `ui.List` |
| `row` | `func(i int)` | 画第 i 行，交给 `ui.List` |
| 返回值 | `ui.Element` | 带容器外观的滚动列表 |
