# Divider

内容之间的分隔线：横向、带文字、虚线，或行内竖向。

```go
display.NewDivider(c)
display.NewDivider(c, display.DividerOptions{Dashed: true})
display.NewDivider(c, display.DividerOptions{Label: "OR"})
```

## shadcn 风格

Divider 是 shadcn Separator：线是 `t.Border` 的 1 DIP 发丝弱墨（不是深色描边），
`Label` 文字用 `t.TextMuted` 并垫 `t.Surface` 弱底，发丝线不穿过文字；字号
`TextSize(t, TextXs)`，间距 `t.Space`。颜色每帧从 `c.Theme()` 取得，明暗与自定义主题
即时跟随。没有 Lobe 固定色板与自造尺寸。

## 行为

- 横向线占满父宽，高度 1 DIP；`Label` 时两侧各一段grow的线与中间文字。
- `Vertical` 画 1 DIP 宽、主题字号高的竖线，居中对齐，用于行内分隔。
- `Dashed` 通过 `Draw` 用 `Painter.Line` 画 4/4 DIP 的虚线，横竖自适应。
- 颜色取 `t.Border`，文字取 `t.TextMuted`，间距取 `t.Space`，每帧从主题重算。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewDivider(c, opts...)` | `*ui.Context, ...DividerOptions` | 构造分隔线；省略 options 为普通横线 |
| `DividerOptions.Vertical` | `bool` | 竖线 |
| `DividerOptions.Dashed` | `bool` | 虚线 |
| `DividerOptions.Label` | `string` | 横线中央文字 |
| `(*Divider).Vertical()` / `.Dashed()` / `.Label(s)` | — | 流式设置 |

## 限制

- 虚线长度固定为 4/4 DIP，不随主题间距缩放；需要不同节奏时可用 `Draw` 自行覆写。
- 竖线高度取主题字号，不提供逐像素高度配置。
