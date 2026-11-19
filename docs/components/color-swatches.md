# ColorSwatches 规格书

> 位置：`controls/colorswatches.go`。

> 状态：**已实现**。参考 Lobe UI 的 ColorSwatches：一排色块选一个颜色，选中的带勾，透明色显示棋盘格。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [Lobe UI ColorSwatches](https://mygo.lobehub.com/components/color-swatches) | RadioGroup + RadioBase、选中环、棋盘格、方向键 |
| `controls/checkbox.go` | 选中态的颜色与勾选墨色派生方式 |
| `icons.Check()` | 选中勾，ISC 授权的 Lucide 图标 |

## 2. 能否实现

- `ui.RadioGroup` + `ui.RadioBase[T comparable]` 绑定 `ui.Color`（可比较）✓。
- 圆/方角由 `Radius(t, RadiusMd)` 与边长一半切换 ✓。
- 透明色用底层 `t.Surface` 加两块 `t.SurfaceHover` 棋盘格表现 ✓。
- 勾选墨色按颜色亮度取黑/白半透明，主题为空槽时回退到 `t.Text` ✓。
- 自定义颜色按钮需要浮层与原生颜色面板，`controls` 不导入 `overlays`；应用可用 `controls.NewColorPicker` 自行组合 ✓。

## 3. 实现到何种程度

圆形/方形、任意边长、每色无障碍名、悬停外圈、选中勾、透明棋盘格、方向键与 radio 组角色。`Square`、`Size`、`Wrap` 的 fluent 方法直接改已建成的元素，下一帧仍跟随主题。

## 4. 项目结构

- `controls/colorswatches.go`；`controls/colorswatches_test.go` 覆盖绑定、名字与尺寸。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewColorSwatches(c *ui.Context, value *ui.Color, swatches ...ColorSwatch) *ColorSwatches` | 建一排色块，选中色写入 `*value` |
| `ColorSwatch{Color, Label}` | 一个色块的颜色与无障碍名 |
| `Size(d float32)` | 色块边长，默认 24 DIP |
| `Square(on bool)` | 方角而非圆形 |
| `Wrap(on bool)` | 允许换行 |

## 6. 限制

- 不内置自定义颜色入口；Lobe 的 `EnableColorPicker` 在这里由应用组合 `controls.NewColorPicker` 实现。
- `ui.RadioGroup` 不向组本身传播 `Changed`；需要知道变化时比较传入的 `*value`。
