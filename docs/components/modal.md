# Modal 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`modal.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Dialog](https://ui.shadcn.com/docs/components/dialog)（`ui/dialog.tsx`） | `DialogTrigger`（`asChild`）· `DialogOverlay`（`bg-black/50`）· `DialogContent`（`rounded-lg border bg-background p-6 shadow-lg max-w-lg` + 右上角 `DialogClose` 的 ✕）· `DialogTitle`/`DialogDescription` · `DialogFooter` | 触发器 ✓；遮罩黑 50% ✓；面板照抄（底色 + 描边 + 圆角 + `Space(6)` 内衬 + 阴影 + 最大宽 512 ✓）；右上角关闭按钮 ✓；标题/描述/页脚由 Options 给 ✓ |
| mygo 的 `ui.DialogBase(c, open, fn func(backdrop, panel))` | 模态：背景罩吃掉外面的按下 ✓；面板居中 ✓；面板有 `RoleDialog` ✓；Escape 关闭 ✓；焦点收进面板 ✓ | 背景罩与面板由库排版 ✓ |
| [Element Plus el-dialog](https://element-plus.org/en-US/component/dialog.html) | `modal`、`draggable`、`fullscreen`、`close-on-click-modal` | 模态 ✓；拖拽/全屏 ✗（框架没有 ✗） |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 背景罩 + 面板 | `ui.DialogBase(c, open *bool, fn func(backdrop Element, panel Element)) Element` ✓ | ✅ |
| 开关 | flag 是 app 的 ✓；Escape 关闭 ✓（框架 ✓） | ✅ |
| 触发器打开 | 框架**不会**代为切换 flag ✗ → 库先建触发器再读 `Clicked()` 代为打开 ✓ | ✅ |
| 焦点 | 框架把焦点收进面板 ✓ | ✅ |
| 无障碍名 | 面板是 `RoleDialog` ✓，`Title` 同时作 `Label` ✓ | ✅ |
| 回调跨帧 | `ui.Local` 存回调与待投递的变化 ✓ | ✅ |
| 拖拽/全屏/非模态 | 框架没有 ✗ | ❌ 后置 ✗ |

## 3. 实现到何种程度

**已实现**
- `NewModal(c, ModalOptions)`：触发器由 `Trigger` 构建，**库先建触发器再建覆盖层**，点击当帧即开 ✓。
- `Body func()` 由库包在 `panel.Children` 里，**调用方不会把内容落到窗口左上角** ✓；`Title`/`Description`/`Footer` 组成结构化面板 ✓。
- `Open`/`SetOpen` ✓；`OnOpenChange` 与 `OnDismiss` 存在 `ui.Local` 里，**晚挂也能收到之前发生的变化** ✓。
- `Title`/`Description`/`Footer`/`ShowClose` 等 fluent 方法写进 Local，**关闭期间调用也有效**，下次打开即呈现 ✓。
- 颜色与尺寸每帧从 `c.Theme()` 取 ✓；状态淡入淡出走 `StateTransition`/`OverlayEnter` ✓。

**已知限制**
- 非模态、拖拽、全屏 ✗（框架没有 ✗）。
- 焦点陷阱由框架兜底，行为细节不可控 ✗。

**实现要点（踩过的坑）**
- 触发器必须在 `ui.DialogBase` **之前**建：点击输入当帧被读走，flag 在面板建之前置位，对话框当帧就开 ✓。
- 面板在覆盖层里建 → 内容必须建在面板里；库统一用 `panel.Children` 包住 `Body` ✓。

## 4. 项目结构是否合理

- 位置：`overlays/` 子包——对话框是「覆盖在窗口上的卡片」✓。
- 依赖：只用 `mygo/ui` + `icons/` + `controls/` ✓。
- `openState`（`open_state.go`）被 Modal/Popover/Drawer/HoverCard 共用，集中处理跨帧回调 ✓。

## 5. API 契约

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
    ShowClose    *bool                // nil=显示右上角 ✕
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
| `ShowClose` | `*bool` | 关掉右上角 ✕；nil 为开 |
| `Open() bool` | 方法 | 当前是否显示 |
| `SetOpen(bool) *Modal` | 方法 | 改 app 的 flag |
| `Dismissed() bool` | 方法 | 本帧是否被 Escape/背景罩关闭 |
| `Title/Description(string) *Modal` | 方法 | 关闭期间也可设，下次打开生效 |
| `Footer(func()) *Modal` | 方法 | 关闭期间也可设 |
| `ShowClose(bool) *Modal` | 方法 | 显示/隐藏 ✕ |
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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`modal.go`）
- [x] 测试（触发器单击打开 ✓、正文落在面板 ✓、晚挂 `OnOpenChange` 触发 ✓、✕ 关闭上报 ✓）
- [ ] 非模态/拖拽/全屏（后置 ✗）
