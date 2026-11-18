# TitleBar 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`titlebar.go`）。自绘窗口标题栏：左侧 logo + 应用名，右侧窗口按钮（Windows 风格）或左侧红绿灯（macOS 风格），整条可拖拽，颜色随主题明暗每帧更新——解决系统标题栏不跟随应用明暗切换的问题。`window/` 子包把它和无框窗口一起封装成一行建窗的 `window.New`（见 [window.md](window.md)）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo `WindowOptions{TitleBarStyle: TitleBarHidden, HideWindowButtons: true}` | 官方路径：隐藏系统标题栏与系统控件，保留阴影/圆角/缩放边，由视图自绘控件并调用 `Window.Minimize/ToggleMaximize/Close`（**不用** `Frameless: true`——那是 Web 路径，Windows 上没有顶部缩放边） |
| mygo `ui.Context.TitleBar()` | 返回 `TitleBar{Height, Left, Right}`：隐藏标题栏但**未**隐藏系统控件时，控件占据的安全区 |
| mygo `Element.DragWindow()` | 拖拽移动窗口；双击由平台处理最大化（`TitleBarDoubleClicked`） |
| mygo `Element.MaximizeButton()` | 把元素标成窗口最大化按钮：Windows 11 悬停出系统贴靠布局 |
| shadcn/ui · Electron 自定义标题栏 | 高度 32 DIP、控件正方形悬停、关闭键悬停红色（Windows 惯例）、无激活时降灰 |
| Windows 11 标题栏度量 | 控件宽 46 DIP、条高 32 DIP |

## 2. 能否实现

- `c.TitleBar()` ✓ → 高度下限与左右安全区（系统控件仍在时留白、不自绘）。
- `Element.DragWindow()` ✓ → 整条拖拽；双击最大化由平台兜底。
- `Element.MaximizeButton()` ✓ → Windows 11 贴靠布局悬停。
- 窗口动作（Minimize / ToggleMaximize / Close）✗ 不在 `ui` 包里——它们在 `mygo.Window` 上。本库只依赖公开 `ui` API，故**动作由 app 以回调传入**（app 持有 `*mygo.Window`，在回调里转调）；`Maximized *bool` 状态同理由 app 持有（每帧 `win.IsMaximized()` 或窗口事件里更新）。
- 窗口未激活 ✓ → `c.WindowActive()`，墨色降灰。
- 明暗/色板跟随 ✓ → 每帧取 `c.Theme()`，按钮悬停色走 `Element.Transition`（动效刻度，Reduce Motion 由框架兜底）。

## 3. 实现到何种程度

**已实现**：

- 条高 `max(32, c.TitleBar().Height)`；`PaddingStart(bar.Left)` / `PaddingEnd(bar.Right)` 安全区。
- 系统控件仍在（`Left > 0` 或 `Right > 0`，即 `TitleBarHidden` 未配 `HideWindowButtons`）时：只留白 + 可拖拽，**不自绘按钮**（回调忽略），避免两套按钮并存；`HideWindowButtons: true` 时安全区为零，自绘按钮。
- **双风格**，`Look` 选项选择（零值 `TitleBarAuto` 按构建平台：darwin → macOS，其余 → Windows；js/wasm 取 Windows）：
  - **Windows**：logo + 应用名在左，图标按钮（Lucide `Minus` / `Square` / `Copy`(还原) / `X`）在右，宽 46、与条同高、直角、透明底：悬停 `SurfaceHover`、按下 `SurfacePressed`，**关闭键悬停/按下为危险色 + 反色墨**（Windows 惯例）；最大化按钮打 `MaximizeButton()`（Win11 贴靠布局悬停）。
  - **macOS**：红绿灯（`#ff5f57` / `#febc2e` / `#28c840`，12 DIP 圆点、间隔 8）在左，**条方向钉 LTR**（RTL 布局下红绿灯仍在左）；悬停圆点向黑压一档并显字形（× / – / 全屏双三角，按灯色加深取墨），未激活整排转灰（明 `#d8d8d8` / 暗 `#4f4f4f`）；logo + 应用名在灯后。
- 某回调为 `nil` 就不出对应按钮；`Maximized *bool` 非 `nil` 时最大化控件在最大化/还原两态间切换（图标与无障碍名都换）。
- 无障碍名走 i18n（`minimize`/`maximize`/`restore`/`close`），窗口未激活时 Windows 控件墨降 `TextMuted`，RTL 下 Windows 风格的左右安全区镜像。
- 易用性：只传 `Name` 也能画（默认组件标）；传 `Icon` 就用它画在主题色方块里，要完全自定义再传 `Logo`；`Logo` 优先于 `Icon`。
- 悬停色带 `stateTransition()` 渐变。

**后置**：按钮 tooltip、系统菜单（右键还原/最小化菜单）、`AlwaysOnTop` 等更多自绘按钮。

## 4. 项目结构是否合理

- 位置：根包 L4（窗口级组合件），只依赖 `ui` + `icons`，不 import `mygo` 根包：窗口句柄与动作留在 app。

## 5. API 契约

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

## 6. 状态

- [x] 调研与实现
- [x] 测试（`titlebar_test.go`：Windows 风格的布局/点击/无按钮形态/还原态/安全区/危险悬停/未激活降灰，macOS 风格的红绿灯顺序与颜色/未激活转灰/RTL 不动）
- [x] showcase「TitleBar」页（两种风格并排）；showcase「Window」页点按钮经 `window.New` 开自定义窗口
- [ ] tooltip、窗口系统菜单、更多自绘按钮（后置）

## 7. 已知限制

- 拖拽、双击最大化、Windows 11 贴靠布局都是平台行为（mygo 的 `DragWindow`/`MaximizeButton`），`ui.Tester` 假后端只能验证布局与点击回调。
- 最大化状态需要 app 自己喂给 `Maximized`（mygo 无公开的"窗口状态变化"视图事件，轮询 `win.IsMaximized()` 即可）；只在标题栏层用时，原生窗口边框的明暗需要 app 调 `mygo.Theme.SetSource`（`window/` 子包已经替 app 做了）。
- 图标取 Lucide 线条风（Square/Copy 带小圆角），与 Windows 11 直角度量略异，但与全库图标一致。
