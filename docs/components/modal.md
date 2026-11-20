# Modal 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
type ModalOptions struct {
    Open         *bool                // app 持有；必填
    Trigger      func() ui.Element    // 建触发器，先于覆盖层
    Title        string               // 标题兼面板无障碍名
    Description  string               // 标题下的弱化说明
    Body         func()               // 库放进 panel.Children
    Footer       func()               // 右对齐操作行
    OnOpenChange func(open bool)      // 开合回调，跨帧保留
    OnDismiss    func()               // Escape/背景罩关闭
    ShowClose    *bool                // nil=显示右上角
}
func NewModal(c *ui.Context, o ModalOptions) *Modal
```

| 属性（方法/字段） | 类型 | 作用 |
|---|---|---|
| `Open` | `*bool` | app 的开关；为真时显示 |
| `Trigger` | `func() ui.Element` | 建触发器；点击即开，当帧生效 |
| `Title` | `string` | 加粗标题，兼面板无障碍名 |
| `Description` | `string` | 标题下的灰色说明 |
| `Body` | `func()` | 面板正文；库包在 `panel.Children` 里 |
| `Footer` | `func()` | 右对齐的操作行 |
| `OnOpenChange` | `func(bool)` | 任意来源的开合都回调；晚挂不丢事件 |
| `OnDismiss` | `func()` | Escape 或点背景罩时回调 |
| `ShowClose` | `*bool` | 关掉右上角；nil 为开 |
| `Open() bool` | 方法 | 当前是否显示 |
| `SetOpen(bool) *Modal` | 方法 | 改 app 的 flag |
| `Dismissed() bool` | 方法 | 本帧是否被 Escape/背景罩关闭 |
| `Title/Description(string) *Modal` | 方法 | 关闭期间也可设，下次打开生效 |
| `Footer(func()) *Modal` | 方法 | 关闭期间也可设 |
| `ShowClose(bool) *Modal` | 方法 | 显示/隐藏 |
| `OnOpenChange(func(bool)) *Modal` | 方法 | 注册回调；晚挂立刻补发之前的改变 |
| `OnDismiss(func()) *Modal` | 方法 | 注册关闭回调 |

用法：

```go
overlaypkg.NewModal(c, overlaypkg.ModalOptions{
    Open:        &app.deleteOpen,
    Trigger:     func() ui.Element { return ctrlpkg.NewButton(c, "删除").Destructive().Element },
    Title:       "删除项目",
    Description: "此操作不能撤销。",
    Body:        func() { ui.Text(c, "确定删除这个项目吗？") },
    Footer: func() {
        if ctrlpkg.NewButton(c, "取消").Ghost().Clicked() { app.deleteOpen = false }
    },
    OnOpenChange: func(open bool) { app.log(open) },
})
```

## 状态与限制

**限制**

- 非模态、拖拽、全屏（框架没有）。
- 焦点陷阱由框架兜底，行为细节不可控。
