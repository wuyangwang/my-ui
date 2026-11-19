# Select 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`select.go`）。

## 1. 参考实现（已核到具体条目）

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Select（v4 registry 源码）](https://github.com/shadcn-ui/ui/blob/main/apps/v4/registry/new-york-v4/ui/select.tsx) | 触发器 `flex w-fit items-center justify-between gap-2 rounded-md border border-input bg-transparent px-3 py-2 text-sm` + `data-[size=default]:h-9` / `data-[size=sm]:h-8` + `data-[placeholder]:text-muted-foreground` + `focus-visible:border-ring` + `aria-invalid:border-destructive` + `transition-[color,box-shadow]`；箭头 `ChevronDown size-4 opacity-50`；弹层 `min-w-[8rem] rounded-md border bg-popover shadow-md p-1`；选项行 `py-1.5 pr-8 pl-2` + 高亮 `focus:bg-accent` + 勾 `absolute right-2` + 禁用 `opacity-50` | 触发器与 Input 同观感；占位文字弱化色；聚焦/错误态换描边色；箭头弱化；弹层有边框、底色、圆角、内衬与**最小宽度**；选项行右侧留勾的位置；高亮用强调色弱化底；禁用变灰 |
| [Element Plus el-select](https://element-plus.org/en-US/component/select.html) | 选项形如 `{value, label, disabled}`（label 缺省即 value）；`clearable`「a clear icon will appear」+ `clear` 事件 + `value-on-clear`；`clear-icon` 默认 CircleClose、`suffix-icon` 默认 ArrowDown；`size` 三档；`placeholder`；`filterable`/`multiple`/`grouping`/`allow-create`/`remote`；`fit-input-width`；`value-on-clear`/`empty-values` | 选项结构 `{Label,Value,Disabled}`（**标签与值可不同**）；**可选清除**：有值时右侧出现清除图标，点击清空并回调；箭头在清除图标之后；尺寸三档；占位文字 |
| [Ant Design Select](https://ant.design/components/select-cn) | `allowClear`、`showSearch`、`optionRender`、`labelInValue`、`popupRender`；示例里用 `style={{width: 120}}` 固定宽度 | 宽度由 app 明确给出（我们给 `Width` 方法 + 默认宽度）；`labelInValue` 对应我们「绑定值、显示标签」 |
| [Lobe MyGo Select](https://mygo.lobehub.com/skills/components/select) | 触发器用 Input 的外观，弹层标记选中项，变体与尺寸跟 Input 一致 | 与我们的目标一致：触发器=Input 观感，弹层打勾 |

**宽度这件事，参考实现的结论**：
- element-plus 自 2.5.0 起默认宽度是 **100%**（填满容器），文档明确说「行内使用时宽度会塌掉，要自己给宽度」，其示例统一写 `style="width: 240px"`。
- shadcn 的触发器是 `w-fit`（随内容收缩），但**弹层**有 `min-w-[8rem]`（128px）兜底。
- 结论：**宽度是 app 的事，但库要给一个不会塌掉的默认值** → 我们取「内容宽 + 默认最小宽度（约 200 DIP）+ `Width(w)` 显式指定」。

**清除图标这件事**：element-plus 的 `clearable` 是显式开关，图标是 CircleClose，点了走 `clear` 事件与 `value-on-clear`；shadcn 没有。→ 我们把它做成**默认有**（用户要求）：有值时出现，点击清空并回调，图标与 Input 的清除按钮一致（普通 ×，不是圆圈 ×）。

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
| **清除按钮的无障碍** | 清除按钮是触发器的**子节点**，而触发器的角色是 `RolePopUpButton`（leaf role）——mygo 的无障碍树不展开 leaf-role 元素的子节点（`ui/access.go:772`），所以清除按钮**不出现在无障碍树里**（可点、可见，读屏用户摸不到） | ⚠️ 已知限制：清除按钮与 Input 的清除按钮同款做法；若要可达，需把清除按钮挪到触发器之外（触发器就窄了，弹层也窄） |

不可行/需规避：不做依赖 `Element.Access(func(*platform.AccessNode))` 的定制（internal 泄漏）。

## 3. 实现到何种程度

**已实现**
- 选项是结构体 `{Label, Value, Disabled}`：**标签与值可以不同**（绑定的是 `Value`，触发器显示 `Label`；`Label` 留空则显示 `Value`）——element-plus 的选项形状 ✓。
- 触发器**与 Input 同一套外观**（边框、底色、圆角、内衬、字号），右侧依次是**清除图标**与下拉箭头（箭头弱化，`opacity-50` 的等价）。
- **宽度**：随内容，但有默认最小宽度（约 200 DIP，`t.Space(50)`），不塌成文字宽；`Width(w)` 可显式指定（element-plus 写死 `style="width: 240px"` 的做法）。
- **清除**：有值时出现清除图标（与 Input 的清除按钮同款：纯 ×、悬停变深、不抢焦点），点击清空绑定值并回调 `OnCleared`。
- **Native 的观感**：Select 的触发器没有「内置外观」（mygo 的 `SelectBase` 给的是裸按钮），所以 Native = **输入框那套外观**（`Surface` 底 + `Border` 描边），不是「什么都不画」——否则触发器会是个没有内衬的空壳（这正是之前的样式错乱）。`Outline` 与它相差在底色（`Background` vs `Surface`）。
- 弹层面板：底色、边框、圆角、内衬、阴影；选项行有内衬；高亮行用强调色弱化底；**选中项带勾**（`icons.Check()`）；禁用项变灰。
- **选项图标**：`SelectOption.Icon *ui.SVG` 非空时，在触发器（选中后）与弹层每行、文字之前画出来，尺寸 `IconSizeXs`；与 Lobe 的选项图标一致。
- **箭头随开合翻转**：触发器右侧的 `ChevronDown` 在弹层打开时用 `AnimateTo` + `MotionNormal` 旋转 180°，关闭时转回；Reduce Motion 由桌面接管。
- **弹层不窄于触发器**：用 `ui.Local` 记住触发器的宽度，`MinWidth` 取所见最宽值，弹层不会比触发器窄。
- 禁用选项（`Disabled: true`）与整个触发器禁用（`Disabled(bool)`）。
- 事件：`OnChanged(func(value string))`、`OnCleared(func())`；查询式 `Value()`、`Open()`、`Changed()`。
- 无障碍：`Label(string)`；触发器是 `RolePopUpButton`，选项是 `RoleOption`。
- 占位文字：未选中时显示（弱化色）。
- 尺寸与变体：`Variant`/`Size` 与 Input 同义。

**已知限制（不是后置，是框架的形状）**
- **可搜索**：弹层的内容由 mygo 的机制接管——它建的选项就是弹层的子节点，app 自建的子节点会被丢弃（实测：过滤框加不进去 ✗）。而触发器是按钮，按钮收不到文字 ✗ → **可搜索下拉是另一个组件**（ComboBox：触发器就是输入框 ✓，与参考实现的组件族一致 ✓：shadcn 的 Combobox、Lobe 的 AutoComplete ✓）。
- **多选**：值变成切片 → 也是另一个组件 ✓。
- **清除按钮的无障碍**：它嵌在触发器内，而触发器的角色是 `RolePopUpButton`（leaf role）→ mygo 的无障碍树不展开 leaf-role 的子节点（`ui/access.go:772`），清除按钮可点可见但读屏摸不到 ✗。

**后置**
- 选项的说明文字、滚动到底的提示与事件（element-plus 的 `end-reached` ✓）、弹层头部/尾部插槽。图标已实现（见上）。

## 4. 项目结构是否合理

- 位置：根包 L3，与 `Input` 同层；取色与尺寸刻度复用 `fieldLookOf`/`fieldPadding`/`fieldFontSize`（与 `field[T]` 基座同源），**不与输入框共用基座**：Select 的触发器里没有文本输入，事件与状态也完全不同。
- 依赖：`myui` → `mygo/ui` + `myui/icons`（箭头与勾各一个）✓。
- 选项的来源：`NewSelect(c, &value, options []SelectOption)`；选项的图标由 `SelectOption.Icon` 直接给出，无需另加构建器。

## 5. API 契约

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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`select.go`）
- [x] 测试（`select_test.go`：标签与值不同、未选中不显示标签、开合与选择、占位文字、禁用不弹、尺寸、选项图标、开合箭头翻转）
- [x] showcase「Select」页
- [ ] 自定义选项内容、可搜索、多选、清空（后置）
