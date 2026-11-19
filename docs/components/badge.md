# Badge 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`badge.go`）。按 Lobe MyGo 审查修复：Status 枚举、Text、ShowZero、Small、Processing 脉冲环。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [shadcn/ui Badge](https://ui.shadcn.com/docs/components/badge) | 文字徽标：`default`/`secondary`/`destructive`/`outline` × 三尺寸 |
| [Element Plus el-tag](https://element-plus.org/zh-CN/component/tag) | 语义色、尺寸档 |
| Lobe MyGo `feedback/badge.go` | `BadgeProps` 的 `Count`/`OverflowCount`/`ShowZero`/`Dot`/`Small`/`Color`/`Status`/`Text`；Processing 用 `ui.Loop` 画脉冲环，无计时器 |
| mygo 的 `ui.Badge(c, text)` | 框架自带胶囊观感（本库自建 pill 以完全控制颜色与尺寸） |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 计数/上限/零值 | `Count(n,max)`、`ShowZero` ✓ | ✅ |
| 状态点 + 文本 | 自建 Row + `Box`/`Text` ✓ | ✅ |
| Processing 脉冲 | `Element.Loop` + `Scale`/`Opacity` + `Border` ✓ | ✅ |
| Small | `Height`/`Padding` ✓ | ✅ |
| 主题跟随 | 每帧 `c.Theme()` ✓ | ✅ |

## 3. 实现到何种程度

- 文字 pill 与状态点两种形态共用一个 root；`Variant`/`Count` 切回 pill，`Status` 切到状态点。
- 状态色按类型取：`Online`→success、`Syncing`→info、`Degraded`→warning、`Offline`→danger、`Idle`→muted、`Processing`→accent；`Color` 可覆盖。
- `Processing` 的环：`dot.Loop("badge-pulse", 1200ms, ui.EaseOut)` 算出 `scale = 1+1.4t`、`opacity = 1-t`，画一个描边圆圈。Reduce Motion 由框架兜底。
- `Count(0,max)` 默认隐藏；`ShowZero(true)` 保留；`max>0 && n>max` 显示 `max+`。
- `Small(true)` 让计数 pill 高 14 DIP（默认 18）。

## 4. 项目结构

- `display/` 子包；共享根包的 `Variant`/`Size` 词表。

## 5. API 契约（逐 prop）

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

## 6. 限制

- `BadgeOptions.Dot` 只把角标缩成小方块，颜色仍由 Variant 决定（默认 Destructive）。
- 状态点的脉冲环是纯描边、不跟随主题以外的状态；Reduce Motion 下框架直接到目标值。

## 7. 状态

- [x] 调研与实现（`badge.go`）
- [x] 测试：变体颜色、尺寸、暗色主题、计数/ShowZero/Small、Status/Text
