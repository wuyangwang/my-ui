# PasswordInput 规格书

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
| **再显示出来** | `Password()` **只会置真**（`ui/textinput.go:294-297`），没有公开的取消隐藏 API ✗ | ⚠️ 只能换一个 key 重建输入框：键换了 → 状态（id）换了 → 编辑器是新的 → 不再隐藏 ✓；值由 app 的 `*string` 保住 ✓，编辑状态（光标、撤销）不保 ✗ |
| 眼睛按钮 | `ui.ButtonBase` + `KeepFocus()` ✓ + `icons.Eye/EyeOff` ✓ | ✅ |
| 有值才出现 | 读 `*value` ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `NewPasswordInput(c, &value)`：一个隐藏所输入内容的字段，末端一个眼睛按钮切换明暗。
- 切换的实现：两个 key（`password-hidden` / `password-shown`）——键换了，输入框是新的，隐藏标志也就没了 ✓；值由 app 持有，不丢 ✓。
- 明暗状态存在字段自己的槽位里（`ui.Local`）✓，跨帧保持 ✓。
- 继承字段基座的能力：`Variant`/`Size`/`Shape`/`Width`/`Placeholder`/`Label`/`Clearable`/`OnChanged` 等 ✓。

**已知代价**：切换时编辑状态（光标位置、撤销栈）重置 ✗（值不丢 ✓）。

**后置**
- 切换后把焦点还给输入框。
- 眼睛按钮的图标与状态色可配。

## 4. 项目结构是否合理

- 位置：根包 L3（变体化控件），与 `Input` 同层；实现放在 `password.go`，共用 `field[T]` 基座 ✓（只有「里面是什么」不同 ✓）。

## 5. API 契约

```go
// NewPasswordInput creates a password field editing *value with the
// native preset.
func NewPasswordInput(c *ui.Context, value *string) *PasswordInput
```

继承自字段基座：`Variant`/`Size`/`Shape`/`Width`/`Placeholder`/`Label`/`Clearable`/`OnChanged`/`OnBlurred`/`OnCleared`/`Control` 等 ✓。

用法：

```go
h := myui.Bind(c)
h.PasswordInput(&app.password).Outline().Label("密码").Placeholder("输入密码")
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`password.go`）
- [x] 测试（`form_test.go`：隐藏时无障碍树不显值、点眼睛后显值）
- [x] showcase「表单」页
- [ ] 切换后还焦点、图标可配（后置）
