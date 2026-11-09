# my-ui 路线图

> "mygo 已有"列指 mygo v0.8.3 现状（证据见 [research.md](research.md)）；"my-ui 补"列为本库产出。条件：基线已验证外部 module 可编译、可发布。

## 优先级排序原则

1. 差异化优先：`ai/` 是 mygo 独有能力（流式 Markdown + 虚拟化列表）之上的库级卖点。
2. 范式先行：先做 tokens + theme + 2~3 个变体控件，把"一行构造器 + 预设"的范式定稿，后续组件照抄。
3. 无风险项优先：瀑布流已明确不做。

## 一期（范式与卖点）

| 项 | mygo 已有 | my-ui 补 |
|---|---|---|
| tokens | `Theme.Space/Rem`、`Color.Hex/Mix/MixOklab` | 色板常量、字号/间距/圆角/动效/内衬刻度体系 |
| 主题 | `LightTheme/DarkTheme`、`SetTheme`、`ui.Themed`、系统外观/强调色自动跟随 | 主题 preset 集、明暗切换入口（跟随/强制）、色板强调色切换入口 |
| 中英文 | `c.Locale()`/`c.Region()`、公开 `i18n` 包（CLDR） | 内置 `en`/`zh-Hans`/`zh-Hant` 小文案表 + 回退链 + app 覆盖 |
| 变体控件 | Button/PrimaryButton 等 uigen themed 包装 | `Button`、`Input`、`Badge` 三件套的 Variant/Size 预设范式 |
| 常用图标 | `ui.Icon`（currentColor 变色、多色 SVG 保留原色） | `icons/`：Lucide typed 注册表 + go generate 管线 |
| 页面骨架 | `Scroll`+`ScrollbarInsets`+`Sticky`+`ScrollState`+`TitleBar()` | `Scaffold`（顶栏安全区/滚区/吸顶吸底 + 底部预留默认不遮挡） |

一期验收：`go build/vet` 三主机 + wasm 全绿；`ui.Tester` 状态测试覆盖三件套 + Scaffold；一个最小 demo 说明页。

## 二期（补桌面细粒度件）

Card / Kbd / Skeleton / EmptyState / Pagination / HoverCard / Drawer(Sheet) / Command Palette(⌘K) / 应用内 Menubar / Table 排序与过滤 / Carousel / AspectRatio——依据：overlays/custom-widgets/web-view 文档有官方手写样例，均为组合封装。

## 三期（AI 包）

`ai/ChatList`（虚拟化气钢布局、typing 指示、贴底跟随、注入条与底部预留整合）、`ai/ModelPicker`（分组 Combobox + 厂商标 + 速度/价格徽标）、`ai/Reasoning`（Segmented 推理档 + temperature 刻度）。依赖一期 tokens/page 的 `ai/` 联动形态，故排三期；也可与二期并行。

## 依赖上游的项

- `Element.Access` 的 internal 泄漏（research.md §五）：修复前 my-ui 不实现依赖该 API 的无障碍定制功能。
- `ui.Themed`/`SetTheme` 语义若上游变动，切换入口随之调整（已锁版，风险受控）。

## 明确不做

- 瀑布流（`ui` 无布局原语，自定义虚拟化成本与收益不匹配）。
- CEF 窗口路径（mygo 范围约定内明确 out of scope）。
- 通过 npm/Bun 发布（native-only，纯 Go module）。
