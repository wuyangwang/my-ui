# Form（my-ui 外观）规格书

> 位置：`forms/` 子包（`forms/form_styled.go`）。字段见 [field.md](field.md)。

> 状态：**已实现**。包裹 mygo 的 `ui.Form`，套上主题外观；字段与校验由 `forms.Field`/根包 `Rule` 承担。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| mygo `ui.Form` | 标签与控件左右对齐的表单容器 | 直接包裹并套主题外观，签名保持一致 |
| [shadcn v4 Field](https://ui.shadcn.com/docs/components/base/field) | `FieldGroup`/`FieldSet`/`FieldLegend`、字段间距 | 字段组用 `forms.FieldGroup`/`FieldSet`；间距由 `Gap` 给 |
| [Element Plus el-form](https://element-plus.org/en-US/component/form.html) | `label-position`、`inline`、`size` | 对齐交给 mygo；`size`/`variant` 用共享词表 |
| Lobe `form.*` | 泛型 `form.New[T]` 状态容器、异步校验、自动保存 | 不做：与 mygo 的即时模式绑定冲突，见 §6 限制 |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 标签右对齐排版 | `ui.Form` 内建 | ✅ |
| 字段/说明/错误 | `forms.NewField`、`ui.Field`、`ui.Fieldset` | ✅ |
| 主题外观 | `c.Theme()` + `StateTransition` | ✅ 每帧取色，随主题/深浅色/高对比/文字缩放变化 |
| 签名兼容 | `func(c *ui.Context, fn func()) ui.Element` | ✅ |

## 3. 实现到何种程度

**已实现**
- `NewForm(c, fn)` 与 `ui.Form(c, fn)` 同签名，调用只需换包名；容器套用主题的间距、圆角与表面。
- `NewFormWith(c, variant, fn)` 提供共享 `Variant` 预设：`Native`（零值，无面）/`Outline`（描边表面）/`Secondary`（弱底）。
- 字段对齐仍由 mygo 的 `ui.Form` 完成；校验与错误由 `forms.Field` + 根包 `Rule` 负责，见 [field.md](field.md)。

**已知限制**
- 不是 Lobe 的泛型 `form.New[T]` 状态容器：那需要反射/泛型状态管理，与 mygo 即时模式绑定约定冲突。需要强类型表单时用 `NewForm` 包住 `NewField` 与自己的状态指针。
- 异步校验、防抖、自动保存不在本组件内：mygo/my-ui 不设计时器；应用可自行在字段外管理。

## 4. 项目结构是否合理

- 位置：`forms/form_styled.go`，与 `form.go`（`Field`）同在 `forms/`；只依赖 `mygo/ui` 与根包 `myui`。
- 分工：Form 只管容器的外观与间距，不碰字段的状态；字段自己管标签/说明/错误与校验。
- 主题：颜色与 metrics 每帧从 `c.Theme()` 取，容器带 `StateTransition()`，跟随 `SetTheme` 与系统外观变化。

## 5. API 契约

| 符号 / 属性 | 类型 | 作用 |
|---|---|---|
| `NewForm` | `func(c *ui.Context, fn func()) ui.Element` | 带 my-ui 外观的表单，签名兼容 `ui.Form`；`fn` 在其中建字段 |
| `NewFormWith` | `func(c *ui.Context, variant Variant, fn func()) ui.Element` | 带共享 `Variant` 预设的表单 |
| `variant`（仅 `NewFormWith`） | `Variant` | `Native`（零值，无面）/ `Outline`（描边表面）/ `Secondary`（弱底） |
| `fn` | `func()` | 建字段的回调，按 mygo 的即时模式每帧执行 |

字段的事件与属性（值指针、`Rules`、`When`、说明、必填星号、错误、`Touched`/`Dirty`/`Valid`）见 [field.md §5](field.md#5-api-契约)。

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

## 6. 状态、限制与进度

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

- [x] 调研与可行性结论（本文）
- [x] 实现（`forms/form_styled.go`）
- [x] `ui.Tester` 测试（`forms/form_styled_test.go`：字段绘制）
- [x] showcase「表单」页
- [ ] 泛型表单状态容器、异步校验（后置，见限制）
