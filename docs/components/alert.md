# Alert 规格书

> 状态：**已实现**（`alert.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Alert](https://ui.shadcn.com/docs/components/alert)（源码 `ui/alert.tsx`） | `Alert`（`rounded-lg border px-4 py-3 text-sm` + `role="alert"`）· `AlertTitle`（`font-medium`）· `AlertDescription`（`text-sm text-muted-foreground`）· 变体（`default` / `destructive`，破坏性的字是危险色）· 图标（`[&>svg]:size-4 text-current`） | 卡片：描边 + 圆角 + `Space(4)/Space(3)` 内衬 + `TextSm` ✓；`role=alert` ✓；标题加粗、描述弱一些 ✓；图标在左边、随文字颜色 ✓；破坏性变体用危险色 ✓ |
| mygo 的部件 | **没有** alert 部件 ✗ | 用基础元素搭 ✓（`ui.Box` + `Row`/`Column`） |
| mygo 的 `ui.RoleAlert` | 无障碍的 live region：内容一出现，读屏就念出来 ✓；字段的错误也是它 ✓ | 提示框用它 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 卡片 | `ui.Box(c).Row()...` + `Background`/`Border`/`Radius`/`Padding` ✓ | ✅ |
| 图标 | `ui.Icon(c, svg)` ✓（取周围文字的颜色 ✓） | ✅ |
| live region | `ui.RoleAlert` ✓ | ✅ |
| 主题变色 | `c.Theme()` 的槽位 + `Transition` ✓ | ✅ |
| 组件库自带的 alert | 没有 ✗ | ❌ 自己搭 ✓ |

## 3. 实现到何种程度

**已实现**
- `myui.NewAlert(c, "注意")` → `.Description("…")`/`.Icon(svg)`/`.Variant(v)` + `.Default()`/`.Destructive()`/`.Success()`/`.Info()`/`.Warning()`。
- 卡片：surface 底 + 描边 + 圆角 + `Space(3)/Space(4)` 内衬 + `TextSm` ✓；图标在左（16 DIP ✓），标题加粗，描述弱一些 ✓。
- 变体给墨色：破坏性是危险色 ✓，success/info/warning 是语义色 ✓（描述用同色九成 ✓，shadcn 的 `text-destructive/90` ✓）。
- `role=alert` ✓ → 读屏在它出现时念出来 ✓。
- 主题变化时颜色渐变（`Transition`）✓。

**已知限制**
- 没有标题/描述以外的部件 ✗（自己搭 ✓）。
- 关闭按钮要 app 自己加 ✓（框架的部件没有 ✗）。
- 一次只显示一条 ✗（多条就多条排下去 ✓）。

## 4. 项目结构是否合理

- 位置：根包 L3——提示框是「页面里的一段消息」✓。
- 依赖：只用 `mygo/ui` ✓。

## 5. API 契约

```go
// NewAlert creates an alert with a title.
func NewAlert(c *ui.Context, title string) *Alert

func (a *Alert) Description(s string) *Alert
func (a *Alert) Icon(svg *ui.SVG) *Alert
func (a *Alert) Variant(v Variant) *Alert
func (a *Alert) Default() *Alert
func (a *Alert) Native() *Alert
func (a *Alert) Destructive() *Alert
func (a *Alert) Success() *Alert
func (a *Alert) Info() *Alert
func (a *Alert) Warning() *Alert
```

用法：

```go
myui.NewAlert(c, "保存前看一眼").
	Description("这条改动会影响所有人。").
	Warning().
	Icon(icons.TriangleAlert())
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`alert.go`）
- [x] 测试（标题与描述 ✓、live region ✓、图标在左 ✓）
- [ ] 自带关闭按钮（后置 ✗）
