# 帧句柄（myui.Bind）规格书

> 状态：**已实现**（`bind.go`）。解决「每个组件都要传 `c`」的噪音，同时不引入包级状态。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo 的 `ui` 包本身 | 每个元素构造器都收 `c *ui.Context`：`ui.Button(c, "Save")`、`ui.Text(c, …)`、`ui.Icon(c, svg)`。`c` 是「这一帧的构建器」，带着当前主题、键空间与帧状态 |
| Flutter / SwiftUI 一类 | 「一个句柄持有当前环境」的写法（`BuildContext`、`Environment`），调用点更短 |

## 2. 能否实现

- `*ui.Context` **每个窗口一个**：`makeContext` 返回 `c.rt.public`（`ui/frame.go:15-20`），同一窗口内指针不变 → **句柄在同一窗口内跨帧、跨页面都成立**，每个窗口各绑一次即可（窗口各有自己的 `*Context`）。
- 每帧的状态在引擎的内部 context 里，构造器经 `c.build()` 取当帧那一份，且**只在帧内有效**（`ui/frame.go:29-35`）→ 组件必须在帧内构造；帧外构造返回空元素（不崩，但什么都不画）。`c.Valid()` 可查当前是否在帧内。
- 键（identity）：`c.Key(key)` 是**改写**上下文——设置「下一个用此上下文建的元素」的键——并返回同一个指针（`ui/frame.go:49-54`）。句柄的用法因此是 `myui.Bind(c.Key("row")).Button("…")`：键落在紧跟其后的那个元素上。
- 多窗口：`mygo.NewWindow` 可多次调用，每个窗口一个 `*Context` → 句柄不可跨窗口使用（否则元素会建到别的窗口里）。
- AGENTS.md 硬规则：颜色与度量来自 `c.Theme()`，不缓存于包级状态。句柄只持有上下文，不缓存颜色 ✓。

结论：**`c` 可以只传一次**（每窗口一次），但组件的**构造**仍要发生在帧内。

## 3. 实现到何种程度

- `myui.Bind(c) UI`：持有该帧的上下文，**每帧重新绑定**，不跨帧保存。句柄是**值类型**，绑定与取用都不分配内存。
- 方法只覆盖 my-ui 自己的组件与主题入口：`Button`、`Input`、`Badge`、`Icon`、`Scaffold`、`Apply`、`Theme`、`Context`。
- **不代理 mygo 的通用件**（`ui.Text`/`ui.Row`/`ui.Column` 等）：那是框架的 API，复制一遍只会让维护面变大。
- **两套入口并存**：`NewXxx(c, …)` 是原语（与 mygo `ui.Button(c, …)` 同形、godoc 可检索、可单独使用），`Bind(c)` 是每帧的糖（少传一个 `c`）。两者等价，随 app 的口味选。
- `Apply`/`NewTheme`/`Scaffold` 本来就是收 `c` 的函数，不随句柄搬家。

## 4. 项目结构是否合理

- 位置：根包 L1/L3 之间的**入口层**（`bind.go`），不引入新包、不引入依赖。
- 句柄是无状态的薄壳：所有样式仍从 `c.Theme()` 派生。
- 生命周期：**一个窗口一个句柄**。同一窗口的多页面（侧栏切换、路由）与多帧共用它；每个窗口各绑一次，句柄不可跨窗口。
- 构造时机：组件仍须在**帧内**构造（帧外构造得到空元素）；`c.Valid()` 可查是否在帧内。

## 5. API 契约

```go
// Bind binds a frame's context to the components of my-ui.
func Bind(c *ui.Context) UI

func (u UI) Context() *ui.Context
func (u UI) Theme() *ui.Theme
func (u UI) Apply(o Options) *ui.Theme
func (u UI) Button(label string) *Button
func (u UI) Input(value *string) *Input
func (u UI) Badge(text string) *Badge
func (u UI) Icon(svg *ui.SVG) *Icon
func (u UI) Scaffold(o ScaffoldOptions)
```

用法（两种等价）：

```go
func view(c *ui.Context) {
	h := myui.Bind(c)
	h.Apply(myui.Options{Appearance: myui.AppearanceDark})
	h.Button("保存").Primary()
}

// 或者
func view(c *ui.Context) {
	myui.Apply(c, myui.Options{Appearance: myui.AppearanceDark})
	myui.NewButton(c, "保存").Primary()
}
```

跨页面的用法（句柄放在 app 的状态里，一个窗口一个）：

```go
type app struct {
	h myui.UI // 这个窗口的句柄
}

func (a *app) view(c *ui.Context) {
	if a.h.Context() == nil {
		a.h = myui.Bind(c)
	}
	a.h.Button("保存").Primary()
}
```

## 6. 状态

- [x] 决策与依据（本文）
- [x] 实现（`bind.go`，值类型句柄）
- [x] 测试（`bind_test.go`：与构造器等价、主题、键、跨帧仍然有效）
- [ ] showcase 示例页（随各组件示例展示）
