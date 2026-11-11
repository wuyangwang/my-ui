# Card 规格书

> 状态：**已实现**（`card.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Card](https://ui.shadcn.com/docs/components/card) | `Card`（`rounded-xl border bg-card shadow`）+ `CardHeader`/`CardTitle`/`CardDescription`/`CardContent`/`CardFooter` | 卡片：surface 底 + 描边 + 圆角 + 内衬 ✓；标题/说明/内容 ✓ |
| mygo | 没有 card 部件 ✗ | 用基础元素搭 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 卡片 | `ui.Column` + 样式方法 ✓ | ✅ |
| 标题/说明/内容 | app 自己建 ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `myui.NewCard(c, "概览")` → `.Description("…")` / `.Content(func(){…})`。
- 卡片：surface 底 + 描边 + 圆角 + `Space(4)` 内衬 ✓；标题 `TextLg` 加粗 ✓；说明弱一些 ✓。

**已知限制**
- 没有头尾分区（`CardHeader`/`CardFooter`）✗（app 自己排 ✓）。
- 没有阴影 ✗（shadcn 有 ✗）。

## 4. 项目结构是否合理

- 位置：根包 L3——卡片是版式件 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
func NewCard(c *ui.Context, title string) *Card
func (card *Card) Description(s string) *Card
func (card *Card) Content(build func()) *Card
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`card.go`）
- [x] 测试（标题、说明与内容 ✓）
