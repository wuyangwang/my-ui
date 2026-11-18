# InputNumber 规格书

> 状态：**已实现**。数字输入：带步进上下按钮、范围与步长的输入控件。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| Element Plus InputNumber | 输入框 + 上下步进按钮，`min/max/step` |
| Ant Design InputNumber | 同样结构，禁用态与尺寸档 |

## 2. 能否实现
- mygo `ui.NumberInput(c, value, lo, hi, step)` 提供原生编辑与步进 ✓；库只需套上共享 `Variant/Size/Width` 观感。

## 3. 实现到何种程度
- `NewInputNumber(c, value *float64, lo, hi, step float64)`；`.Variant`、`.Size`、`.Width(w)`（默认占满父宽）。
- 用 `FrameLookOf` 派生边框/底色，不引入新颜色。

## 4. 项目结构
- `forms/` 子包，复用字段外观。

## 5. API 契约
```go
func NewInputNumber(c *ui.Context, value *float64, lo, hi, step float64) *InputNumber
func (n *InputNumber) Variant(v Variant) *InputNumber
func (n *InputNumber) Size(sz Size) *InputNumber
func (n *InputNumber) Width(w float32) *InputNumber
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
