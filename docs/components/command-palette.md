# Command Palette 规格书

命令面板：一个居中对话框，输入即过滤命令，回车/点击执行。

## API 契约

```go
type Command struct {
    Label string
    Icon *ui.SVG
    Shortcut string
    OnSelect func()
}
func NewCommandPalette(c *ui.Context, open *bool, query *string, commands []Command) ui.Element
```
