# InputOTP 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`input_otp.go`）。验证码输入：一格一字符，填入后半自动跳过、退格回退；填满时回调一次。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [shadcn/ui InputOTP](https://ui.shadcn.com/docs/components/input-otp) | 等宽格子、输入后焦点自动前进、按 pattern 过滤字符（仅数字 / 字母数字） |
| [Lobe MyGo `field.InputOTP`](https://mygo.lobehub.com/skills/components/input-otp) | `InputOTPProps{Length, Alphanumeric, Variant, Size, OnComplete}`；默认 6 格；`OnComplete` 填满时触发一次；24/32/40 三档格子 |
| mygo/ui `TextInputBase` | 一个透明的文本输入接管键盘，格子里只画字符（粘贴整段也能落位） |

## 2. 能否实现（公开 ui API）

| 需要 | 公开 API | 结论 |
|---|---|---|
| 一格一字符的格子 | `ui.Row` + `ui.Text`；格子尺寸用 `t.Space(6/8/10)` ✓ | ✅ |
| 键盘编辑、粘贴整段、焦点 | `ui.TextInputBase(c, value).MaxLength(n)`，透明铺在格子上；`Focus()`/`Focused()` ✓ | ✅ |
| 字符过滤 | `unicode.IsDigit` / `unicode.IsLetter` 在帧首清洗 `*value` ✓ | ✅ |
| 填满回调一次 | `ui.Local` 记录上次完整码，只有新码不同才调用 ✓ | ✅ |
| 外观随主题 | `FrameLookOf(Variant)`、`FrameFontSize`、`StateTransition()` 都在帧首取 `c.Theme()` ✓ | ✅ |

**与 Lobe 的取舍**：Lobe 的格子是固定 24/32/40 DIP；我们按主题的 `Space` 取同款档位（默认间距 4 → 24/32/40），跟随桌面文本与间距设置。mygo 没有公开的逐格焦点 API，改成「一个透明输入 + 可见格子」，焦点态由 `Focused()` 派生的高亮边框表达。

## 3. 实现到何种程度

**已实现**
- 选项是结构体 `OTPOptions`：长度（缺省 6）、是否字母数字、变体、尺寸、完成回调。
- 帧首按 `unicode.IsDigit`（`Alphanumeric` 时再加 `unicode.IsLetter`）清洗 `*value`，最多留 `Length` 个；非法字符（含粘贴进来的）直接丢掉。
- `OnComplete` 在填满那一帧只触发一次：`ui.Local` 记住上次完整码，后续帧码未变就不再调用；清空后重新填一个新码会再次触发。
- 格子尺寸三档：`SizeXs`/`SizeSm` → 24、`SizeMd` → 32、`SizeLg` → 40（主题间距为 4 时）；字号跟 `FrameFontSize`。
- 外观：`Variant` 取 `FrameLookOf`，`Native`/缺省时仍画带边框的面，否则格子会没有形状；当前待填的格子用强调色描边。
- 值以 `*string` 绑定，调用方在同一帧的下一句即可读到。
- 旧构造器 `NewInputOTP(c, value, digits, key...)` 保留为薄兼容层，内部转调 `NewOTP`。

**已知限制（框架的形状，不是后置）**
- 格子本身不可单独 Tab：键盘落在一个透明输入里，Tab 进出整个 OTP；逐格光标位置由框架的文本输入管理。
- 不做逐格自定义渲染（如掩码字符），要把 `*` 当内容就直接填 `*`。

## 4. 项目结构

- `forms/` 子包；依赖 `mygo/ui` 与根包 `myui`（共享词表与 tokens），不引入可选资源。

## 5. API 契约

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

## 6. 状态

- [x] 规格书 / 实现（`input_otp.go`）/ 测试（`input_otp_test.go`：绑定、过滤、字母数字、完成只触发一次、默认 6 位、尺寸档）
