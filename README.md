# my-ui

MyGo 原生 UI（[`package ui`](https://github.com/egoist/mygo)）的第三方组件库（独立 module，`go get` 导入即用）。面向 MyGo 自绘的原生 UI（macOS / Linux / Windows，亦支持 wasm），不依赖 webview。

定位：**主题包 + 语义常量 + 变体化控件预设 + 框架没有的组合组件**。`mygo/ui` 已是完整的立即模式控件层；my-ui 在其公开 API 之上补设计系统与桌面组件（shadcn 式 `Variant/Size` 词汇）。

```go
myui.NewButton(c, "Save").Primary().Sm()
myui.NewInput(c, &app.email).Outline().Clearable()
h := myui.Bind(c)          // 每帧绑定一次，之后不带 c
h.Button("Save").Primary()
```

## 包结构

- 根包 `myui`：tokens 刻度、主题预设（`Apply`/`NewTheme`）、变体化控件（`Button`/`Input`/`TextArea`/`Select`/`ComboBox`/`MultiSelect`/`Badge`/`Icon`）、表单字段层（`Field`/`PasswordInput`）、浮层与反馈（`Tooltip`/`Popover`/`Modal`/`Drawer`/`Toast`/`Alert`）、页面骨架（`Scaffold`）、帧句柄（`Bind`）。
- 二级包按需引入，不 import 不链接：
  - `icons/` —— Lucide（ISC）全量 2,318 个图标，函数式 API（`icons.Save()`），未调用的图标不进二进制；
  - `fonts/` —— 内嵌 OFL 字体与 `ui.RegisterFont` 注册；
  - `ai/` —— ChatList、ModelPicker、Reasoning 等重组件（规划中）。

每个组件的调研与契约见 [`docs/components/`](docs/components)：[button](docs/components/button.md) · [input](docs/components/input.md) · [textarea](docs/components/textarea.md) · [select](docs/components/select.md) · [combobox](docs/components/combobox.md) · [multiselect](docs/components/multiselect.md) · [field](docs/components/field.md) · [password](docs/components/password.md) · [badge](docs/components/badge.md) · [tooltip](docs/components/tooltip.md) · [popover](docs/components/popover.md) · [modal](docs/components/modal.md) · [drawer](docs/components/drawer.md) · [toast](docs/components/toast.md) · [alert](docs/components/alert.md) · [scaffold](docs/components/scaffold.md) · [icons](docs/components/icons.md) · [handle](docs/components/handle.md) · [motion](docs/components/motion.md)。工程结论见 [docs/research.md](docs/research.md)，分层与约定见 [docs/architecture.md](docs/architecture.md)，协作规则见 [AGENTS.md](AGENTS.md)。

## 示例

`examples/showcase`：每个组件一页，侧栏按分组折叠并带搜索框；可切换明暗/色板并用 `ui.Themed` 预览子树主题。

```sh
go run ./examples/showcase
```

或用 mygo CLI（与锁定版本一致）：

```sh
go install github.com/egoist/mygo/cmd/mygo@v0.8.3
mygo dev examples/showcase
mygo build -skip-installer examples/showcase
```

## 版本约定

锁定 mygo **v0.8.3**（`go.mod`）；升级 mygo 走专门提交，同步更新本文件与 [AGENTS.md](AGENTS.md)。my-ui 发版与 mygo 版本一一对应（当前未打 tag）。

## 许可

MIT。内嵌资源按其各自许可再分发（Lucide ISC、simple-icons CC0、字体 OFL）。
