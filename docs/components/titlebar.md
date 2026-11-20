# TitleBar 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

自绘窗口标题栏：左侧 logo + 应用名，右侧窗口按钮（Windows 风格）或左侧红绿灯（macOS 风格），整条可拖拽，颜色随主题明暗每帧更新——解决系统标题栏不跟随应用明暗切换的问题。`window/` 子包把它和无框窗口一起封装成一行建窗的 `window.New`（见 [window.md](window.md)）。

## API 契约

```go
type TitleBarLook uint8

const (
	TitleBarAuto    // 零值：按构建平台取风格（darwin → macOS，其余 → Windows）
	TitleBarWindows // 控件在右的图标按钮，关闭键悬停危险色
	TitleBarMac     // 红绿灯在左，悬停显字形，未激活转灰
)

type TitleBarOptions struct {
	Icon       *ui.SVG  // 品牌标图标（画在强调色方块里，居中）；不传用默认组件标
	Logo       func()   // 自定义品牌标构建器，覆盖 Icon
	Name       string   // 应用名，logo 右侧；空则不画
	Maximized  *bool    // 最大化状态指针：非 nil 时最大化控件两态切换
	OnMinimize func()   // 最小化回调；nil 不出按钮
	OnMaximize func()   // 最大化/还原回调（app 里转调 win.ToggleMaximize）；nil 不出按钮
	OnClose    func()   // 关闭回调（app 里转调 win.Close）；nil 不出按钮
	Look       TitleBarLook
}

func NewTitleBar(c *ui.Context, o TitleBarOptions) *TitleBar
```

`TitleBar` 内嵌 `ui.Element`（即那条可拖拽的 Row），app 可继续链式涂装（`Background`、覆写高度等）。

用法（app 侧）：

```go
win := mygo.NewWindow(mygo.WindowOptions{
	TitleBarStyle:     mygo.TitleBarHidden,
	HideWindowButtons: true,
	Content:           ui.View(appView),
})
// 视图里每帧：
maximized = win.IsMaximized()
ctrlpkg.NewTitleBar(c, ctrlpkg.TitleBarOptions{
	Name:       "MyGo App",
	Maximized:  &maximized,
	OnMinimize: win.Minimize,
	OnMaximize: win.ToggleMaximize,
	OnClose:    win.Close,
})
```

## 已知限制

- 拖拽、双击最大化、Windows 11 贴靠布局都是平台行为（mygo 的 `DragWindow`/`MaximizeButton`），`ui.Tester` 假后端只能验证布局与点击回调。
- 最大化状态需要 app 自己喂给 `Maximized`（mygo 无公开的"窗口状态变化"视图事件，轮询 `win.IsMaximized()` 即可）；只在标题栏层用时，原生窗口边框的明暗需要 app 调 `mygo.Theme.SetSource`（`window/` 子包已经替 app 做了）。
- 图标取 Lucide 线条风（Square/Copy 带小圆角），与 Windows 11 直角度量略异，但与全库图标一致。
