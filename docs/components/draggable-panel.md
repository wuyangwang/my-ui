# DraggablePanel 规格书

> 位置：`layout/` 子包（`layout/draggable_panel.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| Lobe `layout.DraggablePanel` | 接缝缩放、折叠按钮、拖到过小折叠、左右放置 | 全部对应 ✓ |
| mygo | `Element.Dragged` / `DragEnded` | 拖拽用公开 API ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 拖接缝改宽 | `Element.Dragged` | ✅ |
| 折叠/展开 | app 的 `*bool` | ✅ |
| 按钮点击不被吞 | 按钮与接缝是各自的元素 | ✅ |
| 每帧主题色 | `c.Theme()` + `StateTransition` | ✅ |

## 3. 实现到何种程度

- `size *float32` 与 `expanded *bool` 都由 app 持有；nil 时用面板根的 `ui.Local`。
- 接缝悬停/按下变主题强调色，拖动时改宽并按 `MinSize/MaxSize` 夹取。
- `CollapseThreshold` 让拖到过小的面板折叠；折叠后留一条窄握柄，其上的图标按钮可展开并恢复原尺寸。
- `Placement` 支持右侧（默认）与左侧；两侧的内容与接缝顺序镜像。

**已知限制**
- 宽度瞬时跟随指针，没有用 `AnimateTo` 做弹性回弹：折叠时的尺寸跳变由主题色过渡平滑。
- 折叠握柄内只有一个展开图标，不显示被折叠面板的标题。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `PanelPlacement` | `PanelRight`（默认）/ `PanelLeft` |
| `DraggablePanelProps.Placement` | 面板所在边 |
| `DraggablePanelProps.MinSize / MaxSize` | 宽度边界；零 Min 为主题 `Space(20)` |
| `DraggablePanelProps.CollapseThreshold` | 拖到此宽度以下折叠；零禁用 |
| `NewDraggablePanel(c, size, expanded, props, build)` | 构造；`build` 画展开时的内容 |
| `(*DraggablePanel).Expanded()` | 是否展开 |
| `(*DraggablePanel).Size()` | 当前宽度（DIPs） |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（展开内容、折叠后隐藏）
