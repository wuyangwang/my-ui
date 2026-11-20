# Breadcrumb 规格书

面包屑：显示当前页面在层级中的位置，可点回上级。

## API 契约

```go
func NewBreadcrumb(c *ui.Context, items []string, chosen *int) *Breadcrumb
func (b *Breadcrumb) Size(sz Size) *Breadcrumb
```
