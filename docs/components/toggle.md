# Toggle 规格书

> 位置：控件在 [`controls/`](../../controls) 子包；主题/tokens 与文案表在根包 `myui`。
>
> 状态：**已实现**（`controls/toggle_color.go`）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui `Toggle` | 按下态是一等状态；有 variant / size、可带图标、用 `aria-pressed` 告知读屏 |
| mygo/ui `Toggle` | 点击或空格翻转 `*bool`，`Changed` 报告变化，角色是 toggle button |
| my-ui `Button` | 共享 Variant/Size 档位、边框圆角与 StateTransition 动效的基准线 |

## 2. 能否实现

- 按下态、键盘（空格）、角色与 `Changed`：mygo `ui.Toggle` 已提供 ✓
- Variant / Size / 图标 / 禁用：用 `ToggleBase` 之上的样式链与 `Children` 实现 ✓
- 注册式回调：`ui.Local` 存函数，组件每帧构建时读 `Changed` 并调用，晚 attach 也不丢 ✓
- 无障碍名称：`Element.Label` ✓

## 3. 实现到何种程度

- `NewToggle(c, on, label)`：绑定 app 的 `*bool`，内部用 `ui.Toggle` 底座，图标槽 + 文本槽。
- `apply()` 每帧从 `c.Theme()` 取颜色、间距、字号、圆角；按下态用 Accent 的半透明底与 Accent 墨色，Destructive 用 Danger，`Transition(StateTransition())` 让翻转渐变。
- 共享尺寸：Xs/Sm/Md/Lg 决定内边距、字号与高度（28/32/36/40），与 Button 同一词表。
- 回调两种：`OnChanged(fn)` 是帧内链式即时调用；`Handle(fn)` 存进元素的 `ui.Local`，可在任意帧 attach。
- `Changed()`、`Disabled(on)`、`Label(s)`、`Block()` 补齐读状态与常用开关。
- 在 `ToggleGroup` 内仍可当分段使用；`Block()` 用于需要铺满整行的场景。

## 4. 项目结构

- `controls/` 子包，与 Button 同层；dot-import 根包取共享词表，具名导入 `icons/`（示例页用）。
- 只依赖 mygo 公开 `ui` 与根包，不碰 internal，不带资产。

## 5. API 契约

| API | 作用 |
|---|---|
| `func NewToggle(c *ui.Context, on *bool, label string) *Toggle` | 构建绑定 `*on` 的切换按钮 |
| `func (t *Toggle) Variant(v Variant) *Toggle` | 应用共享变体 |
| `func (t *Toggle) Primary/Secondary/Outline/Ghost/Destructive() *Toggle` | 预设变体快捷方法 |
| `func (t *Toggle) Size(s Size) *Toggle` | 应用共享尺寸 |
| `func (t *Toggle) Xs/Sm/Md/Lg() *Toggle` | 尺寸快捷方法 |
| `func (t *Toggle) Icon(svg *ui.SVG) *Toggle` | 在文字前放图标（icon-only 时传空 label + `Label`） |
| `func (t *Toggle) Label(s string) *Toggle` | 无障碍名称 |
| `func (t *Toggle) Disabled(on bool) *Toggle` | 禁用或启用 |
| `func (t *Toggle) Block() *Toggle` | 铺满所在行 |
| `func (t *Toggle) Changed() bool` | 本帧是否翻转 |
| `func (t *Toggle) OnChanged(fn func(on bool)) *Toggle` | 本帧翻转时立即回调（必须同帧链式调用） |
| `func (t *Toggle) Handle(fn func(on bool)) *Toggle` | 注册常驻回调，可在任意帧调用 |

用法：

```go
var bold bool
ctrlpkg.NewToggle(c, &bold, "粗体").Primary().Icon(icons.Bold()).
    Handle(func(on bool) { app.status = fmt.Sprintf("bold=%v", on) })

if t.Changed() { /* 本帧翻转 */ }
```

## 6. 状态

- [x] 规格书（API 表逐项覆盖）
- [x] 实现（`toggle_color.go`）
- [x] 测试（`toggle_test.go`：点击翻转、Changed、两种回调、禁用、明暗主题）
- [x] showcase「Toggle」页（说明用途 + 状态对比 + 读值）

### 已知限制

- `OnChanged` 是即时链式调用，晚一帧 attach 会丢事件；这种场景用 `Handle`。
- 图标与文字的间距取主题 `Space(1)`，不提供单独的图标位置/大小选项（用共享 IconSizeSm）。
