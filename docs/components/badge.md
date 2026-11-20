# Badge 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

提供 Status 枚举、Text、ShowZero、Small、Processing 脉冲环。

## API 契约（逐 prop）

### 类型

| 符号 | 作用 |
|---|---|
| `type BadgeStatus uint8` | 状态枚举 |
| `BadgeOnline` / `BadgeSyncing` / `BadgeDegraded` / `BadgeOffline` / `BadgeIdle` / `BadgeProcessing` | 六个取值，零值为 `BadgeOnline` |

### 构造与方法

| 方法 | 作用 |
|---|---|
| `NewBadge(c, text) *Badge` | 创建文字徽标（Native 预设） |
| `Variant(v Variant) *Badge` | 文字 pill 的预设，并切回 pill 形态 |
| `Primary/Secondary/Default/Native/Outline/Ghost/Destructive/Link/Success/Info/Warning() *Badge` | `Variant(...)` 的快捷方法 |
| `Color(c ui.Color) *Badge` | 覆盖 pill 底色或状态点色；零值清除覆盖 |
| `Size(s Size) *Badge` / `Xs/Sm/Md/Lg()` | 尺寸档（字号 + 内衬） |
| `Count(n, max int) *Badge` | 计数；`max>0` 时超出显示 `max+` |
| `ShowZero(on bool) *Badge` | 零计数是否显示 |
| `Small(on bool) *Badge` | 计数 pill 是否用 14 DIP 小尺寸 |
| `Status(s BadgeStatus) *Badge` | 切到状态点形态，并选颜色；`Processing` 画脉冲环 |
| `Text(s string) *Badge` | 状态点旁的文本，或 pill 的文本 |
| `Closable(closed *bool) *Badge` | 放关闭按钮，写入 flag |

### 角标

| 符号 | 作用 |
|---|---|
| `type BadgeOptions struct` | `Content`、`Dot`、`Count`、`Max`、`Hidden`、`Variant`、`ShowZero`、`Small` |
| `NewBadgeOn(c, anchor, o) *Badge` | 把徽标钉在内容的右上角 |

用法：

```go
display.NewBadge(c, "").Status(BadgeProcessing).Text("同步中")
display.NewBadge(c, "").Count(120, 99).Small(true) // 99+
```

## 限制

- `BadgeOptions.Dot` 只把角标缩成小方块，颜色仍由 Variant 决定（默认 Destructive）。
- 状态点的脉冲环是纯描边、不跟随主题以外的状态；Reduce Motion 下框架直接到目标值。
