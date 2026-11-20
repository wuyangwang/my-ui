# SideNav 规格书

> 位置：`layout/` 子包（`layout/side_nav.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `SideNavProps.Avatar` | 顶部头像槽位 |
| `SideNavProps.TopActions` | 主操作槽位 |
| `SideNavProps.BottomActions` | 尾部操作槽位 |
| `NewSideNav(c, props)` | 构造 |
| `(*SideNav).Width(v)` | 设置轨道宽度 |

## 状态与限制

**限制**

- 操作槽位只做排布，图标按钮由 app 用 `controls.NewButton`/`NewIconAction` 构建。
- 不含展开态；需要可拖可收的导航请用 `DraggableSideNav`。
