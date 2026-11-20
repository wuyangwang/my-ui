# HotkeyInput 规格书

> 位置：`controls/hotkeyinput.go`。

参考 Lobe UI 的 HotkeyInput：聚焦后录制快捷键，拒绝没有修饰键或已被占用的组合。

## shadcn 风格

字段框是 shadcn Input 的观感：`t.Background` 底、1px `t.Border` 发丝边、
`Radius(t, RadiusMd)`；聚焦时描边转 `t.Accent`，错误时转 `t.Danger`，Ghost 变体透明、
Secondary 变体用 `t.Surface`。占位文字 `t.TextMuted`，提示信息 `t.Danger`。键帽沿用
`display.NewKbd` 的胶囊：`t.Surface` 底 + 1px `t.Border` + `Radius(t, RadiusSm)`，
字号 `TextSize(t, TextXs)`。字段最小高度 `t.Space(Spacing2Xl)`（默认主题下 32 DIP），
随主题间距缩放；没有 Lobe 固定色板与自造尺寸。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewHotkeyInput(c *ui.Context, value *string) *HotkeyInput` | 建录制器，写入 `*value` |
| `Placeholder(s string)` | 空值时的提示文字 |
| `Reset(s string)` | 默认快捷键，值不同时显示重设按钮 |
| `AllowClear(on bool)` | 显示清除按钮 |
| `Conflicts(keys ...string)` | 拒绝的快捷键表 |
| `OnConflict(fn func(string))` | 被拒绝时回调 |
| `Messages(invalid, conflict string)` | 两条拒绝消息 |
| `Variant(v Variant)` | 字段框的共享外观 |
| `Disabled(on bool)` | 禁用录制 |
| `Control() ui.Element` | 焦点所在的内层控件 |
| `Value() string` · `Changed()` | 当前值与本帧是否变化 |

## 限制

- 不自动检测 Apple 平台；`mod` 取 Ctrl，`meta` 取 Super，需要平台化命名时由应用改写键名。
- 键名表覆盖字母、数字、常见命名键与 F1–F12；其他按键不能录制。
- 默认消息为英文，用 `Messages` 本地化。
