# Kbd 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`kbd.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Kbd](https://ui.shadcn.com/docs/components/kbd) | `rounded border bg-muted px-1.5 font-mono text-xs` | 小方块 + 描边 + 弱底色 + 等宽字 ✓ |
| mygo | 没有 kbd 部件 ✗ | 用基础元素搭 ✓（`Element.Font("monospace")` ✓） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 等宽字 | `Element.Font("monospace")` ✓ | ✅ |
| 小方块 | `Element.Border`/`Radius`/`Padding` ✓ | ✅ |

## 3. 实现到何种程度

**已实现**
- `displaypkg.NewKbd(c, "Ctrl", "K")`：几个键帽排一行 ✓；弱底色 + 描边 + 圆角 + `TextXs` + 等宽 ✓。

**已知限制**
- 键帽之间没有「+」之类的连接符 ✗（app 自己写 ✓）。

## 4. 项目结构是否合理

- 位置：根包 L3——键帽是排版件 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
func NewKbd(c *ui.Context, keys ...string) *Kbd
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`kbd.go`）
- [x] 测试（键帽在屏 ✓）
