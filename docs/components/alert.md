# Alert 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

提供 Type/Variant 正交、NoIcon、Selectable。

## API 契约（逐 prop）

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

## 限制

- 关闭按钮由库提供；隐藏由 app 的 flag 驱动（下一帧重建时生效）。
- 一次只显示一条，多条自行排布。
- 默认图标按类型在构造时绑定；运行时反复改 `Type` 而不调 `Icon` 时图标随当前类型切换（五个槽同时存在、切换显隐）。
