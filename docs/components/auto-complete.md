# AutoComplete 规格书

自动补全：输入框 + 随输入过滤的建议列表。

## API 契约

```go
func NewAutoComplete(c *ui.Context, value *string, suggestions []string) *AutoComplete
func (a *AutoComplete) Variant(v Variant) *AutoComplete
func (a *AutoComplete) Size(sz Size) *AutoComplete
func (a *AutoComplete) Width(w float32) *AutoComplete
```
