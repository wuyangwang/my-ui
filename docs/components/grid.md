# Grid 规格书

> 位置：`layout/` 子包（`layout/grid.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `NewGrid(c)` | 带主题间距的网格（兼容 `ui.Grid` 签名） |
| `NewGridPanel(c, variant)` | 带表面的网格，`Variant` 决定外观 |
| `AutoGrid(c, state, n, minWidth, height, item)` | 响应式虚拟化网格 |
| `(*Grid).Element` | 可继续用 mygo 的 `Columns/ColumnTracks/Gap` 等方法 |

## 状态与限制

**限制**

- `AutoGrid` 的行高固定（`height` 参数）；不等高瀑布流需自行用 `ui.Grid` 排布。
- `MaxColumns` 未直接暴露：`GridView` 按可用宽度算列数，可在外层容器用 `MaxWidth` 间接限制。
