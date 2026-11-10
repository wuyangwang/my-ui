# ComboBox 规格书

> 状态：**已实现**（`combobox.go`）。可搜索的下拉：触发器就是输入框。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| mygo 的 `ui.ComboboxBase` | 「a field showing the choice, whose text filters the options below it as the user types, those starting with it first, while a click or Down shows them all」；`Item(value)`、`Popup(fn)`、`Chosen()`、`Filtering()`、`Open()`/`SetOpen()`；触发器**就是输入框** ✓ | 触发器用输入框 ✓（所以过滤就发生在触发器里 ✓，不必往弹层塞东西 ✓）；选项与弹层按 Select 的样式 ✓ |
| [Element Plus el-select filterable](https://element-plus.org/en-US/component/select.html) | 过滤规则=**选项 label 含输入值** ✓；`filter-method` 可自定义；`no-match-text`（默认 "No matching data"）/`no-data-text`（默认 "No data"） | 默认按 label 匹配 ✓；空状态文案暂缺 ✓ |
| [shadcn/ui Combobox](https://ui.shadcn.com/docs/components/combobox) | 弹层里一个搜索框 + 选项列表；`CommandInput` 的占位文字 | 搜索框在**触发器**里（mygo 的形状 ✓），不是弹层里 |
| [Lobe MyGo AutoComplete](https://mygo.lobehub.com/skills/components/auto-complete) | 「an Input that suggests options below it as the user types, matching what was typed, with an empty state」 | 同上 ✓；空状态待补 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 触发器即输入框 | `ui.ComboboxBase(c, text *string)` → `ComboboxParts.Input` ✓ | ✅ |
| 选项与弹层 | `Item(value string)`、`Popup(fn func(panel Element))` ✓ | ✅ |
| 选中的值 | `Chosen() (string, bool)` ✓；`Input.Changed()` 报告新的选择 ✓ | ✅ |
| 过滤 | 文字在 app 的 `*string` 里 ✓，选项按 label 匹配 ✓ | ✅ |
| 开关弹层 | `Open()`/`SetOpen()` ✓；输入即开（`Filtering()` ✓） | ✅ |
| 弹层里的其它内容 | 弹层内容由机制接管（见 [select.md](select.md) §2）✗ | ❌ 弹层里只能有选项 ✓ |

## 3. 实现到何种程度

**已实现**
- `NewComboBox(c, text, options)`：触发器是一个带输入框样式的字段 ✓，输入即过滤 ✓（按 label 匹配 ✓），选中即取替文字 ✓。
- 变体与尺寸：与 Input 同一套 `Variant`/`Size`/`Shape` ✓（外观直接画在输入框上 ✓）。
- 分组：选项带 `Group`，弹层按组给标题 ✓（与 Select 一致 ✓）。
- 禁用选项 ✓；选中项打勾 ✓。
- 占位文字、无障碍名 ✓；宽度随内容（默认最小宽度，`Width` 可指定 ✓）。
- 事件：`OnChanged(fn func(value string))`（mygo 的 combobox 在选择后的**下一帧**报告该选择 ✓，库里照此实现 ✓）；查询式 `Chosen()`/`Changed()`/`Open()`。

**已知限制**
- 弹层里放不下过滤框以外的任何内容（mygo 的机制接管弹层内容 ✗）→ 空状态文案、头部/尾部插槽做不了 ✗（**框架限制**，不是后置 ✗）。
- 触发器就是输入框 → **装不下清除按钮**（Input 的清除按钮在外框的槽位里 ✓，这里没有外框 ✗）→ ComboBox 无清除 ✗。

**后置**
- 远程搜索（app 换掉 options 即可 ✓）、自定义匹配规则（我们暴露过滤回调 ✗）。
- 空状态与弹层插槽（需上游支持弹层内容 ✗）。

## 4. 项目结构是否合理

- 位置：根包 L3，与 `Select` 同层；单独成组件而不是 Select 的一个开关 ✓（与参考实现的组件族一致 ✓：shadcn 的 Combobox ✓、Lobe 的 AutoComplete ✓；且 Select 的弹层装不下过滤框 ✗，硬塞进 Select 只会做出半成品 ✓）。
- 依赖：只用 `mygo/ui` + `myui/icons`（勾 ✓）。

## 5. API 契约

```go
// NewComboBox creates a drop-down whose trigger is a text field: what
// is typed narrows the options down.
func NewComboBox(c *ui.Context, text *string, options []SelectOption) *ComboBox

func (b *ComboBox) Variant(v Variant) *ComboBox   // 与 Input 同义
func (b *ComboBox) Size(s Size) *ComboBox       // 与 Input 同义
func (b *ComboBox) Shape(s Shape) *ComboBox     // 与 Input 同义
func (b *ComboBox) Width(w float32) *ComboBox
func (b *ComboBox) Placeholder(s string) *ComboBox
func (b *ComboBox) Label(s string) *ComboBox
func (b *ComboBox) OnChanged(fn func(value string)) *ComboBox
func (b *ComboBox) Chosen() string
func (b *ComboBox) Changed() bool
func (b *ComboBox) Open() bool
```

用法：

```go
h := myui.Bind(c)
h.ComboBox(&app.query, []myui.SelectOption{
	{Label: "简体中文", Value: "zh-Hans"},
	{Label: "English", Value: "en"},
}).Outline().Placeholder("搜索语言…").
	OnChanged(func(v string) { app.lang = v })
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`combobox.go`）
- [x] 测试（过滤、选择、分组）
- [x] showcase「Select」页中的 ComboBox 示例
- [ ] 空状态文案、头部/尾部插槽（框架限制，需上游支持弹层内容）
