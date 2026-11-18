# PasswordInput 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`password.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [Element Plus show-password](https://element-plus.org/en-US/component/input.html) | `show-password`：**只有有值且未禁用时**才出现眼睛按钮；图标**随状态切换**（Eye / EyeOff）；`password-icon` 插槽；暴露 `passwordVisible` | 眼睛按钮在有值时出现 ✓；图标随状态切换 ✓（`icons.Eye` / `icons.EyeOff`） |
| [Ant Design Input.Password](https://ant.design/components/input-cn) | `visibilityToggle`（默认开）；`iconRender: (visible) => node` | 同上 ✓ |
| [Lobe MyGo InputPassword](https://mygo.lobehub.com/skills/components/input-password) | 独立组件，继承 Input 的全部属性，末端一个眼睛按钮切换显隐 | 同上 ✓——独立组件 ✓（Q5/Q10 定的） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 隐藏所输入的内容 | `Element.Password()` ✓（把编辑器的 password 标志置真） | ✅ |
| **再显示出来** | `Password()` 是**按帧标志**：调用的帧隐藏，不调用的帧恢复明文（`ui/textinput.go:290-299`）✓ | ✅ 同一个编辑器、同一个 key，按 `*shown` 决定是否叠加 `Password()`；光标、选区、撤销栈都保留 ✓ |
| 眼睛按钮 | `ui.ButtonBase` + `KeepFocus()` ✓ + `icons.Eye/EyeOff` ✓ | ✅ |
| 有值才出现 | 读 `*value` ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `NewPasswordInput(c, &value)`：一个隐藏所输入内容的字段，末端一个眼睛按钮切换明暗。
- 切换的实现：同一个 key、同一个编辑器，按帧叠加或不叠加 `Password()` ✓；值由 app 持有 ✓，编辑状态（光标、选区、撤销）也保留 ✓。
- 明暗状态存在字段自己的槽位里（`ui.Local`）✓，跨帧保持 ✓；`Visible()` 查询当前是否明文 ✓。
- **切换不改焦点**：眼睛按钮用 `KeepFocus()`，打字接着上次的位置 ✓。
- **眼睛按钮的图标与颜色可配**：`EyeIcons(shown, hidden *ui.SVG)`（显示时/隐藏时各一个）与 `EyeColor(ui.Color)` ✓；两者在**绘制时**读取（`Element.Draw` + `Painter.Icon`）✓，所以链式调用在任何位置都生效 ✓。
- 继承字段基座的能力：`Variant`/`Size`/`Shape`/`Width`/`Placeholder`/`Label`/`Clearable`/`OnChanged` 等 ✓。

**已知限制**：无——显隐切换保留值与编辑状态 ✓。

**后置**：无（组件已可上线 ✓）——后续若加需求，走新增而非后置 ✓。

## 4. 项目结构是否合理

- 位置：根包 L3（变体化控件），与 `Input` 同层；实现放在 `password.go`，共用 `field[T]` 基座 ✓（只有「里面是什么」不同 ✓）。

## 5. API 契约

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

继承自字段基座：`Variant`/`Size`/`Shape`/`Width`/`Placeholder`/`Label`/`Clearable`/`OnChanged`/`OnBlurred`/`OnCleared`/`Control` 等 ✓。

用法：

```go
h := bindpkg.Bind(c)
h.PasswordInput(&app.password).Outline().Label("密码").Placeholder("输入密码")
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`password.go`）
- [x] 测试（`form_test.go`：隐藏时无障碍树不显值、点眼睛后显值）
- [x] showcase「表单」页
- [ ] 切换后还焦点、图标可配（后置）
