# Highlighter

带主题色着色、复制按钮与语言标签的代码块；可选完整头部与折叠。

```go
display.NewHighlighter(c, src, display.HighlighterOptions{
    Language: "go", Variant: display.CodeOutlined, FullFeatured: true,
    FileName: "main.go", Collapsed: &app.collapsed,
})
```

## shadcn 风格

代码块是 shadcn 的 quiet surface：`t.Surface` 弱底、1px `t.Border`（`CodeOutlined`）、
`Radius(t, RadiusMd)`；`Shadow` 用与 `display.NewCard` 同源的 `t.Text.Alpha(0.06)` 软阴影，
不再是描边色的半透明。代码墨色走主题语义槽位（`t.Text`、`t.TextMuted` 与
Success/Warning/Info），弱墨而不固定，没有 Lobe 的代码配色表。复制按钮用
`controls.NewIconAction` 的纯图标 affordance。字号 `TextSize(t, TextXs/TextSm)`、
行距与间距 `t.Space`，不含固定色板与自造像素尺寸。

## 行为

- 源码先去掉首尾空行，再经一个小型分词器着色：字符串、注释（含 `#`/`--`/`/* */`）、数字、关键字与标点，其余取正文墨色。
- 颜色都从当帧 `c.Theme()` 派生（正文、次要墨色与主题语义色槽位），明暗与自定义主题即时跟随。
- `Wrap` 用纵向列换行，否则横向滚动；制表宽 4。
- 头部显示语言或文件名，折叠指针由 app 持有（`Collapsed`），为 nil 时组件自持状态。
- 复制按钮写系统剪贴板，并调用 `OnCopy(src)`；`Actions` 在其前面追加按钮。
- 无计时器、无 goroutine；悬停动作的显隐直接由 `Hovered()`/`FocusWithin()` 决定，不做淡入。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewHighlighter(c, src, o)` | `*ui.Context, string, HighlighterOptions` | 构造代码块 |
| `HighlighterOptions.Language` | `string` | 语言名或扩展名，仅用于着色与标签 |
| `HighlighterOptions.Variant` | `CodeVariant` | `CodeFilled`/`CodeOutlined`/`CodeBorderless` |
| `HighlighterOptions.NoCopy` | `bool` | 去掉复制按钮 |
| `HighlighterOptions.HideLanguage` | `bool` | 去掉语言标签 |
| `HighlighterOptions.Wrap` | `bool` | 换行而非横向滚动 |
| `HighlighterOptions.Shadow` | `bool` | 柔和阴影 |
| `HighlighterOptions.FullFeatured` | `bool` | 加头部与折叠 |
| `HighlighterOptions.FileName` | `string` | 头部标题，替代语言名 |
| `HighlighterOptions.Collapsed` | `*bool` | 折叠指针；nil 时组件自持 |
| `HighlighterOptions.Actions` | `func()` | 复制按钮前的额外动作 |
| `HighlighterOptions.FontSize` | `float32` | 代码字号（DIP），0 用主题字号 |
| `HighlighterOptions.OnCopy` | `func(src string)` | 复制后回调 |

## 限制

- 着色是共享关键字表的小分词器，不是 mygo 编辑器的三十种语言语法；未知语言按纯文本显示，但字符串/注释/数字仍能着色。
- 复制后没有“已复制”的短暂提示（那需要计时器）；用 `OnCopy` 让 app 自己反馈。
- 悬停动作即时显隐，不做 Lobe 的 200 ms 淡入。
