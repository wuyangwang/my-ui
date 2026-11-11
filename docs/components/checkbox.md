# Checkbox 规格书

> 状态：**已实现**（`checkbox.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn/ui Checkbox](https://ui.shadcn.com/docs/components/checkbox) | 16px 方框 + 圆角 + 描边；选中时强调色底 + 白色对勾；`checked`/`disabled` | 框由库画 ✓（`CheckboxBase` + 自绘） |
| mygo 的 `ui.CheckboxBase(c, checked)` | 无外观的基座：点击或空格切换、`Changed()` 报告、`role=checkbox` ✓ | 基座 ✓，外观我们给 ✓ |
| mygo 的 `ui.CheckboxStateBase(c, state)` / `ui.CheckboxGroup` | 半选态（`CheckOn`/`CheckOff`/`CheckMixed`）与「全选」组 ✓ | 暂未包装 ✗（见「已知限制」） |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 勾选/取消 | `ui.CheckboxBase(c, *bool)` ✓ | ✅ |
| 外观自绘 | 基座返回行，框与标签由 app 建 ✓ | ✅ |
| 尺寸 | `Element.Size` ✓ | ✅ |
| 禁用 | `Element.Disabled` ✓ | ✅ |
| 半选/全选组 | `ui.CheckboxStateBase` ✓ / `ui.CheckboxGroup` ✓ | ✅（未包装 ✗） |

## 3. 实现到何种程度

**已实现**
- `myui.NewCheckbox(c, &checked, "标签")`：方框（14/16/18/22 DIP 四档）+ 圆角 + 描边；选中时强调色底 + 对勾 ✓；标签是文字，也是无障碍名 ✓。
- `Size`/`Xs`/`Sm`/`Md`/`Lg` ✓；`OnChanged(func(bool))` ✓；`Disabled(bool)` ✓。

**已知限制**
- 半选态与「全选」组没有包装（框架有 `ui.CheckboxState`/`ui.CheckboxGroup` ✓，直接用即可 ✗）。
- 没有 `Required` 等校验钩子 ✗（用 `Field` 的校验层 ✓）。

## 4. 项目结构是否合理

- 位置：根包 L3——勾选框是变体化控件 ✓。
- 依赖：只用 `mygo/ui` + `icons/` ✓。

## 5. API 契约

```go
func NewCheckbox(c *ui.Context, checked *bool, label string) *Checkbox
func (cb *Checkbox) Size(s Size) *Checkbox
func (cb *Checkbox) Xs() *Checkbox
func (cb *Checkbox) Sm() *Checkbox
func (cb *Checkbox) Md() *Checkbox
func (cb *Checkbox) Lg() *Checkbox
func (cb *Checkbox) OnChanged(fn func(checked bool)) *Checkbox
func (cb *Checkbox) Disabled(disabled bool) *Checkbox
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`checkbox.go`）
- [x] 测试（点击切换 ✓、事件 ✓、禁用 ✓）
- [ ] 半选态与全选组的包装（后置 ✗）
