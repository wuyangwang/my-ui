# FloatingSheet 规格书

> 位置：`layout/` 子包（`layout/floating_sheet.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `FloatingSheetProps.Title` | 拖拽标题 |
| `FloatingSheetProps.SnapPoints` | 吸附高度，≤1 为容器比例 |
| `FloatingSheetProps.MinHeight / RestingHeight / MaxHeight` | 高度边界与初值 |
| `FloatingSheetProps.CloseThreshold` | 低于最低吸附点的多少倍时关闭 |
| `FloatingSheetProps.HeaderActions` | 标题右侧控件；nil 时画关闭图标 |
| `FloatingSheetProps.Inline / Embedded` | 占位/描边外观 |
| `FloatingSheetProps.NoDismiss` | 禁止下拉关闭 |
| `NewFloatingSheet(c, open, props, build)` | 构造；`build` 在面板自己的 `Children` 中调用 |
| `(*FloatingSheet).Open()` | 本帧是否显示 |

## 状态与限制

**限制**

- 吸附比例以 `c.Size()` 的窗口高度作容器高度。浮层放在更小的容器里时，比例是近似值。
- 快速甩动的惯性吸附未做：mygo 的 `DragEnded` 给出速度，但本组件只按当前高度取最近点。
