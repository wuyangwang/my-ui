# 帧句柄（myui.Bind）规格书

> 状态：**已实现**（`bind.go`）。解决「每个组件都要传 `c`」的噪音，同时不引入包级状态。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo 的 `ui` 包本身 | 每个元素构造器都收 `c *ui.Context`：`ui.Button(c, "Save")`、`ui.Text(c, …)`、`ui.Icon(c, svg)`。`c` 是「这一帧的构建器」，带着当前主题、键空间与帧状态 |
| Flutter / SwiftUI 一类 | 「一个句柄持有当前环境」的写法（`BuildContext`、`Environment`），调用点更短 |

## 2. 能否实现

- `*ui.Context` **每帧新建**：`ui.View` 的实现是 `view(makeContext(c))`（`ui/api_gen.go:3120-3121`）。缓存它等于跨帧握着过期句柄。
- 多窗口：`mygo.NewWindow` 可多次调用（mygo 内部维护窗口列表），每个窗口有自己的上下文。包级保存一个上下文会让组件串到别的窗口。
- AGENTS.md 硬规则：颜色与度量来自 `c.Theme()`，不缓存于包级状态。

结论：**`c` 不能省**；能省的是「重复传」。每帧绑定一次即可。

## 3. 实现到何种程度

- `myui.Bind(c) *myui.UI`：持有该帧的上下文，**每帧重新绑定**，不跨帧保存。
- 方法只覆盖 my-ui 自己的组件与主题入口：`Button`、`Input`、`Badge`、`Icon`、`Scaffold`、`Apply`、`Theme`、`Context`。
- **不代理 mygo 的通用件**（`ui.Text`/`ui.Row`/`ui.Column` 等）：那是框架的 API，复制一遍只会让维护面变大。
- 构造器签名不变（`myui.NewButton(c, …)` 仍在），两种写法并存。

## 4. 项目结构是否合理

- 位置：根包 L1/L3 之间的**入口层**（`bind.go`），不引入新包、不引入依赖。
- 句柄是无状态的薄壳：所有样式仍从 `c.Theme()` 派生。

## 5. API 契约

```go
// Bind binds a frame's context to the components of my-ui.
func Bind(c *ui.Context) *UI

func (u *UI) Context() *ui.Context
func (u *UI) Theme() *ui.Theme
func (u *UI) Apply(o Options) *ui.Theme
func (u *UI) Button(label string) *Button
func (u *UI) Input(value *string) *Input
func (u *UI) Badge(text string) *Badge
func (u *UI) Icon(svg *ui.SVG) *Icon
func (u *UI) Scaffold(o ScaffoldOptions)
```

用法：

```go
func view(c *ui.Context) {
	h := myui.Bind(c)
	h.Apply(myui.Options{Appearance: myui.AppearanceDark})
	h.Button("保存").Primary()
}
```

## 6. 状态

- [x] 决策与依据（本文）
- [x] 实现（`bind.go`）
- [ ] showcase 示例页（随各组件示例展示）
