# 帧句柄（bindpkg.Bind）规格书

> 位置：组件在 [`bind/`](../../bind) 子包；主题/tokens 仍在根包 `myui`。

解决「每个组件都要传 `c`」的噪音，同时不引入包级状态。

## API 契约

```go
// Bind binds a frame's context to the components of my-ui.
func Bind(c *ui.Context) UI

func (u UI) Context() *ui.Context
func (u UI) Theme() *ui.Theme
func (u UI) Apply(o myui.Options) *ui.Theme
func (u UI) Button(label string) *ctrlpkg.Button
func (u UI) Input(value *string) *formspkg.Input
func (u UI) Badge(text string) *displaypkg.Badge
func (u UI) Icon(svg *ui.SVG) *ctrlpkg.Icon
```

用法（两种等价）：

```go
func view(c *ui.Context) {
	h := bindpkg.Bind(c)
	h.Apply(myui.Options{Appearance: myui.AppearanceDark})
	h.Button("保存").Primary()
}

// 或者
func view(c *ui.Context) {
	myui.Apply(c, myui.Options{Appearance: myui.AppearanceDark})
	ctrlpkg.NewButton(c, "保存").Primary()
}
```

跨页面的用法（句柄放在 app 的状态里，一个窗口一个）：

```go
type app struct {
	h bindpkg.UI // 这个窗口的句柄
}

func (a *app) view(c *ui.Context) {
	if a.h.Context() == nil {
		a.h = bindpkg.Bind(c)
	}
	a.h.Button("保存").Primary()
}
```
