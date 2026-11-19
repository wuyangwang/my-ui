# my-ui 路线图

> v0.4.0 起组件按功能分组到 `controls/`、`forms/`、`overlays/`、`display/` 子包，根包只留主题/tokens 底座；`myui.NewXxx` 调用需按分组改 import。v0.7.0 按 Lobe MyGo 对照补齐约 30 个组件并优化 Alert/Badge/Tag/Progress/Skeleton/Spin/Tabs/InputOTP/Select/Tooltip/Popover；Progress 量纲改为 0..100，InputOTP 提供 `NewOTP`。v0.6.0 统一 fluent 词表（`Variant/Size/Width/Label/Disabled/OnChanged`），移除重复别名（`Toast.Type`、`Result.Extra`），`bind.Bind` 覆盖全部组件；属破坏性 API 变更。


> 基线 mygo **v0.10.0**（证据见 [research.md](research.md)）。本文是按序实施计划：**每一项交付 = 规格书 + 实现 + `ui.Tester` 测试 + `examples/showcase` 示例页/区块**（适用时）。首个 tag 为 `v0.2.0`，开发中不考虑更早版本的历史兼容。

## 原则

1. **先调研后编码**：每个组件先写规格书 [`docs/components/<name>.md`](components)（参考实现 → 能否实现 → 实现到何种程度 → 结构评估 → API 契约），定稿后才写代码。
2. **范式先行**：先定稿 tokens/theme 与变体控件范式，后续组件照抄。
3. **顺序推进**：按下表从上到下逐项完成；未完成项不跳到后续阶段宣称已覆盖。
4. **差异化后置并行**：`ai/` 等能力在一期范式稳定后推进。

## 一期：范式定稿

| 顺序 | 项目 | 交付内容（规格书 / 实现 / 测试 / 示例） | 状态 |
|---|---|---|---|
| 1 | tokens 刻度 | `tokens.go`：字号/间距/圆角/动效倍率与 `TextSize/Space/Radius`；随主题/刻度的解析测试；showcase 中通过主题与布局体现 | ✅ 已完成 |
| 2 | 主题与取色 | `theme.go`/`palettes.go`/`shades.go`：`Apply`/`NewTheme`、明暗三态、8 套 Palette、系统强调色/高对比/文本缩放、shades 派生；`theme_test`/`shades_test` 覆盖外观、色板、强调色、文本缩放、高对比、子树隔离；showcase「外观」页 + `ui.Themed` 预览 | ✅ 已完成 |
| 3 | Button | 规格书 [components/button.md](components/button.md)；`button.go`：共享 `Variant`/`Size` + `Primary/Outline/Ghost/Destructive/Link` 预设与 hover/pressed/禁用 + 图标/忙碌态；`button_test.go` 覆盖变体绘制与交互、尺寸、禁用、图标、忙碌态；showcase「Button」页 | ✅ 已完成 |
| 4 | Input | 规格书 [components/input.md](components/input.md)；`input.go`：复合结构（外框 + 前后缀/清除/计数槽位 + 内层输入）、清除按钮（`icons.X()` + `KeepFocus`）、事件（变化/失焦/清除/回车）、错误态、`MaxLength`/`Count`；`input_test.go` 覆盖尺寸、禁用/只读、暗色、清除、事件、错误文案、前后缀、计数；showcase「Input」页 | ✅ 已完成 |
| 5 | Badge | 规格书 [components/badge.md](components/badge.md)；`badge.go`：标签/徽标预设（共享 `Variant/Size`）；Tester 覆盖主要变体与明暗主题、尺寸；showcase「Badge」页（含组合示例） | ✅ 已完成 |
| 7 | Icon | 规格书 [components/icons.md](components/icons.md)；`icon.go`：尺寸档（14/20/24 DIP）、语义色、`Label` 无障碍名、`Spin` 旋转；`icons/` 二级包（Lucide 全量 2,318 个图标 + 14 个别名，函数式 API，未用不进二进制）；`icons/icons_test.go` + `icon_test.go`；showcase「Icon」页 | ✅ 已完成 |
| 8 | 帧句柄 | 规格书 [components/handle.md](components/handle.md)；`bind.go`：`bindpkg.Bind(c)` 每帧绑定一次，组件不再逐个传 `c`；不引入包级状态 | ✅ 已完成 |


