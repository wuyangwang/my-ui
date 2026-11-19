# Form（my-ui 外观）规格书

> 位置：`forms/` 子包（`forms/form_styled.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| Lobe `form.*` | 分组、字段、校验、异步、自动保存 | 由已有的 `forms.Field`/`Rule` 体系承担 ✓ |
| mygo `ui.Form` | 标签对齐的表单容器 | 直接包裹并套主题外观 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 标签右对齐排版 | `ui.Form` 内建 | ✅ |
| 字段/说明/错误 | `forms.NewField`、`ui.Field` | ✅ |
| 主题外观 | `c.Theme()` + `StateTransition` | ✅ |
| 签名兼容 | `func(c *ui.Context, fn func()) ui.Element` | ✅ |

## 3. 实现到何种程度

- `NewForm(c, fn)` 与 mygo 的 `ui.Form(c, fn)` 同签名：旧的调用改为包名即可编译。容器套用主题间距、圆角与表面。
- `NewFormWith(c, variant, fn)` 提供共享 `Variant` 预设：`Native`（零值，无面）/ `Outline`（描边表面）/ `Secondary`（弱底）。
- 字段的对齐仍由 mygo 的 `ui.Form` 完成；校验与错误由本库已有的 `forms.Field`/`Rule` 负责。

**已知限制**
- 这不是 Lobe 的泛型 `form.New[T]` 状态容器：那条路线需要反射/泛型状态管理，与 mygo 的即时模式绑定约定冲突。需要强类型表单时，用 `NewForm` 包裹 `NewField` 与自己的状态指针。
- 异步校验与防抖不在本组件内：mygo/my-ui 不设计时器，应用可自行在字段外管理。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `NewForm(c, fn)` | 带 my-ui 外观的表单，签名兼容 `ui.Form` |
| `NewFormWith(c, variant, fn)` | 带共享 `Variant` 预设的表单 |
| `Variant` | 根包共享词表：`Native` / `Secondary` / `Outline` 等 |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（字段绘制）
