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
| 无障碍 | `Element.Role(ui.RoleProgress)` ✓（自绘元素上设角色 ✓） | ✅ |

## 3. 实现到何种程度

**已实现**
- `myui.NewProgress(c, 0.4)`：圆角线 ✓，底色弱 ✓，强调色画到比例处 ✓；`Label(string)` ✓。
- `role=progressbar` 设在自绘元素上 ✓（用 `Element.Role`）✓。
- `Indeterminate()` / 负的 value：一段不确定时长的活动条 ✓（自绘 + `Painter.AnimationFrame`）✓。
- `Value(current, max)`、`Height(h)`、`Variant(v)`（状态色）✓。

**已知限制**
- 无障碍的数值上下限由自绘元素报不出（框架的 `hasRange/accRange` 非导出），角色可读 ✓、具体数值依赖 `Label` 说明。

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
