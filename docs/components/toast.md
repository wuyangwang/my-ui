# Toast 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
type ToastOptions struct {
    Title       string        // 标题，读屏念出
    Description string        // 描述
    Kind        Variant       // Success/Info/Warning/Destructive
    Duration    time.Duration // 0=默认 4s；负数=直到关闭
    ActionLabel string        // 动作按钮文案
    OnAction    func()        // 动作回调，随后关闭
    OnClose     func()        // 关闭回调
}
type Notifier struct{ /* 持有窗口的稳定 Context */ }
func NewNotifier(c *ui.Context) *Notifier
func (n *Notifier) Show(o ToastOptions) string
func (n *Notifier) Close(id string)
func (n *Notifier) Build()                 // 根视图每帧调一次

// 旧链式 API（内部仍用）：
func NewToastViewport(c *ui.Context, opts ...ViewportOption) *ToastViewport
func ViewportPosition(p ToastPosition) ViewportOption
func NewToast(c *ui.Context, title string) *Toast
func (t *Toast) Description/Variant/Success/Info/Warning/Destructive/Timeout/Action/ID/OnClose ...
func (t *Toast) Show() string
```

| 属性（方法/字段） | 类型 | 作用 |
|---|---|---|
| `Title` | `string` | 消息标题，由 `AddToast` 念出 |
| `Description` | `string` | 标题下的描述 |
| `Kind` | `Variant` | 决定图标与语义色（success/info/warning/error） |
| `Duration` | `time.Duration` | 停留时间；0 为 4 秒，负数为不自动关 |
| `ActionLabel` / `OnAction` | `string` / `func()` | 动作按钮；按下后回调并关闭 |
| `OnClose` | `func()` | 关闭时回调 |
| `Show(ToastOptions) string` | 方法 | 显示 toast，返回 id |
| `Close(id string)` | 方法 | 关掉 id（空串关全部） |
| `Build()` | 方法 | 在根视图装自绘视口 |
| `ViewportPosition` | 选项 | 改堆叠角落；默认右下 |

用法：

```go
// 视图里：建一次句柄，每帧 Build：
if s.notifier == nil { s.notifier = overlaypkg.NewNotifier(c) }
s.notifier.Build()

// 动作里（只捕获 notifier）：
s.notifier.Show(overlaypkg.ToastOptions{
    Title: "已保存", Kind: myui.Success, ActionLabel: "撤销",
    OnAction: func() { s.undo() },
})
```

## 状态与限制

**限制**

- 一次最多三条（框架限制）；堆叠不做位移动画。
- toast 内容只能是标题、描述、动作与关闭（框架部件就这些）。
