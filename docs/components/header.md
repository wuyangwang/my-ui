# Header 规格书

> 位置：`layout/` 子包（`layout/header.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| shadcn/ui 顶栏 | Logo/Nav/Actions 三段、Mobile 排列、底部细线与留白 | 全部对应 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 横排三段 | `ui.Row` + `Grow` + `Gap` | ✅ |
| 底部细线 | `BorderWidth` / `BorderColor` | ✅ |
| 吸顶 | `Element.Sticky` | ✅ |

## 3. 实现到何种程度

- 默认高 `Space(16)`（主题 spacing token）；`Mobile` 时 Nav 在最前作菜单位，Logo 居中占余宽。
- 三个槽位都是 builder，点击与状态由 app 管理；`Height` 与 `Sticky` 可链式调用。
- 背景、描边与间距均取 `c.Theme()`，状态变化走 `StateTransition`。

**已知限制**
- 不监听滚动位置：吸顶由 mygo 的 `Sticky` 处理，Header 自身不判断"已离开顶部"。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `HeaderProps.Logo` | 品牌槽位 |
| `HeaderProps.Nav` | 导航槽位；Mobile 时为菜单位 |
| `HeaderProps.Actions` | 尾部操作槽位 |
| `HeaderProps.Mobile` | 切换紧凑排列 |
| `NewHeader(c, props)` | 构造 |
| `(*Header).Height(v)` | 设置高度（DIPs） |
| `(*Header).Sticky()` | 吸顶 |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（三槽位、Mobile 顺序）
