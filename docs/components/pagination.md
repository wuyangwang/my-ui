# Pagination 规格书

> 状态：**已实现**。分页：在当前页指针上翻页，页码可点，超出范围显示省略号。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Pagination | 上一页/下一页 + 页码按钮，当前页为强填充 |
| Element Plus Pagination | 页码窗口与省略号 |

## 2. 能否实现
- 遵循绑定约定：当前页 `*int` 由 app 持有，点击写回；按钮用 `ctrlpkg.NewButton`，左右箭头用 `ctrlpkg.NewIconAction`（图标本身即控件）✓
- 页码窗口：总数 ≤ 7 全显；否则首末页 + 当前页周围两页 + 省略号。

## 3. 实现到何种程度
- `NewPagination(c, current *int, total int) *Pagination`；`.OnChanged(fn)`、`.Size(sz)`；`nil` 指针视为只读。
- 上一页/下一页按钮在边界禁用；无障碍名走 i18n（新增 `previous`/`next` 文案）。
- 后置：每页条数选择、跳转输入框。

## 4. 项目结构
- `display/` 子包（导航件，沿用展示分组）。

## 5. API 契约
```go
func NewPagination(c *ui.Context, current *int, total int) *Pagination
func (p *Pagination) OnChanged(fn func(page int)) *Pagination
func (p *Pagination) Size(sz Size) *Pagination
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
