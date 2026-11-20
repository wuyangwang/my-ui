# Tabs 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

提供 Vertical/Point/Size、独立指示器 + `ElementTransition.Match` 跨项滑移、hover 改文字色。

## API 契约（逐 prop）

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

## 限制

- 键盘导航完全交给 `ui.TabsBase`（方向键/Home/End、禁用项跳过）。
- 指示器滑移依赖框架的 `Match` 链接：同一帧只有一个指示器元素，跨帧身份靠 app 的 `*int` 指针；不要为不同 Tabs 共用同一个指针。
