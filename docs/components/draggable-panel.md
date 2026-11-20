# DraggablePanel 规格书

> 位置：`layout/` 子包（`layout/draggable_panel.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `PanelPlacement` | `PanelRight`（默认）/ `PanelLeft` |
| `DraggablePanelProps.Placement` | 面板所在边 |
| `DraggablePanelProps.MinSize / MaxSize` | 宽度边界；零 Min 为主题 `Space(20)` |
| `DraggablePanelProps.CollapseThreshold` | 拖到此宽度以下折叠；零禁用 |
| `NewDraggablePanel(c, size, expanded, props, build)` | 构造；`build` 画展开时的内容 |
| `(*DraggablePanel).Expanded()` | 是否展开 |
| `(*DraggablePanel).Size()` | 当前宽度（DIPs） |

## 状态与限制

**限制**

- 宽度瞬时跟随指针，没有用 `AnimateTo` 做弹性回弹：折叠时的尺寸跳变由主题色过渡平滑。
- 折叠握柄内只有一个展开图标，不显示被折叠面板的标题。
