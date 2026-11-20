# Button 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

shadcn 式变体词汇的样板：后来的控件照它的形状做。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewButton(c *ui.Context, label string) *Button` | 创建按钮（`label` 留空即只有图标） |
| `Variant(v Variant)` · `Default()` · `Native()` · `Secondary()` · `Primary()` · `Outline()` · `Ghost()` · `Destructive()` · `Link()` · `Success()` · `Info()` · `Warning()` | 变体：底色、边框、文字色与 hover/pressed 三态 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 尺寸：内衬与字号 |
| `Shape(s Shape)` | 形态：`ShapeRect`（默认）/ `ShapeRound`（胶囊）/ `ShapeCircle`（圆形） |
| `Icon(svg *ui.SVG)` | 文字之前的图标（取文字色与字号） |
| `EndIcon(svg *ui.SVG)` | 文字之后的图标 |
| `Busy(on bool)` | 忙碌态：`on` 时转圈并停止响应点击，传 false 恢复图标与点击 |
| `Loading()` | 忙碌态简写，等同 `Busy(true)` |
| `Toggle(pressed *bool)` | 双态按钮（shadcn Toggle）：点击翻转 `*pressed`，无障碍名带按下态 |
| `NewButtonGroup(c *ui.Context, build func()) ui.Element` | 按钮组：把若干按钮排成共享边的一行 |
| `NewToggleGroup` · `NewSegmented` | mygo 签名的按钮组与分段控件，现由 [toggle-group.md](toggle-group.md) / [segmented.md](segmented.md) 各自带 my-ui 外观 |
| `Label(s string)` | 无障碍名（只有图标时必须给） |
| `Width(w float32)` | 宽度（DIP） |
| `Block()` | 占满整行 |

变体词表（与 shadcn 对齐）：

| 变体 | 观感 |
|---|---|
| `Default`（零值） | 控件自带的观感：主题的面 + 描边（mygo 的按钮外观） |
| `Secondary` | 弱填充面：主题的面，无描边 |
| `Primary` | 强调色填充 |
| `Outline` | 背景上的描边面 |
| `Ghost` | 平时无底，悬停才现 |
| `Destructive` | 危险色填充 |
| `Link` | 纯文字，强调色 |
| `Success`/`Info`/`Warning` | 语义色填充（主题层的三种语义色，明暗各一套） |

`Native` 是 `Default` 的旧名，保留以免旧代码失效。

用法：

```go
ctrlpkg.NewButton(c, "保存").Primary()
ctrlpkg.NewButton(c, "删除").Destructive().Sm()
ctrlpkg.NewButton(c, "保存").Icon(icons.Save()).Primary()
ctrlpkg.NewButton(c, "").Icon(icons.Settings()).Label("设置")   // 只有图标
ctrlpkg.NewButton(c, "筛选").Outline().Shape(myui.ShapeRound)
```
