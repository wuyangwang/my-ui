# Popover 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`popover.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Popover](https://ui.shadcn.com/docs/components/popover)（源码 `ui/popover.tsx`） | `PopoverTrigger`（`asChild`）· `PopoverContent`（`Portal` + `bg-popover text-popover-foreground shadow-md rounded-md border p-4 w-72`）· 面板外点击与 Escape 关闭 ✓ | 触发器=锚点 ✓；外观照抄（surface 底 + 描边 + 圆角 + `Space(4)` 内衬 + 阴影 ✓，最大宽 `t.Space(60)` ✓） |
| mygo 的 `ui.PopoverBase(c, anchor, open, fn)` | 面板**挂在锚点上**（`AttachTo(anchor, AnchorBottomStart, AnchorTopStart)`：锚点下方 ✓）；按面板外/Escape 关闭 ✓；非模态 ✓；面板吃掉自己的按下 ✓ | 锚点 + 我们自绘的面板 ✓ |
| [Element Plus el-popover](https://element-plus.org/en-US/component/popover.html) | `virtual-triggering`、`trigger`（click/focus/hover/contextmenu）、`placement`、`popper-style`、`teleported` | 触发方式固定为点击 ✓（框架如此 ✗） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 面板挂在锚点上 | `ui.PopoverBase(c, anchor Element, open *bool, fn func(panel Element)) Element` ✓ | ✅ |
| 开合 | 开关是 app 的 flag ✓；面板外点击/Escape 关闭 ✓（框架处理 ✓） | ✅ |
| 点击锚点打开 | 框架**不会**代为切换 flag ✗ → 由库读锚点的 `Clicked()` 代为打开 ✓ | ✅ |
| 样式 | 面板是一个元素 ✓ | ✅ |
| 面板宽度 | 框架不定 ✗ | ✅（我们自己定） |

## 3. 实现到何种程度

**已实现**
- `NewPopover(c, open, anchor, build)`：锚点由 app 构建（`func() ui.Element` ✓），**点击锚点即打开** ✓（同一帧 ✓）；面板外点击与 Escape 关闭 ✓（框架 ✓，`overlays_test.go` 覆盖 ✓）。
- 面板样式由库给（surface 底 + 描边 + 圆角 + 内衬 + 阴影 ✓），内容由 app 填 ✓。
- `Open()`/`SetOpen(bool)` ✓。

**已知限制**
- 触发方式只有点击（框架如此 ✗；悬停/焦点触发需 app 自己接 ✗）。
- 面板位置由框架定（锚点下方 ✓），不设 `placement` ✗。
- 面板没有自带的头尾/标题（框架的弹层内容由 app 填 ✓，标题要 app 自己写 ✗）。

**实现要点（踩过的坑）**
- 面板在覆盖层里建 → 面板的内容必须建在面板里（`panel.Children(func(){ build(panel) })`），
  否则内容落到窗口左上角 ✗。

## 4. 项目结构是否合理

- 位置：根包 L3——弹层是控件之上的附属物 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
// NewPopover shows a panel over its anchor while open is true.
func NewPopover(c *ui.Context, open *bool, anchor func() ui.Element, build func(panel ui.Element)) *Popover

func (p *Popover) Open() bool
func (p *Popover) SetOpen(open bool) *Popover
```

用法：

```go
overlaypkg.NewPopover(c, &app.open, func() ui.Element {
	return ctrlpkg.NewButton(c, "更多").Element
}, func(panel ui.Element) {
	formspkg.NewField(c, "备注", func() ui.Element { ... })
})
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`popover.go`）
- [x] 测试（点击锚点后面板出现 ✓、内容在面板里 ✓、点击外面关闭 ✓）
- [ ] 悬停/焦点触发、`placement`（后置 ✗）
