# Tag 规格书

> 状态：**已实现**。标签：比 Badge 更轻的元信息标记，可带图标与关闭。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| Element Plus Tag | 纯色弱底 + 文字 + 可关闭图标 |
| shadcn/ui Badge | 共享 Variant/Size 词表 |

## 2. 能否实现
- 用 `FrameLookOf` 取语义色（或回退 Surface），关闭图标用 `NewIconAction` ✓。

## 3. 实现到何种程度
- `TagOptions{Label, Icon, Closable, OnClose}`；`NewTag(c, o)`；`.Variant`、`.Size`。
- `Closed *bool` 指针由 app 持有；关闭按钮调用后写回。

## 4. 项目结构
- `display/` 子包。

## 5. API 契约
```go
type TagOptions struct {
    Label string
    Icon *ui.SVG
    Closed *bool
    OnClose func()
}
func NewTag(c *ui.Context, o TagOptions) *Tag
func (t *Tag) Variant(v Variant) *Tag
func (t *Tag) Size(sz Size) *Tag
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
