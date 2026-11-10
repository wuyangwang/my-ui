# my-ui

MyGo 原生 UI（[`package ui`](https://github.com/egoist/mygo)）的第三方组件库（独立 module，`go get` 导入即用）。面向 MyGo 自绘的原生 UI（macOS / Linux / Windows，亦支持 wasm 构建），不依赖 webview。

定位：**主题包 + 语义常量 + 变体化控件预设 + 框架没有的组合组件**。`mygo/ui` 已是完整的立即模式控件层；my-ui 在其公开 API 之上补设计系统与桌面/AI 应用组件（shadcn 式 `Variant/Size` 预设词汇）。

## 特性目标（对标 shadcn 的词汇与设计系统）

| 方向 | 说明 |
|---|---|
| 多种主题 | 基于系统外观/强调色/高对比/文本缩放派生主题，支持跟随/强制与 `ui.Themed` 子树级主题；详见 [docs/architecture.md](docs/architecture.md) |
| 动画与过渡 | 复用框架弹簧/曲线/AnimatePose/compositor 级动画，库出动效刻度与默认值 |
| 易用 | 一行构造 + 链式预设：`NewButton(c, "Save").Primary().Sm()`，沿用 ui "输入即返回值"绑定约定；也可每帧绑定一次 `h := myui.Bind(c)` 后 `h.Button("Save").Primary()` |
| 常量体系 | 字号/间距/圆角/动效刻度、Palette 色板与取色派生、Lucide 图标集 |
| 细粒度桌面组件 | Card、Kbd、HoverCard、Drawer、Command Palette、Skeleton、EmptyState、Pagination、表格排序等（规划） |
| AI 组件 | ChatList、模型选择器、推理档位等（规划，见二级包 `ai/`） |

## 包结构

- 根包 `myui`：tokens 刻度、主题预设（`Apply`/`NewTheme`）、变体化控件（`Button`/`Input`/`Badge`/`Icon`）、页面骨架（`Scaffold`）、帧句柄（`Bind`）。
- 二级包按需引入，不 import 不链接（遵循 "apps pay for what they use"）：
  - `icons/` —— Lucide（ISC）全量 2,318 个图标，函数式 API（`icons.Save()`），未调用的图标不进二进制；
  - `fonts/`（规划中）—— 内嵌 OFL 字体与 `ui.RegisterFont` 注册；
  - `ai/`（规划中）—— ChatList、ModelPicker、Reasoning 等重组件。

## 当前状态

- **已完成**：tokens 刻度；`Apply`/`NewTheme`（明暗三态 + 8 套 Palette + 系统强调色/高对比/文本缩放）；共享 `Variant`/`Size` 词表（零值 `Native`，`Secondary` 为别名）；`Button`（图标/忙碌态/hover+pressed/禁用）、`Input`（前缀图标如搜索框、后缀、清除按钮、事件、错误态）、`Badge`、`Icon`（Lucide）、`Scaffold`；`icons/` 二级包；`myui.Bind(c)` 每帧句柄；动效约定（状态色渐变、出现渐显）。
- **进行中**：一期后的表单与反馈组件（TextArea、Select、Checkbox、Radio、Switch、Slider、Alert、Tooltip），按 [docs/roadmap.md](docs/roadmap.md) 顺序推进，每个组件含规格书、实现、`ui.Tester` 测试与 showcase 示例。
- **后置**：i18n 文案表、品牌标子包 `icons/brand`、`fonts/`。

每个组件的调研与契约见 [`docs/components/`](docs/components)：[button](docs/components/button.md) · [input](docs/components/input.md) · [badge](docs/components/badge.md) · [scaffold](docs/components/scaffold.md) · [icons](docs/components/icons.md) · [handle](docs/components/handle.md) · [motion](docs/components/motion.md)。工程结论见 [docs/research.md](docs/research.md)，分层与约定见 [docs/architecture.md](docs/architecture.md)，协作规则见 [AGENTS.md](AGENTS.md)。

## 示例

`examples/showcase`：侧栏分页（Button、Input、Badge、Icon、Scaffold、外观），可切换明暗/色板并用 `ui.Themed` 预览子树主题。

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

锁定 mygo **v0.8.3**（`go.mod`）；升级 mygo 走专门提交，同步更新本文件与 [AGENTS.md](AGENTS.md)。my-ui 发版与 mygo 版本一一对应（当前未打 tag，待一期范式定稿）。

## 许可

MIT。内嵌资源按其各自许可再分发（Lucide ISC、simple-icons CC0、字体 OFL）。
