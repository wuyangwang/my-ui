# SortableList 规格书

可排序列表：列表项带拖拽手柄，拖动即改变顺序。

## API 契约

```go
func NewSortableList(c *ui.Context, state *ui.ListState, n int, item func(i int)) *SortableList
```
