# AutoComplete 规格书

> 状态：**已实现**。自动补全：输入框 + 随输入过滤的建议列表。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `Autocomplete` | 输入框与建议列表的完整行为（过滤、选择、键盘） |
| shadcn/ui Command | 弹层面板的边框、圆角与弱墨提示 |

## 2. 能否实现
- `ui.Autocomplete(c, value, suggestions)` ✓；库只加共享 Variant/Size/Width 外观。

## 3. 实现到何种程度
- `NewAutoComplete(c, value *string, suggestions []string)`；`.Variant`、`.Size`、`.Width(w)`（默认占满）。

## 4. 项目结构
- `forms/` 子包。

## 5. API 契约
```go
func NewAutoComplete(c *ui.Context, value *string, suggestions []string) *AutoComplete
func (a *AutoComplete) Variant(v Variant) *AutoComplete
func (a *AutoComplete) Size(sz Size) *AutoComplete
func (a *AutoComplete) Width(w float32) *AutoComplete
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
