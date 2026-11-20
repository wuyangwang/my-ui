# Radio 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

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

## 状态与限制

**限制**

- 没有 `RadioGroup` 的包装（框架的 `ui.RadioGroup` 直接用；共享值靠同一个指针）。
- 方向键的组内移动没有测试覆盖。
