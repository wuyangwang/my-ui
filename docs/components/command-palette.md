# Command Palette 规格书

> 状态：**已实现**。命令面板：一个居中对话框，输入即过滤命令，回车/点击执行。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Command Dialog | 居中面板 + 搜索框 + 结果列表 |
| mygo/ui `DialogBase` | 背景罩、居中面板、Escape/点背景关闭 |

## 2. 能否实现
- `ui.DialogBase` + `ui.TextInput` + 过滤后的按钮列表 ✓；查询字符串 `*string` 由 app 持有。

## 3. 实现到何种程度
- `Command{Label, Icon, Shortcut, OnSelect}`；`NewCommandPalette(c, open, query, commands)`。
- 过滤按 label 大小写不敏感；点击命令执行并关闭面板；空结果显示 i18n 提示。
- 后置：上下键选择、分组命令。

## 4. 项目结构
- `overlays/` 子包。

## 5. API 契约
```go
type Command struct {
    Label string
    Icon *ui.SVG
    Shortcut string
    OnSelect func()
}
func NewCommandPalette(c *ui.Context, open *bool, query *string, commands []Command) ui.Element
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
