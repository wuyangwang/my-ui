# FloatingSheet 规格书

> 位置：`layout/` 子包（`layout/floating_sheet.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| shadcn/ui Sheet/Card | 底部浮层、拖头改高、吸附点、下拉关闭、Inline/Embedded | 全部对应 ✓ |
| mygo | `Element.Dragged` / `DragEnded`、`ui.Overlay` 思路 | 拖拽与抬起事件用公开 API ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 拖头改高 | `Element.Dragged` | ✅ |
| 抬起吸附 | `Element.DragEnded` | ✅ |
| 下拉关闭 | `DragEnded` + 阈值 | ✅ |
| 每帧颜色 | `c.Theme()` | ✅ |
| 取容器高度做吸附比例 | 无公开的容器尺寸事件 | ⚠️ 用窗口高度近似 |

## 3. 实现到何种程度

- 按住头部拖动改高度，`MinHeight`/`MaxHeight` 夹取；松开吸附到最近的 `SnapPoints`。
- 拖到低于 `LowestSnap × CloseThreshold` 时写回 `open = false`。
- `Inline` 让浮层在流中占位，`Embedded` 用描边代替阴影；默认底部绝对定位。
- 高度记忆在 `ui.Local`；状态与颜色每帧从主题取。卡片面用 shadcn zinc token：`t.Surface`、1px `t.Border`、`Radius(t, RadiusXl)` 与轻阴影；`Embedded` 用描边代替阴影。

**已知限制**
- 吸附比例以 `c.Size()` 的窗口高度作容器高度。浮层放在更小的容器里时，比例是近似值。
- 快速甩动的惯性吸附未做：mygo 的 `DragEnded` 给出速度，但本组件只按当前高度取最近点。

## 4. API 契约

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

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（打开/关闭/内容）