## 组件打磨（参考实现对照驱动）

对已实现组件逐个核对 shadcn / element-plus / antd / Lobe MyGo 后定的改进项（见 [docs/components/](components) 各页）：

| 项 | 内容 | 状态 |
|---|---|---|
| 词表重排 | `Default`（零值，控件自带观感）/ `Secondary`（弱填充面）/ `Native`（`Default` 旧名）；`Success`/`Info`/`Warning` 语义色（主题层，明暗各一套） | ✅ 已完成 |
| 尺寸与形态 | 尺寸加 `SizeXs`（4 档）；新增 `Shape` 轴（`ShapeRect`/`ShapeRound`/`ShapeCircle`） | ✅ 已完成 |
| 宽度与占满 | `Width(w)`（Button/Input/TextArea/Select/ComboBox/MultiSelect）、`Block()`（Button）；触发器不铺满整行 | ✅ 已完成 |
| 清除按钮 | 有值即出现、**进 Tab 键序**（antd 做法）、点击保持焦点 | ✅ 已完成 |
| 动效约定 | 规格书 [components/motion.md](components/motion.md)；`motion.go`：状态色渐变（`Element.Transition`）；时长取自动效刻度，Reduce Motion 由框架兜底；姿态动画（`AnimatePose`）实测不可用（元素停在不可见），故不用 | ✅ 已完成 |
| 表单字段层 | 规格书 [components/field.md](components/field.md)；`form.go` + `rules.go`：`NewField(c, 标签, 造控件)` 容器式字段（标签可点击聚焦控件、必填星号、说明 `DescribedBy`、错误 `role=alert` + 控件 `Invalid`）；**校验层**：`Validate(value, rules...)`（被离开后与值变化时检查）+ `Required`/`MinLength`/`MaxLength`/`Pattern`/`Email` 规则 + `Touched`/`Dirty`/`Valid` 状态 | ✅ 已完成 |
| 密码输入 | 规格书 [components/password.md](components/password.md)；`password.go`：独立 `PasswordInput`，眼睛按钮切换明暗（mygo 无公开的取消隐藏 API → 两个 key 切换，值不丢、编辑状态重置），切换后焦点回输入框 | ✅ 已完成 |
| 字数限制与计数 | `MaxLength(n)`（mygo 原生 ✓）+ `Count()`（按 rune 计，显示 "12/100" ✓）；Input/TextArea 共用 | ✅ 已完成 |
| Select 分组 | 选项带 `Group`，弹层按组给标题 ✓（element-plus 的 `el-option-group` ✓） | ✅ 已完成 |
| ComboBox（可搜索下拉） | 规格书 [components/combobox.md](components/combobox.md)；`combobox.go`：触发器就是输入框（mygo 的 `ui.ComboboxBase` ✓），输入即过滤（按 label ✓），选中即取替文字；变体/尺寸/分组/打勾/宽度同 Select ✓ | ✅ 已完成 |
| MultiSelect（多选） | 规格书 [components/multiselect.md](components/multiselect.md)；`multiselect.go`：包 mygo 的 `ui.TokenField`（芯片 + 建议列表 ✓），选中项以芯片呈现、各带移除按钮；值即标签（框架限制 ✗） | ✅ 已完成 |
| 弹层与反馈 | 规格书 [components/tooltip.md](components/tooltip.md) · [components/popover.md](components/popover.md) · [components/modal.md](components/modal.md) · [components/toast.md](components/toast.md) · [components/alert.md](components/alert.md) · [components/drawer.md](components/drawer.md)；`tooltip.go`（提示挂在锚点上、外观照 shadcn、锚点的 description 供读屏）、`popover.go`（点击锚点打开、外面点击/Escape 关闭）、`modal.go`（模态对话框 + 背景罩 + 右上角关闭按钮）、`toast.go`（右下角的通知视口 + 卡片）、`alert.go`（页面里的消息，`role=alert`）、`drawer.go`（从窗口边缘来的面板）；测试 21 组 + showcase「弹层」「Alert」「Toast」「Drawer」页 | ✅ 已完成 |
| 已知限制 | 清除按钮嵌在触发器内 → mygo 的无障碍树不展开 leaf-role 子节点（`ui/access.go:772`），读屏摸不到；密码显隐受限于 mygo；聚焦即校验做不了（无 focus-gained 事件）；弹层内容由框架接管（装不下自定义头尾）；ComboBox 无清除按钮；Tooltip 无小箭头（框架没有）；Popover 只在点击时打开、Modal 不可拖拽/全屏（框架没有）；Toast 无淡入、Drawer 无滑入（框架的姿态动画不可用 ✗） | 📌 已记录 |
| 示例应用外壳 | 组件文档站式布局：顶栏放品牌、组件搜索（过滤侧栏）与版本/明暗切换；侧栏按分组折叠（一次开一组），每页一个图标，选中项为白底药丸 + 左侧强调色条；标题下是该组件的真实说明，变体/状态排成紧凑的演示卡（左侧名称、右侧可交互的组件本身） | ✅ 已完成 |

