# Chart 规格书

> 位置：组件在 [`chart/`](../../chart) 子包；主题/tokens 仍在根包 `myui`。
>

## API 契约

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
