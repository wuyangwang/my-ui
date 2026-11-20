# my-ui 架构

my-ui 是 MyGo 原生 UI（`package ui`）之上的第三方组件库：主题底座 + 语义词表 + 变体化控件预设 + 框架未封装的组合组件。只依赖 mygo 的公开 API，不 fork、不重造控件。

## 1. 分层

```
L4  secondary    icons/ fonts/ ai/              独立二级包，按需 import，不 import 不链接
L3  widgets      controls/ forms/ overlays/ display/   变体化预设与组合组件
L2  theme        theme.go palettes.go shades.go  主题构建、明暗/强调色切换、取色派生
L1  tokens       tokens.go                       刻度：字号/间距/圆角/动效倍率（从 *ui.Theme 解出）
L0  —                                            仅依赖 github.com/egoist/mygo/ui 公开 API
```

依赖单向向下。根包 `myui` 只含 L1/L2 的共享底座（主题、tokens、词表、动效、i18n）；`bind/` 是每帧句柄，`window/` 是无框窗口外壳，组件按功能放在 `controls/`、`forms/`、`overlays/`、`display/`。app 用哪个 import 哪个。

## 2. 主题体系

- 每帧从桌面偏好与 `Options` 派生 `*ui.Theme`，跟随系统明暗、强调色、高对比与文本缩放；无包级颜色缓存、无持久状态。
- `Apply(c, Options{})` 设置窗口主题；`NewTheme(c, opts)` 只构建主题，供 `ui.Themed(c, t)` 包裹子树，不设置窗口主题。
- 默认主题是 shadcn/ui "new-york" zinc 主题的 MyGo 表达：`Background` 白 / `zinc-950`，`Surface` 次级面，`Accent` 近黑，圆角基准 8。
- 强调色三态（base/hover/pressed）由 `shades.go` 从基色派生；色板、系统强调色、语义色 success/warning/info 都在语义槽内变化。
- 主题留空的槽位（`Inverse`、`InverseText`）取框架兜底 `myui.InverseColors`；平台语义色（红绿灯、关闭危险色）按平台色相，由 `trafficLightColors` 逐帧从主题与选项派生。

## 3. 变体化控件

- 共享词表：`Variant`（零值 `Native` 保留底层观感，`Default`、`Secondary` 及其别名）、`Size`（`SizeMd` 零值 / `Xs` / `Sm` / `Lg`）、`Shape`（`ShapeRect` / `ShapeRound` / `ShapeCircle`）。
- 构造器形式：`NewXxx(c, ...).<Variant>().<Size>()`，返回具体控件类型以继续链 ui 方法与读取状态；预设每帧从 `c.Theme()` 派生。
- 预设只映射到 `Theme` 现有槽位，不自造颜色；遵守 ui 的绑定约定（"输入即返回值"），列表态配稳定 `Context.Key`。
- 输入类控件默认占满父宽，`Width(w)` 固定宽度。清除/关闭是带按钮角色的图标控件（`controls.NewIconAction`），hover/按下只改墨色。

## 4. 弹层与动效

- 弹层内容建在 `panel.Children(...)` 里（context 的父是 overlay 层，不是面板）；触发器建在弹层之前，开合 flag 由 app 持有，库只读 `Clicked()`。
- 状态变化走 `Element.Transition`，时长取自动效刻度，Reduce Motion 由框架兜底；组件不带自己的计时器或 goroutine。
- 无障碍走 `Label`/`Description`/`Invalid`/`ErrorMessage`；不依赖 `Element.Access` 的 internal 类型泄漏 API。

## 5. 二级包

| 包 | 内容 |
|---|---|
| `icons` | Lucide 图标，函数式 API，未调用的图标不进二进制 |
| `icons/brand` | Simple Icons 品牌标 |
| `icons/all` | 整套图标名与实例，供图标选择器使用 |
| `fonts` | 内嵌 OFL 字体与 `ui.RegisterFont` 注册 |
| `ai` | ChatList、ModelPicker、Reasoning |

明确不做：CEF 窗口路径；npm/Bun 发布。

## 6. 版本与依赖

- 锁定 mygo **v0.10.0**（`go.mod`）；升级走专门提交，同步 `go.mod` / `README.md` / `AGENTS.md`。
- 当前版本 my-ui **v0.8.1**（`myui.Version`，tag `v0.8.1`），与 mygo 独立编号。
- 纯 Go module，`go get` 发布。

## 7. 测试与平台

- 组件经 `ui.Tester`（假后端、无真窗/GPU）驱动；状态矩阵覆盖默认/hover/pressed/disabled/read-only/明暗主题/RTL。
- 测试不依赖壁钟时间：等条件或事件，不 sleep；按计数而非时长度量。
- 平台矩阵：`go build/vet`（macOS、Linux、Windows）+ `GOOS=js GOARCH=wasm go vet`。无 cgo。

## 8. 示例应用外壳

`examples/showcase` 是组件库示例应用，外壳按组件文档站组织：顶栏放品牌与搜索（过滤侧栏）；侧栏按分组折叠，每页一个图标；标题下写组件的真实用途，变体/状态排成紧凑的演示卡（左侧名称、右侧可交互组件）。外壳只调库公开 API 与 `c.Theme()` 的派生色，不缓存颜色、不带自己的主题状态。

## 9. 已记录的框架限制

以下属 mygo v0.10.0 公开 API 的限制，只记录不列 TODO：

- Select 的清除按钮嵌在 `RolePopUpButton`（leaf role）内，读屏摸不到；按钮本身可见可点。
- Slider 双旋钮区间：`SliderBase` 只绑单值。
- Toast 最多三条、拿不到「已显示时长」（只有 `Left()`）、滑动关闭需手势 API。
- Tooltip 触发方式固定 hover+focus，延迟只能取主题的 `TooltipDelay`。
- Tabs 竖排：`TabsBase.List` 是写死的行，无公开方向切换。
- Progress 无障碍自定义 min/max：`hasRange/accRange` 非导出。
- `Info` 语义色没有主题槽位；Modal 焦点陷阱细节不可控。
