# SideNav 规格书

> 位置：`layout/` 子包（`layout/side_nav.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| Lobe `layout.SideNav` | 58 DIP 图标轨道、顶部头像、上下操作分组 | 全部对应 ✓ |
| mygo | `ui.Column` + 样式方法 | 用基础元素搭 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 固定宽度轨道 | `Width` / `MinHeight` | ✅ |
| 上下贴边 | `Grow(1)` 撑开中段 | ✅ |
| 主题表面与细线 | `c.Theme()` | ✅ |

## 3. 实现到何种程度

- 宽度默认 `Space(15)`（58 DIPs），最小高 `Space(160)`（640 DIPs）。
- `Avatar`、`TopActions`、`BottomActions` 三个槽位；中段 `Grow(1)` 把底部操作推到末端。
- `Width` 可链式覆盖。

**已知限制**
- 操作槽位只做排布，图标按钮由 app 用 `controls.NewButton`/`NewIconAction` 构建。
- 不含展开态；需要可拖可收的导航请用 `DraggableSideNav`。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `SideNavProps.Avatar` | 顶部头像槽位 |
| `SideNavProps.TopActions` | 主操作槽位 |
| `SideNavProps.BottomActions` | 尾部操作槽位 |
| `NewSideNav(c, props)` | 构造 |
| `(*SideNav).Width(v)` | 设置轨道宽度 |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（三槽位绘制）
