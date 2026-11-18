# TextArea 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`textarea.go`）。与 Input 同构，共用 `field[T]` 基座（见 [input.md](input.md) §4）。

## 1. 参考实现

| 来源 | 可借鉴之处 | mygo 下的差异 |
|---|---|---|
| [shadcn/ui Textarea](https://ui.shadcn.com/docs/components/textarea) | 变体/尺寸/状态轴；与 Input 同一套设计语言 | React；无自动高度 |
| [Element Plus el-input type=textarea](https://element-plus.org/zh-CN/component/input) | `rows`、`autosize`（最小/最大行数）、`maxlength` 与字数统计 | Vue 插槽与 v-model |
| [Ant Design Input.TextArea](https://ant.design/components/input-cn) | `autoSize`（按内容增高）、`showCount` | React 受控模型 |
| **mygo 官方 `ui.TextArea`** | 多行输入；`Lines(min, max)` 让它随内容增高、超过上限后滚动（`ui/api_gen.go:5954-5957`）；`ui.TextAreaBase` 是无外观版 | 官方是单控件、单一外观；我们把它做成可预设的复合件 |

## 2. 能否实现

| 能力 | 公开 API | 结论 |
|---|---|---|
| 多行文本与绑定 | `ui.TextArea(c, *string)`；与 `*value` 共享字符串 ✓ | ✅ |
| 随内容增高 | `Element.Lines(min, max)` ✓ | ✅ `Lines(3, 8)` |
| 变体/尺寸/状态 | 与 Input 同一套 `Variant`/`Size`；外框承载外观 ✓ | ✅ |
| 事件 | `Changed()`/`Focused()`/`FocusLeft()` ✓；`OnTextKey` 处理 Ctrl+Enter 等组合键 ✓ | ✅ |
| 错误态 | `Invalid(bool)` + `ErrorMessage(Element)` ✓ | ✅ |
| 清除按钮 | 同 Input：`icons.X()` + `KeepFocus` + `TabStop(false)` ✓ | ✅ |
| 密码 | 多行输入无密码语义 | ❌ 不提供 |

## 3. 实现到何种程度

**已实现**
- 变体与尺寸（同 Input 的共享词表）；`Native` 保留内置观感。
- 占位文字、无障碍名、错误态（`Invalid` + `ErrorText`/`ErrorMessage`）、只读、禁用。
- 组合：`Prefix`/`PrefixIcon`/`Suffix`/`SuffixIcon`、`Clearable()`。
- 事件：`OnChanged`/`OnBlurred`/`OnCleared` 与查询式 `Changed`/`Focused`/`FocusLeft`/`Value`。
- `Lines(min, max)`：随内容增高、超过上限后滚动。
- 动效：聚焦时外框描边渐变为强调色（见 [motion.md](motion.md)）。

**后置**
- 字数限制与计数（`maxlength` 在 mygo 公开 API 里没有对应项，待上游确认）。
- Ctrl/Cmd+Enter 提交：可用 `OnTextKey` 自行实现，未包成方法。
- 自动高度与滚动条外观的更多预设。

## 4. 项目结构是否合理

- 位置：根包 L3，与 `Input` 同层。
- **共用基座**：`field[T]`（`field.go`）是文本字段的公共部分——外框、四个槽位（前后缀/清除/文本）、消息槽、变体与尺寸的取色、事件与查询式接口。泛型参数 `T` 是包裹它的控件类型，所以 `Input` 的链返回 `*Input`、`TextArea` 的链返回 `*TextArea`（各自的类型不丢）。
  - 为什么不复制一份：Input 与 TextArea 的 API 有 30 余个方法重合，复制会让两者日后漂移；泛型基座让新增字段类控件（搜索框、下拉触发器）直接复用。
- 依赖：只用 `mygo/ui` 与 `myui/icons`（清除按钮那一个图标）✓。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewTextArea(c *ui.Context, value *string) *TextArea` | 创建多行文本域，绑定 `*value` |
| `Variant(v Variant)` · `Primary()` · `Secondary()` · `Native()` · `Outline()` · `Ghost()` · `Destructive()` · `Link()` | 变体：作用在外框 |
| `Size(s Size)` · `Sm()` · `Md()` · `Lg()` | 尺寸：外框内衬与字号 |
| `Lines(min, max int)` | 随内容增高（最少 min 行，最多 max 行后滚动） |
| `Placeholder(s string)` | 占位文字 |
| `Label(s string)` | 无障碍名 |
| `Invalid(on bool)` | 错误态：边框转 danger，并置无障碍的无效状态 |
| `ErrorText(s string)` | 错误文案（显示在外框下方） |
| `ErrorMessage(fn func())` | 自定义错误内容 |
| `ReadOnly(on bool)` · `Disabled(on bool)` | 只读 / 禁用 |
| `MaxLength(n int)` | 限制用户输入的长度（mygo 原生 ✓，按 rune 计） |
| `Count()` | 显示当前长度（有上限时显示 "12/100" ✓，按 rune 计） |
| `Prefix(fn func())` · `PrefixIcon(svg)` | 前缀内容 / 前缀图标 |
| `Suffix(fn func())` · `SuffixIcon(svg)` | 后缀内容 / 后缀图标 |
| `Clearable()` | 清除按钮 |
| `OnChanged(fn func(string))` · `OnBlurred(fn func())` · `OnCleared(fn func())` | 事件回调 |
| `Changed() bool` · `Focused() bool` · `FocusLeft() bool` · `Value() string` | 查询式接口 |

用法：

```go
h := bindpkg.Bind(c)
h.TextArea(&app.draft).Outline().Placeholder("写点什么…").
	Lines(3, 8).
	OnChanged(func(v string) { app.draftLen = len(v) })
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 与 Input 共用 `field[T]` 基座
- [x] 实现（`textarea.go`）
- [x] 测试（`textarea_test.go`）
- [x] showcase「TextArea」页
- [ ] 字数限制与计数、提交键（后置）
