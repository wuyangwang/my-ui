# Header 规格书

> 位置：`layout/` 子包（`layout/header.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `HeaderProps.Logo` | 品牌槽位 |
| `HeaderProps.Nav` | 导航槽位；Mobile 时为菜单位 |
| `HeaderProps.Actions` | 尾部操作槽位 |
| `HeaderProps.Mobile` | 切换紧凑排列 |
| `NewHeader(c, props)` | 构造 |
| `(*Header).Height(v)` | 设置高度（DIPs） |
| `(*Header).Sticky()` | 吸顶 |

## 状态与限制

**限制**

- 不监听滚动位置：吸顶由 mygo 的 `Sticky` 处理，Header 自身不判断"已离开顶部"。
