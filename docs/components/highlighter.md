# Highlighter

带主题色着色、复制按钮与语言标签的代码块；可选完整头部与折叠。

```go
display.NewHighlighter(c, src, display.HighlighterOptions{
    Language: "go", Variant: display.CodeOutlined, FullFeatured: true,
    FileName: "main.go", Collapsed: &app.collapsed,
})
```

## 行为

- 源码先去掉首尾空行，再经一个小型分词器着色：字符串、注释（含 `#`/`--`/`/* */`）、数字、关键字与标点，其余取正文墨色。
- 颜色都从当帧 `c.Theme()` 派生（强调色、语义色、次要墨色），明暗与自定义主题即时跟随。
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
