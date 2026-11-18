# my-ui 架构

> 基线：mygo **v0.10.0**（`go.mod` 锁定）。调研证据见 [research.md](research.md)，按序实施计划见 [roadmap.md](roadmap.md)。

## 1. 分层

```
L5  secondary       icons/ fonts/ ai/          二级包（icons/fonts 已交付，ai 规划中），按需 import，不 import 不链接
L3  styled widgets  button.go（已实现）…        变体化预设包装 ui 控件
L2  theme           theme.go palettes.go shades.go 主题构建、明暗/强调色切换、取色派生
L1  tokens          tokens.go                   刻度：字号/间距/圆角/动效倍率（从 *ui.Theme 解出）
L0  —               仅依赖 github.com/egoist/mygo/ui 公开 API
```

- 依赖单向向下。L1/L2 不直接使用 ui 构造器；L3 起持有 `*ui.Context`。
- 根包 `myui` 只含 L1/L2 的共享底座（主题、tokens、词表、动效、i18n）。组件按功能放在 `controls/`、`forms/`、`overlays/`、`display/`；`bind/` 是帧句柄，`window/` 是无框窗口外壳，`icons/`、`fonts/`、`ai/` 是平级独立子包。app 用哪个 import 哪个。

## 2. 主题体系（L2）

### 2.1 明暗切换

框架默认跟随系统外观（浅/深、强调色、高对比、文本缩放）。my-ui 策略：

- **视图根部**：每帧调 `Apply(c, Options{Appearance: ..., Palette: ...})` → `c.SetTheme(t)`，下一帧生效（立即模式）。`Options{}` 也调——默认不是"交给框架"，而是给 shadcn 的默认主题（见 2.2）。
- **跟随系统**（默认）：`AppearanceSystem` 下主题从系统当前明暗派生，并跟随桌面偏好变化。
- **强制浅/深**：`AppearanceLight/Dark` 固定明暗。
- **子树级**：`NewTheme(c, opts)` 产出 `*ui.Theme`，用 `ui.Themed(c, t)` 包裹子树；`NewTheme` 不设置窗口主题，避免子树主题外泄。

所有主题从 `shadcn.go` 的 `shadcnLight()/shadcnDark()` 的**拷贝**派生（遵守 "Change a copy"），库不复写内置主题。该主题是 shadcn/ui "new-york" zinc 主题的 MyGo 表达：`Background` 是白/`zinc-950` 的卡片与页面底，`Surface` 是 `zinc-100`/`zinc-800` 的次级面与侧栏，`Accent`（primary）是近黑/`zinc-50`，圆角基准 8。

### 2.2 主题色（强调色）

- **默认**：shadcn 的 zinc 中性色 + 近黑 primary（`shadcn.go`）；无彩色强调色时保持这套，不是 mygo 自带的蓝。
- **系统强调色**：桌面给了 `c.Preferences().Accent` 且未选色板时，用它覆盖 primary，经 `shades.go` 派生 hover/pressed/ink。
- **应用覆盖**：选择 `PaletteID` 使用 `palettes.go` 中手写的 8 套 light/dark 三态三元组（base/hover/pressed），仅覆盖强调色相关语义槽，其余槽位保持 shadcn 基础主题。
- 语义色：Success/Warning 优先取主题的 `t.Success/t.Warning`（`semanticColor`），跟随自定义主题与明暗；Info 因 mygo 的 `ui.Theme` 无对应槽位，由固定色表给出（框架限制，已记录）。
- 切色与明暗、高对比正交：`NewTheme/Apply` 每帧从桌面偏好与 `Options` 派生 `*ui.Theme`，**无包级缓存、无持久状态**（遵守"never cache colors in package level state"）。

### 2.3 出生与传播

- `Theme` 为固定语义槽平铺结构；派生输入：外观（明/暗/系统）、色板/强调色、桌面偏好（TextScale/HighContrast 等由框架汇入）。
- `Apply`（可变入口）与 `NewTheme`（只构建不设置）每帧按需构建；派生成本低且保证跟随系统变化。

## 3. 中英文（库内文案，后置）

mygo 面向 app 的 i18n 由 `mygo dev/build` 处理 `locales/*.json`；mygo 自身 `internal/l10n` 不可用。库内控件文案极少，后置方案：根包内置 `en` / `zh-Hans` / `zh-Hant` 小消息表，按 `c.Locale()`（BCP-47）解析并回退 `zh-Hant → zh-Hans → en`，允许 app 覆盖；数值/日期格式走公开包 `github.com/egoist/mygo/i18n`。未命中回退英文。

> 状态：消息表与 `Message(c, key)` 已实现（`i18n.go`），见 roadmap。


