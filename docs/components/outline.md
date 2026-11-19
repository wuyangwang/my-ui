# Outline 规格书

> 状态：**已实现**。泛型封装 mygo 的 `ui.Outline[K]`，套 shadcn 的树行外观；开合、键盘与无障碍由 mygo 管。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `ui.Outline[K](c, state, roots, children, row)` | 树视图：`OutlineState[K]` 持 `List` 与 `Open`；`children(item)` 返回 nil 表示叶子；`Generation` 决定何时重新遍历 |
| shadcn/ui `Tree` | 行是轻量的一行文字，缩进、悬停与选中底色跟列表一致 |
| `display.NewList` | 容器涂在底座元素上，圆角/描边/底色每帧取 `c.Theme()` |

## 2. 能否实现

- Go 1.27 的泛型顶层函数可以直接调用 `ui.Outline[K]`，不需要反射，也没有方法集的障碍。
- 请求签名为 `NewOutline(c, state, children, title)`，没有单独的 roots 参数。这里约定：**以 `K` 的零值作为虚拟根**，`children(zero)` 返回的切片就是根节点；`title` 只对真实节点调用。字符串键下 `""` 是自然的虚拟根。
- 限制：零值不可同时作为真实节点；深度树在 `children` 变化时由 app 递增 `state.Generation`。

## 3. 实现到何种程度

- `NewOutline[K comparable](c, state *ui.OutlineState[K], children func(K) []K, title func(K) string) ui.Element`。
- 外观：容器用主题底色、描边与 LG 圆角；每行文字用共享的 `TextSm`，`Title` 为空时退回 `fmt.Sprint(item)`；开合箭头沿用 mygo 的原生箭头，缩进用 `state.Indent`（零值即主题 `Space(4)`）。
- 行底色、悬停、选中与焦点环仍由 mygo 的列表行绘制，本组件不重复画。
- 无计时器、无 goroutine；所有颜色每帧从 `c.Theme()` 取，变化走 `StateTransition()`。

## 4. 项目结构

- `display/` 子包；与 `NewList`、`NewSortableList` 同层，互不改名。

## 5. API 契约

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

## 6. 状态

- [x] 规格书 / 实现 / 测试 / showcase
