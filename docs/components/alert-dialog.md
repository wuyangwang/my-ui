# AlertDialog 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`alert_dialog.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| mygo/ui `ui.AlertDialog(c, open, title, message, buttons...)` | 契约：`*open` 为真时显示；点按钮或 Escape 都把 `*open` 置回 false | 沿用开合语义 ✓ |
| shadcn/ui `AlertDialog` | 面板是一张卡片：描边、圆角、软阴影；标题加粗、说明弱化；按钮右对齐，确认按钮带强调色 | 照抄 ✓ |
| my-ui `overlays.NewModal` | 浮层内容建在 `ui.DialogBase` 的 panel 里；跨帧回调存 `ui.Local` | 复用 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 自绘警告框 | `ui.DialogBase` 的 panel 是元素 ✓；`controls.NewButton` 预设 ✓ | ✅ |
| 类型化按钮回调 | `DialogAction.OnClick func()` ✓ | ✅ |
| 关闭回调跨帧 | `ui.Local` 存回调与待投递的关闭 ✓ | ✅ |
| 无障碍 | `Element.Role(RoleAlertDialog)` + `Label(title)` + `Description(message)` ✓ | ✅ |

不能直接给 mygo 的 `ui.AlertDialog` 换外观：它返回 `int`，按钮由它在浮层里自画，公开 API 无可改样式的句柄。所以用 `ui.DialogBase` 复刻同一契约。

## 3. 实现到何种程度

**已实现**
- `NewAlertDialog(c, AlertDialogOptions) *AlertDialog`：面板用主题底色、描边、XL 圆角与软阴影；标题 `TextLg` 加粗、说明 `TextSm` 弱化 ✓。
- **类型化动作** `[]DialogAction`：每个动作自带 `Label`、`Kind`（按钮预设）与 `OnClick`；**app 不再做索引→标签映射**，删除按钮的回调直接写在动作里 ✓。
- `Kind` 为零值时保留旧默认：非末尾按钮 `Outline`，末尾按钮 `Destructive` 且自动聚焦 ✓。
- `OnDismiss` 存 `ui.Local`，晚挂补发之前的关闭；Escape 触发它 ✓。
- 颜色与尺寸每帧从 `c.Theme()` 取；淡入走 `OverlayEnter`，状态走 `StateTransition` ✓。

**已知限制**
- `ui.DialogBase` 点遮罩也会关闭，而 mygo 的 `ui.AlertDialog` 忽略遮罩点击；这是公开底座的行为 ✓。
- 遮罩点击按 `Dismissed` 处理并回调 `OnDismiss` ✓。

## 4. 项目结构是否合理

- `overlays/` 子包；`controls.NewButton` 提供预设，根包提供 `Radius`/`TextSize` 与语义色 ✓。

## 5. API 契约

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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`alert_dialog.go`）
- [x] 测试（动作 `OnClick` 写回结果 ✓、Escape 调用 `OnDismiss` ✓、晚挂 `OnDismiss` ✓、主题 ✓）
