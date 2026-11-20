# ActionIcon 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题与 tokens 仍在根包 `myui`。

shadcn 式中「只有图标的按钮」，参考 Lobe UI 的 ActionIcon 三种面与尺寸档。

## shadcn 风格

三种面对齐 shadcn Button 的 ghost / secondary / outline：Borderless 为 0 边框透明底，
Filled 用 `t.Surface`，Outlined 用 `t.Background` + 1px `t.Border`。悬停一律
`t.SurfaceHover`、按下 `t.SurfacePressed`，文字 `t.TextMuted` 悬停转 `t.Text`；
`Active` 的选中态与 `Danger` 的危险态都从当帧主题派生。圆角用
`Radius(t, RadiusMd)`（圆形形态为全圆），图标用 `IconSizeXs/Sm/Md/Lg` 档，
间距用 `t.Space`。没有 Lobe 固定色板与自造像素尺寸。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewActionIcon(c *ui.Context, svg *ui.SVG, title string) *ActionIcon` | 建图标按钮，`title` 为无障碍名 |
| `Variant(v ActionIconVariant)` · `Borderless()` · `Filled()` · `Outlined()` | 三种面 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 共享尺寸（28/32/36/40 DIP 方框） |
| `Shape(s Shape)` | 圆角或圆形 |
| `Active(on bool)` | 保持悬停底，用于已切换或已展开状态 |
| `Danger(on bool)` | 图标用危险色 |
| `Label(s string)` | 无障碍名 |
| `Clicked()` | 沿用 `ui.Element` 的点击事件 |

## 限制

- 不自带 tooltip；需要提示时用 `overlays.NewTooltip(c, icon.Element, "…")`。
- 尺寸档决定方框与字形，字形取共享图标尺寸，不提供任意 DIP 字形。
