# my-ui 架构

> 基线：mygo v0.8.3（`go.mod` 锁定）。调研证据见 [research.md](research.md)，组件排期见 [roadmap.md](roadmap.md)。

## 1. 分层

```
L5  secondary       icons/ fonts/ ai/          二级包，按需 import，不 import 不链接
L4  composites      page.go chat 入口等         组合组件（Page 骨架、Overlay 封装）
L3  styled widgets  button.go input.go card.go… 变体化预设包装 ui 控件
L2  theme           theme.go palettes.go        主题预设、明暗/强调色切换、Apply 入口
L1  tokens          tokens.go motion.go inset.go 常量：色板/字号/间距/圆角/动效/内衬
L0  —               仅依赖 github.com/egoist/mygo/ui 公开 API
```

- 依赖单向向下。L1/L2 不用 ui 构造器；L3 起用到 `*ui.Context`。
- 根包 `myui` 只含 L1–L4 的轻量件。`icons`、`fonts`、`ai` 是平级独立子包（不是根包子树），互相不依赖——app 用哪个 import 哪个，未 import 的零字节进入二进制。

## 2. 主题体系（L2）

### 2.1 明暗切换

框架默认跟随系统外观（浅/深自动切换、强调色/高对比/文本缩放随系统）。my-ui 的开关策略：

- 形式为一个可放视图根部的构建时入口（立即模式每帧重跑，因此切换 = 改状态即可，下一帧生效）：
  - 跟随系统（默认，不调 `SetTheme`，交给框架）；
  - 强制浅色 / 强制深色：视图根部调 `c.SetTheme(t)`（`ui` 的 `SetTheme` 作用于窗口根及随后的控件，"unless SetTheme replaced it"）。
- 子树级（预览窗格、暗色工具栏内嵌浅色卡片等）：框在手，`ui.Themed(c, t)` 包住即可，浮层（Popover/菜单/Dialog/Tooltip）继承子树主题。
- 所有主题实例从 `ui.LightTheme()/DarkTheme()` 拷贝派生（框架文档约定"Change a copy"），库不复写内置主题。

### 2.2 主题色（强调色）切换

- 默认接系统强调色（macOS/Windows/GNOME 47+/KDE，`c.Preferences().Accent`），无需代码。
- 应用覆盖：从色板 preset（Oklch/Oklab 下的色相档）派生 `Accent/AccentHover/AccentPressed/AccentText`（`ui.Color` 支持宽色域；`Color.Mix/MixOklab` 推导 hover/pressed 层次），仅覆盖相应语义槽，其余槽位保持当前外观来源。
- 切色与明暗、高对比正交：任何组合都经由同一处派生函数产出 `*ui.Theme`。

### 2.3 出生与传播

- `Theme` 为固定语义槽平铺结构；派生输入只有三轴：外观（明/暗/系统）、色相、桌面偏好（Accent/HighContrast/TextScale/ReduceMotion 由框架自动汇入，库不接管）。
- 参数未变化时复用同一 `*Theme` 指针，避免每帧重建派生（派生本身也在视图帧内做，成本极低，但仍作缓存）。

## 3. 中英文（库内文案）

- mygo 的面向 app 的 i18n 由 `mygo dev/build` 从 app 的 `locales/*.json` 目录生成（独立于本库，本库不加工其生成物）。`internal/l10n`（mygo 自身文案）不可用——是 internal 包。
- 库不推给 app 的文案极少（分页、无结果、清除、颜色拾取器标签等控件内固定的少量词），方案：
  - `myui/strings.go` 内置 `en`、`zh-Hans`、`zh-Hant` 三份小消息表（map 常量），按语言解析取值；
  - 解析键为 `c.Locale()`（BCP-47，公开 API `ui.Context.Locale`），回退链 `zh-Hant → zh-Hans → en`，支持区域后缀（`zh-CN` → `zh-Hans` 语义解析由 mygo `i18n.Match` 同款规则实现，必要时直接调用公开的 `github.com/egoist/mygo/i18n` 包）;
  - 每个 key 允许 app 整表/单 key 覆盖（`Strings` 类型的 pointer map），满足品牌文案自定义；
  - 数值/日期格式不自己造：需要时走公开包 `mygo/i18n`（Unicode CLDR 数据）。
- 中英文之外的语言不阻塞核心（未命中回退英文），目录结构允许后续加。

## 4. 页面骨架（page.go）

组合封装三个官方机制，解决"页面 padding、滚动、底部内容被挡"：

- `Scaffold(c, opts)`：顶部吸收 `c.TitleBar()` 安全区；`ui.Scroll().Grow(1)` 内容区；可选吸顶头 / 吸底栏（`Sticky().Top(0)/Bottom(0)`）。
- 底部预留：滚动内容自动带底部 `Padding`（= 底栏高度 + 边距），滚动条经 `ScrollbarInsets` 内移避开浮条——默认行为就是最后一行不被挡。
- 聊天特化（与 `ai/` 联动）：贴底跟随（`ScrollState` `Y ≥ MaxY → Y=MaxFloat32`）、上滚即停、`ScrollTo` 弹簧回底。

## 5. 变体化控件（L3）

- 构造器形式：`myui.Button(c, &label, myui.Variant.Primary, myui.Size.Md)`——Variant/Size/inset 枚举选择预设，返回 ui 控件的公开类型（如 `*ui.ButtonBase`），其余链式选项原样可用。
- 预设映射到 `Theme` 现有槽位（Surface(surface/hover/pressed)/Accent/Border/Radius/Space/Rem），无自造颜色。
- "输入即返回值"绑定约定不变；列表态一律配 `Context.Key`。

## 6. 二级包

| 包 | 内容 | 资产与许可 |
|---|---|---|
| `icons` | `go generate` 内嵌 Lucide 全集 + 品牌标子集，导出 `*ui.SVG` 常量（typed 注册表，编译期查名） | Lucide ISC；simple-icons CC0；`currentColor` 天然变色 |
| `fonts` | 内嵌 OFL 字体（Inter、JetBrains Mono 等），`init` 时 `ui.RegisterFont` 注册字族；提供字族/字号常量 | 字体 OFL，允许再分发 |
| `ai` | `ChatList`（虚拟化 List + `ui/markdown` Stream/View 气泡）、`ModelPicker`（分组 Combobox + 厂商标）、`Reasoning`（Segmented 档位 + 温度滑杆） | 无资产 |

明确不做：瀑布流。

## 7. 发布与版本

- 锁定所发布 mygo 版本（当前目标 **v0.8.3**）。脚手架期无 mygo 导入时 `go mod tidy` 会清空 require 行，首个组件代码落地即写入并从此保持精确锁版；my-ui 的 release 与 mygo 版本一一对应。
- element API 历史上有破坏性迁移（`ui/migration.md`、`mygo migrate-ui`）：升级 mygo 只走独立 PR，跑全矩阵验证。
- 无 Bun/npm；纯 Go + `go get` 发布。

## 8. 测试与平台

- 每组件经 `ui.Tester`（假装后端、无真窗、无 GPU）驱动：默认/hover/按下/禁用/只读/明暗主题/RTL 状态矩阵。
- `AGENTS.md` 的"Tests never pass or fail on wall-clock time"规则继承。
- 平台矩阵：`go build/vet`（macOS、Linux、Windows 三主机）+ `GOOS=js GOARCH=wasm go vet`（浏览器，mygo 原生 UI 支持项）。
