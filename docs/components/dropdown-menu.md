# DropdownMenu 规格书

下拉菜单：点击触发项后在浮层里显示一组操作。

## API 契约

```go
type MenuItem struct {
    Label string
    Icon *ui.SVG
    Disabled, Separator bool
    OnClick func()
}
func NewDropdownMenu(c *ui.Context, open *bool, trigger func() ui.Element, items []MenuItem) ui.Element
```
