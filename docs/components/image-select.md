# ImageSelect 规格书

> 位置：`forms/` 子包（`forms/image_select.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `ImageSelectItem[T].Value / Label / Icon / Image / ImageFn` | 选项的值、文字、图标与图片 |
| `ImageSelectProps[T].Options` | 选项列表 |
| `ImageSelectProps[T].MinTileWidth` | 卡片最小宽度，默认 `Space(40)` |
| `NewImageSelect[T](c, selected, props)` | 构造 |
| `(*ImageSelect[T]).Value()` | 当前选中值 |

## 状态与限制

**限制**

- 图片比例固定为 16:9（`AspectRatio`）；不同比例的预览需 app 自己裁图。
- 键盘焦点由每张卡自己持有（多 tab  stop），不是单一复合控件的 roving tabindex。
