# Drawer 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
type DrawerOptions struct {
    Open         *bool             // app 持有
    Side         DrawerSide        // 默认 DrawerRight
    Trigger      func() ui.Element // 建触发器，先于覆盖层
    Title        string
    Description  string
    Body         func()            // 库放进 panel.Children
    Footer       func()
    OnOpenChange func(bool)
    OnDismiss    func()
    ShowClose    *bool
    Grow, MaxMain float32          // 默认 3 与 90（w-3/4 max-w-sm）
}
func NewDrawer(c *ui.Context, o DrawerOptions) *Drawer
type DrawerSide uint8
const (
    DrawerRight DrawerSide = iota
    DrawerLeft
    DrawerBottom
    DrawerTop
)
```

| 属性（方法/字段） | 类型 | 作用 |
|---|---|---|
| `Open` | `*bool` | app 的开关 |
| `Side` | `DrawerSide` | 面板从哪条边来；零值为右 |
| `Trigger` | `func() ui.Element` | 建触发器；单击当帧打开 |
| `Title` / `Description` | `string` | 标题（兼无障碍名）/ 灰色说明 |
| `Body` | `func()` | 正文，由库放进 `panel.Children` |
| `Footer` | `func()` | 右对齐操作行 |
| `OnOpenChange` | `func(bool)` | 开合回调，跨帧保留、晚挂补发 |
| `OnDismiss` | `func()` | 遮罩点击/Escape 关闭 |
| `ShowClose` | `*bool` | 关掉；nil 为开 |
| `Grow, MaxMain` | `float32` | 主轴占比与 spacing 上限 |
| `Open() bool` | 方法 | 当前是否显示 |
| `SetOpen(bool) *Drawer` | 方法 | 改 app 的 flag |
| `Dismissed() bool` | 方法 | 本帧是否被遮罩/Escape 关闭 |
| `Side(DrawerSide) *Drawer` | 方法 | 改边；关闭期间也有效 |
| `Size(grow, max float32) *Drawer` | 方法 | 改占比与上限 |
| `Title/Description(string) *Drawer` | 方法 | 关闭期间也可设 |
| `Footer/ShowClose` | 方法 | 操作行与 |
| `OnOpenChange/OnDismiss` | 方法 | 注册回调；补发之前的改变 |

用法：

```go
overlaypkg.NewDrawer(c, overlaypkg.DrawerOptions{
    Open:    &app.settingsOpen,
    Side:    overlaypkg.DrawerRight,
    Trigger: func() ui.Element { return ctrlpkg.NewButton(c, "设置").Element },
    Title:   "设置",
    Body:    func() { ui.Text(c, "在这里调整偏好") },
})
```

## 状态与限制

**限制**

- 没有滑入滑出的姿态动画（见 [motion.md](motion.md)）。
- 遮罩点击与 Escape 都算 `Dismissed`，但 不算（与 mygo 对话框一致）。
