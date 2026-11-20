# Tooltip 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
// NewTooltip shows a tip with a text over the anchor.
func NewTooltip(c *ui.Context, anchor ui.Element, text string, options ...TooltipOptions) *Tooltip
// NewTooltipContent shows a tip built by build over the anchor.
func NewTooltipContent(c *ui.Context, anchor ui.Element, build func(tip ui.Element), options ...TooltipOptions) *Tooltip

type TooltipOptions struct {
	Side   TooltipSide
	Hotkey string
}
```

`TooltipOptions` 在构造时给值（首帧即可见），下面这些方法给同样的值，可链：

| 属性（方法） | 作用 |
|---|---|
| `Side(s TooltipSide) *Tooltip` | 提示出现在锚点的哪一侧：`TooltipTop`（默认）/`TooltipBottom`/`TooltipLeft`/`TooltipRight`；空间不足时框架翻到对侧 |
| `Hotkey(hotkey string) *Tooltip` | 文字后显示快捷键胶囊，写法同 `display.Hotkey`（如 `"mod+k"`）；`mod` 在 macOS 显示 `Cmd`、其他平台显示 `Ctrl`；单字母大写。只对 `NewTooltip` 的文字提示生效 |
| `Content(build func(tip ui.Element)) *Tooltip` | 替换提示内容；`build` 必须用 `tip.Children(func(){ ... })` 把内容建在提示里 |

用法：

```go
overlaypkg.NewTooltip(c, ctrlpkg.NewButton(c, "帮助").Element, "这是帮助文字").
	Hotkey("mod+k").Side(overlaypkg.TooltipBottom)
```

## 状态与限制

**限制**

- 提示的小箭头（框架没有，shadcn 有）。
- 触发方式只有悬停与聚焦（框架如此；点击触发需 app 自己接）。
- 只有文字内容（回调里能建更多，API 未开放）。
