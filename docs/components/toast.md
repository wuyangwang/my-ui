# Toast 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`toast.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Toast](https://ui.shadcn.com/docs/components/toast)（底层 sonner） | `ToastProvider`（`position`/`duration`/`closeButton`）· `Toast`（卡片）· `ToastTitle`/`ToastDescription`/`ToastAction`/`ToastClose` · 四类 type | 视口 ✓（默认右下角，可改）✓；卡片 + 标题/描述/动作/关闭 ✓；类型给图标与语义色 ✓ |
| mygo 的 `ui.ToastViewportBase` / `ui.ToastBase` / `c.AddToast` / `c.CloseToast` | 视口由 app 建并取代主题观感 ✓；字段一一对应 ✓；最多三条 ✓；指针停在上面时时间停住 ✓；`AddToast` 自己念出标题与描述 ✓ | 字段对应 ✓；`Notifier.Show` 直接走 `AddToast` ✓ |
| [Element Plus el-message](https://element-plus.org/en-US/component/message.html) | `type`、`duration`、`showClose`、`icon`、`offset` | 类型 ✓；停留时间 ✓；关闭按钮 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 显示/关闭一条 | `Context.AddToast(ui.Toast) string` / `Context.CloseToast(id)` ✓ | ✅ |
| 自绘视口/单条 | `ui.ToastViewportBase(c, fn)` / `ui.ToastBase(c, t)` ✓ | ✅ |
| 停留时间/动作 | `Toast.Timeout` / `Toast.Action`+`OnAction` ✓ | ✅ |
| 淡入淡出 | `Element.Jump`+`AnimateTo` 给淡入 ✓；`ToastParts.Left()` 给剩余时间，自己算淡出 ✓ | ✅（旧文档说「淡入做不了」已过时） |
| 稳定句柄 | `Context` 是窗口的稳定句柄，`Notifier` 存它即可 ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `Notifier`（`NewNotifier(c)`）：`Show(ToastOptions) string` 包装 `c.AddToast`，`Close(id)` 包装 `c.CloseToast`，`Build()` 装自绘视口 ✓。
- app 把 `*Notifier` 存在自己状态里，**点击处理器只捕获 notifier，不捕获每帧的 `*ui.Context`**；`Build()` 在根视图调一次 ✓。
- `ToastOptions`：`Title`/`Description`/`Kind`/`Duration`/`ActionLabel`/`OnAction`/`OnClose` ✓。
- **类型→语义色映射**（`toastSemanticColor`）：success/info/warning 走根包 `SemanticColor(t, …)`，error 走 `t.Danger`，每帧从 `c.Theme()` 取 ✓。
- 淡入：新 toast 首帧 `Jump` 到 0，再 `AnimateTo` 到 1（走 `MotionNormal`、`NonMotion`，尊重 Reduce Motion）✓；剩余时间不足一 token 时淡出 ✓。
- `ToastViewport` 的位置可由 `ViewportPosition` 改 ✓；保留 `NewToast`/`NewToastViewport` 作为链式小封装 ✓。

**已知限制**
- 一次最多三条（框架限制 ✓）；堆叠不做位移动画 ✗。
- toast 内容只能是标题、描述、动作与关闭（框架部件就这些 ✓）。

## 4. 项目结构是否合理

- 位置：`overlays/` 子包——toast 是「窗口上的一层反馈」✓。
- 依赖：`mygo/ui` + `icons/` + `controls/` ✓。

## 5. API 契约

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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`toast.go`）
- [x] 测试（标题/描述 ✓、关闭按钮 ✓、动作 ✓、超时自走 ✓、`Notifier.Show` 不捕获 context ✓、`Notifier.Close` ✓）
- [ ] 堆叠位移动画（后置 ✗）
