# my-ui 路线图

> 基线 mygo **v0.8.3**（证据见 [research.md](research.md)）。本文是按序实施计划：**每一项交付 = 规格书 + 实现 + `ui.Tester` 测试 + `examples/showcase` 示例页/区块**（适用时）。当前无 tag，开发中不考虑历史兼容。

## 原则

1. **先调研后编码**：每个组件先写规格书 [`docs/components/<name>.md`](components)（参考实现 → 能否实现 → 实现到何种程度 → 结构评估 → API 契约），定稿后才写代码。
2. **范式先行**：先定稿 tokens/theme 与变体控件范式，后续组件照抄。
3. **顺序推进**：按下表从上到下逐项完成；未完成项不跳到后续阶段宣称已覆盖。
4. **差异化后置并行**：`ai/` 等能力在一期范式稳定后推进。

## 一期：范式定稿（当前阶段）

| 顺序 | 项目 | 交付内容（规格书 / 实现 / 测试 / 示例） | 状态 |
|---|---|---|---|
| 1 | tokens 刻度 | `tokens.go`：字号/间距/圆角/动效倍率与 `TextSize/Space/Radius`；随主题/刻度的解析测试；showcase 中通过主题与布局体现 | ✅ 已完成 |
| 2 | 主题与取色 | `theme.go`/`palettes.go`/`shades.go`：`Apply`/`NewTheme`、明暗三态、8 套 Palette、系统强调色/高对比/文本缩放、shades 派生；`theme_test`/`shades_test` 覆盖外观、色板、强调色、文本缩放、高对比、子树隔离；showcase「外观」页 + `ui.Themed` 预览 | ✅ 已完成 |
| 3 | Button | 规格书 [components/button.md](components/button.md)；`button.go`：共享 `Variant`/`Size`（零值 `Native`，`Secondary` 为别名）+ `Primary/Outline/Ghost/Destructive/Link` 预设与 hover/pressed/禁用；`button_test.go` 覆盖变体绘制与交互、尺寸、禁用；showcase「Button」页 | ✅ 已完成 |
| 4 | Input | 规格书 [components/input.md](components/input.md)；`input.go`：复合结构（外框 + 前后缀/清除槽位 + 内层输入）、`Variant`/`Size`、清除按钮（`icons.X()` + `KeepFocus` + `TabStop(false)`）、事件（变化/失焦/清除/回车）、错误态；`input_test.go` 覆盖尺寸、禁用/只读、暗色、清除、事件、错误文案、前后缀；showcase「Input」页 | ✅ 已完成 |
| 5 | Badge | 规格书 [components/badge.md](components/badge.md)；`badge.go`：标签/徽标预设（共享 `Variant/Size`）；Tester 覆盖主要变体与明暗主题、尺寸；showcase「Badge」页（含组合示例） | ✅ 已完成 |
| 6 | Scaffold | 规格书 [components/scaffold.md](components/scaffold.md)；`scaffold.go`：页面布局骨架（`TitleBar()` 安全区、顶栏、滚动内容区、底栏与底部预留/`ScrollbarInsets`）；Tester 覆盖基础布局与标题栏安全区；showcase「Scaffold」页 | ✅ 已完成 |
| 7 | Icon | 规格书 [components/icons.md](components/icons.md)；`icon.go`：尺寸档（14/20/24 DIP）、语义色、`Label` 无障碍名、`Spin` 旋转；`icons/` 二级包（Lucide 全量 2,318 个图标 + 14 个别名，函数式 API，未用不进二进制）；`icons/icons_test.go` + `icon_test.go`；showcase「Icon」页 | ✅ 已完成 |
| 8 | 帧句柄 | 规格书 [components/handle.md](components/handle.md)；`bind.go`：`myui.Bind(c)` 每帧绑定一次，组件不再逐个传 `c`；不引入包级状态 | ✅ 已完成 |
| 9 | 动效约定 | 规格书 [components/motion.md](components/motion.md)；`motion.go`：状态色渐变（`Element.Transition`）与出现渐显（`AnimatePose`）；Button/Input/Badge 用上；时长取自动效刻度，Reduce Motion 由框架兜底 | ✅ 已完成 |

