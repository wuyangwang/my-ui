# Progress 规格书

> 状态：**已实现**（`progress.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Progress](https://ui.shadcn.com/docs/components/progress) | `h-2 w-full overflow-hidden rounded-full bg-primary/20` + 里面的 `bg-primary` | 细圆角线 ✓；底色弱 ✓；走到的那段是强调色 ✓ |
| mygo 的 `ui.Progress(c, value)` | 自带主题外观的进度条 ✓ | 我们的自己画（shadcn 的外观 ✓） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 画一条自定义的条 | `Element.Draw` + `Painter.Fill` ✓ | ✅ |
| 无障碍 | `role=progressbar` 由框架的基座给 ✓（我们用普通 Box ✗） | ⚠️ 部分 |

## 3. 实现到何种程度

**已实现**
- `myui.NewProgress(c, 0.4)`：2 DIP 高的圆角线 ✓，底色弱 ✓，强调色画到比例处 ✓；`Label(string)` ✓。

**已知限制**
- 没有用框架的 `role=progressbar` ✗（我们自绘，普通元素 ✗；读屏读到的是名字与角色 ✗）——可改用 `ui.Progress` 的基座 ✗（框架没有公开的基座 ✗）。
- 没有不确定时长的动画（`indeterminate`）✗。

## 4. 项目结构是否合理

- 位置：根包 L3 ✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
func NewProgress(c *ui.Context, value float64) *Progress
func (p *Progress) Label(label string) *Progress
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`progress.go`）
- [x] 测试（在屏 ✓）
- [ ] `role=progressbar`、不确定时长（后置 ✗）
