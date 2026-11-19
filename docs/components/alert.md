# Alert 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`alert.go`）。按 Lobe MyGo 审查修复：Type/Variant 正交、NoIcon、Selectable。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Alert](https://ui.shadcn.com/docs/components/alert) | 卡片：`rounded-lg border px-4 py-3 text-sm` + `role="alert"`；标题加粗、描述弱化 | 圆角卡片、图标在左、描述在下 ✓ |
| Lobe MyGo `feedback/alert.go` | `AlertType`（Info/Success/Warning/Error/Secondary）与 `AlertVariant`（Soft/Outlined/Plain）两个正交轴；Soft 用主题色淡底 + 淡描边；`NoIcon`、`Selectable`、`Closed` | 全部照做 ✓ |
| mygo 的 `ui.RoleAlert` | live region：一出现读屏就念 | 默认角色 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 卡片/淡底/淡描边 | `Background`、`MixOklab`、`InsetShadow`、`Border` ✓ | ✅ |
| 图标/文字选择 | `ui.Icon`、`Selectable`、`TextColor` ✓ | ✅ |
| 关闭按钮 | `controls.NewIconAction`（读 `Clicked()`）✓ | ✅ |
| 无障碍 | `Role(ui.RoleAlert)` / `Role(ui.RoleStatus)` ✓ | ✅ |
| 主题跟随 | 每帧从 `c.Theme()` 取色、`Transition(StateTransition())` ✓ | ✅ |

## 3. 实现到何种程度

- Type 决定色调与默认图标，Variant 决定表面，二者正交。
- Soft：`t.Surface.MixOklab(tone, 0.06)` 淡底 + `InsetShadow(0,0,0,1, tone.Alpha(0.16))` 淡描边。
- Outlined：透明底 + `t.Border` 描边；Plain：无底无描边、内衬更紧。
- 五个类型的图标在构造时各建一个并只显示当前类型的一个；`Icon(svg)` 覆盖默认图标。`NoIcon` 隐藏图标槽。
- `Selectable` 让标题和描述可被选中复制；`Closed(flag)` 放关闭按钮，写入 flag 并 `Hidden(true)`。
- 颜色每帧从 `c.Theme()` 取，无包级缓存、无计时器/goroutine。

## 4. 项目结构

- `display/` 子包；只用公开 `ui` API 与根包共享词表，不 import mygo 的 internal 包。

## 5. API 契约（逐 prop）

### 类型

| 符号 | 作用 |
|---|---|
| `type AlertType uint8` | 消息种类；决定色调与默认图标 |
| `AlertInfo` / `AlertSuccess` / `AlertWarning` / `AlertError` / `AlertSecondary` | 五个取值，零值为 `AlertInfo` |
| `type AlertVariant uint8` | 表面；与 Type 正交 |
| `AlertSoft` / `AlertOutlined` / `AlertPlain` | 三个取值，零值为 `AlertSoft` |

### 构造与方法

| 方法 | 作用 |
|---|---|
| `NewAlert(c *ui.Context, title string) *Alert` | 创建提示；默认 Info + Soft |
| `Type(t AlertType) *Alert` | 设置消息种类（色调 + 默认图标） |
| `Variant(v AlertVariant) *Alert` | 设置表面：Soft / Outlined / Plain |
| `Description(s string) *Alert` | 标题下的描述；空串清除 |
| `Icon(svg *ui.SVG) *Alert` | 覆盖默认图标；nil 恢复类型图标 |
| `NoIcon(on bool) *Alert` | 隐藏（或显示）左侧图标 |
| `Selectable(on bool) *Alert` | 标题与描述可被选中复制 |
| `Closed(flag *bool) *Alert` | 放关闭按钮，写入 flag 并隐藏提示 |
| `Closable(flag *bool) *Alert` | `Closed` 的旧名 |
| `Action(build func()) *Alert` | 右侧操作区（shadcn 的 AlertAction） |
| `Role(r ui.Role) *Alert` | 无障碍角色，默认 `RoleAlert`，可改 `RoleStatus` |
| `Info()` / `Success()` / `Warning()` / `Error()` / `Secondary()` | `Type(...)` 的快捷方法 |
| `Destructive()` | `Error()` 的旧名 |
| `Soft()` / `Outlined()` / `Plain()` | `Variant(...)` 的快捷方法 |
| `Default()` / `Native()` | 恢复 Info + Soft |

用法：

```go
display.NewAlert(c, "出错了").
    Description("网络没有连上。").
    Error().                 // Type
    Variant(AlertOutlined).  // Variant
    Selectable(true)
```

## 6. 限制

- 关闭按钮由库提供；隐藏由 app 的 flag 驱动（下一帧重建时生效）。
- 一次只显示一条，多条自行排布。
- 默认图标按类型在构造时绑定；运行时反复改 `Type` 而不调 `Icon` 时图标随当前类型切换（五个槽同时存在、切换显隐）。

## 7. 状态

- [x] 调研与实现（`alert.go`）
- [x] 测试：标题/描述、live region、Type×Variant、NoIcon/Selectable、关闭
