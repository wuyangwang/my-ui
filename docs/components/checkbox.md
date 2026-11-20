# Checkbox 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
func NewCheckbox(c *ui.Context, checked *bool, label string) *Checkbox
func (cb *Checkbox) Size(s Size) *Checkbox
func (cb *Checkbox) Xs() *Checkbox
func (cb *Checkbox) Sm() *Checkbox
func (cb *Checkbox) Md() *Checkbox
func (cb *Checkbox) Lg() *Checkbox
func (cb *Checkbox) OnChanged(fn func(checked bool)) *Checkbox
func (cb *Checkbox) Disabled(disabled bool) *Checkbox
```

## 状态与限制

**限制**

- 半选态与「全选」组没有包装（框架有 `ui.CheckboxState`/`ui.CheckboxGroup`，直接用即可）。
- 没有 `Required` 等校验钩子（用 `Field` 的校验层）。