组合官方机制解决顶栏安全区、滚动与底部遮挡：

- 底部预留：滚动内容自动带底部 padding（= 底栏高度 + 边距），滚动条经 `ScrollbarInsets` 内移，默认最后一行不被挡。
- （后续）与 `ai/` 联动的贴底跟随/回底。

## 4. 变体化控件（L3）

- 共享词表：`Variant`（零值 `Native` = 不涂预设、保留底层控件原生观感；`Secondary` 为 `Native` 的 shadcn 别名；另有 `Primary/Outline/Ghost/Destructive/Link`）、`Size`（`SizeMd` 零值 / `Sm` / `Lg`）。
- 构造器形式（代码为准）：`NewXxx(c, ...).<Variant>().<Size>()`，链式应用预设，返回具体控件类型以继续链 ui 方法与读取状态；预设每帧从 `c.Theme()` 派生颜色与尺寸。
- 预设映射到 `Theme` 现有槽位（Accent/Surface/Border/Radius/Space/Rem 等），无自造颜色。
- 遵守 ui 绑定约定（"输入即返回值"）；列表态一律配 `Context.Key`。

## 5. 二级包

| 包 | 内容 | 资产与许可 | 状态 |
|---|---|---|---|
| `icons` | Lucide + 品牌标 typed `*ui.SVG`，`go generate` 管线 | Lucide ISC；simple-icons CC0 | ✅ 已交付 |
| `fonts` | 内嵌 OFL 字体，`ui.RegisterFont` 注册字族与常量 | OFL | ✅ 已交付 |
| `ai` | `ChatList`、`ModelPicker`、`Reasoning` | 无资产 | 规划中 |

明确不做：CEF 窗口路径；npm/Bun 发布。

## 6. 发布与版本

- 锁定 mygo **v0.10.0**（`go.mod`）；升级走专门提交，同步 `go.mod` / `README.md` / `AGENTS.md`。
- 当前版本 my-ui **v0.5.0**（`myui.Version`，tag `v0.5.0`），与 mygo 独立编号；发布用 git tag 标记。element API 有破坏性迁移先例（`mygo migrate-ui`），mygo 升级只走独立 PR 并跑全矩阵验证。
- 纯 Go module，`go get` 发布。

## 7. 测试与平台

- 组件经 `ui.Tester`（假后端、无真窗/GPU）驱动；目标状态矩阵：默认/hover/pressed/disabled/read-only/明暗主题/RTL（按控件适用性逐步覆盖）。
- `AGENTS.md` 的 "Tests never pass or fail on wall-clock time" 规则继承。
- 平台矩阵：`go build/vet`（macOS、Linux、Windows）+ `GOOS=js GOARCH=wasm go vet`。

## 8. 示例应用外壳（showcase）

`examples/showcase` 的定位是**组件库示例应用**，外壳按组件文档站组织，而不是后台产品：

- **顶栏**（`appBar`）：品牌区与侧栏同为 256px，下面用一条 1px 线分隔；搜索框与明暗切换在右侧。品牌与版本只出现一次。
- **侧栏**（`sidebar`）：16rem、与页面同底色、1px 分隔线；分组标题为小号弱墨文字，条目 36 DIP、圆角 6；选中项是强调色淡底 + 强调色文字；点击当前分组可显式合起，不被当前页的归属覆盖。
- **页面**：标题、分组徽标与 `pages.go` 里逐页写的 `desc`（组件的真实用途），下面一条分隔线，再进入滚动区。
- **演示卡**（`card` / `sectionNote`）：白底 + 细边框 + 双层柔和阴影，圆角取 `RadiusLg`；变体页用三列等高小卡，左侧名称/说明、右侧就是该变体的可交互组件，读起来是一张变体表而非一堆空盒。

外壳只调库公开 API 与 `c.Theme()` 的派生色，不缓存颜色、不带自己的主题状态；换主题、换色板都走每帧的 `Apply`。

## 9. 已记录的框架限制（不可在本库修）

组件审查后确认属 mygo v0.10.0 公开 API 的限制，只记录不列 TODO：

- Select 的清除按钮嵌在 `RolePopUpButton`（leaf role）内，读屏摸不到；按钮本身可见可点。
- Slider 双旋钮区间：`SliderBase` 只绑单值。
- Toast 最多三条、拿不到「已显示时长」（只有 `Left()`）、滑动关闭需手势 API。
- Tooltip 触发方式固定 hover+focus，延迟只能取主题的 `TooltipDelay`。
- Tabs 竖排：`TabsBase.List` 是写死的行，无公开方向切换。
- Progress 无障碍自定义 min/max：`hasRange/accRange` 非导出。
- `Info` 语义色没有主题槽位；Modal 焦点陷阱细节不可控。
