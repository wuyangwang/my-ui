# Tooltip 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`tooltip.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Tooltip](https://ui.shadcn.com/docs/components/tooltip)（源码 `ui/tooltip.tsx`） | `TooltipProvider`（`delayDuration` 延迟）· `TooltipTrigger`（`asChild` 触发器）· `TooltipContent`（`Portal` + `bg-foreground text-background rounded-md px-3 py-1.5 text-xs shadow-md`，`sideOffset`）· `TooltipPrimitive.Arrow`（小箭头） | 延迟取主题 ✓；触发器=锚点 ✓；外观照抄（主题翻转 `InverseColors` + `Space(1.5)/Space(3)` + `TextXs` + 圆角 + 阴影 ✓）；**箭头缺** ✗（框架没有 ✗） |
| mygo 的 `ui.TooltipBase(c, anchor, fn)` | 提示挂在**锚点**上（默认 `AttachTo(anchor, AnchorTop, AnchorBottom)`：上面放不下就放下面 ✓）；指针停留/键盘聚焦时出现 ✓；延迟由主题的 `TooltipDelay` 给 ✓；按下锚点则提示消失 ✓；`IgnoreInput` 让指针穿过提示 ✓ | 锚点 + 我们自绘的提示卡 ✓ |
| [Element Plus el-tooltip](https://element-plus.org/en-US/component/tooltip.html) | `content`/`raw-content`、`placement`（12 向）、`effect`（dark/light）、`show-after`、`trigger`（hover/click/focus）、`popper-class` | 位置与效果由框架给 ✓；触发方式固定为悬停+聚焦 ✓（框架如此 ✗） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 提示挂在锚点上 | `ui.TooltipBase(c, anchor Element, fn func(tip Element)) Element` ✓ | ✅ |
| 出现时机与延迟 | 指针停留/键盘聚焦 ✓，延迟取主题 ✓ | ✅ |
| 外观 | 提示是一个元素 ✓，我们自己排 ✓ | ✅ |
| 读屏念出提示 | `Element.Description(s)` ✓（框架的 `Tooltip` 也是这么做的 ✓） | ✅ |
| 箭头 | 框架的提示没有箭头 ✗ | ⚠️ 缺（框架限制 ✗） |
| 自定义内容 | 回调里能建任意内容 ✓（现在是文字 ✗） | ✅（可扩展） |

## 3. 实现到何种程度

**已实现**
- `NewTooltip(c, anchor, text)`：提示挂在锚点上 ✓，外观照 shadcn（`Inverse` 底 + `InverseText` 字、`Space(1.5)/Space(3)` 内衬、`TextXs`、圆角、阴影 ✓），最大宽度 `t.Space(60)` ✓。
- `Hotkey("mod+k")`：文字后画一排快捷键胶囊（`display.NewKbd(...).Xs()` ✓）；`mod` 按 `runtime.GOOS` 映射为 macOS 的 `Cmd`、其他平台的 `Ctrl` ✓；胶囊仍建在 `tip.Children` 里，不落窗口角 ✓（`tooltip_test.go` 的 `TestTooltipShowsItsHotkey` 覆盖 ✓）。
- `Side(s)`：复用 `tooltipAnchors` 的锚点映射，提示可落在上/下/左/右 ✓。
- 出现/消失由框架的时机与主题的延迟决定 ✓；按下锚点则消失 ✓；Escape 关闭 ✓。
- 提示文字同时成为锚点的 **description** ✓ → 读屏念出提示 ✓（`overlays_test.go` 的 `TestTooltipDescribesItsAnchor` 覆盖 ✓）。

**已知限制**
- 提示的小箭头 ✗（框架没有 ✗，shadcn 有 ✗）。
- 触发方式只有悬停与聚焦（框架如此 ✗；点击触发需 app 自己接 ✗）。
- 只有文字内容（回调里能建更多，API 未开放 ✗）。

**实现要点（踩过的坑）**
- `ui.TooltipBase` 的回调里，`c` 的父节点是**覆盖层**，不是提示本身 → 提示里的文字必须用
  `tip.Children(func(){ ui.Text(c, text) })` 建在提示里，否则文字落到窗口左上角 ✗。
  （回归测试 `TestTooltipShowsOverItsAnchor` 断言提示在锚点上方、不在窗口角上 ✓。）
- 主题的 `Inverse`/`InverseText` **留空时是零值**（框架内部自己兜底成 `Text`/`Background`）→
  直接读这两个字段会拿到透明色，提示只剩一个空框 ✗。用 `myui.InverseColors(th)` 取兜底后的颜色 ✓
  （回归测试 `TestTooltipDrawsItsText` 看像素 ✓）。

## 4. 项目结构是否合理

- 位置：根包 L3（变体化控件）——提示是「控件之上的附属物」✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`tooltip.go`）
- [x] 测试（指针停在锚点上提示出现 ✓、提示在锚点上方 ✓、锚点的 description ✓、Hotkey 文本出现且在提示内 ✓，构造 options 与 fluent 两种方式 ✓）
- [x] API 表（TooltipOptions/Side/Hotkey/Content）已补 ✓
- [ ] 箭头（框架限制 ✗）、自定义内容的 API（后置 ✗）
