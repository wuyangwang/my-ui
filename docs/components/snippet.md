# Snippet

单行代码（通常是命令），带前缀与复制按钮；过长时横向滚动而不换行。

```go
display.NewSnippet(c, "go get github.com/wuyangwang/my-ui", display.SnippetOptions{
    Language: "sh", Prefix: "$",
})
```

## shadcn 风格

与 Highlighter 同一套 shadcn 代码块观感：`t.Surface` 弱底、1px `t.Border`
（`CodeOutlined`）、`Radius(t, RadiusMd)`，软阴影用 `t.Text.Alpha(0.06)`；分词颜色取
当帧主题的语义槽位与次要墨色，是弱墨代码色而非 Lobe 固定配色。复制按钮用
`controls.NewIconAction`。字号 `TextSize(t, TextXs/TextSm)`，间距 `t.Space`，
高度 `t.Space(10)` 的主题步长，不含固定色板与自造像素尺寸。

## 行为

- `Prefix` 显示在代码前（如 `$`），但复制时只包含 `src`。
- 着色复用 Highlighter 的小型分词器，颜色来自当帧主题。
- `Variant` 选择淡底、带框或无面；`NoCopy` 去掉复制按钮。
- 复制写系统剪贴板并调用 `OnCopy(src)`；没有计时器或 goroutine。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewSnippet(c, src, o)` | `*ui.Context, string, SnippetOptions` | 构造单行代码 |
| `SnippetOptions.Language` | `string` | 着色语言提示 |
| `SnippetOptions.Prefix` | `string` | 代码前的前缀，不参与复制 |
| `SnippetOptions.Variant` | `CodeVariant` | `CodeFilled`/`CodeOutlined`/`CodeBorderless` |
| `SnippetOptions.NoCopy` | `bool` | 去掉复制按钮 |
| `SnippetOptions.Shadow` | `bool` | 柔和阴影 |
| `SnippetOptions.FontSize` | `float32` | 代码字号（DIP），0 用主题字号 |
| `SnippetOptions.Label` | `string` | 无障碍名称，默认 `Code snippet` |
| `SnippetOptions.OnCopy` | `func(src string)` | 复制后回调 |

## 限制

- 使用与 Highlighter 相同的简化分词器，不是完整语法高亮。
- 复制后不显示“已复制”的状态提示（无计时器），由 `OnCopy` 交给 app。
