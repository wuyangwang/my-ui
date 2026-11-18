# Result 规格书

> 状态：**已实现**。结果页：一次操作（提交、支付、删除）完成后的状态页。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| Ant Design Result | success / warning / error / info 四种状态，图标 + 标题 + 描述 + 额外动作 |
| shadcn/ui Alert | 语义色与图标风格 |

## 2. 能否实现
- `ui.Icon` + `SemanticColor` + `Danger` ✓；图标尺寸用根包 `IconSizeLg`。
- 状态文案由 app 传入，库只提供图标与语义色，不内置固定文案。

## 3. 实现到何种程度
- `ResultStatus`（Success/Warning/Info/Error）；`ResultOptions{Status, Title, Description, Extra}`；`NewResult(c, o)`。
- 图标在主题色圆形底上，标题 `Text2Xl` 粗体，描述弱墨，`Extra` 动作槽。

## 4. 项目结构
- `display/` 子包。

## 5. API 契约
```go
type ResultStatus uint8
const (ResultSuccess ResultStatus = iota; ResultWarning; ResultInfo; ResultError)
type ResultOptions struct {
    Status ResultStatus
    Title, Description string
    Extra func()
}
func NewResult(c *ui.Context, o ResultOptions) ui.Element
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
