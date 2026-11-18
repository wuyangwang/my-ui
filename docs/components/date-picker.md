# DatePicker 规格书

> 状态：**已实现**。日期选择：输入框 + 日历弹层，日期指针由 app 持有。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `DateInput` | 日期编辑与日历弹层（底层是 `Calendar`） |
| shadcn/ui Calendar | 面板的边框、圆角与状态色 |

## 2. 能否实现
- `ui.DateInput(c, date *time.Time)` ✓；库加共享 Variant/Size/Width 外观。

## 3. 实现到何种程度
- `NewDatePicker(c, date *time.Time)`；`.Variant`、`.Size`、`.Width(w)`（默认占满）；弹层行为交给 mygo。

## 4. 项目结构
- `forms/` 子包。

## 5. API 契约
```go
func NewDatePicker(c *ui.Context, date *time.Time) *DatePicker
func (d *DatePicker) Variant(v Variant) *DatePicker
func (d *DatePicker) Size(sz Size) *DatePicker
func (d *DatePicker) Width(w float32) *DatePicker
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
