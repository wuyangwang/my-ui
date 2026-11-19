# Popover 规格书

> 位置：组件在 [`overlays/`](../../overlays) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`popover.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Popover](https://ui.shadcn.com/docs/components/popover)（`ui/popover.tsx`） | `PopoverTrigger`（`asChild`）· `PopoverContent`（`Portal` + `rounded-md border bg-popover shadow-md p-4`）· 面板外点击与 Escape 关闭 | 触发器=锚点 ✓；外观照抄（底色 + 描边 + 圆角 + `Space(4)` 内衬 + 阴影 ✓，最大宽 240 ✓） |
| mygo 的 `ui.PopoverBase(c, anchor, open, fn)` | 面板挂在锚点上（`AttachTo`）；点外面/Escape 关闭 ✓；非模态 ✓ | 锚点 + 自绘面板 ✓ |
| [Element Plus el-popover](https://element-plus.org/en-US/component/popover.html) | `trigger`（click/focus/hover）· `placement` · `teleported` | 触发方式固定为点击 ✓（框架如此 ✗） |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 面板挂在锚点上 | `ui.PopoverBase(c, anchor Element, open *bool, fn func(panel Element)) Element` ✓ | ✅ |
| 开合 | flag 是 app 的 ✓；点外面/Escape 关闭 ✓（框架 ✓） | ✅ |
| 点击锚点开合 | 框架不会代为切换 flag ✗ → 库先建锚点再读 `Clicked()` 切换 ✓ | ✅ |
| 样式/宽度 | 面板是元素 ✓；宽度自己定 ✓ | ✅ |
| 回调跨帧 | `ui.Local` 存回调与待投递的变化 ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `NewPopover(c, PopoverOptions)`：`Trigger` 由 app 构建，**库先建触发器再建浮层**，点击当帧开合 ✓。
- `Body func()` 由库包在 `panel.Children` 里 ✓；`Title`/`Description`/`Footer` 组成结构化面板 ✓。
- `Side`（Options 的 `*PopoverSide` 或 fluent 方法）：复用 `tooltipAnchors` 的映射做 `AttachTo`，默认 `PopoverBottom`，空间不足时框架翻到对侧 ✓。
- `OnOpenChange`/`OnDismiss` 跨帧保留，晚挂不丢事件 ✓；fluent 方法关闭期间也有效 ✓。

**已知限制**
- 触发方式只有点击（框架如此 ✗）。
- 只有四向位置（`PopoverTop/Bottom/Left/Right`），无 `sideOffset`/align 细调 ✗（框架不提供 ✗）。

**实现要点（踩过的坑）**
- 面板在覆盖层里建 → 内容必须建在面板里；库统一用 `panel.Children` 包住 `Body` ✓。

## 4. 项目结构是否合理

- 位置：`overlays/` 子包——弹层是控件之上的附属物 ✓。
- 依赖：只用 `mygo/ui` ✓；跨帧状态与 Modal 共用 `openState` ✓。

## 5. API 契约

```go
type PopoverOptions struct {
    Open         *bool             // app 持有
    Trigger      func() ui.Element // 建锚点，先于浮层
    Title        string            // 标题兼面板无障碍名
    Description  string            // 弱化说明
    Body         func()            // 库放进 panel.Children
    Footer       func()            // 右对齐操作行
    OnOpenChange func(bool)        // 开合回调，跨帧保留
    OnDismiss    func()            // 点外面/Escape 关闭
    Side         *PopoverSide      // nil=PopoverBottom
    ShowClose    *bool             // nil=显示 ✕
}
func NewPopover(c *ui.Context, o PopoverOptions) *Popover
type PopoverSide = TooltipSide
const (
    PopoverTop PopoverSide = iota
    PopoverBottom
    PopoverLeft
    PopoverRight
)
```

| 属性（方法/字段） | 类型 | 作用 |
|---|---|---|
| `Open` | `*bool` | app 的开关 |
| `Trigger` | `func() ui.Element` | 建锚点；点击切换开合，当帧生效 |
| `Title` / `Description` | `string` | 面板标题 / 灰色说明 |
| `Body` | `func()` | 面板正文，由库放进 `panel.Children` |
| `Footer` | `func()` | 右对齐操作行 |
| `OnOpenChange` | `func(bool)` | 任意来源的开合；晚挂不丢事件 |
| `OnDismiss` | `func()` | 点外面或 Escape 关闭 |
| `Side` | `*PopoverSide` | 面板出现在锚点哪一侧；nil 为下侧 |
| `ShowClose` | `*bool` | 关掉 ✕；nil 为开 |
| `Open() bool` | 方法 | 当前是否显示 |
| `SetOpen(bool) *Popover` | 方法 | 改 app 的 flag |
| `Dismissed() bool` | 方法 | 本帧是否被点外面/Escape 关闭 |
| `Side(PopoverSide) *Popover` | 方法 | 改锚点方向，关闭期间也有效 |
| `Title/Description(string) *Popover` | 方法 | 关闭期间也可设 |
| `Footer/ShowClose` | 方法 | 操作行与 ✕ |
| `OnOpenChange(func(bool)) *Popover` | 方法 | 注册开合回调；补发之前的改变 |
| `OnDismiss(func()) *Popover` | 方法 | 注册关闭回调 |

用法：

```go
side := overlaypkg.PopoverBottom
overlaypkg.NewPopover(c, overlaypkg.PopoverOptions{
    Open: &app.moreOpen,
    Trigger: func() ui.Element { return ctrlpkg.NewButton(c, "更多").Element },
    Body:    func() { ui.Text(c, "面板里的内容") },
    Side:    &side,
})
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`popover.go`）
- [x] 测试（单击锚点打开 ✓、正文落在面板 ✓、点外面关闭 ✓、Side 决定上/下 ✓、晚挂回调 ✓）
- [ ] 悬停/焦点触发、`sideOffset`/align 细调（后置 ✗）
