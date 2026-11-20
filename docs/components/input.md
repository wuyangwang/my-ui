# Input 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

实现见 `input.go`，测试见 `input_test.go`，示例见 `examples/showcase` 的「Input」页。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewInput(c *ui.Context, value *string) *Input` | 创建文本输入框，绑定 `*value` |
| `Variant(v Variant)` · `Primary()` · `Secondary()` · `Native()` · `Outline()` · `Ghost()` · `Destructive()` · `Link()` | 变体：作用在外框（边框/底色/圆角） |
| `Size(s Size)` · `Sm()` · `Md()` · `Lg()` | 尺寸：外框内衬与字号 |
| `Placeholder(s string)` | 占位文字 |
| `Label(s string)` | 无障碍名 |
| `Invalid(on bool)` | 错误态：边框转 danger，并置无障碍的无效状态 |
| `ErrorText(s string)` | 错误文案（显示在外框下方，接入无障碍的错误关系） |
| `ErrorMessage(fn func())` | 自定义错误内容（显示在外框下方） |
| `ReadOnly(on bool)` | 只读 |
| `Disabled(on bool)` | 禁用 |
| `MaxLength(n int)` | 限制用户输入的长度（mygo 原生，按 rune 计） |
| `Count()` | 显示当前长度（有上限时显示 "12/100"，按 rune 计） |
| `Password()` | 隐藏所输入的内容 |
| `Prefix(fn func())` | 前缀内容（图标、文字、控件皆可） |
| `PrefixIcon(svg *ui.SVG)` | 前缀图标（取文字色与字号） |
| `Suffix(fn func())` | 后缀内容 |
| `SuffixIcon(svg *ui.SVG)` | 后缀图标 |
| `Clearable()` | 清除按钮：有值且可编辑时出现，点击清空并保持焦点；纯图标，无按钮外观，悬停变深 |
| `OnChanged(fn func(string))` | 值在本帧变化时回调 |
| `OnBlurred(fn func())` | 焦点在本帧离开时回调 |
| `OnCleared(fn func())` | 清除按钮清空时回调 |
| `OnSubmit(fn func())` | 回车时回调 |
| `Changed() bool` | 值在本帧是否变化 |
| `Focused() bool` | 是否持有焦点 |
| `FocusLeft() bool` | 焦点是否在本帧离开 |
| `Value() string` | 当前的值 |

用法：

```go
h := bindpkg.Bind(c)
h.Input(&app.email).Outline().Placeholder("邮箱").
	Invalid(app.emailError != "").ErrorText(app.emailError).
	OnChanged(func(v string) { app.email = v })

// 搜索框：前缀放大镜 + 可清除
h.Input(&app.query).Outline().PrefixIcon(icons.Search()).
	Placeholder("搜索").Clearable()
```
