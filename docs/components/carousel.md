# Carousel 规格书

轮播：一次显示一张内容，左右翻页 + 圆点指示。

## API 契约

```go
func NewCarousel(c *ui.Context, index *int, count int, build func(i int)) *Carousel
func (c *Carousel) Arrows(on bool) *Carousel
func (c *Carousel) Dots(on bool) *Carousel
```
