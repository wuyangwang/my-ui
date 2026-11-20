# DraggableSideNav 规格书

> 位置：`layout/` 子包（`layout/draggable_side_nav.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `SideNavState.Expanded` | 是否展开 |
| `SideNavState.Collapse` | 折叠进度 0..1 |
| `SideNavState.RailWidth` | 轨道宽度（DIPs） |
| `DraggableSideNavProps.MinWidth / MaxWidth` | 宽度边界 |
| `DraggableSideNavProps.Header / Footer` | 接收 `SideNavState` 的槽位 |
| `NewDraggableSideNav(c, width, expanded, props, body)` | 构造；`body` 收 `SideNavState` |
| `(*DraggableSideNav).Expanded() / Width()` | 读取当前状态 |

## 状态与限制

**限制**

- 收起动画只靠主题色过渡与宽度变化；`Collapse` 在展开/收起间取端点值，不随拖拽连续插值。
- 只有左侧一种放置。
- 轨道面用 shadcn token：`t.Surface` + 1px `t.Border`，边缘标签用 `Radius(t, RadiusMd)`，悬停边缘变 `t.Accent`；颜色每帧从 `c.Theme()` 读取。
