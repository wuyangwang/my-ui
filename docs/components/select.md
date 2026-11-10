# Select 规格书

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
- 禁用选项（`Disabled: true`）与整个触发器禁用（`Disabled(bool)`）。
- 事件：`OnChanged(func(value string))`、`OnCleared(func())`；查询式 `Value()`、`Open()`、`Changed()`。
- 无障碍：`Label(string)`；触发器是 `RolePopUpButton`，选项是 `RoleOption`。
- 占位文字：未选中时显示（弱化色）。
- 尺寸与变体：`Variant`/`Size` 与 Input 同义。

**后置**
- 自定义选项内容（图标、分组、说明文字）。
- 可搜索（`ui.Combobox` 的路子）、远程搜索。
- 多选、标签折叠。
- 弹层滚动到底的提示（shadcn 的上下箭头按钮）。

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
	Group    string
	Disabled bool
}

// NewSelect creates a drop-down choosing one of options into *selected.
func NewSelect(c *ui.Context, selected *string, options []SelectOption) *Select

func (s *Select) Variant(v Variant) *Select   // 与 Input 同义
func (s *Select) Size(s Size) *Select         // 与 Input 同义
func (s *Select) Sm() *Select
func (s *Select) Md() *Select
func (s *Select) Lg() *Select
func (s *Select) Width(w float32) *Select     // 显式宽度（默认随内容，最小约 200 DIP）
func (s *Select) Placeholder(s string) *Select // 未选中时显示
func (s *Select) Label(s string) *Select       // 无障碍名
func (s *Select) Disabled(on bool) *Select
func (s *Select) OnChanged(fn func(value string)) *Select
func (s *Select) OnCleared(fn func()) *Select   // 清除图标被点时
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
