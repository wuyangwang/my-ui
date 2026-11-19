# Grid 规格书

> 位置：`layout/` 子包（`layout/grid.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| shadcn/ui 栅格版式 | 等宽列、按宽度自动增减列、设置上限 | 自动增减 ✓（用 `AutoGrid`） |
| mygo | `ui.Grid`、`ui.GridView`、`Track.Min` | 公开 API 足够 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 等宽列网格 | `ui.Grid` | ✅ |
| 响应式自动列数 | `ui.GridView`，按 `minWidth` 计算列数 | ✅ |
| 虚拟化大列表 | `ui.GridView` | ✅ |

## 3. 实现到何种程度

- `NewGrid(c)` 保留 mygo `ui.Grid(c)` 签名，套用主题间距 `Space(2)`；`NewGridPanel(c, variant)` 提供带表面的结构化版本及 `Variant` 预设。
- `AutoGrid(c, state, n, minWidth, height, item)` 包裹虚拟化的 `GridView`，按宽度放入至少 `minWidth` 的等宽列，滚动条用 `ShadcnScrollbar`。
- `NewGridPanel` 的卡片面用 shadcn zinc token：`t.Surface`、`t.Border`（1px）、`Radius(t, RadiusLg)`；`Outline`/`Secondary` 只换面，不改排布。

**已知限制**
- `AutoGrid` 的行高固定（`height` 参数）；不等高瀑布流需自行用 `ui.Grid` 排布。
- `MaxColumns` 未直接暴露：`GridView` 按可用宽度算列数，可在外层容器用 `MaxWidth` 间接限制。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `NewGrid(c)` | 带主题间距的网格（兼容 `ui.Grid` 签名） |
| `NewGridPanel(c, variant)` | 带表面的网格，`Variant` 决定外观 |
| `AutoGrid(c, state, n, minWidth, height, item)` | 响应式虚拟化网格 |
| `(*Grid).Element` | 可继续用 mygo 的 `Columns/ColumnTracks/Gap` 等方法 |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（子元素、自动网格）
