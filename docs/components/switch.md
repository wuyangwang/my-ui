# Switch 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
func NewSwitch(c *ui.Context, on *bool) *Switch
func (s *Switch) OnChanged(fn func(on bool)) *Switch
func (s *Switch) Label(label string) *Switch
func (s *Switch) Disabled(disabled bool) *Switch
```

## 状态与限制

**限制**

- 开关没有文字（shadcn 也没有），名字由 app 给。
- 没有尺寸档（开关的大小固定，框架的也是）。
