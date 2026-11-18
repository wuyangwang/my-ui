# Switch 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`switch.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Switch](https://ui.shadcn.com/docs/components/switch) | 36×20 的胶囊 + 16 的旋钮；开时强调色底 + 白旋钮；`checked`/`disabled` | 由库画 ✓（`SwitchBase` + 自绘） |
| mygo 的 `ui.SwitchBase(c, on)` | 无外观的基座：点击切换、`Changed()` 报告、`role=switch` ✓ | 基座 ✓，外观我们给 ✓ |
| mygo 的 `Element.Animate(key, target, d)` | 一个值随时间走到目标值 ✓（框架自己的开关用它走旋钮 ✓） | 旋钮走位 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 开关 | `ui.SwitchBase(c, *bool)` ✓ | ✅ |
| 旋钮走位 | `Element.Animate` ✓ + `Element.Draw` ✓ | ✅ |
| 外观 | 自己画 ✓ | ✅ |
| 禁用 | `Element.Disabled` ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `ctrlpkg.NewSwitch(c, &on)`：36×20 的胶囊 + 16 的旋钮 ✓；开时强调色底 ✓；旋钮 140ms 走位 ✓（`Animate` ✓，随桌面的 reduce motion 由框架兜底 ✓）。
- `OnChanged(func(bool))` ✓；`Label(string)`（无障碍名 ✓，开关本身没有文字 ✓）；`Disabled(bool)` ✓。

**已知限制**
- 开关没有文字（shadcn 也没有 ✓），名字由 app 给 ✓。
- 没有尺寸档 ✗（开关的大小固定 ✓，框架的也是 ✓）。

## 4. 项目结构是否合理

- 位置：根包 L3 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
func NewSwitch(c *ui.Context, on *bool) *Switch
func (s *Switch) OnChanged(fn func(on bool)) *Switch
func (s *Switch) Label(label string) *Switch
func (s *Switch) Disabled(disabled bool) *Switch
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`switch.go`）
- [x] 测试（点击切换 ✓、事件 ✓、禁用 ✓）
