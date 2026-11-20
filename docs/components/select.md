# Select 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
// SelectOption is one option of a Select: the label the trigger and the
// popup show, and the value the app holds.
type SelectOption struct {
	Label    string
	Value    string
	Group    string
	Icon     *ui.SVG   // 非空时画在文字之前（触发器与弹层，IconSizeXs）
	Disabled bool
}

// NewSelect creates a drop-down choosing one of options into *selected.
func NewSelect(c *ui.Context, selected *string, options []SelectOption) *Select
func NewSelectClear(c *ui.Context, selected *string, options []SelectOption, clearable bool) *Select

func (s *Select) Variant(v Variant) *Select    // 与 Input 同义
func (s *Select) Size(s Size) *Select          // 与 Input 同义
func (s *Select) Xs() *Select
func (s *Select) Sm() *Select
func (s *Select) Md() *Select
func (s *Select) Lg() *Select
func (s *Select) Shape(shape Shape) *Select
func (s *Select) Width(w float32) *Select      // 显式宽度（默认随内容，最小约 200 DIP）
func (s *Select) Placeholder(s string) *Select // 未选中时显示
func (s *Select) Label(s string) *Select       // 无障碍名
func (s *Select) Disabled(on bool) *Select
func (s *Select) Invalid(on bool) *Select
func (s *Select) EmptyText(text string) *Select
func (s *Select) OnChanged(fn func(value string)) *Select
func (s *Select) OnCleared(fn func()) *Select   // 清除图标被点时
func (s *Select) Value() string
func (s *Select) Open() bool
func (s *Select) SetOpen(on bool) *Select
func (s *Select) Highlight(value string) *Select
func (s *Select) Changed() bool
func (s *Select) Control() ui.Element
```

| API | 作用 |
|---|---|
| `NewSelect(c, selected, options)` | 构造下拉，默认可清除；值绑定 `*selected`，触发器显示选项标签 |
| `NewSelectClear(..., clearable)` | `clearable=false` 时不建清除按钮（构造期决定，之后改无效） |
| `SelectOption.Label` / `Value` | 显示文案 / 绑定的值；`Label` 留空则显示 `Value` |
| `SelectOption.Group` | 分组名；同一组的选项上方显示一次组名 |
| `SelectOption.Icon` | 非空时画在文字之前（`IconSizeXs`），触发器（选中后）与弹层每行都画 |
| `SelectOption.Disabled` | 禁用该项：变灰且不可选 |
| `Variant(v)` / `Primary()` … `Warning()` | 触发器外观，与 Input 同义 |
| `Size(s)` / `Xs()` `Sm()` `Md()` `Lg()` | 触发器的内衬、字号与高度档 |
| `Shape(shape)` | 触发器圆角：直角 / 胶囊 / 圆 |
| `Width(w)` | 显式触发器宽度；不设时随内容但有最小宽度 |
| `Placeholder(s)` / `EmptyText(text)` | 未选中时的占位文字 / 空列表文案 |
| `Label(s)` / `Disabled(on)` / `Invalid(on)` | 无障碍名 / 禁用 / 错误态描边 |
| `OnChanged(fn)` / `OnCleared(fn)` | 值变化 / 清除按钮被点时的回调 |
| `Value()` / `Open()` / `SetOpen(on)` / `Changed()` / `Highlight(v)` / `Control()` | 查询值与状态、开合弹层、移动高亮、取描述用的控件元素 |

用法：

```go
h := bindpkg.Bind(c)
h.Select(&app.lang, []formspkg.SelectOption{
	{Label: "简体中文", Value: "zh-Hans"},
	{Label: "English", Value: "en"},
	{Value: "日本語"},
}).Outline().Placeholder("选择语言").
	OnChanged(func(v string) { app.lang = v })
```
