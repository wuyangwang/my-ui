# Slider 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
func NewSlider(c *ui.Context, value *float64, lo, hi float64) *Slider
func (s *Slider) Step(step float64) *Slider
func (s *Slider) OnChanged(fn func(value float64)) *Slider
func (s *Slider) Label(label string) *Slider
func (s *Slider) Disabled(disabled bool) *Slider
```

用法：

```go
ctrlpkg.NewSlider(c, &app.volume, 0, 100).Step(10).Label("音量")
```

## 状态与限制

**限制**

- 没有竖向的包装（框架的 `Element.Vertical()`）。
- 区间滑杆（两个旋钮）没有（框架没有）。
- 没有刻度与数值标签（shadcn 也没有）。
