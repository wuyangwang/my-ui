# Window 规格书

> 状态：**已实现**（`window/window.go`）。无框窗口外壳：`window.New` 一行建窗——mygo 无框窗口（`TitleBarHidden + HideWindowButtons`）+ 根包的 `myui.NewTitleBar` + 每帧主题应用 + 窗口动作接线，app 不再自己拼装这几样。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [titlebar.md](titlebar.md) | 标题栏本体与双风格（Windows / macOS），窗口动作以回调传入 |
| mygo `WindowOptions{TitleBarStyle: TitleBarHidden, HideWindowButtons: true}` | 无框路径：隐藏系统标题栏与控件，**保留**阴影、圆角与四周缩放边（含顶边）。**不用** `Frameless: true`——那是 Web 页面路径（`--app-region: drag`），Windows 上没有顶部缩放边 |
| mygo `Window.Minimize / ToggleMaximize / Close / IsMaximized / SetMinimumSize / Center` | 窗口动作与尺寸约束，本包负责接线 |

## 2. 能否实现

- 建窗 + 无框选项 ✓ → `mygo.NewWindow` 返回 `*mygo.Window`，本包转手给 app。
- 动作接线 ✓ → 标题栏回调转调 `win.Minimize/ToggleMaximize/Close`；最大化状态每帧 `win.IsMaximized()` 读回（视图在窗口线程上跑）。
- 主题 ✓ → `Theme func() myui.Options` 每帧取 app 的明暗/色板，`myui.Apply` 先于标题栏与内容，两者同帧同色。
- 暗色白边 ✓ → 三层一起做：① shell 的列显式画 `t.Background`，盖住 mygo 根视图按系统浅色画的底色；② 窗口自身底色首帧 `light-dark(#ffffff, #09090b)`、之后每帧随主题 `Background` 调 `win.SetBackgroundColor`（可用 `BackgroundColor` 覆盖）；③ **原生窗口边框**（Windows DWM 画的 1px 边）按 `Theme` 的 appearance 调 `mygo.Theme.SetSource`（light/dark/system）——`myui.Apply` 只改视图，不会给边框换色。测试覆盖四边像素与边框外观切换。
- 无窗口渲染 ✓ → `render` 把窗口动作收进 `actions`（窗口出现前为 no-op），`ui.Tester` 可直接画 shell。

## 3. 实现到何种程度

**已实现**：`Options{Title, Width/Height, MinWidth/MinHeight, Maximized, Center, Logo, Name, Look, Theme, View}`；标题栏占顶条，`View` 在剩余空间里构建；`New` 返回 `*mygo.Window`（想拦截关闭、开第二窗口的 app 拿它继续操作）。

**后置**：关闭拦截钩子（`OnShouldClose`）、多窗口管理、窗口状态持久化（mygo `StateKey` 透传）、菜单栏接入。

## 4. 项目结构是否合理

- 位置：**二级包 `window/`**，是全库唯一 import mygo 根包（`mygo.NewWindow` 等）的地方——窗口是 app 建在它上面的外壳，动作在 `mygo.Window` 上而不在 `ui` 包里；根包 `myui` 仍只依赖公开 `ui` API。
- 依赖方向：`window → mygo` + `myui` + `ui`；不反向。

## 5. API 契约

```go
type Options struct {
	Title            string
	Width, Height    int
	MinWidth, MinHeight int
	Maximized        bool
	Center           bool
	Icon               *ui.SVG             // 品牌标图标；不传则用默认组件标
	Logo               func(c *ui.Context) // 自定义品牌标构建器（覆盖 Icon）
	Name               string              // 应用名，标右侧；同时是默认窗口标题
	Look               myui.TitleBarLook   // Auto 按构建平台取 Windows/macOS 风格
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

## 6. 状态

- [x] 调研与实现
- [x] 测试（`window/window_test.go`：标题栏在上内容在下、无 View 也有栏、`Theme` 喂给 app 视图、三个控件经 `actions` 触发、窗口底色随明暗切换、暗色四边无白边、原生边框外观随主题切换、默认 Title 取 Name）
- [x] showcase「Window」页：点按钮用 `window.New` 开一个自定义窗口（只传 `Icon` + `Name` + View/Theme）；品牌标改用标题栏内置绘制（`Icon: icons.Component()`），切主题时同步 `mygo.Theme.SetSource`
- [ ] 关闭拦截、窗口状态持久化（后置）
