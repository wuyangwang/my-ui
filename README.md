# my-ui

MyGo 原生 UI（[`package ui`](https://github.com/egoist/mygo)）的第三方组件库（独立 module，`go get` 导入即用）。面向 MyGo 自绘的原生 UI（macOS / Linux / Windows，亦支持 wasm），不依赖 webview。

定位：**主题包 + 语义常量 + 变体化控件预设 + 框架没有的组合组件**。`mygo/ui` 已是完整的立即模式控件层；my-ui 在其公开 API 之上补设计系统与桌面组件（shadcn 式 `Variant/Size` 词汇）。默认主题即 shadcn/ui 的 zinc 主题（白底卡片、近黑 primary，`Apply(c, Options{})` 也会应用），仍逐帧跟随系统明暗与桌面偏好。

```go
import (
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/bind"
	"github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/forms"
)

controls.NewButton(c, "Save").Primary().Sm()   // 控件在分组子包
forms.NewInput(c, &app.email).Outline().Clearable()
h := bind.Bind(c)                              // 每帧绑定一次，之后不带 c
h.Button("Save").Primary()
```

## 包结构

- 根包 `myui`：主题/tokens 底座——`Apply`/`NewTheme`、`Variant`/`Size`/`Shape` 词表、`TextSize`/`Space`/`Radius`、动效与 i18n。
- 组件按功能分组到子包，按需引入：
  - `controls/` —— 按钮、图标、勾选/单选/开关/滑杆、自绘标题栏（`ctrlpkg.NewButton`、`ctrlpkg.NewTitleBar` 等）；mygo 原样提供的 `NewToggleGroup`/`Segmented` 直接重导出；
  - `forms/` —— 输入框、多行输入、密码、下拉、可搜索下拉、多选、字段层与表单（`formspkg.NewInput` 等）；输入类组件默认占满父宽，`Width(170)` 固定宽度；
  - `overlays/` —— 提示、浮层、模态、抽屉、通知（`overlaypkg.NewTooltip` 等）；
  - `display/` —— 消息、徽标、卡片、标签页、进度、键帽（`displaypkg.NewAlert` 等）；
  - `bind/` —— 每帧帧句柄 `bind.Bind(c)`；
  - `window/` —— 无框窗口外壳：`window.Open(icon, name, view)` 或 `window.New` 一行建窗（mygo `TitleBarHidden + HideWindowButtons` + 内置 `TitleBar` + 主题/窗口边框/底色 + 动作接线），只给图标和名字即可，其余走默认；是唯一 import mygo 根包的子包；
  - mygo/ui 里 my-ui 既没有、也没规划要做的控件才从对应分组重导出（`display.NewText`、`forms.NewTimeInput`、`controls.NewRangeSlider` 等）；my-ui 已有的控件（Button、Input、Select 等）和路线图里规划中的（Table、Spin、DatePicker 等）都不导出 mygo 的同款；
  - `icons/` —— Lucide（ISC）全量 2,318 个图标，函数式 API（`icons.Save()`），未调用的图标不进二进制；
  - `fonts/` —— 内嵌 OFL 字体与 `ui.RegisterFont` 注册；
  - `ai/` —— ChatList、ModelPicker、Reasoning 等重组件（规划中）。

图标用根包的 `IconSizeXs/Sm/Md/Lg` 四档刻度；清除、关闭这类图标本身就是控件（`controls.NewIconAction`，shadcn/Element Plus 的做法），hover/按下只换图标墨色，颜色与过渡由主题的 `StateTransition` 统一派生。

用法：主题从根包取，控件从分组子包取：

```go
t := myui.Apply(c, myui.Options{Appearance: myui.AppearanceDark})
formspkg.NewInput(c, &name).Outline().Width(170)
```

每个组件的调研与契约见 [`docs/components/`](docs/components)：[button](docs/components/button.md) · [input](docs/components/input.md) · [textarea](docs/components/textarea.md) · [select](docs/components/select.md) · [combobox](docs/components/combobox.md) · [multiselect](docs/components/multiselect.md) · [field](docs/components/field.md) · [password](docs/components/password.md) · [checkbox](docs/components/checkbox.md) · [radio](docs/components/radio.md) · [switch](docs/components/switch.md) · [slider](docs/components/slider.md) · [badge](docs/components/badge.md) · [tooltip](docs/components/tooltip.md) · [popover](docs/components/popover.md) · [modal](docs/components/modal.md) · [drawer](docs/components/drawer.md) · [toast](docs/components/toast.md) · [alert](docs/components/alert.md) · [card](docs/components/card.md) · [tabs](docs/components/tabs.md) · [progress](docs/components/progress.md) · [kbd](docs/components/kbd.md) · [titlebar](docs/components/titlebar.md) · [window](docs/components/window.md) · [skeleton](docs/components/skeleton.md) · [empty-state](docs/components/empty-state.md) · [result](docs/components/result.md) · [pagination](docs/components/pagination.md) · [breadcrumb](docs/components/breadcrumb.md) · [input-number](docs/components/input-number.md) · [search-bar](docs/components/search-bar.md) · [hover-card](docs/components/hover-card.md) · [dropdown-menu](docs/components/dropdown-menu.md) · [context-menu](docs/components/context-menu.md) · [spin](docs/components/spin.md) · [table](docs/components/table.md) · [auto-complete](docs/components/auto-complete.md) · [date-picker](docs/components/date-picker.md) · [carousel](docs/components/carousel.md) · [tag](docs/components/tag.md) · [menubar](docs/components/menubar.md) · [command-palette](docs/components/command-palette.md) · [sortable-list](docs/components/sortable-list.md) · [input-otp](docs/components/input-otp.md) · [upload](docs/components/upload.md) · [rangeslider](docs/components/range-slider.md) · [stepslider](docs/components/step-slider.md) · [stepper](docs/components/stepper.md) · [rating](docs/components/rating.md) · [meter](docs/components/meter.md) · [colorpicker](docs/components/color-picker.md) · [calendar](docs/components/calendar.md) · [timeinput](docs/components/time-input.md) · [editabletext](docs/components/editable-text.md) · [icons](docs/components/icons.md) · [handle](docs/components/handle.md) · [motion](docs/components/motion.md)。工程结论见 [docs/research.md](docs/research.md)，分层与约定见 [docs/architecture.md](docs/architecture.md)，协作规则见 [AGENTS.md](AGENTS.md)。

## 自定义窗口 + 橙色主题

`window.New` 把无框窗口、自绘标题栏（左 logo + 应用名，右窗口按钮）、主题、窗口边框/底色和动作接线一次做完。只需给图标、名字、主题和视图：

```go
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
	"github.com/wuyangwang/my-ui/window"
)

func main() {
	mygo.App.WhenReady(func() {
		window.New(window.Options{
			Title:  "My App",
			Name:   "My App",
			Icon:   icons.Component(), // 也可传自己的 SVG；不传用默认组件标
			Width:  1100,
			Height: 700,
			Theme: func() myui.Options {
				return myui.Options{
					Appearance: myui.AppearanceDark, // System / Light / Dark
					Palette:    myui.PaletteOrange,   // 橙色强调色
				}
			},
			View: func(c *ui.Context) {
				t := c.Theme() // 已经是橙色主题
				ui.Column(c).Fill().Padding(t.Space(6)).Children(func() {
					ui.Text(c, "Hello, my-ui").FontSize(myui.TextSize(t, myui.Text2Xl))
				})
			},
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
```

标题栏风格默认按平台（macOS 红绿灯 / Windows 右侧按钮），也可 `Look: controls.TitleBarLookMac` 强制；想自定义整个品牌区就传 `Logo: func(c *ui.Context) { ... }` 代替 `Icon`。`Theme` 每帧返回当前选项，切换明暗/色板时标题栏、内容和原生窗口边框一起换色。

## 示例

`examples/showcase`：按组件文档站的样式排布——顶栏放品牌、组件搜索（直接过滤侧栏）与版本/明暗切换；侧栏按分组折叠，每个组件一页，标题下写明该组件的用途；变体与状态以紧凑的演示卡排列，可点的组件本身就在卡里。可切换明暗/色板并用 `ui.Themed` 预览子树主题。

```sh
go run ./examples/showcase
```

或用 mygo CLI（与锁定版本一致）：

```sh
go install github.com/egoist/mygo/cmd/mygo@v0.10.0
mygo dev examples/showcase
mygo build -skip-installer examples/showcase
```

## 版本约定

当前版本 **my-ui v0.5.1**（`myui.Version`，git tag `v0.5.0`）。锁定 mygo **v0.10.0**（`go.mod`）；升级 mygo 走专门提交，同步更新本文件与 [AGENTS.md](AGENTS.md)。my-ui 与 mygo 各自独立编号，发布用 git tag 标记。

## 许可

MIT。内嵌资源按其各自许可再分发（Lucide ISC、simple-icons CC0、字体 OFL）。
