# my-ui 调研纪要

本文记录立项前针对 mygo 的 `package ui` 的实机调研结论（调研基线：mygo v0.8.3），全部有代码或文档证据支撑，供实现阶段回查。

## 一、总体结论

`ui` 本身已是完整的立即模式控件层（fluent 链式 API + "输入即返回值"绑定），自带 54+ 控件与 `ui.Theme` 主题体系；`internal/uigen` 为每个 `XxxBase` 生成带主题外观的公开包装（`ui/api_gen.go`）。my-ui 的正确定位：**主题包 + 语义常量 + 变体化预设 + 框架未封装的组合组件**，不重造控件、不 fork。

## 二、分项调研

### 1. 主题（结论：地基极强，库做主题包）

- 默认主题自动跟随系统外观（浅/深）、强调色、高对比、文本缩放、Reduce Motion、指针形态（`docs/ui/styling.md` "The desktop's preferences"），经 `Handler.ThemeChanged` 流入。
- 窗口级 `c.SetTheme(t)`；子树级 `ui.Themed(c, t)` 可嵌套，内部打开的 Popover/菜单/Dialog/Tooltip 继承该主题；官方示例即"主题选择器并排预览浅深"（`styling.md` §A theme for part of the window）。
- `Theme.Space(n)` / `Theme.Rem(n)` 是间距、字号的既有换算语义。
- 限制：`Theme` 是固定语义槽位的平铺结构（约 25 色 + 指标），预设只能在语义槽内变体（用 `Color.Mix`/`MixOklab` 从 seed 色推导 hover/pressed 等衍生槽位）。

### 2. 动画（结论：框架体系强于 shadcn 的 CSS transition，库只包预设）

- `Element.Animate/AnimateTo/Animate`：弹簧（`ui.Smooth`/`Snappy`/`Bouncy`，SwiftUI 风格 Duration+Bounce），精确求解、保留速度；曲线含 `CubicBezier`（`docs/ui/motion.md`）。
- `AnimateColor`（sRGB/Oklab）、`AnimatePose`（仅重绘不重跑视图，compositor 级，成本最低；`card.Hovered()` 悬浮抬升是官方例子）。
- 内建 transitions、列表行重排动画、`Blur`/`Opacity`/`ClipInset` 擦入/渐变遮罩/`shader.Effect`/glass 插件。
- 三端都尊重系统 Reduce Motion（`Animate` 直达目标）。

### 3. 易用性

- 立即模式：`ui.Button(c, &label)` 返回指针承载输入，下一句可读；组件库提供一步式构造器封掉易忘事项（`Key()` 身份、主线程、变体配色）。
- 文档完备：`docs/ui/` 70+ 篇，每控件一篇。
- 测试：`ui.Tester`（可指定暗色/Reduce Motion 起步），不开真窗、无 GPU 依赖，CI 三平台 + wasm 可行。

### 4. 独立第三方包（已实机验证 ✅）

- 临时模块 `example.com/uiext`（`require github.com/egoist/mygo v0.8.3`）从 Go 代理正常拉取，`go build`/`go vet` 通过；外部成功使用 `ui.View`/`ui.Button`/`ui.Icon`+`TextColor`/`c.Theme()`。mygo 已有完整 semver 历史（v0.1.0 → v0.8.3）在代理上。
- mygo 仓库单 module；供外部消费的公开注入点是 `mygo.WindowOptions{Content: ui.View(fn)}`（`content.go`）。
- API 稳定性：存在 `ui/migration.md` 与 `mygo migrate-ui`（element API 历史上有破坏性迁移），必须锁版本。

### 5. 图标（`ui.Icon` 原生支持变色）

- `ui.Icon` 渲染 SVG，`stroke="currentColor"` 的图标（Lucide/Heroicons/Tabler/Material Symbols 原文件即如此）自动跟随 `TextColor`；`FontSize`/`Size` 控制尺寸；`Rotate`/`Loop` 可作旋转指示。多色 SVG（品牌标）保留原色（`ui/svg_test.go`、`ui/markdown` 的图标实现）。
- `Painter.Icon` 在自绘回调中可用。
- `plugins/sfsymbols` 是 macOS 专属（其他平台显式不可用），只能作为可选平台增强，不进基础包。

### 6. 字体

