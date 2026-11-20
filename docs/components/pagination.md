# Pagination 规格书

分页：在当前页指针上翻页，页码可点，超出范围显示省略号。

## API 契约

```go
func NewPagination(c *ui.Context, current *int, total int) *Pagination
func (p *Pagination) OnChanged(fn func(page int)) *Pagination
func (p *Pagination) Size(sz Size) *Pagination  // 下一帧生效，不重建按钮
```
