# Form（my-ui 外观）规格书

> 位置：`forms/` 子包（`forms/form_styled.go`）。字段见 [field.md](field.md)。

包裹 mygo 的 `ui.Form`，套上主题外观；字段与校验由 `forms.Field`/根包 `Rule` 承担。

## API 契约

| 符号 / 属性 | 类型 | 作用 |
|---|---|---|
| `NewForm` | `func(c *ui.Context, fn func()) ui.Element` | 带 my-ui 外观的表单，签名兼容 `ui.Form`；`fn` 在其中建字段 |
| `NewFormWith` | `func(c *ui.Context, variant Variant, fn func()) ui.Element` | 带共享 `Variant` 预设的表单 |
| `variant`（仅 `NewFormWith`） | `Variant` | `Native`（零值，无面）/ `Outline`（描边表面）/ `Secondary`（弱底） |
| `fn` | `func()` | 建字段的回调，按 mygo 的即时模式每帧执行 |

字段的事件与属性（值指针、`Rules`、`When`、说明、必填星号、错误、`Touched`/`Dirty`/`Valid`）见 [field.md](field.md#api-契约)。

用法：

```go
forms.NewFormWith(c, myui.Outline, func() {
    forms.NewField(c, "邮箱", &app.email, func() *forms.Input {
        return forms.NewInput(c.Key("email"), &app.email).Outline()
    }, forms.FieldOptions[string]{
        Required: true,
        Rules: []myui.Rule{
            myui.Required("请填写邮箱"),
            myui.Email("邮箱格式不对"),
        },
    })

    forms.NewField(c, "语言", &app.lang, func() *forms.ComboBox {
        return forms.NewComboBox(c.Key("lang"), forms.ComboOptions{
            Value: &app.lang, Options: langOptions, Label: "语言",
        }).Outline()
    }, forms.FieldOptions[string]{})
})
```

## 状态与限制

| 状态 | 表现 |
|---|---|
| 默认（`Native`） | 无背景/边框，仅保留间距 |
| 描边（`Outline`） | `Surface` 表面 + 1px `Border` + 大圆角 |
| 弱底（`Secondary`） | `Surface` 表面 + 大圆角，无边框 |
| 字段级禁用/只读/无效/空 | 由其中的 `Field` 决定，Form 不重复实现 |
| 加载中（loading） | 不适用：Form 没有异步状态 |

**限制**
- `NewFormWith` 的 `variant` 只在构造时决定表面；要随主题切换请每帧传当前值。
- 字段的对齐是 mygo `ui.Form` 的行为；my-ui 不覆盖它。
- 不是 Lobe 的泛型 `form.New[T]` 状态容器：那需要反射/泛型状态管理，与 mygo 即时模式绑定约定冲突。需要强类型表单时用 `NewForm` 包住 `NewField` 与自己的状态指针。
- 异步校验、防抖、自动保存不在本组件内：mygo/my-ui 不设计时器；应用可自行在字段外管理。
