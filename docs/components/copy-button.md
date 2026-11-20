# CopyButton 规格书

> 位置：`controls/copybutton.go`。

参考 Lobe UI 的 CopyButton：点击复制文本，随后短暂显示对勾与「已复制」。

## shadcn 风格

按钮复用 ActionIcon 的三种 shadcn 面（ghost / secondary / outline），外观全部取当帧
`c.Theme()` 的 zinc 槽位：`t.Background`/`t.Surface` 弱底、1px `t.Border`、
`t.SurfaceHover` 悬停、墨色 `t.TextMuted` → `t.Text`。圆角用 `Radius(t, RadiusMd)`，
尺寸走共享 `Size`（28/32/36/40 DIP），字形走 `IconSize` 档，间距用 `t.Space`。
复制反馈是框架动画的时长，不引入颜色或尺寸常量。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewCopyButton(c *ui.Context, content string) *CopyButton` | 建按钮，点击复制 `content` |
| `Variant(v ActionIconVariant)` · `Borderless()` · `Filled()` · `Outlined()` | 三种面 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 共享尺寸 |
| `Labels(copy, copied string)` | 复制态与已复制态的无障碍名 |
| `Copied() bool` | 对勾是否还在显示 |

## 限制

- 不展示浮动提示；需要时用 `overlays.NewTooltip` 包裹。
- 两秒反馈时长固定为框架动画常量；应用若要别的时长需自行以 `AnimateTo` 组合。