**一期验收（已达成）**：`Button/Input/Badge/Scaffold/Icon` 均已交付实现 + `ui.Tester` 测试 + showcase 演示；本地 `go build/vet/test ./...` 通过。跨平台/wasm 矩阵（macOS/Linux/Windows + `GOOS=js GOARCH=wasm go vet`）待在对应环境/CI 验证。

## 组件打磨（参考实现对照驱动）

对已实现组件逐个核对 shadcn / element-plus / antd / Lobe MyGo 后定的改进项（见 [docs/components/](components) 各页）：

| 项 | 内容 | 状态 |
|---|---|---|
| 词表重排 | `Default`（零值，控件自带观感）/ `Secondary`（弱填充面）/ `Native`（`Default` 旧名）；`Success`/`Info`/`Warning` 语义色（主题层，明暗各一套） | ✅ 已完成 |
| 尺寸与形态 | 尺寸加 `SizeXs`（4 档）；新增 `Shape` 轴（`ShapeRect`/`ShapeRound`/`ShapeCircle`） | ✅ 已完成 |
| 宽度与占满 | `Width(w)`（Button/Input/TextArea/Select）、`Block()`（Button） | ✅ 已完成 |
| 清除按钮 | 有值即出现、**进 Tab 键序**（antd 做法）、点击保持焦点 | ✅ 已完成 |
| 表单字段层 | 规格书 [components/field.md](components/field.md)；`form.go` + `rules.go`：`NewField(c, 标签, 造控件)` 容器式字段（标签可点击聚焦控件、必填星号、说明 `DescribedBy`、错误 `role=alert` + 控件 `Invalid`）；**校验层**：`Validate(value, rules...)`（被离开后与值变化时检查）+ `Required`/`MinLength`/`MaxLength`/`Pattern`/`Email` 规则 + `Touched`/`Dirty`/`Valid` 状态 | ✅ 已完成 |
| 密码输入 | 规格书 [components/password.md](components/password.md)；`password.go`：独立 `PasswordInput`，眼睛按钮切换明暗（mygo 无公开的取消隐藏 API → 两个 key 切换，值不丢、编辑状态重置） | ✅ 已完成 |
| 字数限制与计数 | `MaxLength(n)`（mygo 原生 ✓）+ `Count()`（按 rune 计，显示 "12/100" ✓）；Input/TextArea 共用 | ✅ 已完成 |
| Select 分组 | 选项带 `Group`，弹层按组给标题 ✓（element-plus 的 `el-option-group` ✓） | ✅ 已完成 |
| ComboBox（可搜索下拉） | 规格书 [components/combobox.md](components/combobox.md)；`combobox.go`：触发器就是输入框（mygo 的 `ui.ComboboxBase` ✓），输入即过滤（按 label ✓），选中即取替文字；变体/尺寸/分组/打勾/宽度同 Select ✓ | ✅ 已完成 |
| MultiSelect（多选） | 值变成切片 + 标签形态 | ⏳ 待做 |
| 已知限制 | 清除按钮嵌在触发器内 → mygo 的无障碍树不展开 leaf-role 子节点（`ui/access.go:772`），读屏摸不到；密码显隐受限于 mygo；聚焦即校验做不了（无 focus-gained 事件） | 📌 已记录 |

## 一期后：表单与反馈补齐（按序）

在一期范式定稿后，优先补齐高频输入与反馈组件（优先基于 mygo 已有控件做 `Variant/Size/状态` 预设包装）。每项交付 = **规格书 + 实现 + `ui.Tester` 测试 + showcase 示例**。

