# Text

按 my-ui 排版样式显示的文本，带语义色、标题层级、强调线与行数截断。

```go
display.NewText(c, "保存成功").Success()
display.NewText(c, "详情", display.TextOptions{Heading: 2}).Strong()
display.NewText(c, "说明文字", display.TextOptions{Type: display.TextSecondary, Rows: 1})
```

## 行为

- `Type` 选择语义墨色：默认、次要、成功、警告、危险、信息；`Variant` 是同一组角色的共享词表别名。
- `Heading`（1–5）按主题字号的 2.5/2/1.5/1.25/1 倍设置字号并加粗、行高 1.25；`FontSize` 可覆盖。
- `Strong`/`Italic`/`Underline`/`Delete` 控制字重、斜体与线饰；`Code` 用等宽字体；`Mark` 用主题强调色高亮。
- `Disabled` 用次要墨色并给出禁止光标；`Rows` 限制行数（1 行省略号，多行 `MaxLines`）。
- 颜色、字号与间距每帧从 `c.Theme()` 取得，跟随主题、文字缩放与高对比度。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewText(c, s, opts...)` | `*ui.Context, string, ...TextOptions` | 构造文本；省略 options 为正文 |
| `TextOptions.Type` | `TextType` | 语义色角色 |
| `TextOptions.Variant` | `myui.Variant` | 共享角色词表（`Type` 为默认时生效） |
| `TextOptions.Heading` | `int` | 标题层级 1–5 |
| `TextOptions.Size` | `myui.Size` | 共享尺寸词表 |
| `TextOptions.FontSize` | `float32` | 覆盖字号（DIP） |
| `TextOptions.Strong/Italic/Underline/Delete` | `bool` | 字重、斜体、下划线、删除线 |
| `TextOptions.Code` | `bool` | 等宽字体 |
| `TextOptions.Mark` | `bool` | 强调色高亮 |
| `TextOptions.Disabled` | `bool` | 禁用态 |
| `TextOptions.Rows` | `int` | 行数上限 |
| `(*Text).Size(s)` | `myui.Size` | 设置共享尺寸 |
| `(*Text).Variant(v)` / `.Secondary()` / `.Success()` / `.Warning()` / `.Danger()` / `.Info()` | — | 语义色快捷方法 |
| `(*Text).Heading(n)` / `.Strong()` / `.Italic()` / `.Code()` / `.Mark()` / `.Disabled(off)` / `.Rows(n)` | — | 流式设置 |

## 限制

- 不实现 Lobe 的 `Shiny` 扫光动画：那需要计时器，桌面“减少动态效果”也难以统一，故省略。
- `Text` 是独立类型，与 `ui.Text` 同名的方法通过嵌入 `ui.Element` 保留；`Italic()`、`Disabled()` 被本组件覆盖以同时更新选项。
