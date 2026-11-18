# Slider 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`slider.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Slider](https://ui.shadcn.com/docs/components/slider) | 2px 的轨道 + 16 的旋钮；轨道到旋钮的一段是强调色；`min`/`max`/`step` | 由库画 ✓（`SliderBase` + 自绘） |
| mygo 的 `ui.SliderBase(c, value, lo, hi)` | 无外观的基座：拖动取值、方向键与 Home/End 移动、`Changed()` 报告、`role=slider` ✓；`Element.Step(step)` ✓ | 基座 ✓，外观我们给 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 取值 | `ui.SliderBase` ✓ | ✅ |
| 步进 | `Element.Step(float64)` ✓ | ✅ |
| 外观 | `Element.Draw` ✓ 自己画 ✓ | ✅ |
| 竖向 | `Element.Vertical()` ✓（框架的 ✓） | ✅（未包装 ✗） |

## 3. 实现到何种程度

**已实现**
- `ctrlpkg.NewSlider(c, &value, lo, hi)`：轨道（2 DIP）+ 旋钮（16 DIP，带阴影 ✓）；轨道到旋钮的一段是强调色 ✓。
- `Step(float64)` ✓；`OnChanged(func(float64))` ✓；`Label(string)` ✓；`Disabled(bool)` ✓。

**已知限制**
- 没有竖向的包装 ✗（框架的 `Element.Vertical()` ✓）。
- 区间滑杆（两个旋钮）没有 ✗（框架没有 ✗）。
- 没有刻度与数值标签 ✗（shadcn 也没有 ✓）。

## 4. 项目结构是否合理

- 位置：根包 L3 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`slider.go`）
- [x] 测试（点击取值 ✓、方向键 ✓、禁用 ✓）
- [ ] 竖向、区间滑杆（后置 ✗）
