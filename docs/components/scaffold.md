# Scaffold 规格书

> 状态：**已实现**（`scaffold.go`）。注意与**骨架屏 Skeleton** 区分：Scaffold 是页面布局骨架（顶栏安全区/内容区/底栏），Skeleton 是加载占位（尚未实现，见 roadmap）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [shadcn/ui Sidebar / ScrollArea](https://ui.shadcn.com/docs/components/scroll-area) | 页面外壳：固定头尾 + 滚动内容区；滚动条内移避开浮条 |
| [Element Plus el-container](https://element-plus.org/zh-CN/component/container) | `header`/`main`/`footer`/`aside` 的分区布局；`el-scrollbar` 的内边距 |
| mygo 官方文档 | `c.TitleBar()`（隐藏标题栏窗口的控件安全区）、`ui.Scroll` + `ScrollbarInsets`（滚动条内移）、`Sticky().Top/Bottom`（吸顶吸底）、`ScrollState`/`TrackScroll`/`ScrollTo`（贴底跟随） |

## 2. 能否实现

- `c.TitleBar()` 返回 `TitleBar{Height, Left, Right}` ✓ → 顶部安全区与左右留白。
- `ui.Scroll` + `ScrollbarInsets(v ...float32)` ✓ → 滚动条内移避开底栏。
- `ui.Sticky().Top(0)/Bottom(0)` ✓ → 吸顶/吸底。
- 底部预留：滚动内容加底部 `Padding`（底栏高度 + 边距）✓，底栏高度由 app 给出（`BottomInset`）——立即模式下底栏是 app 的内容，库无法预先量出它的高度，故**高度由 app 传入**。
- 拖拽窗口：`ui.Row(...).DragWindow()` ✓（顶栏区域）。

## 3. 实现到何种程度

**已实现**：`ScaffoldOptions{UseTitleBar, Top, Bottom, BottomInset, Content}`；顶栏安全区（含左右留白）、顶栏内容、滚动内容区、底栏、底部预留与滚动条内移。

**后置**：贴底跟随与回底弹簧（与 `ai/ChatList` 联动）、吸顶吸底的可配置项、更多分区（侧栏折叠）、网格/分栏布局。

## 4. 项目结构是否合理

- 位置：根包 L4（组合组件），依赖 L3 之下的 ui 能力。
- 只做组合，不发明布局：分区由 `Top`/`Bottom`/`Content` 三个构建器交给 app，库负责安全区、滚动与底部预留。
- 命名：`Scaffold` = 布局骨架；**骨架屏另名 `Skeleton`**（roadmap 二期），两者在文档中互相注明。

## 5. API 契约

```go
type ScaffoldOptions struct {
	UseTitleBar bool       // 预留窗口控件区（c.TitleBar），可拖拽
	Top         func()     // 顶部栏内容
	Bottom      func()     // 底部栏内容
	BottomInset float32    // 滚动内容的底部预留（默认按底栏存在与否推导）
	Content     func()     // 滚动内容
}

func Scaffold(c *ui.Context, o ScaffoldOptions)
```

用法：

```go
myui.Scaffold(c, myui.ScaffoldOptions{
	UseTitleBar: true,
	Top:         func() { /* 顶栏 */ },
	Bottom:      func() { /* 底栏 */ },
	BottomInset: 56,
	Content:     func() { /* 内容 */ },
})
```

## 6. 状态

- [x] 调研与实现
- [x] 测试（`scaffold_test.go`：区域渲染、标题栏安全区偏移）
- [x] showcase「Scaffold」页
- [ ] 贴底跟随、吸顶吸底可配置、侧栏折叠（后置）
- [ ] 骨架屏 `Skeleton`（roadmap 二期，勿与本组件混淆）
