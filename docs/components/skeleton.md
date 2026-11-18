# Skeleton 规格书

> 状态：**已实现**。加载占位：在数据尚未到达时先画出内容的大致形状，避免布局跳动。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Skeleton | 用 `Surface` 色块替代文字与图片，圆角随主题 |
| Ant Design Skeleton | `Lines` 行数、头像圆点、按钮块 |

## 2. 能否实现
- `ui.Row`/`ui.Box` + `Background(t.Surface)` + `Radius` ✓
- 动效走 `Element.Transition`/`Animate`，Reduce Motion 由框架兜底；v1 只做静态占位，不做脉冲。

## 3. 实现到何种程度
- `NewSkeleton(c, lines int)`：默认 3 行；`.Circle()` 头像圆；`.Width(w)` 固定宽度（默认占满）；`.Size(SizeXs/Sm/Md/Lg)` 控制条高（10/12/16/20 DIP）。
- 后置：脉冲动画、自定义条形布局。

## 4. 项目结构
- `display/` 子包，展示件分组。

## 5. API 契约
```go
func NewSkeleton(c *ui.Context, lines int) *Skeleton
func (s *Skeleton) Circle() *Skeleton
func (s *Skeleton) Width(w float32) *Skeleton
func (s *Skeleton) Size(sz Size) *Skeleton
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
