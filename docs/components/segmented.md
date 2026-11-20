# Segmented 规格书

> 位置：`controls/segmented.go`。

参考 Lobe UI 的 Segmented：几选一的轨道，指示器滑向选中的那一段；签名兼容 mygo 的 `ui.Segmented`。

## shadcn 风格

轨道是 shadcn Tabs 的弱底发丝边：Filled 用 `t.Surface` + 1px `t.Border`，Outlined
保留窗口底色 + 1px `t.Border`；选中的那一段用 `t.Accent.Alpha(0.10)` 的弱强调底，
其余文字 `t.TextMuted`、选中与悬停转 `t.Text`。圆角用 `Radius(t, RadiusMd/Sm)`，
内边距与间距用 `t.Space`，字号 `TextSize(t, TextXs…TextLg)`，滑行时长走
`MotionNormal`。没有 Lobe 固定色板与自造像素尺寸。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewSegmented(c *ui.Context, selected *int, labels ...string) *Segmented` | 建分段控件，选中的下标写入 `*selected` |
| `Variant(v SegmentedVariant)` · `Filled()` · `Outlined()` | 两种轨道 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 共享尺寸 |
| `Block(on bool)` | 占满整行并平分宽度 |
| `Label(name string)` | 整组无障碍名 |
| `Changed() bool` | 本帧是否换了段 |

## 限制

- 只接受文字标签；图标段可用 `ui.SegmentedBase` 自行构建，本构造器不支持。
- 滑行来自框架 transition，减少动态效果时按框架约定瞬间到位。
