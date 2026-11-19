# Image

带圆角边框的图片，点击打开覆盖窗口的预览。边框、占位与预览面板只用当前帧 `c.Theme()` 的 shadcn 槽位，圆角走 `Radius` 词表。

```go
display.NewImage(c, src, display.ImageOptions{
    Size: 160, Alt: "封面", Variant: myui.Outline,
})
```

## 行为

- `Width`/`Height` 或 `Size` 定尺寸；都不设时按图片自身尺寸、最多占满父宽。
- `Contain` 让整张图放进盒子，否则裁切填满。
- `Loading` 显示 `t.SurfaceHover` 占位块，`Failed`（或 `src == nil`）显示 `t.TextMuted` 图标占位（图标尺寸为 `IconSizeLg`）。
- `Alt` 写入无障碍名称；`Actions` 在悬停（或 `AlwaysShowActions`）时显示在右上角，动作按钮由调用方用 `controls.NewIconAction` 构建。
- 可点击的图片打开预览层；预览遮罩是**从主题派生的中性暗罩**：由 `previewScrim(t)` 按明暗取 `t.Text`/`t.Background` 再压 alpha，不使用固定黑色。预览面板用 `t.Background` + 1px `t.Border` + `Radius(t, RadiusXl)`。
- `OnOpen`/`OnClose` 由 app 观察开关；颜色与圆角每帧从 `c.Theme()` 取得，无计时器与 goroutine。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewImage(c, src, o)` | `*ui.Context, ui.ImageSource, ImageOptions` | 构造图片 |
| `ImageOptions.Width/Height/Size` | `float32` | 盒子尺寸；`Size` 同时设两者 |
| `ImageOptions.Contain` | `bool` | 完整放入而非裁切填满 |
| `ImageOptions.Variant` | `myui.Variant` | `Default` 淡底、`Outline` 边框、`Ghost` 无面 |
| `ImageOptions.Loading` | `bool` | 加载占位 |
| `ImageOptions.Failed` | `bool` | 失败占位 |
| `ImageOptions.NoPreview` | `bool` | 禁用点击预览 |
| `ImageOptions.Alt` | `string` | 无障碍名称 |
| `ImageOptions.Actions` | `func()` | 右上角动作区内容（用 `controls.NewIconAction`） |
| `ImageOptions.AlwaysShowActions` | `bool` | 一直显示动作区 |
| `ImageOptions.OnOpen` / `.OnClose` | `func()` | 预览打开/关闭回调 |

## 限制

- 预览是简单的大图面板，不带 Lobe 的滚轮/捏合缩放、拖拽平移、画廊切换与旋转；需要这些能力时由 app 在 `Actions` 中自行提供。
- 图片解码由调用方完成（`ui.NewBitmap`、`ui.DecodeImage` 等），组件只接收 `ui.ImageSource`。
- 预览的开关是组件内部状态（`ui.Local`），不暴露给 app 写入；`OnOpen`/`OnClose` 只作通知。
