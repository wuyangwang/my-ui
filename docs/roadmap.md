# my-ui 路线图

> 基线 mygo **v0.8.3**（证据见 [research.md](research.md)）。本文是按序实施计划：**每一项交付 = 实现 + `ui.Tester` 测试 + `examples/showcase` 示例页/区块**（适用时）。当前无 tag，开发中不考虑历史兼容。

## 原则

1. **范式先行**：先定稿 tokens/theme 与变体控件范式，后续组件照抄。
2. **顺序推进**：按下表从上到下逐项完成；未完成项不跳到后续阶段宣称"已覆盖"。
3. **差异化后置并行**：`ai/` 等能力在一期范式稳定后推进。

## 一期：范式定稿（当前阶段）

| 顺序 | 项目 | 交付内容（实现 / 测试 / 示例） | 状态 |
|---|---|---|---|
| 1 | tokens 刻度 | `tokens.go`：字号/间距/圆角/动效倍率与 `TextSize/Space/Radius`；随主题/刻度的解析测试；showcase 中通过主题与布局体现 | ✅ 已完成 |
| 2 | 主题与取色 | `theme.go`/`palettes.go`/`shades.go`：`Apply`/`NewTheme`、明暗三态、8 套 Palette、系统强调色/高对比/文本缩放、shades 派生；`theme_test`/`shades_test` 覆盖外观、色板、强调色、文本缩放、高对比、子树隔离；showcase「外观」页 + `ui.Themed` 预览 | ✅ 已完成 |
| 3 | Button | `button.go`：共享 `Variant`/`Size`（零值 `Native`，`Secondary` 为别名）+ `Primary/Outline/Ghost/Destructive/Link` 预设与 hover/pressed/禁用；`button_test.go` 覆盖变体绘制与交互、尺寸、禁用；showcase「Button」页 | ✅ 已完成 |
| 4 | Input | `input.go`：文本输入的变体/尺寸/状态预设（对齐共享 `Variant`/`Size` 语义，含占位/禁用/只读等适用状态），绑定遵循 ui 输入约定；Tester 覆盖变体/尺寸、禁用/只读、暗色主题；showcase 新增「Input」页 | ✅ 已完成 |
| 5 | Badge | `badge.go`：标签/徽标预设（共享 `Variant/Size`）；Tester 覆盖主要变体与明暗主题、尺寸；showcase 新增「Badge」页（含组合示例） | ✅ 已完成 |
| 6 | Scaffold | `scaffold.go`：页面骨架组合（`TitleBar()` 安全区、滚动内容区、底部栏与底部预留/`ScrollbarInsets`）；Tester 覆盖基础布局与标题栏安全区；showcase 提供 Scaffold 骨架示例页 | ✅ 已完成 |

**一期验收（已达成）**：`Button/Input/Badge/Scaffold` 均已交付实现 + `ui.Tester` 测试 + showcase 演示；本地 `go build/vet/test ./...` 通过。跨平台/wasm 矩阵（macOS/Linux/Windows + `GOOS=js GOARCH=wasm go vet`）待在对应环境/CI 验证。

## 一期后：表单与反馈补齐（按序）

在一期范式定稿后，优先补齐高频输入与反馈组件（优先基于 mygo 已有控件做 `Variant/Size/状态` 预设包装）。每项交付 = **实现 + `ui.Tester` 测试 + showcase 示例**。

| 顺序 | 组件 | 要点 | 状态 |
|---|---|---|---|
| 7 | TextArea | 多行输入，`Variant/Size`、占位/禁用/只读、随内容增长（适用时） | ⏳ 未开始 |
| 8 | Select | 单选下拉，触发器与弹层样式预设、选中态、禁用 | ⏳ 未开始 |
| 9 | Checkbox | 勾选/半选态、标签与组布局、禁用与键盘可达 | ⏳ 未开始 |
| 10 | Radio | 单选组、方向与禁用态 | ⏳ 未开始 |
| 11 | Switch | 开关预设与动效跟随主题/Reduced Motion | ⏳ 未开始 |
| 12 | Slider | 单值滑杆（必要时扩展区间），步进/禁用/状态样式 | ⏳ 未开始 |
| 13 | Alert | 语义类型（info/success/warn/error 等映射主题槽）、标题/描述/可关闭 | ⏳ 未开始 |
| 14 | Tooltip | 文本提示预设，延迟/位置封装（基于 ui 能力） | ⏳ 未开始 |

> 后续可按需插入：`InputPassword`、`InputNumber`、`SearchBar`、`Segmented`、`ToggleGroup`、`AutoComplete`、`DatePicker`、`InputOTP`、`Upload`、`Form` 等。

## 二期：导航、展示与桌面细粒度组件（按需排期）

按优先级逐项推进，每项同样要求 **实现 + Tester 测试 + showcase 示例**。

| 类别 | 组件 | 状态 |
|---|---|---|
| 导航/覆盖层 | Tabs、Breadcrumb、Pagination、Drawer(Sheet)、Modal/Dialog、Popover、DropdownMenu、ContextMenu、Accordion/Collapsible、Command Palette、应用内 Menubar | ⏳ 未开始（Pagination/Drawer 已在早期候选中） |
| 展示/布局 | Card/Block、Divider、Avatar、Tag、Kbd/Hotkey、Skeleton、EmptyState、Result、Progress、Spin、HoverCard、Carousel、AspectRatio | ⏳ 未开始（部分已在早期候选中） |
| 数据/复杂 | Table（样式与排序/过滤预设）、SortableList 等 | ⏳ 未开始 |

## 三期：AI 二级包（`ai/`）

`ChatList`（虚拟化流式气泡、贴底跟随等）、`ModelPicker`、`Reasoning`。依赖一期范式与页面骨架形态，可与二期部分并行；每项含实现、测试与示例。当前 `ai/` 仅占位。

## 后置（随需推进）

| 项目 | 说明 | 状态 |
|---|---|---|
| i18n 文案表 | 根包内置 `en`/`zh-Hans`/`zh-Hant` 小消息表 + `c.Locale()` 回退链 + app 覆盖（控件内固定文案） | ⏳ 后置 |
| icons 二级包 | Lucide + 品牌标 typed `*ui.SVG` 注册表与 `go generate` 管线（`icons/`） | ⏳ 后置，规划中无代码 |
| fonts 二级包 | OFL 字体内嵌与 `ui.RegisterFont` 注册（`fonts/`） | ⏳ 后置，规划中无代码 |

## 依赖上游

- `Element.Access` 的 internal 类型泄漏（见 [research.md](research.md) §五）：修复前不实现依赖该 API 的无障碍定制。
- `ui.Themed` / `SetTheme` 语义若上游变动，随锁版升级调整。

## 明确不做

- CEF 窗口路径（超出当前范围）。
- 通过 npm/Bun 发布（纯 Go module）。
