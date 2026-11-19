# Chart 规格书

> 位置：组件在 [`chart/`](../../chart) 子包；主题/tokens 仍在根包 `myui`。
>
> 状态：**已实现**（`chart.go`）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo `ui/chart`（v0.10.0：`chart.go`、`scale.go`、`shape.go`） | **无外观的底层积木**：`Plot` 画坐标区与四周刻度，`NewLinear`/`NewBand` 做比例，`Rect`/`Point`/`RuleX` 放标记，`Line`/`Area`/`Sector` 出路径，`Pie` 分扇形，`Select` 用指针与方向键选中；标记按百分比定位，随区域缩放 |
| shadcn/ui 的图表观感 | 主色作第一序列，弱底 `Surface` 作块面，边框色 `Border` 作网格，次要文字 `TextMuted` 作刻度 |
| d3 / Swift Charts | 刻度取 1/2/5 × 10ⁿ 的圆整值（`Linear.Nice`/`Ticks`），折线可选平滑（`CurveMonotone`）与阶梯 |

## 2. 能否实现

- 坐标轴与网格：`mchart.Plot(c, Axes{Left, Bottom}, build)` ✓；网格在 `area.Draw` 里用 `Painter.Line` 画。
- 柱/线/面/饼：`Rect`、`Line`、`Area`、`Pie`+`Sector` 都是公开 API ✓；本包把它们组装成预设，颜色由 my-ui 色板给出。
- 选中态：`mchart.Select(c, area, index *int, xs, describe)` ✓（状态指针由 app 持有，指针离开写 -1，方向键移动）；本包的 `Options.Selected` 直接透传。
- 主题每帧：本包每次构造都调用 `c.Theme()` 重新生成色板；上游标记是普通元素，可用 `Transition(myui.StateTransition())` 让数据变化时颜色渐变 ✓。
- 不能做：mygo 没有「一行出图」的上层组件，也没有图例/动画/缩放的内建件；图例在饼图预设里用文本与色块自绘，缩放交互没有（框架没有）。

## 3. 实现到何种程度

**已实现**：

- `NewPalette(t)`：从帧主题取色（主色 → Info → Success → Warning → Danger 循环作序列色，`Border` 作网格，`TextMuted` 作刻度，`Surface` 作底）；没有包级颜色缓存。
- 四个即用预设：`NewBars`（柱）、`NewLine`（折线 + 淡填充 + 数据点）、`NewArea`（面积）、`NewPie`（饼 + 色块图例），都接受 `[]float64`，空值也安全。
- 可配置版 `Bars`/`Line`/`Area`/`Pie` + `Options`：类别标签、`Selected *int` 选中指针、高度、曲线插值、无障碍名与描述。
- 柱与折线的选中态画参考线与数据点；选中指针由 app 持有，离开区域写 -1。
- 无障碍：图表的 `Label`/`Description` 写上去，选中值由 `Select` 的 `describe` 读给读屏。
- 无计时器、无 goroutine、无包级状态；动效走 `myui.StateTransition()` 与框架 Reduce Motion。

**后置**：多序列堆叠、时间轴 `NewTime`/对数轴 `NewLog` 的预设、缩放/平移、图例点击过滤——需要时作为新组件加入，不在预设里留半成品。

## 4. 项目结构是否合理

- 位置：`chart/` 是独立子包，与 `markdown/`、`ai/` 一样按需引入；根包不导入它，它只引入 `ui`、`ui/chart` 与根包 `myui`（具名导入，避免与根包 `Options` 撞名）。
- 依赖：仅 mygo 公开 API 与根包；不碰 internal，不带字体/图标资产。
- 状态：数值、标签与选中索引都在 app 一侧，本包不跨帧持有。

## 5. API 契约

| API | 作用 |
|---|---|
| `type Palette struct { Series []ui.Color; Grid, Axis, Surface ui.Color }` | 一帧的图表色板：序列色、网格、刻度、底色 |
| `func (p Palette) Color(i int) ui.Color` | 取第 i 个序列色，循环使用，空色板返回透明 |
| `func NewPalette(t *ui.Theme) Palette` | 从 my-ui 主题现建色板，不缓存 |
| `type Options struct` | 预设的可选项（下列六个字段） |
| `Options.Labels []string` | 类别轴名称，缺的按 1、2、3… 补齐 |
| `Options.Selected *int` | 选中索引指针；写回当前项，离开写 -1 |
| `Options.Height float32` | 绘图高度；0 取 45 个主题间距单位 |
| `Options.Curve mchart.Curve` | 折线插值；0 为直线 |
| `Options.Label string` | 无障碍名称，空为 "Chart" |
| `Options.Description string` | 无障碍描述 |
| `func NewBars(c *ui.Context, values []float64) ui.Element` | 即用柱状图（零 `Options`） |
| `func NewLine(c *ui.Context, values []float64) ui.Element` | 即用折线图（零 `Options`） |
| `func NewArea(c *ui.Context, values []float64) ui.Element` | 即用面积图（零 `Options`） |
| `func NewPie(c *ui.Context, values []float64) ui.Element` | 即用饼图与图例（零 `Options`） |
| `func Bars(c *ui.Context, values []float64, o Options) ui.Element` | 带标签/选中/高度/名称的柱状图 |
| `func Line(c *ui.Context, values []float64, o Options) ui.Element` | 带插值与选中的折线图 |
| `func Area(c *ui.Context, values []float64, o Options) ui.Element` | 带插值与选中的面积图 |
| `func Pie(c *ui.Context, values []float64, o Options) ui.Element` | 带标签与高度的饼图 |

用法：

```go
chartpkg.NewBars(c, []float64{18, 32, 24, 41})

idx := -1 // app 持有
chartpkg.Line(c, sales, chartpkg.Options{
    Labels:   months,
    Selected: &idx,
    Curve:    mchart.CurveMonotone,
})
```

## 6. 状态

- [x] 调研与规格书（按 mygo v0.10.0 实际 API 对齐）
- [x] 实现（`chart.go`：色板 + 四个即用预设 + 四个可配置预设）
- [x] 测试（`chart_test.go`：色板取色、四种预设绘制、空数据、标签与曲线、选中指针、主题切换）
- [x] showcase「Chart」页（`page_chart.go`，函数 `chartPage`，由上层注册）
