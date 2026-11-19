# Field 规格书

> 位置：组件在 [`forms/`](../../forms) 子包（`forms/form.go`）；主题/tokens 与 `Rule` 在根包 `myui`。

> 状态：**已实现**。表单字段的组合层：标签 + 控件 + 说明 + 错误，外加校验与 `Touched`/`Dirty`/`Valid` 状态。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn v4 Field](https://ui.shadcn.com/docs/components/base/field) | `Field`、`FieldLabel`（可点击聚焦、必填星号）、`FieldDescription`（弱化色）、`FieldError`（多条 `role="alert"`） | 取这四件；标签可点聚焦；错误是活区域；控件被标记为无效 |
| [Element Plus el-form-item](https://element-plus.org/en-US/component/form.html) | 星号、`validate-status`、`rules`、失焦/变更校验 | 星号在标签后；校验色取 danger；时机由 `When` 配置 |
| [Ant Design Form](https://ant.design/components/form-cn) | `status="error"`、`help`、`rules` | 校验色只有 danger；`help` 对应 `Description` |
| [Lobe MyGo Form](https://mygo.lobehub.com/skills/components/form) | 初始值、`IsDirty`、`Touched`、异步校验 | 保留 `Touched`/`Dirty`/`Valid`；不做异步校验（见 §6 限制） |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 标签可点击聚焦控件 | `Element.Focus()`、`ui.ButtonBase().TabStop(false)` | ✅ |
| 控件的无障碍名与描述 | `Element.Label`、`Element.DescribedBy` | ✅ |
| 错误与控件的关系 | `Element.ErrorMessage`、`Element.Invalid`、`Role(ui.RoleAlert)` | ✅ |
| 跨帧状态 | `ui.Local(e, key, init)` | ✅ 状态挂在控件节点上 |
| 「刚获得焦点」事件 | 无，只有 `Focused()` 状态 | ⚠️ 由 `Focused()` 的跨帧边沿算出 |
| 自动拿到控件的内层元素 | `ui.Element` 是结构体，不是接口 | ⚠️ `build` 用推断的泛型返回类型：返回控件本体时通过 `Control()` 自动取内层，返回裸 `ui.Element` 时直接使用 |

## 3. 实现到何种程度

**已实现**
- `NewField[T, B](c, label, value, build, o)`：一次调用给出值、规则、时机、说明与必填标记；不再有第二处 `Validate(value, rules...)`。
- `build` 可以返回 `*Input`/`*Select`/`*ComboBox` 等控件本体——字段自动调用它的 `Control()`，调用者不必记得写 `.Control()`；返回裸 `ui.Element` 也能用。
- 规则逐条检查，第一条不通过的文案成为错误；内置 `Required`/`MinLength`/`MaxLength`/`Pattern`/`Email`。
- `When` 决定**何时显示**错误：`OnBlur`（默认）/`OnChange`/`OnFocus`/`Always`。值本身**每帧都检查**，所以 `Valid()` 不依赖 `Touched()`：已填且合法的值在首次失焦前就是 `true`。
- `Touched()`（失焦过）、`Dirty()`（与初始值不同）、`Valid()`（规则通过）。
- `ErrorText(s)`、`Errors(msgs...)`、`Description(s)`、`Required(on)`、`Invalid(on)` 均可单独调用，且即时渲染。

**已知限制**
- `Rule` 是 `func(value string) string`；非字符串的值经 `fmt.Sprint` 转成文本再检查（见 §5）。
- 没有异步校验/防抖：mygo 与 my-ui 都不设计时器或后台任务碰 UI 状态。
- 「聚焦即校验」靠跨帧边沿模拟，不是框架事件。

## 4. 项目结构是否合理

- 位置：`forms/form.go`，与 `Input`/`Select` 同在 `forms/` 子包；只依赖 `mygo/ui` 与根包 `myui`，不引入内部包。
- 分工：字段只编排标签/说明/错误并建立无障碍关系，不复制控件的样式；值由 app 持有，校验状态由字段持有。
- 泛型：`Field[T]` 存 app 的值指针；Go 方法不能泛型，所以 `bind.UI.Field` 只提供 `*string` 的 shim，其它类型直接调 `forms.NewField[T]`。

## 5. API 契约

| 符号 / 属性 | 类型 | 作用 |
|---|---|---|
| `NewField[T, B]` | `(c *ui.Context, label string, value *T, build func() B, o FieldOptions[T]) *Field[T]` | 创建字段；`value` 或 `o.Value` 给出值指针（只给一次），`build` 返回控件本体或裸元素 |
| `FieldOptions.Value` | `*T` | app 持有的值指针；非空时覆盖 `value` 参数 |
| `FieldOptions.Rules` | `[]Rule` | 校验规则，按顺序检查，第一条失败的文案显示 |
| `FieldOptions.When` | `ValidateOn` | 错误显示时机；零值 `OnBlur` |
| `FieldOptions.Description` | `string` | 控件下的弱化说明文字 |
| `FieldOptions.Required` | `bool` | 标签后加 danger 色星号 |
| `Description(s)` | `func(string) *Field[T]` | 设置说明并与控件建立 `DescribedBy` |
| `ErrorText(s)` | `func(string) *Field[T]` | 设置一条错误（`role="alert"`，控件标记无效） |
| `Errors(msgs...)` | `func(...string) *Field[T]` | 设置一组错误，逐条 `role="alert"` |
| `Required(on)` | `func(bool) *Field[T]` | 切换必填星号 |
| `Invalid(on)` | `func(bool) *Field[T]` | 手动标记控件无效，标签转 danger |
| `Touched()` | `func() bool` | 字段是否被离开过 |
| `Dirty()` | `func() bool` | 值是否与初始值不同 |
| `Valid()` | `func() bool` | 规则是否通过（不要求 `Touched`） |
| `OnFocused(fn)` | `func(func()) *Field[T]` | 控件获得焦点的那一帧调用 `fn` |
| `Control()` | `func() ui.Element` | 字段描述的那个元素 |

`ValidateOn` 取值：

| 常量 | 作用 |
|---|---|
| `OnBlur`（零值） | 失焦后显示错误 |
| `OnChange` | 值变化后显示错误 |
| `OnFocus` | 获得焦点时显示错误 |
| `Always` | 每帧都显示错误 |

规则（根包 `myui`）：

| 规则 | 作用 |
|---|---|
| `Required(msg)` | 值不可为空 |
| `MinLength(n, msg)` / `MaxLength(n, msg)` | rune 长度上下限 |
| `Pattern(re, msg)` | 值须匹配正则 |
| `Email(msg)` | 值须是邮箱地址 |

用法（值只传一次，控件本体直接返回，无需 `.Control()`）：

```go
forms.NewField(c.Key("email"), "邮箱", &app.email, func() *forms.Input {
    return forms.NewInput(c.Key("in-email"), &app.email).Outline()
}, forms.FieldOptions[string]{
    Description: "我们只用它发通知",
    Required:    true,
    When:        forms.OnBlur,
    Rules: []myui.Rule{
        myui.Required("请填写邮箱"),
        myui.Email("邮箱格式不对"),
    },
})
```

非字符串值：

```go
forms.NewField(c, "数量", &app.count, func() *forms.InputNumber {
    return forms.NewInputNumber(c, &app.count, 0, 100, 1)
}, forms.FieldOptions[float64]{
    Rules: []myui.Rule{myui.Required("请填写数量")},
})
```

## 6. 状态、限制与进度

| 状态 | 表现 |
|---|---|
| 默认 | 标签 + 控件 + 说明；无错误 |
| 禁用 / 只读 | 由控件的 `Disabled`/`ReadOnly` 决定，字段不重复实现 |
| 无效（invalid） | 控件取 danger 边框，标签转 danger，错误 `role="alert"` |
| 空（empty） | `Required` 规则失败时为空值显示错误；`Valid()` 为 `false` |
| 加载中（loading） | 不适用：本组件没有异步校验（见限制） |

**限制**
- 无异步校验、防抖、跨字段联动重校验；应用可在字段外自行管理，或在值变化后重建字段。
- `Rule` 只接受字符串；非字符串值经 `fmt.Sprint` 转换，复杂规则请自行把值格式化后再比较。
- `Valid()` 每帧反映规则结果，但错误**文案**只在 `When` 指定的时机出现——这是有意的，避免表单一打开就满屏红字。

- [x] 调研与可行性结论（本文）
- [x] 实现（`forms/form.go`）
- [x] `ui.Tester` 测试（`forms/form_test.go`：标签/说明/错误、聚焦、变更/失焦校验、填充值失焦前即 `Valid`、单次指针）
- [x] showcase「表单」页
- [ ] 异步校验与跨字段联动（后置）
