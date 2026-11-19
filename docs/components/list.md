# List 规格书

> 状态：**已实现**。把 mygo 的虚拟列表放进一张带圆角、描边和底色的容器里；滚动、选择、键盘与无障碍仍归 mygo。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `ui.List(c, state, n, row)` | 虚拟列表：只建可见行；`ListState` 管位置、选择、排序键与滚动 |
| shadcn/ui `Card` | 容器是主题底色、一像素描边、圆角与内边距，随明暗主题换色 |
| `display.NewTable` | 直接把容器样式涂在 mygo 底座元素上，避免包一层后丢失高度约束 |

## 2. 能否实现

- `ui.List` 返回 `ui.Element`，其上是完整的 `Element` 样式链，可以直接给滚动容器加 `Background`/`Border`/`Radius`/`Padding`。
- 选择态的悬停与选中底色由 mygo 的列表行自己画，本组件不覆盖；只在容器层加外观。

## 3. 实现到何种程度

- `NewList(c, state *ui.ListState, n int, row func(i int)) ui.Element`：样式取自当帧的 `c.Theme()`，圆角用共享的 `RadiusLg`，描边与底色用 `Border`/`Background`，行与容器的变化走 `StateTransition()`。
- 不做新状态：指针 `state` 原样交给 `ui.List`，调用方仍可读 `Selected`、`ScrollTo`，`ListState.Key` 继续给行一个稳定的键。
- 与 `NewSortableList` 并存：后者在每行加拖拽手柄，本组件只加容器外观，互不调用、互不改名。

## 4. 项目结构

- `display/` 子包，与 `sortable_list.go`、`table.go` 同层。

## 5. API 契约

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

## 6. 状态

- [x] 规格书 / 实现 / 测试 / showcase
