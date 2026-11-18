# Breadcrumb 规格书

> 状态：**已实现**。面包屑：显示当前页面在层级中的位置，可点回上级。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Breadcrumb | 文字项 + `ChevronRight` 分隔，当前项为普通文字，其余可点 |
| Element Plus Breadcrumb | `el-breadcrumb-item` + `separator` |

## 2. 能否实现
- `ui.Row` + `ctrlpkg.NewButton` + `icons.ChevronRight` ✓；当前页指针 `*int` 由 app 持有（绑定约定）。

## 3. 实现到何种程度
- `NewBreadcrumb(c, items []string, chosen *int)`；`.Size(sz)`；当前项为强调色文字，其他项弱墨且可点，点击写回 `*chosen`。
- 后置：自定义分隔符、下拉式面包屑。

## 4. 项目结构
- `display/` 子包（导航展示件）。

## 5. API 契约
```go
func NewBreadcrumb(c *ui.Context, items []string, chosen *int) *Breadcrumb
func (b *Breadcrumb) Size(sz Size) *Breadcrumb
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
