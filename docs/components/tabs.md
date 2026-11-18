# Tabs 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`tabs.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Tabs](https://ui.shadcn.com/docs/components/tabs) | `TabsList`（`bg-muted rounded-lg p-1`）+ `TabsTrigger`（选中的那个是 `bg-background shadow`）+ `TabsContent` | 一排药丸 ✓；选中的那个在自己的面上 ✓；内容只有开着的那块建 ✓ |
| mygo 的 `ui.TabsBase(c, selected *int, n int) TabsParts` | 无外观的基座：`List`、`Tab(i)`、键盘在组内移动、`Changed()` ✓ | 基座 ✓，外观我们给 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 标签组 | `ui.TabsBase` ✓ | ✅ |
| 选中项 | app 的 `*int` ✓ | ✅ |
| 内容 | app 自己建 ✓（只建选中的那块 ✓） | ✅ |

## 3. 实现到何种程度

**已实现**
- `displaypkg.NewTabs(c, &selected, displaypkg.Tab{Label: "概览", Content: func(){…}}, …)`：标签行（弱底色 + 圆角 + 内衬 ✓），选中的那个在自己的面上（带一点阴影 ✓）；只有开着的那块内容被建 ✓。
- `OnChanged(func(int))` ✓；`Tab.Icon` 放图标 ✓、`Tab.Disabled` 禁用 ✓（框架键盘导航自动跳过禁用项 ✓）；`Line()` 切到 shadcn 的下划线样式 ✓。

**已知限制**
- 标签在侧边的竖排 ✗（`TabsBase.List` 是框架写死的行，无公开方向切换 ✗）。

## 4. 项目结构是否合理

- 位置：根包 L3——标签页是导航件 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
type Tab struct {
	Label   string
	Content func()
}

func NewTabs(c *ui.Context, selected *int, tabs ...Tab) *Tabs
func (tb *Tabs) OnChanged(fn func(index int)) *Tabs
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`tabs.go`）
- [x] 测试（选中的内容 ✓、切换 ✓）
