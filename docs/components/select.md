# Select 规格书

> 状态：**已实现**（`select.go`）。

## 1. 参考实现

| 来源 | 可借鉴之处 | mygo 下的差异 |
|---|---|---|
| [shadcn/ui Select](https://ui.shadcn.com/docs/components/select) | 触发器**长得像输入框**（边框、底色、圆角），右侧一个箭头；弹层是浮起的面板，选项行有高亮与选中标记 | React 组合件；状态在 React 里 |
| [Element Plus el-select](https://element-plus.org/zh-CN/component/select) | `multiple`、`disabled`、`clearable`、`filterable`（可搜索）、分组、自定义选项模板 | Vue 插槽 |
| [Ant Design Select](https://ant.design/components/select-cn) | 触发器与弹层的尺寸/宽度联动、`dropdownRender`、状态色 | React |
| [Lobe MyGo Select](https://mygo.lobehub.com/skills/components/select) | 触发器用 Input 的外观、弹层标记选中项、尺寸与变体跟 Input 一致 | 与我们的目标最接近 |
| **mygo 官方 `ui.SelectBase`** | 官方文档即给出「自绘」的用法：`sel.Trigger`（一个裸按钮）与 `sel.Popup(func(panel))`（弹层面板，选项用 `sel.Item(value)` 造）——**触发器与弹层都可自己排版** | 官方只给骨架，样式要 app 自己写；这正是我们要包的部分 |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 需要 | 公开 API | 结论 |
|---|---|---|
| 触发器 + 弹层的骨架 | `ui.SelectBase[T](c, selected)` → `SelectParts[T]{Trigger, …}`；`Popup(fn func(panel Element)) Element` ✓ | ✅ |
| 选项行 | `SelectParts.Item(value T) Element` ✓；行有 `Highlighted()` 状态 ✓ | ✅ |
| 选中与切换 | 触发器的 `Changed()` 报告新的选择 ✓；`Open()`/`SetOpen(bool)` 报告与设置开合 ✓ | ✅ |
| 键盘 | 上下键开合与移动高亮、Enter/Space 选择、Escape 关闭、首字母跳转 ✓ | ✅ |
| 无障碍 | 触发器是 `RolePopUpButton`，选项是 `RoleOption`，`Label` 命名 ✓ | ✅ |
| 禁用选项 | `Item` 的行可 `Disabled(true)` ✓ | ✅ |
| 图标（箭头、勾） | `ui.Icon` + 我们的 `icons` 包 ✓ | ✅ |
| 尺寸/变体 | 与 Input 同一套 `Variant`/`Size`；取色走 `fieldLookOf` ✓ | ✅ |

不可行/需规避：不做依赖 `Element.Access(func(*platform.AccessNode))` 的定制（internal 泄漏）。

## 3. 实现到何种程度

**已实现**
- 选项是结构体 `{Label, Value, Disabled}`：**标签与值可以不同**（绑定的是 `Value`，触发器显示 `Label`；`Label` 留空则显示 `Value`）——这是 shadcn/antd 的写法。
- 触发器**与 Input 同一套外观**（外框、边框、底色、圆角、内衬、字号），右侧一个下拉箭头（`icons.ChevronDown()`），点击开合；开合状态由 mygo 维护。
- 弹层面板：底色、边框、圆角、内衬；选项行有内衬；高亮行用主题的强调色弱化底；**选中项带勾**（`icons.Check()`）。
- 禁用选项（`Disabled: true`）与整个触发器禁用（`Disabled(bool)`）。
- 事件：`OnChanged(func(value string))`；查询式 `Value()`、`Open()`、`Changed()`。
- 无障碍：`Label(string)`；触发器是 `RolePopUpButton`，选项是 `RoleOption`。
- 占位文字：未选中时显示（弱化色）。
- 尺寸与变体：`Variant`/`Size` 与 Input 同义。

**后置**
- 自定义选项内容（图标、分组、说明文字）。
- 可搜索（`ui.Combobox` 的路子）。
- 多选、清空（clearable）。
- 弹层的宽度与触发器对齐、滚动、选项的虚拟化。

## 4. 项目结构是否合理

- 位置：根包 L3，与 `Input` 同层；取色与尺寸刻度复用 `fieldLookOf`/`fieldPadding`/`fieldFontSize`（与 `field[T]` 基座同源），**不与输入框共用基座**：Select 的触发器里没有文本输入，事件与状态也完全不同。
- 依赖：`myui` → `mygo/ui` + `myui/icons`（箭头与勾各一个）✓。
- 选项的来源：`NewSelect(c, &value, options []string)` 先收字符串选项（够用且好测）；自定义选项内容等后置，届时加 `Items(...)` 之类的构建器。

## 5. API 契约

```go
// SelectOption is one option of a Select: the label the trigger and the
// popup show, and the value the app holds.
type SelectOption struct {
	Label    string
	Value    string
	Disabled bool
}

// NewSelect creates a drop-down choosing one of options into *selected.
func NewSelect(c *ui.Context, selected *string, options []SelectOption) *Select

func (s *Select) Variant(v Variant) *Select   // 与 Input 同义
func (s *Select) Size(s Size) *Select         // 与 Input 同义
func (s *Select) Sm() *Select
func (s *Select) Md() *Select
func (s *Select) Lg() *Select
func (s *Select) Placeholder(s string) *Select // 未选中时显示
func (s *Select) Label(s string) *Select       // 无障碍名
func (s *Select) Disabled(on bool) *Select
func (s *Select) OnChanged(fn func(value string)) *Select
func (s *Select) Value() string
func (s *Select) Open() bool
func (s *Select) Changed() bool
```

用法：

```go
h := myui.Bind(c)
h.Select(&app.lang, []myui.SelectOption{
	{Label: "简体中文", Value: "zh-Hans"},
	{Label: "English", Value: "en"},
	{Value: "日本語"},
}).Outline().Placeholder("选择语言").
	OnChanged(func(v string) { app.lang = v })
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`select.go`）
- [x] 测试（`select_test.go`：标签与值不同、未选中不显示标签、开合与选择、占位文字、禁用不弹、尺寸）
- [x] showcase「Select」页
- [ ] 自定义选项内容、可搜索、多选、清空（后置）
