# Field 规格书

> 状态：**已实现**（`form.go`）。表单字段的组合层：标签 + 控件 + 说明 + 错误。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| [shadcn v4 Field](https://ui.shadcn.com/docs/components/base/field)（10 个部件） | `Field`（`orientation` vertical/horizontal/responsive，`data-invalid`，`role="group"`，错误时整组文字转 danger）；`FieldLabel`（可点击聚焦控件、必填星号）；`FieldDescription`（弱化色）；`FieldError`（多条消息成列表，`role="alert"`，接受校验器的 issues）；`FieldGroup`/`FieldSet`/`FieldLegend`/`FieldContent`/`FieldSeparator`/`FieldTitle` | 取 **Field + Label + Description + Error** 四件（Q13 定的最小集）；标签可点击聚焦；必填星号；错误 `role="alert"` 且控件被标记为无效 |
| [Element Plus el-form / el-form-item](https://element-plus.org/en-US/component/form.html) | `label-suffix`、`hide-required-asterisk`、`require-asterisk-position`（默认 left）、`label-position`（默认 right）、`status-icon`、`scroll-to-error`、`inline-message`、`validate-status` | 星号在标签后（element-plus 的默认位置）；错误是控件的边框与文案（我们把边框交给控件的 `Invalid`） |
| [Ant Design Form](https://ant.design/components/form-cn) | `status="error"｜"warning"` 两种校验色；`help` 文案；`rules` 校验规则 | 校验色只有 danger（语义色 Success/Warning 在控件层可用 ✗） |
| [Lobe MyGo Form](https://mygo.lobehub.com/skills/components/form) | 表单持有初始值、`IsDirty`/`Reset`、每字段的 touched/error/running；失焦校验、异步校验与取消、依赖字段重校验 | **不实现校验层**（Q1 定的范围）：规则与时机留给 app |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 标签可点击聚焦控件 | `Element.Focus()` ✓；标签做成 `ui.ButtonBase`（无外观）+ `TabStop(false)`（不进 Tab 序）✓ | ✅ |
| 控件的无障碍名 | `Element.Label(string)` ✓ | ✅ 字段把标签文字给控件 ✓ |
| 说明与控件的关系 | `Element.DescribedBy(el)` ✓ | ✅ |
| 错误与控件的关系 | `Element.ErrorMessage(el)` ✓ + `Element.Invalid(bool)` ✓ | ✅ |
| 错误被读出 | `Element.Role(ui.RoleAlert)` ✓（`RoleAlert` 是活区域 ✓） | ✅ |
| 必填标记 | 标签里的一个 `*` 文本 ✓ | ✅ |
| 校验规则/时机 | 无公开的「刚获得焦点」事件（只有 `Focused()` 状态 ✗）；`MaxLength(n)` ✓ 只限长不计数 | ⚠️ 校验层不做（Q1）；字数限制可做（见 input.md 后置项） |

## 3. 实现到何种程度

**已实现**
- `NewField(c, label, build)`：标签（可点击聚焦控件）+ 控件 + 说明 + 错误，依次排开。
- `Description(s)`：控件下的说明（弱化色），并与控件建立 `DescribedBy` 关系 ✓。
- `Error(s)`：控件下的错误（danger 色、`role="alert"`），并与控件建立 `ErrorMessage` 关系、控件标记为无效 ✓。
- `Required(on)`：标签后加一个星号 ✓。
- `Invalid(on)`：控件标记为无效，标签转 danger 色 ✓。
- `Control()`：字段描述的那个元素（控件本身，或控件的文本）✓。

**分工**：控件保留 `Label(string)`（无障碍名，控件自己的责任 ✓）；错误文案与说明归字段 ✓（控件的 `ErrorText` 已在 Q9 决定中去掉 ✗ → 待办）。

**后置**
- 校验层（规则、时机、dirty/touched、异步校验与取消、依赖重校验）——按 Q1 留给 app 或后续。
- 横向排列的字段组（shadcn 的 `orientation="horizontal"`）、字段集与图例（`FieldSet`/`FieldLegend`）、多条错误消息成列表。
- 必填标记的位置可配（element-plus 的 `require-asterisk-position`）。

## 4. 项目结构是否合理

- 位置：根包 L4（组合件），在 `field.go` 的字段基座之上；只依赖 `mygo/ui` ✓。
- 与控件的关系：字段不复制控件的样式，只编排标签/说明/错误并建立无障碍关系 ✓。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewField(c *ui.Context, label string, build func() ui.Element) *Field` | 创建字段；`build` 造控件并返回它描述的那个元素 |
| `Description(s string) *Field` | 控件下的说明，并与控件建立 `DescribedBy` 关系 |
| `Error(s string) *Field` | 控件下的错误（`role="alert"`），并把控件标记为无效 |
| `Required(on bool) *Field` | 标签上的必填星号 |
| `Invalid(on bool) *Field` | 控件无效，标签转 danger 色 |
| `Control() ui.Element` | 字段描述的那个元素 |

用法：

```go
h := myui.Bind(c)
h.Field("邮箱", func() ui.Element {
	return h.Input(&app.email).Outline().Control()
}).Description("我们只用它发通知").Required(true).Error(app.emailError)
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`form.go`）
- [x] 测试（`form_test.go`：标签/说明/错误在屏、错误是 alert、标签聚焦控件）
- [x] showcase「表单」页
- [ ] 校验层、横向字段组、字段集、多条错误（后置）
