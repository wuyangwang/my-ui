# Input 规格书

> 状态：**首轮已实现**（变体/尺寸/清除/事件/前后缀/错误态）。后置项见 §3。实现见 `input.go`，测试见 `input_test.go`，示例见 `examples/showcase` 的「Input」页。

## 1. 参考实现

| 来源 | 可借鉴之处 | mygo 下的差异 |
|---|---|---|
| [shadcn/ui Input](https://ui.shadcn.com/docs/components/input) + [Input Group](https://ui.shadcn.com/docs/components/input-group) | 变体/尺寸/状态轴；前缀后缀与清除按钮作为**组合件**（Input Group）而非输入框本身；清除按钮在有值且可编辑时出现 | React/CSS 变量；组合件拆成独立组件 |
| [Element Plus el-input](https://element-plus.org/zh-CN/component/input) | `clearable`、`prefix-icon`/`suffix-icon`、`show-password`、`formatter`/`parser`；事件 `focus`/`blur`/`change`/`input`/`clear` | Vue 插槽与 v-model 双绑定；事件是 DOM 事件对象 |
| [Ant Design Input](https://ant.design/components/input-cn) | `allowClear`、`prefix`/`suffix`、`status="error"`、受控/非受控两种模式 | React 受控组件模型 |
| **mygo 官方 `ui.SearchField`** | 框架自己的“带放大镜与清除按钮的输入框”：清除按钮与输入框是**兄弟节点**，画在一个 field 里；清除后 `*query = ""` 并 `in.Focus()` 保持焦点 | 这是本库最直接的参照：结构照它，样式按我们的预设 |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 文本输入与绑定 | `ui.TextInput(c, *string)`；输入框与 `*value` 共享字符串，app 改 `*value` 即同步编辑器（`ui/textinput.go:52-57`） | ✅ 清空只需 `*value = ""` |
| 值变化事件 | `Element.Changed()`（本帧值是否变化） | ✅ `OnChanged` |
| 失焦事件 | `Element.FocusLeft()`（焦点自上帧起离开本元素） | ✅ `OnBlurred` |
| 获得焦点事件 | 无「刚获得焦点」的公开 API，只有 `Focused()` 状态 | ⚠️ 只提供状态，不提供事件（记录为限制） |
| 回车 | `Element.OnTextKey(func(ui.TextKey) bool)`，`TextKey{Key, Mods}` | ✅ `OnSubmit` |
| 组合（前后缀/清除按钮） | `Element.Children(fn)` **立即执行**；只有 fn 内新建的元素才成为子节点 → 用**固定槽位**（构造时建好空槽，之后用 `Children` 填充）实现 | ✅ 槽位方案可行 |
| 清除按钮不抢焦点 | `Element.KeepFocus()`（按下时焦点留在输入框）、`Element.TabStop(false)`（不进 Tab 序）、`Element.Label("清除")` | ✅ |
| 错误态 | `Element.Invalid(bool)`（无障碍状态）、`Element.ErrorMessage(Element)`（错误文案关系）；边框走主题 `Danger` | ✅ |
| 只读/禁用/密码 | `Element.ReadOnly(bool)`、`Element.Disabled(bool)`、`Element.Password()` | ✅（密码显隐切换需 app 状态，后置） |
| 无障碍 | `Element.Label/Description/Invalid/ErrorMessage` | ✅；**不做**依赖 `Element.Access(func(*platform.AccessNode))` 的定制（internal 泄漏，见 [../research.md](../research.md) §五） |

## 3. 实现到何种程度

**首轮（已实现）**
- 变体：`Native`（零值，保留 ui 内置观感）/`Secondary`（别名）/`Primary`/`Outline`/`Ghost`/`Destructive`/`Link`，作用在**外框**（边框、底色、圆角），内层输入框自身透明无边框。
- 尺寸：`SizeSm/Md/Lg` → 外框内衬与字号（走主题 `Space`/`Rem` 刻度）。
- 组合：`Prefix(func())`、`Suffix(func())`、`Clearable()`（清除按钮用 `icons.X()`，有值且可编辑时出现，点击清空并保持焦点）。
- 事件：`OnChanged(func(string))`、`OnBlurred(func())`、`OnCleared(func())`、`OnSubmit(func())`；查询式 `Changed()`/`Focused()`/`FocusLeft()`/`Value()`。
- 错误态：`Invalid(bool)`（边框转 danger + 无障碍状态）、`ErrorText(string)`（框下错误文案，接入 `ErrorMessage`）、`ErrorMessage(func())`（自定义错误内容）。
- 状态：`Disabled(bool)`、`ReadOnly(bool)`、`Password()`、`Placeholder(string)`、`Label(string)`。

**后置**
- 密码显隐切换（需 app 持有「显示中」状态）。
- 字数限制与计数、formatter/parser（格式化后再绑定）。
- `InputOTP`、`DatePicker`、`InputNumber`、`SearchBar`（独立组件，走同一套预设）。

## 4. 项目结构是否合理

- 位置：根包 L3（变体化控件），与 `Button`/`Badge`/`Icon` 同层。
- 依赖：`myui` → `mygo/ui` + `myui/icons`（只用清除按钮那一个图标，实测对二进制影响约等于零）。
- 入口：既可 `myui.NewInput(c, &v)`，也可用每帧句柄 `myui.Bind(c).Input(&v)`（见 [../../docs/components/handle.md](../../docs/components/handle.md) 备注）。
- 组合件不另立 `InputGroup`：清除与前后缀都是「输入框的一部分」，同源同预设；shadcn 拆开是因为 CSS 组合方便，Go 里槽位方案已经解决。

## 5. API 契约

```go
// NewInput creates a text input editing *value with the native preset.
func NewInput(c *ui.Context, value *string) *Input

func (in *Input) Variant(v Variant) *Input   // Native/Secondary/Primary/Outline/Ghost/Destructive/Link
func (in *Input) Size(s Size) *Input         // SizeSm/Md/Lg
func (in *Input) Sm() *Input
func (in *Input) Md() *Input
func (in *Input) Lg() *Input

func (in *Input) Placeholder(s string) *Input
func (in *Input) Label(s string) *Input
func (in *Input) Invalid(on bool) *Input
func (in *Input) ErrorText(s string) *Input
func (in *Input) ErrorMessage(fn func()) *Input
func (in *Input) ReadOnly(on bool) *Input
func (in *Input) Disabled(on bool) *Input
func (in *Input) Password() *Input

func (in *Input) Prefix(fn func()) *Input
func (in *Input) Suffix(fn func()) *Input
func (in *Input) Clearable() *Input

func (in *Input) OnChanged(fn func(string)) *Input
func (in *Input) OnBlurred(fn func()) *Input
func (in *Input) OnCleared(fn func()) *Input
func (in *Input) OnSubmit(fn func()) *Input

func (in *Input) Changed() bool
func (in *Input) Focused() bool
func (in *Input) FocusLeft() bool
func (in *Input) Value() string
```

用法：

```go
h := myui.Bind(c)
h.Input(&app.email).Outline().Placeholder("邮箱").
	Invalid(app.emailError != "").ErrorText(app.emailError).
	OnChanged(func(v string) { app.email = v })
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 复合结构（外框 + 槽位 + 内层输入）
- [x] 清除按钮（图标、`KeepFocus`、`TabStop(false)`、无障碍名）
- [x] 事件（变化/失焦/清除/回车）与查询式接口
- [x] 前后缀与错误态
- [x] 测试（`input_test.go`：尺寸、禁用/只读、暗色、清除、事件、错误文案、前后缀）
- [x] showcase「Input」页
- [ ] 密码显隐、字数统计、formatter/parser（后置）