## 一期后：表单与反馈补齐（按序）

在一期范式定稿后，优先补齐高频输入与反馈组件（优先基于 mygo 已有控件做 `Variant/Size/状态` 预设包装）。每项交付 = **规格书 + 实现 + `ui.Tester` 测试 + showcase 示例**。

| 顺序 | 组件 | 要点 | 状态 |
|---|---|---|---|
| 10 | TextArea | 规格书 [components/textarea.md](components/textarea.md)；`textarea.go`：与 Input 同构（共用 `field[T]` 基座），多行文本域，随内容增高（`Lines(min,max)`）、清除/前后缀/错误态/事件、`MaxLength`/`Count`；`textarea_test.go` 覆盖尺寸、增高、事件、禁用只读、错误文案、清除；showcase「TextArea」页 | ✅ 已完成 |
| 11 | Select | 规格书 [components/select.md](components/select.md)；`select.go`：下拉选择，选项为 `{Label,Value,Disabled}` 结构（标签与值可不同），触发器与 Input 同外观，弹层面板与选项高亮、选中打勾、占位文字、禁用；`select_test.go` 覆盖标签/值、开合选择、占位、禁用、尺寸；showcase「Select」页 | ✅ 已完成 |
| 12 | Checkbox | 规格书 [components/checkbox.md](components/checkbox.md)；`checkbox.go`：勾选框（方框 + 圆角 + 对勾），四档尺寸、事件、禁用 | ✅ 已完成 |
| 13 | Radio | 规格书 [components/radio.md](components/radio.md)；`radio.go`：单选（圆环 + 圆点），泛型值，四档尺寸、事件、禁用 | ✅ 已完成 |
| 14 | Switch | 规格书 [components/switch.md](components/switch.md)；`switch.go`：开关（胶囊 + 旋钮），旋钮走位（`Animate`）、事件、无障碍名、禁用 | ✅ 已完成 |
| 15 | Slider | 规格书 [components/slider.md](components/slider.md)；`slider.go`：滑杆（轨道 + 旋钮），步进、事件、无障碍名、禁用 | ✅ 已完成 |
| 16 | Alert | 规格书 [components/alert.md](components/alert.md)；`alert.go`：页面里的一段消息（卡片 + 图标 + 标题 + 描述），变体给墨色（破坏性/成功/信息/警告），`role=alert` 让读屏念出来；`alert_test.go` 覆盖标题描述/live region/图标 | ✅ 已完成 |
| 17 | Toast | 规格书 [components/toast.md](components/toast.md)；`toast.go`：窗口右下角的通知（视口 + 卡片），类型给图标与颜色，动作与关闭按钮，停留时间与淡出；`toast_test.go` 覆盖显示/关闭/动作/超时；showcase「Toast」页 | ✅ 已完成 |
| 18 | Drawer | 规格书 [components/drawer.md](components/drawer.md)；`drawer.go`：从窗口边缘来的一块面板（用 `ui.Overlay` 搭），触发器打开、点外面或 Escape 收起，模态（焦点在里面）；`drawer_test.go` 覆盖开合与位置；showcase「Drawer」页 | ✅ 已完成 |
| 19 | TitleBar | 规格书 [components/titlebar.md](components/titlebar.md)；`titlebar.go`：自绘标题栏，Windows 风格（左 logo+应用名、右窗口按钮、关闭键悬停危险色）或 macOS 风格（左红绿灯、未激活转灰），`Look` 选项选择（Auto 按平台）；整条 `DragWindow`、最大化按钮打 `MaximizeButton`；`titlebar_test.go` 覆盖两种风格；showcase「TitleBar」页 | ✅ 已完成 |
| 20 | Window（无框窗口外壳） | 规格书 [components/window.md](components/window.md)；`window/` 子包：`window.New` 一行建窗（`TitleBarHidden + HideWindowButtons` + `TitleBar` + 每帧主题 + 动作接线），返回 `*mygo.Window`；`window_test.go` 覆盖布局/主题/动作触发；showcase「Window」页点按钮开演示窗口 | ✅ 已完成 |

