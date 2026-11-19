# Tag 规格书

> 状态：**已实现**。按 Lobe MyGo 审查修复：任意 `Color`、`TagVariant`、`Round`、关闭时 `Hidden(true)` 仍调 `OnClose`。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| Lobe MyGo `display/tag.go` | `TagVariant`（Filled/Outlined/Borderless/Solid）、`Color`、`Round`；关闭时 `*closed=true; Hidden(true); OnClose()` |
| Element Plus Tag | 弱底 + 文字 + 可关闭图标 |
| shadcn/ui Badge | 共享 Variant/Size 词表 |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 任意色 / 四种表面 | `Background`、`Border`、`TextColor`、`MixOklab`、`Alpha`、`OnAccent` ✓ | ✅ |
| 圆角 | `Radius(999)` ✓ | ✅ |
| 关闭 + 隐藏 + 回调 | `NewIconAction`、`Hidden`、`ui.Local` ✓ | ✅ |

## 3. 实现到何种程度

- `Color` 取任意 `ui.Color`；零值回退到共享语义 `Variant`，再回退到主题中性墨色。
- `TagVariant`：`Filled`（12% 淡底）、`Outlined`（6% 底 + 45% 描边）、`Borderless`（仅文字）、`Solid`（整块底色 + `OnAccent` 可读文字）。
- `Round` 用 `Radius(999)` 全圆角。
- 关闭图标按下：写 `*Closed`、用 `ui.Local` 记住 closed 并 `Hidden(true)`、**仍然调用 `OnClose`**。
- 颜色每帧从 `c.Theme()` 取。

## 4. 项目结构

- `display/` 子包。

## 5. API 契约（逐 prop）

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

## 6. 限制

- `Variant` 与 `Color` 同时存在时 `Color` 优先；语义色只在 `Color` 为零值时使用。
- 关闭状态通过 `ui.Local` 跨帧保留，即使 app 没有读 `Closed` 也不会重新出现。

## 7. 状态

- [x] 规格书 / 实现 / 测试（关闭 + OnClose、四种 TagVariant、Round、Color）
