# Carousel 规格书

> 状态：**已实现**。轮播：一次显示一张内容，左右翻页 + 圆点指示。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Carousel | 左右箭头按钮 + 底部圆点 |
| Element Plus Carousel | 索引指针由外部持有 |

## 2. 能否实现
- 左右箭头用 `ctrlpkg.NewIconAction`（图标本身即控件）✓；圆点用带无障碍名的小按钮 ✓。

## 3. 实现到何种程度
- `NewCarousel(c, index *int, count int, build func(i int)) *Carousel`；`.Arrows(on bool)`、`.Dots(on bool)`；默认两者都开。
- 索引越界自动夹到 [0,count-1]；翻页写回指针。

## 4. 项目结构
- `display/` 子包。

## 5. API 契约
```go
func NewCarousel(c *ui.Context, index *int, count int, build func(i int)) *Carousel
func (c *Carousel) Arrows(on bool) *Carousel
func (c *Carousel) Dots(on bool) *Carousel
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
