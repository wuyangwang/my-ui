# Tag 规格书

提供 任意 `Color`、`TagVariant`、`Round`、关闭时 `Hidden(true)` 仍调 `OnClose`。

## API 契约（逐 prop）

### 类型

| 符号 | 作用 |
|---|---|
| `type TagVariant uint8` | 标签表面 |
| `TagFilled` / `TagOutlined` / `TagBorderless` / `TagSolid` | 四个取值，零值为 `TagFilled` |
| `type TagOptions struct` | `Label`、`Icon`、`Closed`、`OnClose`、`Color`、`Round`、`Size` |

### 构造与方法

| 方法 | 作用 |
|---|---|
| `NewTag(c, o TagOptions) *Tag` | 创建标签 |
| `Variant(v Variant) *Tag` | 共享语义色（无 `Color` 时生效） |
| `TagVariant(v TagVariant) *Tag` | 设置表面：Filled / Outlined / Borderless / Solid |
| `Color(c ui.Color) *Tag` | 任意色；零值清除 |
| `Round(on bool) *Tag` | 全圆角开关 |
| `Size(sz Size) *Tag` | 字号档（Xs/Sm/Md/Lg） |

用法：

```go
display.NewTag(c, display.TagOptions{Label: "新功能", Color: ui.Hex("#7c3aed"), Round: true}).
    TagVariant(TagSolid)
```

## 限制

- `Variant` 与 `Color` 同时存在时 `Color` 优先；语义色只在 `Color` 为零值时使用。
- 关闭状态通过 `ui.Local` 跨帧保留，即使 app 没有读 `Closed` 也不会重新出现。
