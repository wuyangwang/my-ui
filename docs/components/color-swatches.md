# ColorSwatches 规格书

> 位置：`controls/colorswatches.go`。

参考 Lobe UI 的 ColorSwatches：一排色块选一个颜色，选中的带勾，透明色显示棋盘格。

## shadcn 风格

色块的颜色是数据本身，其余全部取当帧 `c.Theme()` 的 shadcn zinc 槽位：未选中的环用
`t.Border`、选中环用 `t.Text`、悬停外圈用 `t.Text`，透明块的棋盘格用 `t.Surface` 与
`t.SurfaceHover`，勾的墨色由色块自身经 `OnAccent` 派生。圆角用 `Radius(t, RadiusMd)`
（圆形为边长一半），默认边长是 `t.Space(SpacingXl)`（默认主题下 24 DIP），随主题间距
缩放。不含 Lobe 固定色板与自造尺寸。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewColorSwatches(c *ui.Context, value *ui.Color, swatches ...ColorSwatch) *ColorSwatches` | 建一排色块，选中色写入 `*value` |
| `ColorSwatch{Color, Label}` | 一个色块的颜色与无障碍名 |
| `Size(d float32)` | 色块边长，默认 24 DIP |
| `Square(on bool)` | 方角而非圆形 |
| `Wrap(on bool)` | 允许换行 |

## 限制

- 不内置自定义颜色入口；Lobe 的 `EnableColorPicker` 在这里由应用组合 `controls.NewColorPicker` 实现。
- `ui.RadioGroup` 不向组本身传播 `Changed`；需要知道变化时比较传入的 `*value`。
