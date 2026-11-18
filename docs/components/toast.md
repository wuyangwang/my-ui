# Toast 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`toast.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Toast](https://ui.shadcn.com/docs/components/toast)（源码 `ui/toast.tsx`，底层 sonner） | `ToastProvider`（视口：`position="bottom-right"`、`duration`、`closeButton`）· `Toast`（卡片）· `ToastTitle`/`ToastDescription`/`ToastAction`/`ToastClose` · 类型（success/error/warning/info） | 视口在**右下角** ✓（框架的视口是底部居中，我们自己摆 ✓）；卡片 + 标题/描述/动作/关闭 ✓；类型给图标与颜色 ✓ |
| mygo 的 `ui.ToastViewportBase(c, fn)` / `ui.ToastBase(c, t)` / `c.AddToast` / `c.CloseToast` | 视口由 app 建 ✓（我们的视口**取代**主题自带的观感 ✓）；toast 的字段（`Title`/`Description`/`Type`/`Timeout`/`Action`/`OnAction`/`OnClose`/`Data`）✓；最多三条 ✓；指针停在上面时时间停住 ✓；框架自带淡入淡出 ✗（自定义视口后没有 ✗） | 字段一一对应 ✓；我们的视口自己画 ✓；淡出自己补 ✓ |
| [Element Plus el-message](https://element-plus.org/en-US/component/message.html) | `type`（info/success/warning/error）、`duration`、`showClose`、`icon`、`customClass`、`offset` | 类型给图标与颜色 ✓；停留时间 ✓；关闭按钮 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 显示一条通知 | `Context.AddToast(ui.Toast) string` ✓ | ✅ |
| 关掉一条 | `Context.CloseToast(id)` ✓ | ✅ |
| 自己画视口 | `ui.ToastViewportBase(c, fn func(viewport Element, toasts []Toast)) Element` ✓ | ✅ |
| 自己画一条 | `ui.ToastBase(c, t) ToastParts`（`Root`/`ActionButton()`/`CloseButton()`/`Left()`）✓ | ✅ |
| 停留时间 | `Toast.Timeout` ✓；超时自动关 ✓ | ✅ |
| 动作按钮 | `Toast.Action` + `OnAction` ✓（框架代为关掉 ✓） | ✅ |
| 淡入淡出 | 框架的淡入淡出只在自己的观感里 ✗；`ToastParts.Left()` 给出还剩多久 ✓（淡出可做 ✓，淡入做不了 ✗） | ⚠️ 半 |

## 3. 实现到何种程度

**已实现**
- `overlaypkg.NewToastViewport(c)`：窗口的 toast 视口，**每帧在视图里调一次** ✓；它取代主题自带的观感 ✓。toast 停在**右下角** ✓（shadcn 的位置 ✓），卡片：surface 底 + 描边 + 圆角 + 内衬 + 阴影 ✓。
- `overlaypkg.NewToast(c, "已保存")` → `.Description(...)`/`.Success()`/`.Info()`/`.Warning()`/`.Destructive()`/`.Timeout(...)`/`.Action("撤销", fn)`/`.ID(...)`/`.OnClose(...)` → `.Show()` 返回 id ✓。
- 每种类型一个图标与颜色（success/info/warning/error）✓；标题加粗、描述弱一些 ✓；动作按钮 ✓；关闭按钮（无障碍名「关闭」）✓；时间将尽时淡出 ✓。

**已知限制**
- 淡入没有 ✗（框架的淡入只在自己的观感里 ✗；`ToastParts` 只给出「还剩多久」，没有「已经过了多久」 ✗）。
- 视口固定在右下角 ✗（框架的视口是一个铺满窗口的列，位置由 app 摆 ✓，但我们的摆法写死 ✓）。
- 一次最多三条 ✓（框架的限制 ✓）；堆叠不做动画位移 ✗。
- toast 的内容只能是标题、描述、动作与关闭 ✗（框架的部件就这些 ✓）。

## 4. 项目结构是否合理

- 位置：根包 L3——toast 是「窗口上的一层反馈」✓。
- 依赖：只用 `mygo/ui` + `icons/` ✓。

## 5. API 契约

```go
// NewToastViewport builds the window's toasts with the library's look.
func NewToastViewport(c *ui.Context) *ToastViewport

// NewToast creates a toast with a title, to show with Show.
func NewToast(c *ui.Context, title string) *Toast

func (t *Toast) Description(s string) *Toast
func (t *Toast) Type(v Variant) *Toast
func (t *Toast) Success() *Toast
func (t *Toast) Info() *Toast
func (t *Toast) Warning() *Toast
func (t *Toast) Destructive() *Toast
func (t *Toast) Timeout(d time.Duration) *Toast
func (t *Toast) Action(label string, fn func()) *Toast
func (t *Toast) ID(s string) *Toast
func (t *Toast) OnClose(fn func()) *Toast
func (t *Toast) Show() string
```

用法：

```go
// 视图里，每帧一次：
overlaypkg.NewToastViewport(c)

// 动作里：
overlaypkg.NewToast(c, "已保存").Description("改动已经写入。").Success().Show()
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`toast.go`）
- [x] 测试（标题与描述 ✓、关闭按钮 ✓、动作 ✓、超时自走 ✓、视口取代主题观感 ✓）
- [ ] 淡入、堆叠动画（后置 ✗）
