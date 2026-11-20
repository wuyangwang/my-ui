# InputOTP 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

验证码输入：一格一字符，填入后半自动跳过、退格回退；填满时回调一次。

## API 契约

```go
type OTPOptions struct {
    Length       int                // 码的长度；<=0 时为 6
    Alphanumeric bool               // true 时字母也算有效字符
    Variant      myui.Variant       // 格子的外观，同 Input
    Size         myui.Size          // 格子尺寸档：24/32/40
    OnComplete   func(code string)  // 填满时调用一次
}

func NewOTP(c *ui.Context, value *string, o OTPOptions) *InputOTP
func NewInputOTP(c *ui.Context, value *string, digits int, key ...string) *InputOTP // 兼容旧接口
func (o *InputOTP) Width(w float32) *InputOTP
```

| API | 作用 |
|---|---|
| `NewOTP(c, value, o)` | 构造 OTP，编辑 `*value`；按 `o` 过滤、着色、定档，并在填满时回调 |
| `NewInputOTP(c, value, digits, key...)` | 旧构造器：`Length=digits`、缺省外观；`key` 给同父下的多个 OTP 定作用域 |
| `OTPOptions.Length` | 码长；`<=0` 取 6 |
| `OTPOptions.Alphanumeric` | 允许字母（`unicode.IsLetter`），否则只留数字（`unicode.IsDigit`） |
| `OTPOptions.Variant` | 格子外观，取 `FrameLookOf`；缺省/Native 画带边框的面 |
| `OTPOptions.Size` | 格子边长档：Xs/Sm → 24、Md → 32、Lg → 40（默认间距下） |
| `OTPOptions.OnComplete` | 码填满那一帧调用一次，传入完整码；额外帧不重复调用 |
| `Width(w)` | 把每个格子的宽度设为 `w` DIP；不设时按尺寸档（24/32/40） |

用法：

```go
forms.NewOTP(c, &code, forms.OTPOptions{
    Length: 6,
    OnComplete: func(code string) { verified = "Verified " + code },
}).Label("Code")

forms.NewOTP(c, &invite, forms.OTPOptions{
    Length: 4, Alphanumeric: true, Variant: myui.Secondary,
}).Label("Invite")
```
