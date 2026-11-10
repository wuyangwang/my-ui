# my-ui 架构

> 基线：mygo **v0.8.3**（`go.mod` 锁定）。调研证据见 [research.md](research.md)，按序实施计划见 [roadmap.md](roadmap.md)。

## 1. 分层

```
L5  secondary       icons/ fonts/ ai/          二级包（规划中），按需 import，不 import 不链接
L4  composites      (Scaffold 等，待实现)       组合组件：页面骨架、Overlay 封装
L3  styled widgets  button.go（已实现）…        变体化预设包装 ui 控件
L2  theme           theme.go palettes.go shades.go 主题构建、明暗/强调色切换、取色派生
L1  tokens          tokens.go                   刻度：字号/间距/圆角/动效倍率（从 *ui.Theme 解出）
L0  —               仅依赖 github.com/egoist/mygo/ui 公开 API
```

- 依赖单向向下。L1/L2 不直接使用 ui 构造器；L3 起持有 `*ui.Context`。
- 根包 `myui` 只含 L1–L4 的轻量件。`icons`、`fonts`、`ai` 是平级独立子包（当前仅占位 README，尚无代码），互不依赖——app 用哪个 import 哪个。

## 2. 主题体系（L2）

### 2.1 明暗切换

框架默认跟随系统外观（浅/深、强调色、高对比、文本缩放）。my-ui 策略：

- **跟随系统**（默认）：不调 `SetTheme`，交给框架。
- **强制浅/深**：视图根部每帧调 `Apply(c, Options{Appearance: ...})` → `c.SetTheme(t)`，下一帧生效（立即模式）。
- **子树级**：`NewTheme(c, opts)` 产出 `*ui.Theme`，用 `ui.Themed(c, t)` 包裹子树；`NewTheme` 不设置窗口主题，避免子树主题外泄。

所有主题从 `ui.LightTheme()/DarkTheme()` 的**拷贝**派生（遵守 "Change a copy"），库不复写内置主题。

### 2.2 主题色（强调色）

- **默认**：使用系统强调色（`c.Preferences().Accent`），经 `shades.go` 派生 hover/pressed/ink。
- **应用覆盖**：选择 `PaletteID` 使用 `palettes.go` 中手写的 8 套 light/dark 三态三元组（base/hover/pressed；源自内置主题强调色及色相邻居），仅覆盖强调色相关语义槽，其余槽位保持当前外观来源。
- 切色与明暗、高对比正交：`NewTheme/Apply` 每帧从桌面偏好与 `Options` 派生 `*ui.Theme`，**无包级缓存、无持久状态**（遵守"never cache colors in package level state"）。

### 2.3 出生与传播

- `Theme` 为固定语义槽平铺结构；派生输入：外观（明/暗/系统）、色板/强调色、桌面偏好（TextScale/HighContrast 等由框架汇入）。
- `Apply`（可变入口）与 `NewTheme`（只构建不设置）每帧按需构建；派生成本低且保证跟随系统变化。

## 3. 中英文（库内文案，后置）

mygo 面向 app 的 i18n 由 `mygo dev/build` 处理 `locales/*.json`；mygo 自身 `internal/l10n` 不可用。库内控件文案极少，后置方案：根包内置 `en` / `zh-Hans` / `zh-Hant` 小消息表，按 `c.Locale()`（BCP-47）解析并回退 `zh-Hant → zh-Hans → en`，允许 app 覆盖；数值/日期格式走公开包 `github.com/egoist/mygo/i18n`。未命中回退英文。

> 状态：未实现，已移出一期（见 roadmap）。

## 4. 页面骨架（Scaffold，待实现）

组合官方机制解决顶栏安全区、滚动与底部遮挡：

- `Scaffold(...)`：顶部吸收 `c.TitleBar()` 安全区；`ui.Scroll().Grow(1)` 内容区；可选吸顶/吸底（`Sticky().Top/Bottom`）。
- 底部预留：滚动内容自动带底部 padding（= 底栏高度 + 边距），滚动条经 `ScrollbarInsets` 内移，默认最后一行不被挡。
- （后续）与 `ai/` 联动的贴底跟随/回底。

## 5. 变体化控件（L3）

- 共享词表：`Variant`（零值 `Native` = 不涂预设、保留底层控件原生观感；`Secondary` 为 `Native` 的 shadcn 别名；另有 `Primary/Outline/Ghost/Destructive/Link`）、`Size`（`SizeMd` 零值 / `Sm` / `Lg`）。
- 构造器形式（代码为准）：`NewXxx(c, ...).<Variant>().<Size>()`，链式应用预设，返回具体控件类型以继续链 ui 方法与读取状态；预设每帧从 `c.Theme()` 派生颜色与尺寸。
- 预设映射到 `Theme` 现有槽位（Accent/Surface/Border/Radius/Space/Rem 等），无自造颜色。
- 遵守 ui 绑定约定（"输入即返回值"）；列表态一律配 `Context.Key`。

## 6. 二级包（规划中）

| 包 | 计划内容 | 资产与许可 | 状态 |
|---|---|---|---|
| `icons` | Lucide + 品牌标 typed `*ui.SVG` 注册表，`go generate` 管线 | Lucide ISC；simple-icons CC0 | 规划中，无代码 |
| `fonts` | 内嵌 OFL 字体，`ui.RegisterFont` 注册字族与常量 | OFL | 规划中，无代码 |
| `ai` | `ChatList`、`ModelPicker`、`Reasoning` | 无资产 | 规划中，无代码 |

明确不做：CEF 窗口路径；npm/Bun 发布。

## 7. 发布与版本

- 锁定 mygo **v0.8.3**（`go.mod`）；升级走专门提交，同步 `go.mod` / `README.md` / `AGENTS.md`。
- 发版与 mygo 版本一一对应；当前未打 tag，待一期范式定稿。element API 有破坏性迁移先例（`mygo migrate-ui`），升级只走独立 PR 并跑全矩阵验证。
- 纯 Go module，`go get` 发布。

## 8. 测试与平台

- 组件经 `ui.Tester`（假后端、无真窗/GPU）驱动；目标状态矩阵：默认/hover/pressed/disabled/read-only/明暗主题/RTL（按控件适用性逐步覆盖）。
- `AGENTS.md` 的 "Tests never pass or fail on wall-clock time" 规则继承。
- 平台矩阵：`go build/vet`（macOS、Linux、Windows）+ `GOOS=js GOARCH=wasm go vet`。