| 21 | Skeleton | 规格书 [components/skeleton.md](components/skeleton.md)；`display/skeleton.go`：加载占位条，`Circle()` 头像、`Width(w)`、`Size(sz)`；`skeleton_test.go` 覆盖绘制与变体；showcase「状态件」页 | ✅ 已完成 |
| 22 | EmptyState | 规格书 [components/empty-state.md](components/empty-state.md)；`display/empty_state.go`：图标 + 标题 + 说明 + 动作槽；`empty_state_test.go`；showcase「状态件」页 | ✅ 已完成 |
| 23 | Result | 规格书 [components/result.md](components/result.md)；`display/result.go`：四种状态的图标与语义色 + 标题/说明/动作；`result_test.go`；showcase「状态件」页 | ✅ 已完成 |
| 24 | Pagination | 规格书 [components/pagination.md](components/pagination.md)；`display/pagination.go`：页码指针绑定、省略号窗口、上一页/下一页图标；`pagination_test.go`；showcase「状态件」页 | ✅ 已完成 |

| 25 | Breadcrumb | 规格书 [components/breadcrumb.md](components/breadcrumb.md)；`display/breadcrumb.go`：层级路径，当前项强调墨色、其余可点写回索引；`breadcrumb_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 26 | InputNumber | 规格书 [components/input-number.md](components/input-number.md)；`forms/input_number.go`：mygo `NumberInput` + 共享 Variant/Size/Width；`input_number_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 27 | SearchBar | 规格书 [components/search-bar.md](components/search-bar.md)；`forms/search_bar.go`：共享 Frame 基座 + 放大镜 + 清除图标；`search_bar_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 28 | HoverCard | 规格书 [components/hover-card.md](components/hover-card.md)；`overlays/hover_card.go`：`ui.PopoverBase` + hover 判定与帧数宽限；`hover_card_test.go`；showcase「更多组件」页 | ✅ 已完成 |

| 29 | DropdownMenu | 规格书 [components/dropdown-menu.md](components/dropdown-menu.md)；`overlays/dropdown_menu.go` + `menu_common.go`：包 mygo `MenuBase`，shadcn 面板/条目外观；`dropdown_menu_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 30 | ContextMenu | 规格书 [components/context-menu.md](components/context-menu.md)；`overlays/context_menu.go`：包 mygo `ContextMenuBase`，共用菜单条目外观；`context_menu_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 31 | Spin | 规格书 [components/spin.md](components/spin.md)；`display/spin.go`：`NewIcon(...).Spin()`，尺寸/Variant 预设，动效由框架兜底；`spin_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 32 | Table | 规格书 [components/table.md](components/table.md)；`display/table.go`：包 mygo `Table`（ListState/排序/选择仍由框架提供），shadcn 边框与弱墨列头容器；`table_test.go`；showcase「更多组件」页 | ✅ 已完成 |

