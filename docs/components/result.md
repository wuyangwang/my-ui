# Result 规格书

结果页：一次操作（提交、支付、删除）完成后的状态页。

## API 契约

```go
type ResultStatus uint8
const (ResultSuccess ResultStatus = iota; ResultWarning; ResultInfo; ResultError)
type ResultOptions struct {
    Status ResultStatus
    Title, Description string
    Action func()
}
func NewResult(c *ui.Context, o ResultOptions) ui.Element
```
