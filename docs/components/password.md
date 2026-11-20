# PasswordInput 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
// NewPasswordInput creates a password field editing *value with the
// native preset.
func NewPasswordInput(c *ui.Context, value *string) *PasswordInput

// EyeIcons sets the icons of the button: the one it shows while the
// text shows, and the one it shows while the text is hidden.
func (p *PasswordInput) EyeIcons(shown, hidden *ui.SVG) *PasswordInput

// EyeColor sets the color of the button's icon.
func (p *PasswordInput) EyeColor(c ui.Color) *PasswordInput
```

继承自字段基座：`Variant`/`Size`/`Shape`/`Width`/`Placeholder`/`Label`/`Clearable`/`OnChanged`/`OnBlurred`/`OnCleared`/`Control` 等。

用法：

```go
h := bindpkg.Bind(c)
h.PasswordInput(&app.password).Outline().Label("密码").Placeholder("输入密码")
```

## 状态与限制

**限制**

**已知限制**：无——显隐切换保留值与编辑状态。
