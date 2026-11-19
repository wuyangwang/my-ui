# HotkeyInput 规格书

> 位置：`controls/hotkeyinput.go`。

> 状态：**已实现**。参考 Lobe UI 的 HotkeyInput：聚焦后录制快捷键，拒绝没有修饰键或已被占用的组合。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [Lobe UI HotkeyInput](https://mygo.lobehub.com/components/hotkey-input) | 录制、Placeholder、ResetValue、AllowClear、Conflicts、OnConflict |
| mygo `ui.InputEvent` / `ui.Modifiers` / `ui.Key` | 键盘事件的公开 API |
| `controls/iconaction.go` | 重置与清除按钮复用同一图标 affordance |

## 2. 能否实现

- `HandleInput` 在 `InputKeyDown` 上取修饰键与按键；`Tab`/`Shift-Tab` 放行让焦点离开 ✓。
- 快捷键写成 `mod+shift+k` 的名字串；`mod` 是主键修饰（Apple 与 Windows 的命名差异留给应用的 `FormatHotkey` 组合），本组件取 Ctrl 为 mod、Super 为 meta ✓。
- 冲突表、无效提示、重设值、清除按钮都从 app 状态或 fluent 参数设置 ✓。
- 重置/清除按钮由各自的 fluent 方法填进独立插槽，避免同帧重复追加子节点 ✓。

## 3. 实现到何种程度

录制、修饰键检查、冲突拒绝与回调、无效提示、重设与清除按钮、可本地化的两条消息、占位文字、字段边框随焦点/错误换色。键帽用主题 Surface 与边框绘制。

## 4. 项目结构

- `controls/hotkeyinput.go`；`controls/hotkeyinput_test.go` 覆盖录制、冲突、无效键、重设与清除、名字格式化。

## 5. API 契约

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

## 6. 限制

- 不自动检测 Apple 平台；`mod` 取 Ctrl，`meta` 取 Super，需要平台化命名时由应用改写键名。
- 键名表覆盖字母、数字、常见命名键与 F1–F12；其他按键不能录制。
- 默认消息为英文，用 `Messages` 本地化。
