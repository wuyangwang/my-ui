# ComboBox 规格书

> 位置：组件在 [`forms/`](../../forms) 子包（`forms/combobox.go`）；主题/tokens、i18n 文案在根包 `myui`。

> 状态：**已实现**。可搜索的下拉：触发器是输入框，选中的值(app-owned)与输入的查询(internal/app-owned)是**两份独立的状态**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| mygo `ui.ComboboxBase` | `Input` 即触发器；`Popup(fn)`、`Item(value)`、`Chosen()`、`Open()`/`SetOpen()`、`Filtering()` | 直接用公开 API；过滤发生在输入框里 ✓ |
| [Element Plus el-select filterable](https://element-plus.org/en-US/component/select.html) | `filter-method`、`no-match-text`、`clearable`、选中后显示 label | 默认按 label 大小写无关匹配；`EmptyText` 默认走 i18n；`Clearable` 的清除图标替换箭头 |
| [shadcn/ui Combobox](https://ui.shadcn.com/docs/components/combobox) | 搜索框 + 选项列表 + 空状态 | 搜索框就是触发器（mygo 的形状），不往弹层塞第二个输入框 |
| [Lobe MyGo AutoComplete](https://mygo.lobehub.com/skills/components/auto-complete) | 输入即过滤、空状态 | 同上；空状态文案在弹层里渲染 |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 触发器即输入框 | `ui.ComboboxBase(c, text *string)` | ✅ |
| 选项与弹层 | `Item(value string)`、`Popup(fn func(panel Element))` | ✅ |
| 选中的值 | `Chosen() (string, bool)`（只在构建帧有效） | ✅ 复制进 `ui.Local`，跨帧可读 |
| 查询与值分离 | ComboboxBase 只绑定一个 `*string` | ✅ `Text` 是查询，`Value` 是选择，构造时分别持有 |
| 清除图标 | `controls.NewIconAction`（图标按钮角色） | ✅ 放在输入框外的箭头槽里，替换 chevron |
| 回调晚绑定 | `ui.Local` 存回调 + 待处理选择 | ✅ 构造后（甚至隔帧）再 `OnChanged` 也能收到 |
| 弹层里的自定义内容 | 弹层内容由 mygo 接管 | ⚠️ 只能放选项和空状态文案，见 §6 限制 |

## 3. 实现到何种程度

**已实现**
- `NewComboBox(c, o ComboOptions)`：`Value` 是选中的值（app 持有），`Text` 是查询（可选，nil 则用内部 slot）。
- 输入即过滤（默认大小写无关的 label 匹配，`Filter`/`Filter()` 可替换）；分组标题、禁用选项、选中打勾。
- `Clearable`：有选中值时清除图标替换箭头；点击同时清空 `Value` 与 `Text`，并上报空值。
- `EmptyText`：无匹配时显示；未给时走根包 i18n 的 `Message(c, "noResults")`，未收录该键则回退英文 "No results found"。
- `Chosen()`（最后一次选择）、`Selected()`（当前 `Value`）、`Text()`（当前查询）均跨帧可读。
- `OnChanged(fn)` 通过 `ui.Local` 注册，选择或清除时触发，晚绑定有效。
- `Variant`/`Size`/`Shape`/`Width`/`Disabled`/`Invalid`/`Open`/`SetOpen` 全部保留；颜色与尺寸每帧取自 `c.Theme()`。
- `Control()` 返回输入框，可作为 `Field` 的控件。

**已知限制**
- 弹层内容由 mygo 机制接管：除选项与空状态文案外放不进头部/尾部插槽（框架限制）。
- `Text` 是输入框的文本：选中后写入选中项的 label（和 mygo 自己的 ComboBox 一致），所以「查询」在选中后被替换；真正的选中值始终在 `Value` 里。

## 4. 项目结构是否合理

- 位置：`forms/combobox.go`，与 `Select` 同层；单独成组件而不是给 Select 加开关（Select 的弹层装不下过滤框）。
- 依赖：`mygo/ui` + 根包 `myui` + `myui/controls`（清除图标按钮）+ `myui/icons`（chevron、check、x）；无内部包、无 cgo。
- 状态：`Value`/`Text` 由 app 通过指针持有；`chosen`/`pending`/回调/过滤器/空文案存在 `ui.Local`，随输入框节点存活。

## 5. API 契约

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

## 6. 状态、限制与进度

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

- [x] 调研与可行性结论（本文）
- [x] 实现（`forms/combobox.go`）
- [x] `ui.Tester` 测试（`forms/combobox_test.go`：过滤、选择、分组、清除、空文案、晚绑定回调）
- [x] showcase「可搜索」页
- [ ] 弹层插槽（需上游支持弹层内容）
