# ActionIconGroup 规格书

> 位置：`controls/actionicongroup.go`。

参考 Lobe UI 的 ActionIconGroup：一组共用底面的图标动作，返回被点中的 Key。

## shadcn 风格

组容器与项沿用 ActionIcon 的 shadcn 形态：Filled 是 `t.Surface` 弱底、Outlined 是
`t.Background` + 1px `t.Border`、Borderless 透明且悬停转 `t.SurfaceHover`。间距与内边距
用 `t.Space`，圆角用 `Radius(t, RadiusMd)`。每项由 `NewIconAction` 承载悬停/按下墨色，
颜色均取当帧 `c.Theme()` 的 zinc 槽位；不含 Lobe 固定色板或自造尺寸。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewActionIconGroup(c *ui.Context, chosen *string, items []ActionIconGroupItem) *ActionIconGroup` | 建组，点击的 Key 写入 `chosen` |
| `ActionIconGroupItem{Key, Icon, Label, Danger, Disabled}` | 一项的键、图标、名字、危险色、禁用 |
| `Variant(v ActionIconGroupVariant)` · `Filled()` · `Outlined()` · `Borderless()` | 三种面 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 图标的共享尺寸 |
| `Vertical(on bool)` | 竖排 |
| `Disabled(on bool)` | 整组禁用 |
| `Chosen() string` | 本帧点中的 Key，未点中为 `""` |

## 限制

- 不内置 More 菜单与阴影预设；需要时用 `overlays` 的菜单组件组合。
- `Chosen()` 只反映当前帧，跨帧读取应看传入的 `chosen` 指针。
