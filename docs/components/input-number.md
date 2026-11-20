# InputNumber 规格书

数字输入：带步进上下按钮、范围与步长的输入控件。

## API 契约

```go
func NewInputNumber(c *ui.Context, value *float64, lo, hi, step float64) *InputNumber
func (n *InputNumber) Variant(v Variant) *InputNumber
func (n *InputNumber) Size(sz Size) *InputNumber
func (n *InputNumber) Width(w float32) *InputNumber
```
