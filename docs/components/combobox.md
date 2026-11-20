# ComboBox 规格书

> 位置：组件在 [`forms/`](../../forms) 子包（`forms/combobox.go`）；主题/tokens、i18n 文案在根包 `myui`。

可搜索的下拉：触发器是输入框，选中的值(app-owned)与输入的查询(internal/app-owned)是**两份独立的状态**。

## API 契约

`ComboOptions`（一次配置完）：

| 属性 | 类型 | 作用 |
|---|---|---|
| `Value` | `*string` | app 持有的选中值；选择时写入，清除时置 `""` |
| `Text` | `*string` | 输入框文本与查询；nil 时用组件内部 slot |
| `Options` | `[]SelectOption` | 弹层里的选项（label/value/group/icon/disabled） |
| `Clearable` | `bool` | 有选中值时用清除图标替换箭头 |
| `EmptyText` | `string` | 无匹配时的文案；空则用 i18n 默认 |
| `Placeholder` | `string` | 输入框为空时的占位文字 |
| `Label` | `string` | 无障碍名 |
| `Filter` | `func(SelectOption, string) bool` | 自定义匹配；nil 用默认 label 匹配 |
| `OnChanged` | `func(value string)` | 选择或清除时的回调；可稍后用 fluent 方法覆盖 |

构造与方法：

| 符号 | 作用 |
|---|---|
| `NewComboBox(c *ui.Context, o ComboOptions) *ComboBox` | 创建可搜索下拉 |
| `Variant(v)` / `Primary()` / `Secondary()` / `Default()` / `Native()` / `Outline()` / `Ghost()` / `Destructive()` / `Link()` / `Success()` / `Info()` / `Warning()` | 与 Input 同一套共享预设 |
| `Size(s)` / `Xs()` / `Sm()` / `Md()` / `Lg()` | 文本与内边距尺度 |
| `Shape(s)` | 圆角/胶囊/圆形 |
| `Width(w)` | 固定外框宽度；0 则填满父容器 |
| `Placeholder(s)` / `Label(s)` | 占位文字 / 无障碍名 |
| `Disabled(on)` | 禁用触发器 |
| `Invalid(on)` | 外框取 danger 边框 |
| `Open()` / `SetOpen(on)` | 查询/设置弹层开关 |
| `Clearable` 行为 | 由 `ComboOptions.Clearable` 决定；有值时出现清除图标 |
| `EmptyText(s)` / `Filter(fn)` | 运行时改空文案 / 匹配规则 |
| `OnChanged(fn)` | 注册回调；晚绑定会收到待处理的选择 |
| `Chosen()` / `Selected()` / `Text()` | 最后一次选择 / 当前选中值 / 当前查询 |
| `Changed()` | 输入框在本帧是否变化 |
| `Control()` | 返回输入框，供 `Field` 使用 |

用法：

```go
forms.NewComboBox(c.Key("lang"), forms.ComboOptions{
    Value:     &app.lang,
    Options: []forms.SelectOption{
        {Label: "简体中文", Value: "zh-Hans"},
        {Label: "English", Value: "en"},
        {Group: "更多", Label: "日本語", Value: "ja"},
    },
    Clearable:   true,
    Placeholder: "搜索语言…",
    Label:       "语言",
    OnChanged:   func(v string) { app.lastEvent = "chosen: " + v },
}).Outline()
```

晚绑定回调：

```go
combo := forms.NewComboBox(c, forms.ComboOptions{Value: &app.lang, Options: opts})
combo.OnChanged(func(v string) { app.langChanged(v) }) // 构造之后才挂也能收到
```

与 Field 组合：

```go
forms.NewField(c, "语言", &app.lang, func() *forms.ComboBox {
    return forms.NewComboBox(c.Key("cb"), forms.ComboOptions{
        Value: &app.lang, Options: opts, Label: "语言",
    }).Outline()
}, forms.FieldOptions[string]{Required: true})
```

## 状态与限制

| 状态 | 表现 |
|---|---|
| 默认 | 输入框 + 箭头；输入即过滤 |
| 禁用（disabled） | 输入框不可编辑，弹层不打开 |
| 无效（invalid） | 外框取 danger 边框 |
| 空（empty） | 无匹配时显示 `EmptyText`；`Clearable` 清空后 `Selected()`/`Text()` 都为空 |
| 加载中（loading） | 不适用：远程搜索由 app 换掉 `Options` 实现，组件不设计时器 |

**限制**
- 弹层内容由 mygo 接管：不能放头部/尾部插槽；空文案是唯一允许的非选项内容。
- `Chosen()` 是「最后一次选择」，清除后为空；`Selected()` 始终读 `Value`。
- `Filter` 只决定哪些选项显示，不改变 `Value`；过滤掉已选项后打勾不显示，但值仍在。
- `Text` 是输入框的文本：选中后写入选中项的 label（和 mygo 自己的 ComboBox 一致），所以「查询」在选中后被替换；真正的选中值始终在 `Value` 里。
