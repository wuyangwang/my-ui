# Radio 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`radio.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui RadioGroup](https://ui.shadcn.com/docs/components/radio-group) | 圆环 + 选中时中间的圆点；一组共用一个值；`RadioGroup`/`RadioItem`/`RadioItemIndicator` | 圆环由库画 ✓（`RadioBase` + 自绘） |
| mygo 的 `ui.RadioBase[T](c, selected *T, value T)` | 无外观的基座：点击选中、`Changed()` 报告、`role=radio` ✓；`ui.RadioGroup(c, fn)` 给一组 ✓ | 基座 ✓，外观我们给 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 单选 | `ui.RadioBase[T]` ✓（泛型，值可以是任何可比类型 ✓） | ✅ |
| 外观自绘 | 基座返回行，圆与标签由 app 建 ✓ | ✅ |
| 组 | `ui.RadioGroup` ✓ | ✅（共享值由 app 的指针天然做到 ✓） |
| 键盘 | 方向键在组内移动 ✓（框架的基座 ✓） | ✅ |

## 3. 实现到何种程度

**已实现**
- `ctrlpkg.NewRadio(c, &selected, value, "标签")`：圆环（14/16/18/22 DIP）+ 描边；选中时强调色环 + 中间的点 ✓；标签 ✓。
- `Size` 四档 ✓；`OnChanged(func(T))` ✓；`Disabled(bool)` ✓。

**已知限制**
- 没有 `RadioGroup` 的包装 ✗（框架的 `ui.RadioGroup` ✓ 直接用 ✓；共享值靠同一个指针 ✓）。
- 方向键的组内移动没有测试覆盖 ✗。

## 4. 项目结构是否合理

- 位置：根包 L3 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
func NewRadio[T comparable](c *ui.Context, selected *T, value T, label string) *Radio[T]
func (r *Radio[T]) Size(s Size) *Radio[T]
func (r *Radio[T]) OnChanged(fn func(value T)) *Radio[T]
func (r *Radio[T]) Disabled(disabled bool) *Radio[T]
```

用法：

```go
ctrlpkg.NewRadio(c, &app.size, "large", "大")
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`radio.go`）
- [x] 测试（选中 ✓、换选 ✓、事件 ✓）
