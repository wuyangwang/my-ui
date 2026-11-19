# HoverCard 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`hover_card.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui HoverCard](https://ui.shadcn.com/docs/components/hover-card)（`ui/hover-card.tsx`） | `HoverCardTrigger` + `HoverCardContent`（`Portal` 定位的卡片） | 触发器 + 浮层面板 ✓ |
| [Element Plus el-popover](https://element-plus.org/en-US/component/popover.html)（`trigger="hover"`） | 悬停打开，指针移到面板上保持打开 | 悬停判定 + 几帧宽限 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 浮层定位与层级 | `ui.PopoverBase(c, anchor, open, fn)` ✓ | ✅ |
| 悬停判定 | `Element.Hovered()` ✓ | ✅ |
| 停留宽限 | `ui.Local` 存帧计数 ✓（无计时器） | ✅ |
| 回调跨帧 | `ui.Local` 存回调与待投递变化 ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `NewHoverCard(c, HoverCardOptions)`：`Trigger` 由库先于浮层构建 ✓；`Body func()` **由库包在 `panel.Children` 里**，调用方不必知道面板的父层是覆盖层 ✓。
- 触发区或面板 `Hovered()` 时保持打开；两者都离开后数帧（帧计数，非计时器）再关闭 ✓。
- `OnOpenChange` 跨帧保留，晚挂补发之前的改变 ✓；颜色与尺寸每帧从 `c.Theme()` 取 ✓。

**已知限制**
- 只有悬停触发，没有点击/焦点触发（组件定位如此 ✗）。
- 宽限是固定帧数，不可配 ✗。

**实现要点（踩过的坑）**
- 旧版把挂 `panel.Children` 的责任留给调用方，内容会落到窗口左上角；新版由库统一包住 `Body` ✓。

## 4. 项目结构是否合理

- 位置：`overlays/` 子包 ✓。
- 依赖：`mygo/ui` + 根包 tokens；与其余浮层共用 `openState` ✓。

## 5. API 契约

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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`hover_card.go`）
- [x] 测试（悬停打开、移开数帧后关闭 ✓、正文落在面板 ✓、晚挂回调 ✓）
