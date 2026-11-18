# Table 规格书

> 状态：**已实现**。表格：列头 + 行单元，带排序与选择状态（由 mygo `ListState` 提供）。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `Table` | `TableColumn`/`ListState`、列宽/排序/选择由框架提供 |
| shadcn/ui Table | 圆角边框容器、弱墨列头、行分隔线 |

## 2. 能否实现
- `ui.Table(c, state, columns, n, cell)` ✓；库只加容器外观、字号与默认宽度。

## 3. 实现到何种程度
- `NewTable(c, state, columns, n, cell) *Table`；`.Variant(v)`、`.Size(sz)`、`.Width(w)`（默认占满）。
- 后置：自定义空态、斑马纹预设。

## 4. 项目结构
- `display/` 子包。

## 5. API 契约
```go
func NewTable(c *ui.Context, state *ui.ListState, columns []ui.TableColumn, n int, cell func(row, col int)) *Table
func (t *Table) Variant(v Variant) *Table
func (t *Table) Size(sz Size) *Table
func (t *Table) Width(w float32) *Table
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
