# Window 规格书

无框窗口外壳：`window.New` 一行建窗——mygo 无框窗口（`TitleBarHidden + HideWindowButtons`）+ 根包的 `ctrlpkg.NewTitleBar` + 每帧主题应用 + 窗口动作接线，app 不再自己拼装这几样。

## API 契约

```go
type Options struct {
	Title            string
	Width, Height    int
	MinWidth, MinHeight int
	Maximized        bool
	Center           bool
	StateKey         string              // mygo 窗口位置/大小的持久化 key；空则不保存
	Icon               *ui.SVG             // 品牌标图标；不传则用默认组件标
	Logo               func(c *ui.Context) // 自定义品牌标构建器（覆盖 Icon）
	Name               string              // 应用名，标右侧；同时是默认窗口标题
	Look               ctrlpkg.TitleBarLook   // Auto 按构建平台取 Windows/macOS 风格
	Theme              func() myui.Options // 每帧的主题选项；nil 用默认
	BackgroundColor    string              // 窗口自身底色（mygo CSS 语法）；空则每帧取主题 Background
	View               func(c *ui.Context) // 标题栏之下的内容
}

func New(o Options) *mygo.Window
func Open(icon *ui.SVG, name string, view func(c *ui.Context)) *mygo.Window // 最简形式
```

用法（只给图标和名字也能开窗口；Title 默认取 Name，Look 按平台，主题取默认）：

```go
mygo.App.WhenReady(func() {
	window.Open(icons.Component(), "MyGo App", appView)
	// 或：window.New(window.Options{Icon: icons.Component(), Name: "MyGo App", View: appView})
})

// 主题由 app 管时，Theme 每帧返回当前选项：
window.New(window.Options{
	Icon:  icons.Component(),
	Name:  "MyGo App",
	Theme: func() myui.Options { return myui.Options{Appearance: app.appearance} },
	View:  appView,
})
```
