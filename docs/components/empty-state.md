# EmptyState 规格书

> 状态：**已实现**。空状态：列表或页面没有内容时的说明与引导。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Empty | 图标 + 标题 + 说明 + 动作，居中成组 |
| Ant Design Empty | 简要说明与下一步动作放在同一区域 |

## 2. 能否实现
- `ui.Column` 居中、`ui.Icon` 统一 `IconSizeLg`、`controls.NewButton` 动作槽 ✓
- 图标由调用方传 `*ui.SVG`（`icons/` 子包），库不嵌入图标集。

## 3. 实现到何种程度
- `EmptyStateOptions{Icon, Title, Description, Action}`；`NewEmptyState(c, o)` 返回 ui.Element。
- 图标在主题色圆形底上，标题 `TextLg` 粗体，说明 `TextSm` 弱墨，动作在下方。

## 4. 项目结构
- `display/` 子包。

## 5. API 契约
```go
type EmptyStateOptions struct {
    Icon *ui.SVG
    Title, Description string
    Action func()
}
func NewEmptyState(c *ui.Context, o EmptyStateOptions) ui.Element
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
