# SortableList 规格书

> 状态：**已实现**。可排序列表：列表项带拖拽手柄，拖动即改变顺序。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `ListState.Reorder` | 平台原生拖拽排序回调 `func(rows []int, to int)` |
| shadcn/ui Sortable | 每项有可见手柄，拖拽时显示落点 |

## 2. 能否实现
- `ui.List` + `ListState.Reorder` ✓；顺序数据由 app 持有，回调里重排切片。

## 3. 实现到何种程度
- `NewSortableList(c, state *ui.ListState, n int, item func(i int)) *SortableList`；
- 每项渲染一个拖拽手柄图标（`icons.GripVertical`，`IconSizeSm`）和 app 的行内容；手柄之前留 `Space(1)` 间距。
- 后置：手柄可见性选项、禁用某些行拖动。

## 4. 项目结构
- `display/` 子包。

## 5. API 契约
```go
func NewSortableList(c *ui.Context, state *ui.ListState, n int, item func(i int)) *SortableList
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
