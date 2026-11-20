# ContextMenu 规格书

右键菜单：给目标元素附加一组操作。

## API 契约

```go
func NewContextMenu(c *ui.Context, open *bool, target ui.Element, items []MenuItem) ui.Element
```
