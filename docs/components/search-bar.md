# SearchBar 规格书

搜索框：放大镜图标 + 输入框 + 清除，用于筛选与搜索。

## API 契约

```go
func NewSearchBar(c *ui.Context, query *string) *SearchBar
func (s *SearchBar) Placeholder(v string) *SearchBar
func (s *SearchBar) Width(w float32) *SearchBar
```
