# Menubar 规格书

应用内菜单栏：横排多个下拉菜单，每个共用 DropdownMenu 的外观。

## API 契约

```go
type MenubarMenu struct {
    Label string
    Open *bool
    Items []MenuItem
}
func NewMenubar(c *ui.Context, menus []MenubarMenu) *Menubar
```
