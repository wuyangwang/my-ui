# AlertDialog 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
type DialogAction struct {
    Label   string   // 按钮文案
    Kind    Variant  // 按钮预设；零值=旧默认（末尾 Destructive）
    OnClick func()   // 按钮回调，在主线程、点击当帧
}
type AlertDialogOptions struct {
    Open      *bool          // app 持有
    Title     string         // 标题兼面板无障碍名
    Message   string         // 弱化说明兼无障碍描述
    Actions   []DialogAction // 类型化动作
    OnDismiss func()         // Escape/遮罩关闭；跨帧保留
}
func NewAlertDialog(c *ui.Context, o AlertDialogOptions) *AlertDialog
```

| 属性（方法/字段） | 类型 | 作用 |
|---|---|---|
| `Open` | `*bool` | app 的开关 |
| `Title` | `string` | 加粗标题，兼面板无障碍名 |
| `Message` | `string` | 弱化说明，兼面板无障碍描述 |
| `Actions` | `[]DialogAction` | 类型化按钮；`Label` 文案、`Kind` 预设、`OnClick` 回调 |
| `OnDismiss` | `func()` | Escape/遮罩关闭；晚挂补发 |
| `Open() bool` | 方法 | 当前是否显示 |
| `Dismissed() bool` | 方法 | 本帧是否被关闭 |
| `OnDismiss(func()) *AlertDialog` | 方法 | 注册关闭回调；补发之前的关闭 |

用法：

```go
overlaypkg.NewAlertDialog(c, overlaypkg.AlertDialogOptions{
    Open:    &app.confirmOpen,
    Title:   "删除项目？",
    Message: "此操作不能撤销。",
    Actions: []overlaypkg.DialogAction{
        {Label: "取消", OnClick: func() { app.result = "cancelled" }},
        {Label: "删除", Kind: myui.Destructive, OnClick: app.delete},
    },
    OnDismiss: func() { app.result = "dismissed" },
})
```

## 状态与限制

**限制**

- `ui.DialogBase` 点遮罩也会关闭，而 mygo 的 `ui.AlertDialog` 忽略遮罩点击；这是公开底座的行为。
- 遮罩点击按 `Dismissed` 处理并回调 `OnDismiss`。
