# AlertDialog 规格书

> 状态：**已实现**。桌面风格的系统警告框：标题、说明与一行按钮，按钮走 myui 预设，最后一个按钮默认为 Destructive，结果由 `OnResult` 交回 app。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `ui.AlertDialog(c, open, title, message, buttons...)` | 契约：`*open` 为真时显示；返回被点按钮的索引，`ui.AlertDismissed`（-2）表示 Escape 取消，-1 表示本帧没有结果；点中按钮或取消都会把 `*open` 置回 false |
| shadcn/ui `AlertDialog` | 面板是一张卡片：描边、圆角、底色、软阴影；标题加粗、说明弱化；按钮右对齐，确认按钮带强调色 |
| my-ui `overlays.NewModal` | 浮层内容建在 `ui.DialogBase` 的 panel 里；背景遮罩、面板配色、`OverlayEnter` 的淡入都来自 `c.Theme()` 与共享 token |

## 2. 能否实现

- 不能直接给 mygo 的 `ui.AlertDialog` 换外观：它返回的是 `int` 而不是 `ui.Element`，按钮由它自己在浮层里画，公开 API 里没有可以改样式的句柄。
- 因此用公开的 `ui.DialogBase` 复刻同一套契约：浮层里自己排标题、说明和 myui `Button`，保留索引语义与 `ui.AlertDismissed`，并把外观换成 shadcn 风格。
- `OnResult` 在结果产生的那一帧同步调用；没有结果时不调用。

## 3. 实现到何种程度

- `NewAlertDialog(c, open, AlertDialogOptions) ui.Element`：面板用主题底色、描边、XL 圆角与软阴影，标题 `TextLg` 加粗、说明 `TextSm` 弱化，按钮右对齐可换行。
- 非最后一个按钮用 `Button.Outline()`；最后一个用 `Button.Destructive()` 并 `AutoFocus()`，Enter 即点中它。
- 颜色与尺寸每帧从 `c.Theme()` 取；状态淡入淡出走 `StateTransition()`，不加计时器或 goroutine。
- 面板的名字是标题，读屏的说明是 message（`ui.Element.Label` / `Description`），角色为 `ui.RoleAlertDialog`。
- 已知差异：`ui.DialogBase` 点遮罩也会关闭，而 mygo 的 `ui.AlertDialog` 忽略遮罩点击；这一点是公开底座的行为，不能在不自己重画浮层的前提下关掉。

## 4. 项目结构

- `overlays/` 子包；`controls.NewButton` 提供按钮预设，根包提供 `Radius`/`TextSize`/`StateTransition`/`OverlayEnter`。

## 5. API 契约

```go
type AlertDialogOptions struct {
    Title    string
    Message  string
    Buttons  []string
    OnResult func(i int)
}

func NewAlertDialog(c *ui.Context, open *bool, o AlertDialogOptions) ui.Element
```

| 属性 | 类型 | 作用 |
|---|---|---|
| `open` | `*bool` | app 持有的开关；为真时显示，按钮或关闭后置回 false |
| `o.Title` | `string` | 加粗标题，同时作为面板的无障碍名称 |
| `o.Message` | `string` | 标题下的弱化说明，同时作为面板的无障碍描述 |
| `o.Buttons` | `[]string` | 按钮文案，按序排列；最后一个默认 Destructive 并自动聚焦 |
| `o.OnResult` | `func(i int)` | 有结果时回调：按钮索引，或 `ui.AlertDismissed`（-2） |
| 返回值 | `ui.Element` | 浮层面板；关闭期间是无效的 Element |

## 6. 状态

- [x] 规格书 / 实现 / 测试 / showcase
