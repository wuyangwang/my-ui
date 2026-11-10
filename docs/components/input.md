# Input 规格书

> 状态：**首轮已实现**（变体/尺寸/清除/事件/前后缀/错误态）。后置项见 §3。实现见 `input.go`，测试见 `input_test.go`，示例见 `examples/showcase` 的「Input」页。

## 1. 参考实现

| 来源 | 可借鉴之处 | mygo 下的差异 |
|---|---|---|
| [shadcn/ui Input](https://ui.shadcn.com/docs/components/input) + [Input Group](https://ui.shadcn.com/docs/components/input-group) | 变体/尺寸/状态轴；前缀后缀与清除按钮作为**组合件**（Input Group）而非输入框本身；清除按钮在有值且可编辑时出现 | React/CSS 变量；组合件拆成独立组件 |
| [Element Plus el-input](https://element-plus.org/zh-CN/component/input) | `clearable`、`prefix-icon`/`suffix-icon`（**图标**插槽）、`show-password`、`formatter`/`parser`；事件 `focus`/`blur`/`change`/`input`/`clear` | Vue 插槽与 v-model 双绑定；事件是 DOM 事件对象 |
| [Ant Design Input](https://ant.design/components/input-cn) | `allowClear`、`prefix`/`suffix`、`status="error"`、受控/非受控两种模式 | React 受控组件模型 |
| **mygo 官方 `ui.SearchField`** | 框架自己的“带放大镜与清除按钮的输入框”：清除按钮与输入框是**兄弟节点**，画在一个 field 里；清除后 `*query = ""` 并 `in.Focus()` 保持焦点 | 这是本库最直接的参照：结构照它，样式按我们的预设；官方版是固定外观，我们做成可预设的槽位 |
| mygo 官方 `ui.Icon` | 图标随文字色（`currentColor`）与字号变化，正合输入框里的图标 | 我们的槽位直接用它 |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 文本输入与绑定 | `ui.TextInput(c, *string)`；输入框与 `*value` 共享字符串，app 改 `*value` 即同步编辑器（`ui/textinput.go:52-57`） | ✅ 清空只需 `*value = ""` |
| 值变化事件 | `Element.Changed()`（本帧值是否变化） | ✅ `OnChanged` |
| 失焦事件 | `Element.FocusLeft()`（焦点自上帧起离开本元素） | ✅ `OnBlurred` |
| 获得焦点事件 | 无「刚获得焦点」的公开 API，只有 `Focused()` 状态 | ⚠️ 只提供状态，不提供事件（记录为限制） |
| 回车 | `Element.OnTextKey(func(ui.TextKey) bool)`，`TextKey{Key, Mods}` | ✅ `OnSubmit` |
| 组合（前后缀/清除按钮） | `Element.Children(fn)` **立即执行**；只有 fn 内新建的元素才成为子节点 → 用**固定槽位**（构造时建好空槽，之后用 `Children` 填充）实现 | ✅ 槽位方案可行 |
| 图标插槽 | `ui.Icon(c, svg)`；图标取所在处的文字色与字号（`currentColor`） | ✅ `PrefixIcon`/`SuffixIcon` |
| 搜索框形态 | 前缀放大镜 + 输入 + 清除按钮，即官方 `ui.SearchField` 的三段结构 | ✅ 用槽位拼出，样式随我们的预设 |
| 清除按钮不抢焦点 | `Element.KeepFocus()`（按下时焦点留在输入框）、`Element.TabStop(false)`（不进 Tab 序）、`Element.Label("清除")` | ✅ |
| 错误态 | `Element.Invalid(bool)`（无障碍状态）、`Element.ErrorMessage(Element)`（错误文案关系）；边框走主题 `Danger` | ✅ |
| 只读/禁用/密码 | `Element.ReadOnly(bool)`、`Element.Disabled(bool)`、`Element.Password()` | ✅（密码显隐切换需 app 状态，后置） |
| 无障碍 | `Element.Label/Description/Invalid/ErrorMessage` | ✅；**不做**依赖 `Element.Access(func(*platform.AccessNode))` 的定制（internal 泄漏，见 [../research.md](../research.md) §五） |

## 3. 实现到何种程度

**首轮（已实现）**
- 变体：`Native`（零值，保留 ui 内置观感）/`Secondary`（别名）/`Primary`/`Outline`/`Ghost`/`Destructive`/`Link`，作用在**外框**（边框、底色、圆角），内层输入框自身透明无边框。
- 尺寸：`SizeSm/Md/Lg` → 外框内衬与字号（走主题 `Space`/`Rem` 刻度）。
- 组合：`Prefix(fn())`、`PrefixIcon(svg)`、`Suffix(fn())`、`SuffixIcon(svg)`、`Clearable()`（清除按钮用 `icons.X()`，**纯图标无按钮外观**，有值且可编辑时出现，点击清空并保持焦点）。
- 事件：`OnChanged(func(string))`、`OnBlurred(func())`、`OnCleared(func())`、`OnSubmit(func())`；查询式 `Changed()`/`Focused()`/`FocusLeft()`/`Value()`。
- 错误态：`Invalid(bool)`（边框转 danger + 无障碍状态）、`ErrorText(string)`（框下错误文案，接入 `ErrorMessage`）、`ErrorMessage(func())`（自定义错误内容）。
- 状态：`Disabled(bool)`、`ReadOnly(bool)`、`Password()`、`Placeholder(string)`、`Label(string)`。
- 动效（见 [motion.md](motion.md)）：聚焦时外框描边渐变为强调色、错误态转 danger 色；清除按钮与错误文案直接出现（姿态动画在本框架不可用，见动效规格书）。

**后置**
- 密码显隐切换（需 app 持有「显示中」状态）。
- 字数限制与计数、formatter/parser（格式化后再绑定）。
- `InputOTP`、`DatePicker`、`InputNumber`、`SearchBar`（独立组件，走同一套预设）。

## 4. 项目结构是否合理

- 位置：根包 L3（变体化控件），与 `Button`/`Badge`/`Icon` 同层。
- **共用基座**：`field[T]`（`field.go`）是文本字段的公共部分——外框、四个槽位（前后缀/清除/文本）、消息槽、变体与尺寸的取色、事件与查询式接口。泛型参数 `T` 是包裹它的控件类型，所以 `Input` 的链返回 `*Input`、`TextArea` 的链返回 `*TextArea`（各自的类型不丢）。
  - 为什么不复制一份：Input 与 TextArea 的 API 有 30 余个方法重合，复制会让两者日后漂移；泛型基座让新增字段类控件（搜索框、下拉触发器）直接复用。
- 依赖：`myui` → `mygo/ui` + `myui/icons`（只用清除按钮那一个图标，实测对二进制影响约等于零）。
- 入口：既可 `myui.NewInput(c, &v)`，也可用每帧句柄 `myui.Bind(c).Input(&v)`（见 [handle.md](handle.md)）。
- 组合件不另立 `InputGroup`：清除与前后缀都是「输入框的一部分」，同源同预设；shadcn 拆开是因为 CSS 组合方便，Go 里槽位方案已经解决。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewInput(c *ui.Context, value *string) *Input` | 创建文本输入框，绑定 `*value` |
| `Variant(v Variant)` · `Primary()` · `Secondary()` · `Native()` · `Outline()` · `Ghost()` · `Destructive()` · `Link()` | 变体：作用在外框（边框/底色/圆角） |
| `Size(s Size)` · `Sm()` · `Md()` · `Lg()` | 尺寸：外框内衬与字号 |
| `Placeholder(s string)` | 占位文字 |
| `Label(s string)` | 无障碍名 |
| `Invalid(on bool)` | 错误态：边框转 danger，并置无障碍的无效状态 |
| `ErrorText(s string)` | 错误文案（显示在外框下方，接入无障碍的错误关系） |
| `ErrorMessage(fn func())` | 自定义错误内容（显示在外框下方） |
| `ReadOnly(on bool)` | 只读 |
| `Disabled(on bool)` | 禁用 |
| `MaxLength(n int)` | 限制用户输入的长度（mygo 原生 ✓，按 rune 计） |
| `Count()` | 显示当前长度（有上限时显示 "12/100" ✓，按 rune 计） |
| `Password()` | 隐藏所输入的内容 |
| `Prefix(fn func())` | 前缀内容（图标、文字、控件皆可） |
| `PrefixIcon(svg *ui.SVG)` | 前缀图标（取文字色与字号） |
| `Suffix(fn func())` | 后缀内容 |
| `SuffixIcon(svg *ui.SVG)` | 后缀图标 |
| `Clearable()` | 清除按钮：有值且可编辑时出现，点击清空并保持焦点；纯图标，无按钮外观，悬停变深 |
| `OnChanged(fn func(string))` | 值在本帧变化时回调 |
| `OnBlurred(fn func())` | 焦点在本帧离开时回调 |
| `OnCleared(fn func())` | 清除按钮清空时回调 |
| `OnSubmit(fn func())` | 回车时回调 |
| `Changed() bool` | 值在本帧是否变化 |
| `Focused() bool` | 是否持有焦点 |
| `FocusLeft() bool` | 焦点是否在本帧离开 |
| `Value() string` | 当前的值 |

用法：

```go
h := myui.Bind(c)
h.Input(&app.email).Outline().Placeholder("邮箱").
	Invalid(app.emailError != "").ErrorText(app.emailError).
	OnChanged(func(v string) { app.email = v })

// 搜索框：前缀放大镜 + 可清除
h.Input(&app.query).Outline().PrefixIcon(icons.Search()).
	Placeholder("搜索").Clearable()
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 复合结构（外框 + 槽位 + 内层输入）
- [x] 清除按钮（图标、`KeepFocus`、`TabStop(false)`、无障碍名、无按钮外观）
- [x] 事件（变化/失焦/清除/回车）与查询式接口
- [x] 前后缀（构建器与图标）与错误态
- [x] 动效（聚焦描边渐变；清除按钮与错误文案直接出现）
- [x] 测试（`input_test.go`：尺寸、禁用/只读、暗色、清除、事件、错误文案、前后缀、前缀图标、清除按钮只有图标、聚焦描边取强调色）
- [x] showcase「Input」页（含搜索框：前缀放大镜 + 可清除）
- [ ] 密码显隐、字数统计、formatter/parser（后置）
- [ ] 退场渐隐（见 [motion.md](motion.md) 后置项）
