# DatePicker 规格书

日期选择：输入框 + 日历弹层，日期指针由 app 持有。

## API 契约

```go
func NewDatePicker(c *ui.Context, date *time.Time) *DatePicker
func (d *DatePicker) Variant(v Variant) *DatePicker
func (d *DatePicker) Size(sz Size) *DatePicker
func (d *DatePicker) Width(w float32) *DatePicker
```
