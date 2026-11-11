# icons（二级包）规格书

> 状态：**设计定稿，待实现**。本文是实现前的调研与契约，实现进度见 [../roadmap.md](../roadmap.md)。

## 1. 参考实现

| 来源 | 可借鉴之处 | 不可照搬之处 |
|---|---|---|
| [Lucide 官方](https://lucide.dev/guide)（1.70.0，ISC） | 图标集本体；24×24 viewBox、纯描边、`stroke="currentColor"`；官方强调 tree-shaking「只打包用到的图标」 | 官方只发 npm 包，且**不含品牌 logo**（官方声明不做品牌图标） |
| [Lobe MyGo 的 `lucide` 包](https://mygo.lobehub.com/skills/components/icon) | **每个图标一个 Go 函数**（`lucide.Save()`），首次调用才解析，未调用的图标不进二进制 | 其包名为 `lucide`（与来源同名），我们统一放 `icons` 包 |
| [shadcn/ui](https://ui.shadcn.com/docs/components/input) | 图标作为组合件的 slot 位置（前缀/后缀/清除按钮）；`currentColor` 随文字色 | React 语境（JSX/CSS 变量），与 Go 立即模式无对应 |
| [Element Plus](https://element-plus.org/zh-CN/component/input) | `prefix-icon` / `suffix-icon` / `clear-icon` 插槽化；清除图标可自定义 | Vue 插槽 + 全局 icon 组件体系 |

许可事实：Lucide 主许可 **ISC**；其中源自 Feather 的图标为 **MIT**，上游 LICENSE 同时收录两段，署名必须完整保留。

## 2. 能否实现（mygo v0.8.3 公开 API 约束）

| 需要的能力 | mygo 公开 API | 结论 |
|---|---|---|
| 解析 SVG | `ui.ParseSVG([]byte)` / `ui.MustParseSVG([]byte) *ui.SVG` | ✅ 支持 paths/basic shapes/groups/use/symbol/transforms/strokes/dashes/opacity/clip/mask；不支持 text/image/pattern/marker/filter——Lucide 图标不含这些，无需预处理 |
| 描边与主题色 | `ui.Icon(c, svg)` 走 `currentColor`，随 `TextColor` 变色 | ✅ Lucide 的 `stroke="currentColor"` 原生生效 |
| 尺寸 | `ui.Icon(...).FontSize(n)` / `.Size(w,h)` | ✅ |
| 按需装载（不进二进制） | **函数内 `//go:embed`**（mygo `ui/svg.go` 文档即此模式）+ 链接器丢弃未引用函数 | ✅ 采用 Lobe 同款：每图标一个函数，函数内 embed |
| 惰性解析 | 首次调用时解析并缓存（`sync.OnceValue` / 包内缓存） | ✅ 不在 `init` 全量解析 |
| 旋转/加载动效 | `ui.Icon(...).Rotate()` / `.Loop()` | ✅（`myui.Icon` 侧再包一层） |
| 无障碍名 | `ui.Element.Label(string)` | ✅ |

不可行/需规避：不得依赖 `Element.Access(func(*platform.AccessNode))`（internal 类型泄漏，见 [../research.md](../research.md) §五），故图标的无障碍只走 `Label`。

## 3. 实现到何种程度

**首轮必做**
- 全量 Lucide 图标（上游 `icons/*.svg` 原样入库，不改内容，便于 diff 与升级）。
- **函数式 API**：`icons.Save() *ui.SVG`——每个图标一个导出函数，SVG 内嵌在各自函数里的**包级变量**中（`//go:embed svg/save.svg` → `var svgSave []byte`）。
- **常用别名表**（我们自维护，独立文件生成）：`icons.Close()`→`x`、`icons.Trash()`→`trash-2`、`icons.Loading()`→`loader-circle`、`icons.Warning()`→`triangle-alert`、`icons.Error()`→`circle-alert`、`icons.Success()`→`circle-check` 等；与上游重名的别名不生成（以上游为准）。
- 生成管线：`icons/gen` 下的 Go 生成器 + `//go:generate` 指令，产物 `icons_gen.go`、`icons_gen.go` 的别名文件 `aliases_gen.go` 入库。
- 许可留痕：`icons/LICENSE`（Lucide ISC 全文，含 Feather MIT 段落）+ `icons/README.md` 记录版本、来源 URL、引入批次。
- 测试：抽样图标可解析可渲染、别名指向正确目标、解析只做一次。

**体积（实测，Windows amd64，与只导入 `mygo/ui` 的对照程序相比）**

| 方案 | 1 个图标 | 30 个图标 | 0 个图标（仅 import） |
|---|---|---|---|
| **每图标独立 embed 变量（采用）** | **+4.1 KB** | **+27.6 KB**（≈850 B/图标） | +512 B |
| 整目录 `embed.FS`（放弃） | +1.29 MB | +1.29 MB | 0 B |

链接器会丢弃未被引用的函数与其 embed 数据，所以**整目录 `embed.FS` 方案下引用任一图标即嵌入整套 1.29 MB**；每图标独立 embed 变量才真正做到「未用零字节」。整目录方案另需 1.29 MB 常驻，已放弃。

根包 `myui` 只引用清除按钮的 `icons.X()` 一个图标，实测对示例程序的二进制影响约等于零（15,981,056 vs 15,981,568 字节，链接器裁掉其余 2,317 个图标函数）。

**后置**
- 品牌标（simple-icons, CC0）→ 独立子包 `icons/brand`，不与 Lucide 混放 ✅ 已实现（3,462 个品牌标）。
- 整套图标的名字清单 → `icons/all`（`Names` 与 `ByName`），给图标选择器用 ✅ 已实现。
- 填充版（lucide-solid）、更多别名（后置）。

## 4. 项目结构是否合理

```
my-ui/
├── icons/                  # 二级包：只依赖 mygo/ui，不依赖根包 myui
│   ├── icons.go            # 包文档 + 解析缓存
│   ├── icons_gen.go        # 生成产物：每个图标一个函数与一个 embed 变量
│   ├── aliases_gen.go      # 生成产物：常用别名
│   ├── svg/                # 上游 Lucide SVG 原样入库（2,322 个文件，1.1 MB）
│   ├── gen/                # 生成器（独立 main 包）
│   ├── LICENSE             # Lucide ISC（含 Feather MIT 说明）
│   └── README.md
├── icon.go                 # 根包 myui.Icon：尺寸/语义色/无障碍名预设（L3）
└── icons/                  # 根包不 import icons 子包以外的东西
```

- **依赖方向**：`icons` → `mygo/ui`（单向）；`icons` 不 import 根包 `myui`；根包 `myui` 只引用 `icons` 中的**一个**图标（清除按钮），实测代价可忽略。
- **根包 `myui.Icon(c, svg)`**：只依赖 `*ui.SVG`，提供尺寸档（`SizeSm/Md/Lg` → 14/20/24 DIP）、语义色（`Variant`）、`Label` 无障碍名、`Spin` 旋转。
- **为什么图标不进根包**：2,322 个 SVG 资源必须隔离在二级包（AGENTS.md 硬规则：根包无 `go:embed`）。根包只借用一个图标，且已实测无代价。

## 5. API 契约

```go
package icons

// Save returns the Lucide "save" icon.
func Save() *ui.SVG

// Close is an alias of X.
func Close() *ui.SVG
```

```go
package myui

// NewIcon creates an icon with the native preset.
func NewIcon(c *ui.Context, svg *ui.SVG) *Icon

func (i *Icon) Variant(v Variant) *Icon   // Primary/Destructive/Link 给语义色
func (i *Icon) Size(s Size) *Icon         // SizeSm/Md/Lg → 14/20/24 DIP
func (i *Icon) Sm() *Icon
func (i *Icon) Md() *Icon
func (i *Icon) Lg() *Icon
func (i *Icon) Label(s string) *Icon      // 无障碍名
func (i *Icon) Spin() *Icon               // 旋转（ui 的 Loop/Rotate）
```

用法：

```go
myui.NewIcon(c, icons.Save()).Md().Label("保存")
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] vendor Lucide 1.70.0 SVG 入库 + LICENSE
- [x] 生成器 + 生成产物（2,318 个图标 + 14 个别名）
- [x] 别名表
- [x] 根包 `myui.Icon`
- [x] 测试（icons 包 + 根包 icon_test.go）
- [x] showcase「Icon」页
- [x] 品牌标子包 `icons/brand`（已实现，3,462 个）
