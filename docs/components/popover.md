# Popover 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
type PopoverOptions struct {
    Open         *bool             // app 持有
    Trigger      func() ui.Element // 建锚点，先于浮层
    Title        string            // 标题兼面板无障碍名
    Description  string            // 弱化说明
    Body         func()            // 库放进 panel.Children
    Footer       func()            // 右对齐操作行
    OnOpenChange func(bool)        // 开合回调，跨帧保留
    OnDismiss    func()            // 点外面/Escape 关闭
    Side         *PopoverSide      // nil=PopoverBottom
    ShowClose    *bool             // nil=显示
}
func NewPopover(c *ui.Context, o PopoverOptions) *Popover
type PopoverSide = TooltipSide
const (
    PopoverTop PopoverSide = iota
    PopoverBottom
    PopoverLeft
    PopoverRight
)
```

| 属性（方法/字段） | 类型 | 作用 |
|---|---|---|
| `Open` | `*bool` | app 的开关 |
| `Trigger` | `func() ui.Element` | 建锚点；点击切换开合，当帧生效 |
| `Title` / `Description` | `string` | 面板标题 / 灰色说明 |
| `Body` | `func()` | 面板正文，由库放进 `panel.Children` |
| `Footer` | `func()` | 右对齐操作行 |
| `OnOpenChange` | `func(bool)` | 任意来源的开合；晚挂不丢事件 |
| `OnDismiss` | `func()` | 点外面或 Escape 关闭 |
| `Side` | `*PopoverSide` | 面板出现在锚点哪一侧；nil 为下侧 |
| `ShowClose` | `*bool` | 关掉；nil 为开 |
| `Open() bool` | 方法 | 当前是否显示 |
| `SetOpen(bool) *Popover` | 方法 | 改 app 的 flag |
| `Dismissed() bool` | 方法 | 本帧是否被点外面/Escape 关闭 |
| `Side(PopoverSide) *Popover` | 方法 | 改锚点方向，关闭期间也有效 |
| `Title/Description(string) *Popover` | 方法 | 关闭期间也可设 |
| `Footer/ShowClose` | 方法 | 操作行与 |
| `OnOpenChange(func(bool)) *Popover` | 方法 | 注册开合回调；补发之前的改变 |
| `OnDismiss(func()) *Popover` | 方法 | 注册关闭回调 |

用法：

```go
side := overlaypkg.PopoverBottom
overlaypkg.NewPopover(c, overlaypkg.PopoverOptions{
    Open: &app.moreOpen,
    Trigger: func() ui.Element { return ctrlpkg.NewButton(c, "更多").Element },
    Body:    func() { ui.Text(c, "面板里的内容") },
    Side:    &side,
})
```

## 状态与限制

**限制**

- 触发方式只有点击（框架如此）。
- 只有四向位置（`PopoverTop/Bottom/Left/Right`），无 `sideOffset`/align 细调（框架不提供）。
