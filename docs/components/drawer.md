# Drawer 规格书

> 状态：**已实现**（`drawer.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Sheet](https://ui.shadcn.com/docs/components/sheet)（源码 `ui/sheet.tsx`） | `SheetTrigger` · `SheetOverlay`（`bg-black/50`）· `SheetContent`（`fixed z-50 flex flex-col gap-4 bg-background shadow-lg`，右侧是 `inset-y-0 right-0 h-full w-3/4 border-l sm:max-w-sm`）· `SheetClose`（右上角 ✕，`absolute top-4 right-4`）· `SheetHeader`/`SheetFooter`/`SheetTitle`/`SheetDescription` · 四条边 | 触发器 ✓；遮罩黑 50% ✓；面板取窗口的四分之三、最多一张卡片宽 ✓；右上角关闭按钮 ✓；四条边 ✓（右/左/下） |
| mygo 的 `ui.Overlay(c, fn)` | 在窗口之上建元素 ✓（`Absolute`/`Left`/`Top` 定位 ✓，或 `AttachTo` ✓）；它之上的元素拿走焦点，关闭时还回去 ✓ | 抽屉用它搭 ✓ |
| mygo 的 `Element.Modal()` | 做成对话框的背景罩：焦点进到里面、Tab 在里面转、外面的快捷键等、读屏只看得到它 ✓ | 遮罩用它 ✓ |
| mygo 的 `Element.OverlayShortcut(mods, key)` | 覆盖层里的快捷键 ✓ | Escape 收起 ✓ |
| [Element Plus el-drawer](https://element-plus.org/en-US/component/drawer.html) | `direction`（rtl/ltr）、`size`、`append-to-body`、`modal`、`show-close`、`before-close` | 边 ✓；大小固定 ✓（shadcn 的比例 ✓）；关闭前的钩子 ✗（框架没有 ✗） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 浮在窗口上的一层 | `ui.Overlay(c, fn)` ✓ | ✅ |
| 铺满窗口的遮罩 + 边缘的面板 | `Absolute`/`Left`/`Top`/`Right`/`Bottom`/`Width`/`MaxWidth`/`Grow` ✓ | ✅ |
| 模态（焦点在里面） | `Element.Modal()` ✓ | ✅ |
| 点外面关闭 | 遮罩可点 ✓ | ✅ |
| Escape 关闭 | `Element.OverlayShortcut(0, ui.KeyEscape)` ✓ | ✅ |
| 抽屉自己的部件 | 没有 ✗ | ❌ 自己搭 ✓ |

## 3. 实现到何种程度

**已实现**
- `myui.NewDrawer(c, &open, myui.DrawerRight, trigger, build)`：触发器由 app 构建 ✓，点击即打开（同一帧 ✓）；遮罩（黑 40% ✓）吃掉外面的点击 ✓；面板在窗口的边 ✓，取四分之三的room、最多 360 DIP 宽 ✓（shadcn 的 `w-3/4 max-w-sm` ✓），其余留给遮罩 ✓（窗口后面露出来 ✓）；面板：surface 底 + 描边 + 圆角 + 阴影 + `Space(4)` 内衬 ✓；右上角 ✕ 关闭按钮 ✓；Escape 关闭 ✓。
- 边：`DrawerRight`/`DrawerLeft`/`DrawerBottom` ✓。
- 面板是模态的：焦点在里面、Tab 在里面转 ✓（框架的 `Modal()` ✓）。

**已知限制**
- 抽屉没有滑入滑出的动画 ✗（框架的姿态动画不可用 ✗，见 [motion.md](motion.md)）。
- 没有上边 ✗（shadcn 有 ✗；要做的话再加 ✓）。
- 遮罩 40% 而不是 50% ✗（与对话框一致 ✓，可改 ✓）。
- 尺寸是固定的比例与上限 ✗（不做 `size` 属性 ✗）。

## 4. 项目结构是否合理

- 位置：根包 L3——抽屉是「盖在窗口边上的一块面板」✓。
- 依赖：只用 `mygo/ui` + `icons/` ✓。

## 5. API 契约

```go
// DrawerSide is the edge of the window a drawer comes from.
type DrawerSide uint8

const (
	DrawerRight DrawerSide = iota
	DrawerLeft
	DrawerBottom
)

// NewDrawer shows a panel over the edge of the window while open is true.
func NewDrawer(c *ui.Context, open *bool, side DrawerSide, trigger func() ui.Element, build func(panel ui.Element)) *Drawer

func (d *Drawer) Open() bool
func (d *Drawer) SetOpen(open bool) *Drawer
```

用法：

```go
myui.NewDrawer(c, &app.open, myui.DrawerRight, func() ui.Element {
	return myui.NewButton(c, "设置").Element
}, func(panel ui.Element) {
	ui.Text(c, "设置").Bold()
	myui.NewButton(c, "面板里的按钮").Outline()
})
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`drawer.go`）
- [x] 测试（触发器打开 ✓、面板在窗口的边 ✓、✕ 关闭 ✓、点外面关闭 ✓）
- [ ] 滑入滑出动画、上边（后置 ✗）
