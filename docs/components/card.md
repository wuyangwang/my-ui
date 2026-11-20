# Card 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
func NewCard(c *ui.Context, title string) *Card
func (card *Card) Description(s string) *Card
func (card *Card) Content(build func()) *Card
```

## 状态与限制

**限制**

- 没有头尾分区（`CardHeader`/`CardFooter`）（app 自己排）。
- 没有阴影（shadcn 有）。
