# Tabs 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`tabs.go`）。按 Lobe MyGo 审查修复：Vertical/Point/Size、独立指示器 + `ElementTransition.Match` 跨项滑移、hover 改文字色。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| Lobe MyGo `nav/tabs.go` | `TabsVariant`（Rounded/Square/Point）、`Size`、`Vertical`；指示器是独立元素 + `ElementTransition{Match: tabsMatch{selected}}` 跨项滑移；hover/选中改 `TextColor` | 全部照做（用 mygo 等价 API）✓ |
| [shadcn/ui Tabs](https://ui.shadcn.com/docs/components/tabs) | 药丸托盘 + 选中项自己的面；下划线 | Pills / Line ✓ |
| mygo 的 `ui.TabsBase(c, selected, n)` | 无外观基座：键盘、选中状态、`Changed()` ✓ | 基座 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 指示器滑移 | `ui.ElementTransition{Duration, Ease, Match}` + `Box.Transition` ✓ | ✅ |
| 竖排 | `Element.Column()` / `Element.Row()` ✓ | ✅ |
| hover 文字色 | `Element.Hovered()` + `TextColor` + `Transition(StateTransition())` ✓ | ✅ |
| 无障碍 | `ui.TabsBase` ✓ | ✅ |

## 3. 实现到何种程度

- 选中指示器是建在当前选中项里的**独立元素**，带 `Transition(ui.ElementTransition{Duration: MotionNormal, Ease: ui.EaseOut, Match: tabsMatch{selected}})`。`Match` 挂在 app 的 `*int` 上，跨帧稳定，框架据此把新位置的指示器与上一帧的配对并滑过去。**不用 `AnimatePose`**（本库实测会让元素停在不可见状态）。
- `Line()` 下划线（竖排时在右侧）、`Point()` 圆点、`Vertical()` 竖排并把面板放到侧边、`Size(sz)` 改高度/内衬/字号。
- hover / 选中改文字色，`StateTransition()` 让颜色渐变；禁用项文字弱化。
- 只建选中项的内容；`OnChanged(fn)` 在 `list.Changed()` 为真时回调。

## 4. 项目结构

- `display/` 子包；只用公开 `ui` API。

## 5. API 契约（逐 prop）

### 类型

| 符号 | 作用 |
|---|---|
| `type Tab struct` | `Label`、`Icon *ui.SVG`、`Disabled bool`、`Content func()` |

### 构造与方法

| 方法 | 作用 |
|---|---|
| `NewTabs(c, selected *int, tabs ...Tab) *Tabs` | 创建；只建选中项的内容 |
| `Line() *Tabs` | 下划线样式 |
| `Point() *Tabs` | 圆点样式 |
| `Vertical() *Tabs` | 竖排标签，面板在侧边 |
| `Size(s Size) *Tabs` | 高度/内衬/字号档 |
| `OnChanged(fn func(index int)) *Tabs` | 选中变化时回调 |

用法：

```go
display.NewTabs(c, &selected,
    display.Tab{Label: "概览", Content: func() { ui.Text(c, "今天") }},
    display.Tab{Label: "设置", Content: func() { ui.Text(c, "参数") }},
).Vertical().Point().Size(myui.SizeSm)
```

## 6. 限制

- 键盘导航完全交给 `ui.TabsBase`（方向键/Home/End、禁用项跳过）。
- 指示器滑移依赖框架的 `Match` 链接：同一帧只有一个指示器元素，跨帧身份靠 app 的 `*int` 指针；不要为不同 Tabs 共用同一个指针。

## 7. 状态

- [x] 调研与实现（`tabs.go`）
- [x] 测试：内容切换、竖排 + Point + Size、OnChanged