| 顺序 | 组件 | 要点 | 状态 |
|---|---|---|---|
| 10 | TextArea | 规格书 [components/textarea.md](components/textarea.md)；`textarea.go`：与 Input 同构（共用 `field[T]` 基座），多行文本域，随内容增高（`Lines(min,max)`）、清除/前后缀/错误态/事件；`textarea_test.go` 覆盖尺寸、增高、事件、禁用只读、错误文案、清除；showcase「TextArea」页 | ✅ 已完成 |
| 11 | Select | 规格书 [components/select.md](components/select.md)；`select.go`：下拉选择，选项为 `{Label,Value,Disabled}` 结构（标签与值可不同），触发器与 Input 同外观，弹层面板与选项高亮、选中打勾、占位文字、禁用；`select_test.go` 覆盖标签/值、开合选择、占位、禁用、尺寸；showcase「Select」页 | ✅ 已完成 |
| 12 | Checkbox | 勾选/半选态、标签与组布局、禁用与键盘可达 | ⏳ 未开始 |
| 13 | Radio | 单选组、方向与禁用态 | ⏳ 未开始 |
| 14 | Switch | 开关预设与动效跟随主题/Reduced Motion | ⏳ 未开始 |
| 15 | Slider | 单值滑杆（必要时扩展区间），步进/禁用/状态样式 | ⏳ 未开始 |
| 16 | Alert | 语义类型（info/success/warn/error 等映射主题槽）、标题/描述/可关闭 | ⏳ 未开始 |
| 17 | Tooltip | 文本提示预设，延迟/位置封装（基于 ui 能力） | ⏳ 未开始 |

> 后续可按需插入：`InputPassword`（含显隐切换）、`InputNumber`、`SearchBar`、`Segmented`、`ToggleGroup`、`AutoComplete`、`DatePicker`、`InputOTP`、`Upload`、`Form` 等。

## 二期：导航、展示与桌面细粒度组件（按需排期）

按优先级逐项推进，每项同样要求 **规格书 + 实现 + Tester 测试 + showcase 示例**。

| 类别 | 组件 | 状态 |
|---|---|---|
| 导航/覆盖层 | Tabs、Breadcrumb、Pagination、Drawer(Sheet)、Modal/Dialog、Popover、DropdownMenu、ContextMenu、Accordion/Collapsible、Command Palette、应用内 Menubar | ⏳ 未开始 |
| 展示/布局 | Card/Block、Divider、Avatar、Tag、Kbd/Hotkey、**Skeleton（骨架屏：加载占位；与 Scaffold 的页面布局骨架不是一回事）**、EmptyState、Result、Progress、Spin、HoverCard、Carousel、AspectRatio | ⏳ 未开始 |
| 数据/复杂 | Table（样式与排序/过滤预设）、SortableList 等 | ⏳ 未开始 |

## 三期：AI 二级包（`ai/`）

`ChatList`（虚拟化流式气泡、贴底跟随等）、`ModelPicker`、`Reasoning`。依赖一期范式与页面骨架形态，可与二期部分并行；每项含规格书、实现、测试与示例。当前 `ai/` 仅占位。

## 后置（随需推进）

| 项目 | 说明 | 状态 |
|---|---|---|
| i18n 文案表 | 根包内置 `en`/`zh-Hans`/`zh-Hant` 小消息表 + `c.Locale()` 回退链 + app 覆盖（控件内固定文案） | ⏳ 后置 |
| 品牌标 | simple-icons（CC0）→ `icons/brand` 子包，与 Lucide 分开 | ⏳ 后置 |
| fonts 二级包 | OFL 字体内嵌与 `ui.RegisterFont` 注册（`fonts/`） | ⏳ 后置，规划中无代码 |

## 依赖上游

- `Element.Access` 的 internal 类型泄漏（见 [research.md](research.md) §五）：修复前不实现依赖该 API 的无障碍定制。
- `ui.Themed` / `SetTheme` 语义若上游变动，随锁版升级调整。

## 明确不做

- CEF 窗口路径（超出当前范围）。
- 通过 npm/Bun 发布（纯 Go module）。
