# TextArea 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

与 Input 同构，共用 `field[T]` 基座。

## API 契约

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
| `MaxLength(n int)` | 限制用户输入的长度（mygo 原生，按 rune 计） |
| `Count()` | 显示当前长度（有上限时显示 "12/100"，按 rune 计） |
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
