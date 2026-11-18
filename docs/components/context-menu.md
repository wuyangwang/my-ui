# ContextMenu 规格书

> 状态：**已实现**。右键菜单：给目标元素附加一组操作。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `ContextMenuBase` | `target` + `open` + `Popup` 构建菜单，行为与 `MenuBase` 一致 |
| shadcn/ui ContextMenu | 与 DropdownMenu 相同的面板与条目录 |

## 2. 能否实现
- `ui.ContextMenuBase(c, target, open)` ✓；与 DropdownMenu 共用条目构建器。

## 3. 实现到何种程度
- `NewContextMenu(c, open, target, items) ui.Element`；条目选项与 DropdownMenu 相同。

## 4. 项目结构
- `overlays/` 子包。

## 5. API 契约
```go
func NewContextMenu(c *ui.Context, open *bool, target ui.Element, items []MenuItem) ui.Element
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
