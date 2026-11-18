# Modal 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`modal.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Dialog](https://ui.shadcn.com/docs/components/dialog)（源码 `ui/dialog.tsx`） | `DialogTrigger`（`asChild`）· `DialogOverlay`（`bg-black/50`）· `DialogContent`（`rounded-lg border bg-background p-6 shadow-lg max-w-lg` + 右上角 `DialogClose` 的 ✕）· `DialogTitle`/`DialogDescription`（无障碍名）· `DialogFooter` | 触发器 ✓；遮罩黑 50% ✓；面板照抄（`Background` 底 + 描边 + 圆角 + `Space(6)` 内衬 + 阴影 + 最大宽 `t.Space(128)`=512 ✓）；右上角关闭按钮 ✓；标题/描述由 app 写，面板用 `panel.Label(...)` 命名 ✓ |
| mygo 的 `ui.DialogBase(c, open, fn func(backdrop, panel))` | 模态：背景罩吃掉外面的按下 ✓；面板居中 ✓；面板有 `Role(RoleDialog)` ✓；Escape 关闭 ✓；焦点收进面板 ✓ | 背景罩与面板由库排版 ✓ |
| [Element Plus el-dialog](https://element-plus.org/en-US/component/dialog.html) | `modal`（模态/非模态）、`draggable`、`fullscreen`、`append-to-body`、`close-on-click-modal`、`destroy-on-close` | 模态 ✓；拖拽/全屏 ✗（框架没有 ✗） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 背景罩 + 面板 | `ui.DialogBase(c, open *bool, fn func(backdrop Element, panel Element)) Element` ✓ | ✅ |
| 开关 | flag 是 app 的 ✓；Escape 关闭 ✓（框架 ✓） | ✅ |
| 触发器打开 | 框架**不会**代为切换 flag ✗ → 由库读触发器的 `Clicked()` 代为打开 ✓ | ✅ |
| 焦点 | 框架把焦点收进面板 ✓ | ✅ |
| 无障碍名 | 面板是 `RoleDialog` ✓，`panel.Label(...)` 命名 ✓ | ✅ |
| 拖拽/全屏/非模态 | 框架没有 ✗ | ❌ 后置 ✗ |

## 3. 实现到何种程度

**已实现**
- `NewModal(c, open, trigger, build)`：**触发器由 app 构建**（`func() ui.Element` ✓），点击即打开（同一帧 ✓）；背景罩（黑 50% ✓）与面板（`Background` 底 + 描边 + 圆角 + `Space(6)` 内衬 + 阴影 ✓，最大宽 512 ✓）；右上角一个 ✕ 关闭按钮 ✓（`icons.X()`，无障碍名「关闭」✓）；Escape 与点击背景罩关闭 ✓（框架 ✓）。
- 内容由 app 填 ✓；面板即无障碍的对话框，用 `panel.Label("删除项目")` 命名 ✓。
- `Open()`/`SetOpen(bool)` ✓。

**已知限制**
- 非模态、拖拽、全屏 ✗（框架没有 ✗）。
- 面板内容由 app 填，没有自带的头尾/标题栏 ✗（要 app 自己写 ✗）。
- 焦点陷阱由框架兜底，行为细节不可控 ✗。

**实现要点（踩过的坑）**
- 触发器必须在 `ui.DialogBase` **之前**建：点击的输入在当帧就被读走，flag 在面板建之前置位，
  对话框当帧就开 ✓（反之要等下一帧 ✗）。
- 面板在覆盖层里建 → 内容必须建在面板里（`panel.Children(func(){ build(panel) })`），
  否则内容落到窗口左上角 ✗。

## 4. 项目结构是否合理

- 位置：根包 L3——对话框是「覆盖在窗口上的卡片」✓。
- 依赖：只用 `mygo/ui` + `icons/` ✓。

## 5. API 契约

```go
// NewModal shows a dialog over the window while open is true.
func NewModal(c *ui.Context, open *bool, trigger func() ui.Element, build func(panel ui.Element)) *Modal

func (m *Modal) Open() bool
func (m *Modal) SetOpen(open bool) *Modal
```

用法：

```go
overlaypkg.NewModal(c, &app.open, func() ui.Element {
	return ctrlpkg.NewButton(c, "删除").Destructive().Element
}, func(panel ui.Element) {
	panel.Label("删除项目")
	ui.Text(c, "确定删除这个项目吗？")
	ctrlpkg.NewButton(c, "取消").Ghost()
	ctrlpkg.NewButton(c, "删除").Destructive()
})
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`modal.go`）
- [x] 测试（触发器打开 ✓、内容在面板里 ✓、✕ 关闭 ✓、Escape 关闭 ✓）
- [ ] 非模态/拖拽/全屏（后置 ✗）
