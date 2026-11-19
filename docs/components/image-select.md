# ImageSelect 规格书

> 位置：`forms/` 子包（`forms/image_select.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| Lobe `choice.ImageSelect` | 图片单选、方向键选择、选中高亮、任意可比较值 | 全部对应 ✓ |
| shadcn/ui 卡片选择 | 选中圈、悬停加深 | ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 图片 | `ui.Image` + `ImageSource` | ✅ |
| 单选状态指针 | 泛型 `*T` | ✅ |
| 方向键 | `Element.Shortcut` | ✅ |
| 悬停/选中色 | `c.Theme()` + `StateTransition` | ✅ |

## 3. 实现到何种程度

- 泛型 `ImageSelect[T comparable]`：`selected *T` 由 app 持有；nil 时用 `ui.Local` 保存选中值。
- 点选写回指针；方向键在聚焦时移动选择。选中卡的圈为主题强调色，悬停/按下加深为文字色。
- `ImageFn` 支持惰性生成图片；行会在放不下时换行。

**已知限制**
- 图片比例固定为 16:9（`AspectRatio`）；不同比例的预览需 app 自己裁图。
- 键盘焦点由每张卡自己持有（多 tab  stop），不是单一复合控件的 roving tabindex。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `ImageSelectItem[T].Value / Label / Icon / Image / ImageFn` | 选项的值、文字、图标与图片 |
| `ImageSelectProps[T].Options` | 选项列表 |
| `ImageSelectProps[T].MinTileWidth` | 卡片最小宽度，默认 `Space(40)` |
| `NewImageSelect[T](c, selected, props)` | 构造 |
| `(*ImageSelect[T]).Value()` | 当前选中值 |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（点选写回指针）
