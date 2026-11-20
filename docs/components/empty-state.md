# EmptyState 规格书

空状态：列表或页面没有内容时的说明与引导。

## API 契约

```go
type EmptyStateOptions struct {
    Icon *ui.SVG
    Title, Description string
    Action func()
}
func NewEmptyState(c *ui.Context, o EmptyStateOptions) ui.Element
```
