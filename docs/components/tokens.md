# Tokens 规格书

my-ui 的共享设计刻度：主题色、字号、间距、圆角、图标尺寸与动效时长，全部从 `*ui.Theme` 解析，逐帧跟随主题与桌面偏好。

## API 契约

| 符号 | 作用 |
|---|---|
| `TextSize(t,n)` | 主题字号的 n 倍 |
| `Space(t,n)` | 主题间距的 n 单位 |
| `Radius(t,s)` | 主题圆角的 s 倍 |
| `IconSizeXs/Sm/Md/Lg` | 12/14/16/24 DIP 图标刻度 |
| `MotionFast/Normal/Slow` | 120/200/320ms 动效时长 |
| `Apply(c,o)` / `NewTheme(c,o)` | 把主题装进帧/子树 |
| `InverseColors(t)` | 主题留空时的反色兜底 |
