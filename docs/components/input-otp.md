# InputOTP 规格书

> 状态：**已实现**。验证码输入：一格一字符，自动跳到下一格，退格回上一格。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui InputOTP | 等宽格子、输入后焦点自动前进、退格回退 |
| mygo/ui `TextInput` | 原生文本编辑与键盘事件 |

## 2. 能否实现
- 每格一个 `ui.TextInput`，绑到 app 的字符串切片；用 `OnTextKey`/`Changed` 处理输入与退格（mygo 无公开的 focus API，焦点跳转需在下一帧用 `Focus()` 请求）。
- 若焦点跳转在 mygo 上做不到，写成“每格独立、可 Tab 切换、自动去空格”的降级形态，并把限制写进文档。

## 3. 实现到何种程度
- `NewInputOTP(c, value *string, digits int) *InputOTP`；`.Width(w)`；值以 `*string` 绑定。
- 格子统一用字段圆角与边框，输入后写回字符串。

## 4. 项目结构
- `forms/` 子包。

## 5. API 契约
```go
func NewInputOTP(c *ui.Context, value *string, digits int) *InputOTP
func (o *InputOTP) Width(w float32) *InputOTP
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
