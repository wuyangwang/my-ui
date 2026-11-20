# HoverCard 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
type HoverCardOptions struct {
    Open         *bool             // app 持有；nil 时用组件自己的稳定状态
    Trigger      func() ui.Element // 建悬停目标，先于浮层
    Body         func()            // 库放进 panel.Children
    OnOpenChange func(bool)        // 开合回调，跨帧保留
    Key          string            // 多个卡片同父且无 Open 时区分身份
}
func NewHoverCard(c *ui.Context, o HoverCardOptions) *HoverCard
```

| 属性（方法/字段） | 类型 | 作用 |
|---|---|---|
| `Open` | `*bool` | app 的开关；nil 时用组件自带状态 |
| `Trigger` | `func() ui.Element` | 建悬停目标；先于浮层构建 |
| `Body` | `func()` | 面板正文，由库放进 `panel.Children` |
| `OnOpenChange` | `func(bool)` | 开合回调；晚挂补发之前的改变 |
| `Key` | `string` | 多个卡片共用父级且无 `Open` 时区分身份 |
| `Open() bool` | 方法 | 当前是否显示 |
| `OnOpenChange(func(bool)) *HoverCard` | 方法 | 注册开合回调 |

用法：

```go
overlaypkg.NewHoverCard(c, overlaypkg.HoverCardOptions{
    Open:    &app.cardOpen,
    Trigger: func() ui.Element { return ctrlpkg.NewButton(c, "悬停我").Element },
    Body:    func() { ui.Text(c, "卡片内容") },
})
```

## 状态与限制

**限制**

- 只有悬停触发，没有点击/焦点触发（组件定位如此）。
- 宽限是固定帧数，不可配。