- `ui.Font{Family, Size}`；`Theme.Font` 字族字符串（空 = 系统字体）；`"monospace"` 解析系统等宽。
- `ui.RegisterFont(data []byte, family string)`（`ui/window.go:40`）：内嵌 TTF 注册字族 → my-ui `fonts/` 二级包的实现依据（OFL 字体允许再分发）。

### 7. AI 组件（欠缺点：仅瀑布流）

- ChatList：`ui.List` 为锚点式虚拟化列表，行高动态可变、行可增删（`docs/architecture.md` "Lists"）；`ui/markdown` 专为 LLM 逐 token 流式回复构建（增量解析成本与回复长度恒定、半成品语法按终态显示、文本淡入）。
- 模型选择器：`Combobox` + `Avatar`/`Icon` + `ComboboxParts.Chosen` 选中态可包。
- 推理档位/温度：`Segmented` + `RangeSlider`/`Stepper` 可包。
- 瀑布流：`ui` 无等价布局原语（GridView 为等行高网格，List 为线性），当前不列入计划。

### 8. 页面骨架（padding / 滚动 / 底部遮挡）

- `ui.Scroll` + 内容 `Padding` 防贴边；`ScrollbarInsets` CSS 式内移滚动条避开浮动条（`scroll.md` L216-224 官方示例：内容从浮条下滚过、末行不被挡）。
- `Sticky().Top(0)/Bottom(0)`（`position: sticky` 语义）；`ScrollState`/`TrackScroll`/`ScrollTo` 弹簧滚动与"贴底跟随、上滚即停"官方模式。
- `c.TitleBar()` 返回隐藏标题栏窗口的控件安全区（红绿灯/窗口按钮），官方示例喂给 `Padding` 与 `DragWindow()`；macOS 另有 Vibrancy。

## 三、mygo 已有控件清单（节选，见 ui/api_gen.go）

Base 内核 + uigen 生成 themed 包装：Accordion、AlertDialog、Autocomplete、Avatar、Badge、Breadcrumbs、Button/PrimaryButton、Calendar、Checkbox(组、State)、Collapsible/Accordion、ColorPicker、ColorWell、Combobox、ContextMenu、DataDragOver/DropData、DateInput、DecimalInput、Dialog、Divider、EditableText、Field/Fieldset、FindBar、Form、Grid/GridView、Icon/Image/Video、Link、MenuBase/MenuButton、Meter、Modal、NativeView、NumberInput、Overlay、Popover、Progress、RadioGroup、RangeSlider、Rating、RichText、Scroll/SearchField、Segmented、Select、Sidebar、Slider/StepSlider、Split、Spinner、Stepper、Switch、Table、Tabs、Text/Textf、TextArea、TextInput、Themed、TimeInput、Toast、Toggle/ToggleGroup、TokenField、Toolbar、Tooltip、Tree、WebView……
另含重型件：`ui/chart`、`ui/editor`（代码编辑器）、`ui/markdown`、`ui/richtext`、shader/glass。

## 四、上 lib 可补的桌面件（按序计划见 roadmap.md）

Card、Kbd、HoverCard、应用内 Menubar、Drawer/Sheet、Command Palette、Skeleton、EmptyState、Pagination、Carousel、AspectRatio、表格排序/过滤等均为组合封装，有官方样例可循。

## 五、上游问题（建议向上游反馈）

`ui/api_gen.go:3358` `Element.Access(fn func(n *platform.AccessNode))` 公开签名引用 `github.com/egoist/mygo/internal/platform` 类型：本仓库（外部 module）无法引用该方法，涉及元件的定制无障碍会受限。规避原则：my-ui 不实现依赖该 API 的功能；待上游将类型换成 `ui` 包内公开类型后跟进。

## 六、风险表

| 风险 | 缓解 |
|---|---|
| mygo element API 破坏性迁移（有先例） | go.mod 精确锁版；发版与 mygo 版本一一对应；迁移用 `mygo migrate-ui` |
| 上游 Access internal 泄漏 | 规避实现；向上游提 issue |
| 内嵌资产体积（字体/图标） | 全部隔离在二级包，不 import 不链接 |
| 性能红线 | 沿用 mygo "apps pay only for what they use"；热路径不缓存颜色；动效走 AnimatePose（仅重绘） |
