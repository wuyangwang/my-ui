# Drawer 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`drawer.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Sheet](https://ui.shadcn.com/docs/components/sheet)（`ui/sheet.tsx`） | `SheetTrigger` · `SheetOverlay`（`bg-black/50`）· `SheetContent`（`fixed flex flex-col gap-4 bg-background shadow-lg`，右侧是 `w-3/4 sm:max-w-sm`）· `SheetClose`（✕）· 四条边 | 触发器 ✓；遮罩黑 50% ✓；面板取 3/4、最多一卡片宽 ✓；✕ ✓；四条边 ✓ |
| mygo 的 `ui.Overlay(c, fn)` | 在窗口之上建元素 ✓（`Absolute`/`Left`/`Top` 定位 ✓）；其上元素拿走焦点 ✓ | 抽屉用它搭 ✓ |
| mygo 的 `Element.Modal()` | 做成模态背景：焦点、Tab、读屏都限在罩内 ✓ | 遮罩用它 ✓ |
| mygo 的 `Element.OverlayShortcut` | 覆盖层里的快捷键 ✓ | Escape 收起 ✓ |
| [Element Plus el-drawer](https://element-plus.org/en-US/component/drawer.html) | `direction`、`size`、`append-to-body`、`modal`、`show-close` | 边 ✓；大小由 `Size` 给 ✓；关闭前钩子 ✗（框架没有 ✗） |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 浮在窗口上的一层 | `ui.Overlay(c, fn)` ✓ | ✅ |
| 铺满窗口的遮罩 + 边缘面板 | `Absolute`/`Left`/`Top`/`Right`/`Bottom`/`Grow`/`MaxWidth` ✓ | ✅ |
| 模态 | `Element.Modal()` ✓ | ✅ |
| 点外面/Escape 关闭 | 遮罩可点 + `OverlayShortcut(0, KeyEscape)` ✓ | ✅ |
| 回调跨帧 | `ui.Local` 存回调与待投递变化 ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `NewDrawer(c, DrawerOptions)`：`Trigger` 由库先于覆盖层构建，单击当帧开 ✓；遮罩黑 50% 吃掉外面点击 ✓；面板在指定边 ✓（右/左/下/上）✓；`Size` 自定义占比与上限 ✓。
- `Body func()` 由库包在 `panel.Children` 里 ✓；`Title`/`Description`/`Footer` 结构化 ✓。
- **关闭期间不再提前返回**：`Title`/`Footer`/`Side`/`Size` 等 fluent 方法写进 Local，下次打开即生效（修掉旧版 silent no-op）✓。
- `OnOpenChange`/`OnDismiss` 跨帧保留，晚挂不丢事件 ✓。
- 颜色与尺寸每帧从 `c.Theme()` 取 ✓。

**已知限制**
- 没有滑入滑出的姿态动画 ✗（见 [motion.md](motion.md)）。
- 遮罩点击与 Escape 都算 `Dismissed`，但 ✕ 不算（与 mygo 对话框一致）✓。

**实现要点（踩过的坑）**
- 旧版在 `!open` 时直接 `return`，导致随后调用的 fluent 方法作用在 `Parts==nil` 上而失效；新版总是先建状态和触发器，只有覆盖层按开合构建 ✓。

## 4. 项目结构是否合理

- 位置：`overlays/` 子包——抽屉是「盖在窗口边上的一块面板」✓。
- 依赖：`mygo/ui` + `icons/` + `controls/`；与其余浮层共用 `openState` ✓。

## 5. API 契约

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
| `ShowClose` | `*bool` | 关掉 ✕；nil 为开 |
| `Grow, MaxMain` | `float32` | 主轴占比与 spacing 上限 |
| `Open() bool` | 方法 | 当前是否显示 |
| `SetOpen(bool) *Drawer` | 方法 | 改 app 的 flag |
| `Dismissed() bool` | 方法 | 本帧是否被遮罩/Escape 关闭 |
| `Side(DrawerSide) *Drawer` | 方法 | 改边；关闭期间也有效 |
| `Size(grow, max float32) *Drawer` | 方法 | 改占比与上限 |
| `Title/Description(string) *Drawer` | 方法 | 关闭期间也可设 |
| `Footer/ShowClose` | 方法 | 操作行与 ✕ |
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

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`drawer.go`）
- [x] 测试（触发器打开 ✓、面板在边 ✓、✕ 关闭 ✓、点外面关闭 ✓、fluent 关闭期间有效 ✓、开合上报 ✓）
- [ ] 滑入滑出姿态动画（后置 ✗）
