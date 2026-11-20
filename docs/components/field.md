# Field 规格书

> 位置：组件在 [`forms/`](../../forms) 子包（`forms/form.go`）；主题/tokens 与 `Rule` 在根包 `myui`。

表单字段的组合层：标签 + 控件 + 说明 + 错误，外加校验与 `Touched`/`Dirty`/`Valid` 状态。

## API 契约

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

## 状态与限制

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
- 「聚焦即校验」靠跨帧边沿模拟，不是框架事件。
