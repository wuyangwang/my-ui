# DraggableSideNav 规格书

> 位置：`layout/` 子包（`layout/draggable_side_nav.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| Lobe `layout.DraggableSideNav` | 边缘改宽、收成轨道、`SideNavState`（Expanded/Collapse/RailWidth） | 全部对应 ✓ |
| mygo | `Element.Dragged` / `DragEnded` | 拖拽用公开 API ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 拖边改宽 | `Element.Dragged` | ✅ |
| 收起/展开 | app 的 `*bool` | ✅ |
| 把折叠进度交给内容 | `SideNavState` 回调 | ✅ |
| 点击边缘标签不误触拖拽 | 标签与边缘是各自的元素 | ✅ |

## 3. 实现到何种程度

- `width *float32` 与 `expanded *bool` 由 app 持有；nil 时用 `ui.Local`。
- 边缘标签在悬停或收起时显现，点击切换；拖到距最小宽三分之一以内收起，并记录可恢复的宽度。
- `Header`、`Footer` 与正文都收到 `SideNavState`：`Expanded` 决定画什么，`Collapse` 从 0 到 1，`RailWidth` 给列表做居中图标的轨道宽。

**已知限制**
- 收起动画只靠主题色过渡与宽度变化；`Collapse` 在展开/收起间取端点值，不随拖拽连续插值。
- 只有左侧一种放置（与 Lobe 一致）。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `SideNavState.Expanded` | 是否展开 |
| `SideNavState.Collapse` | 折叠进度 0..1 |
| `SideNavState.RailWidth` | 轨道宽度（DIPs） |
| `DraggableSideNavProps.MinWidth / MaxWidth` | 宽度边界 |
| `DraggableSideNavProps.Header / Footer` | 接收 `SideNavState` 的槽位 |
| `NewDraggableSideNav(c, width, expanded, props, body)` | 构造；`body` 收 `SideNavState` |
| `(*DraggableSideNav).Expanded() / Width()` | 读取当前状态 |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（正文绘制、状态回传）