| 33 | AutoComplete | 规格书 [components/auto-complete.md](components/auto-complete.md)；`forms/auto_complete.go`：包 mygo `Autocomplete` + 共享 Variant/Size/Width；`auto_complete_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 34 | DatePicker | 规格书 [components/date-picker.md](components/date-picker.md)；`forms/date_picker.go`：包 mygo `DateInput` + 共享外观；`date_picker_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 35 | Carousel | 规格书 [components/carousel.md](components/carousel.md)；`display/carousel.go`：索引指针绑定、左右图标与圆点；`carousel_test.go`；showcase「更多组件」页 | ✅ 已完成 |
| 36 | Tag | 规格书 [components/tag.md](components/tag.md)；`display/tag.go`：轻量标签、语义弱底、可关闭图标；`tag_test.go`；showcase「更多组件」页 | ✅ 已完成 |

| 37 | Menubar | 规格书 [components/menubar.md](components/menubar.md)；`overlays/menubar.go`：多个 mygo `MenuBase` 共用菜单外观；`menubar_test.go`；showcase「Menubar」页 | ✅ 已完成 |
| 38 | Command Palette | 规格书 [components/command-palette.md](components/command-palette.md)；`overlays/command_palette.go`：`DialogBase` + 输入过滤命令；`command_palette_test.go`；showcase「Command Palette」页 | ✅ 已完成 |

| 39 | SortableList | 规格书 [components/sortable-list.md](components/sortable-list.md)；`display/sortable_list.go`：包 mygo `ListState.Reorder`，每行一个拖拽手柄；`sortable_list_test.go`；showcase「SortableList」页 | ✅ 已完成 |
| 40 | InputOTP | 规格书 [components/input-otp.md](components/input-otp.md)；`forms/input_otp.go`：一格一字符、自动跳格与退格回退；`input_otp_test.go`；showcase「InputOTP」页 | ✅ 已完成 |
| 41 | Upload | 规格书 [components/upload.md](components/upload.md)；`forms/upload.go`：拖放区 + `DroppedFiles()` 路径回调（不做文件选择器，见限制）；`upload_test.go`；showcase「Upload」页 | ✅ 已完成 |

| 42 | RangeSlider | 双端滑杆，包 mygo RangeSlider + 共享圆角；包 mygo 构造器并加 shadcn 外观；测试与 showcase「RangeSlider」页 | ✅ 已完成 |
| 43 | StepSlider | 带刻度的滑杆，包 mygo StepSlider；包 mygo 构造器并加 shadcn 外观；测试与 showcase「StepSlider」页 | ✅ 已完成 |
| 44 | Stepper | 上下步进器，带 Variant/Size/Width；包 mygo 构造器并加 shadcn 外观；测试与 showcase「Stepper」页 | ✅ 已完成 |
| 45 | Rating | 星级评分，包 mygo Rating；包 mygo 构造器并加 shadcn 外观；测试与 showcase「Rating」页 | ✅ 已完成 |
| 46 | Meter | 量度计，包 mygo Meter 的状态色；包 mygo 构造器并加 shadcn 外观；测试与 showcase「Meter」页 | ✅ 已完成 |
| 47 | ColorPicker | 颜色色块，包 mygo ColorPicker 并加字段边框；包 mygo 构造器并加 shadcn 外观；测试与 showcase「ColorPicker」页 | ✅ 已完成 |
| 48 | Calendar | 月视图日历，包 mygo Calendar + 卡片外观；包 mygo 构造器并加 shadcn 外观；测试与 showcase「Calendar」页 | ✅ 已完成 |
| 49 | TimeInput | 时间输入，包 mygo TimeInput + 字段外观；包 mygo 构造器并加 shadcn 外观；测试与 showcase「TimeInput」页 | ✅ 已完成 |
| 50 | EditableText | 就地重命名文本，包 mygo EditableText；包 mygo 构造器并加 shadcn 外观；测试与 showcase「EditableText」页 | ✅ 已完成 |

| 51 | Accordion | 包 mygo 构造器并加 my-ui 主题外观；规格书 / 测试 / showcase「Accordion」页 | ✅ 已完成 |
| 52 | Collapsible | 包 mygo 构造器并加 my-ui 主题外观；规格书 / 测试 / showcase「Collapsible」页 | ✅ 已完成 |
| 53 | DecimalInput | 包 mygo 构造器并加 my-ui 主题外观；规格书 / 测试 / showcase「DecimalInput」页 | ✅ 已完成 |
| 54 | SplitButton | 包 mygo 构造器并加 my-ui 主题外观；规格书 / 测试 / showcase「SplitButton」页 | ✅ 已完成 |
| 55 | Toggle | 包 mygo 构造器并加 my-ui 主题外观；规格书 / 测试 / showcase「Toggle」页 | ✅ 已完成 |
| 56 | ColorWell | 包 mygo 构造器并加 my-ui 主题外观；规格书 / 测试 / showcase「ColorWell」页 | ✅ 已完成 |

