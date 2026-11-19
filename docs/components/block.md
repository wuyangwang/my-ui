# Block

带圆角的面块，作为网格里的卡片式容器；可点击时是按钮，悬停与按下的颜色走共享状态过渡。

```go
b := display.NewBlock(c, display.BlockOptions{Clickable: true, Label: "Open"}).
    Label("打开").Outline()
b.Children(func() { ui.Text(c, "内容") })
if b.Clicked() { /* ... */ }
```

## shadcn 风格

Block 是 shadcn Card 的弱底形态：表面 `t.Background`/`t.Surface`，`Outline` 加 1px
`t.Border`，`Ghost` 为 0 边框；`Default` 静止是 `t.Surface`、悬停转 `t.SurfaceHover`
（不再以悬停色做静止底），所有颜色取当帧 `c.Theme()` 的 zinc 槽位。圆角
`Radius(t, RadiusLg)`，`Shadow` 用与 `display.NewCard` 同源的 `t.Text.Alpha` 两层软阴影，
状态变化经 `StateTransition()` 淡入；间距用 `t.Space`，不含 Lobe 固定色板与自造尺寸。

## 行为

- `Clickable` 时用 `ui.ButtonBase`，否则用 `ui.Box`；点击回调 `OnChanged`，也可直接读返回元素的 `Clicked()`。
- `Variant` 使用根包共享词表：`Default` 淡底、`Outline` 带边框、`Ghost` 无面、`Secondary` 主题面。
- 悬停只改变底色，颜色经 `StateTransition()` 过渡；`Shadow` 加两层柔和阴影。
- `Disabled` 透传给元素，禁用时不触发 `OnChanged`。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewBlock(c, o)` | `*ui.Context, BlockOptions` | 构造块 |
| `BlockOptions.Variant` | `myui.Variant` | 共享变体词表 |
| `BlockOptions.Clickable` | `bool` | 让块成为按钮 |
| `BlockOptions.Shadow` | `bool` | 柔和阴影 |
| `BlockOptions.Label` | `string` | 无障碍名称 |
| `BlockOptions.Disabled` | `bool` | 禁用点击 |
| `BlockOptions.OnChanged` | `func()` | 点击时调用 |
| `(*Block).Variant(v)` | `myui.Variant` | 设置变体 |
| `(*Block).Outline()` / `.Ghost()` | — | 变体快捷方法 |

`Clickable` 在 `BlockOptions` 中构造时设置；流式方法只调整外观与状态，不重建元素。
| `(*Block).Shadow(on)` | `bool` | 设置阴影 |
| `(*Block).Label(s)` | `string` | 设置无障碍名称 |
| `(*Block).Disabled(off)` | `bool` | 设置禁用 |
| `(*Block).OnChanged(fn)` | `func()` | 设置点击回调 |

## 限制

- 点击是瞬时动作，没有“选中”状态；需要持续状态时由 app 持有指针并据此改变 `Variant`。
- 阴影与圆角取自主题，不提供逐边配置。
