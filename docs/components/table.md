# Table 规格书

表格：列头 + 行单元，带排序与选择状态（由 mygo `ListState` 提供）。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewTable(c *ui.Context, state *ui.ListState, columns []ui.TableColumn, n int, cell func(row, col int)) *Table` | 创建 n 行表格；`cell` 每帧填一个单元，排序/选择/列宽由 mygo 的 `ListState` 持有 |
| `Variant(v Variant) *Table` | 共享变体的外框与底色 |
| `Size(sz Size) *Table` | 字号档：`SizeXs/Sm/Lg`（默认 `TextSm` 档） |
| `Width(w float32) *Table` | 宽度（DIP）；0 时占满父宽 |