> 后续可按需插入：`InputNumber`、`SearchBar`、`Segmented`、`ToggleGroup`、`AutoComplete`、`DatePicker`、`InputOTP`、`Upload`、`Form` 等。

## 二期：导航、展示与桌面细粒度组件（按需排期）

按优先级逐项推进，每项同样要求 **规格书 + 实现 + Tester 测试 + showcase 示例**。

| 类别 | 组件 | 状态 |
|---|---|---|
| 展示/布局 | Skeleton ✅、EmptyState ✅、Result ✅、Pagination ✅；Divider/Avatar（mygo 自带，`display/` 重导出 ✓）；Skeleton、HoverCard、Carousel、AspectRatio 等 | ✅ 已完成（Skeleton/EmptyState/Result/Pagination） |
| 导航/覆盖层 | Breadcrumb ✅、HoverCard ✅、DropdownMenu ✅、ContextMenu ✅、Menubar ✅、Command Palette ✅；Tabs 已做 | ✅ 已完成（六项） |
| 数据/复杂 | Table ✅、SortableList ✅ | ✅ 已完成（两项） |

## 三期：AI 二级包（`ai/`）

`ChatList`（虚拟化流式气泡、贴底跟随等）、`ModelPicker`、`Reasoning`。依赖一期范式与页面骨架形态，可与二期部分并行；每项含规格书、实现、测试与示例。

| 组件 | 规格书 | 实现 | 测试 | 示例 | 状态 |
|---|---|---|---|---|---|
| ChatList | [chatlist.md](components/chatlist.md) | `ai/chatlist.go` | `ai/chatlist_test.go` | showcase「ChatList」 | ✅ |
| ModelPicker | [model-picker.md](components/model-picker.md) | `ai/modelpicker.go` | `ai/modelpicker_test.go` | showcase「ModelPicker」 | ✅ |
| Reasoning | [reasoning.md](components/reasoning.md) | `ai/reasoning.go` | `ai/reasoning_test.go` | showcase「Reasoning」 | ✅ |

三项都提供 `bind.Bind(c)` 句柄方法。调研 shadcn AI Elements / Vercel AI Elements / assistant-ui 后，Prompt Input、Tool、Task、Sources 等只是普通控件的组合，本库已有等价件，暂不单列。

## 后置（随需推进）

| 项目 | 说明 | 状态 |
|---|---|---|
| i18n 文案表 | 根包内置 `en`/`zh-Hans`/`zh-Hant` 小消息表 + `c.Locale()` 回退链 + app 覆盖（控件内固定文案）：`i18n.go` 的 `Message(c, key)` 与 `SetMessages(locale, msgs)`；控件里的固定文案（清除、关闭、显示/隐藏密码）走它 | ✅ 已完成 |
| 品牌标 | simple-icons（CC0）→ `icons/brand` 子包，与 Lucide 分开：3,462 个品牌标，函数式 API，名字取品牌自己的拼写 | ✅ 已完成 |
| 整套图标的名字清单 | `icons/all` 子包：`Names` 与 `ByName(name)`，给图标选择器用（带上整套图标，只有要整套时才 import） | ✅ 已完成 |
| fonts 二级包 | OFL 字体内嵌与 `ui.RegisterFont` 注册（`fonts/`） | ⏸ 暂不做：用系统自带的字体（`fonts/README.md` 保留规划） |

## 依赖上游

- `Element.Access` 的 internal 类型泄漏（见 [research.md](research.md) §五）：修复前不实现依赖该 API 的无障碍定制。
- `ui.Themed` / `SetTheme` 语义若上游变动，随锁版升级调整。

## 明确不做

- CEF 窗口路径（超出当前范围）。
- 通过 npm/Bun 发布（纯 Go module）。
