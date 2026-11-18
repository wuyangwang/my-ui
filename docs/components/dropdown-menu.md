# DropdownMenu 规格书

> 状态：**已实现**。下拉菜单：点击触发项后在浮层里显示一组操作。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `MenuBase` | 菜单按钮 + `Popup` 里的 `Item`/`CheckItem`/`Separator` 与 `Highlighted()`/`Clicked()` |
| shadcn/ui DropdownMenu | 面板 `Background`/`Border`/`Radius`/阴影，条目 hover 淡底，可带图标 |

## 2. 能否实现
- `ui.MenuBase(c, open)` ✓；条目着色用主题 token，图标用 `IconSizeSm`。

## 3. 实现到何种程度
- `NewDropdownMenu(c, open, trigger, items)`；`MenuItem{Label, Icon, Disabled, Separator, OnClick}`；点击写回 `*open` 并调 `OnClick`。
- 后置：复选/单选条目、子菜单（mygo `CheckItem`/`RadioItem`/`Submenu` 已有，后续加选项）。

## 4. 项目结构
- `overlays/` 子包。

## 5. API 契约
```go
type MenuItem struct {
    Label string
    Icon *ui.SVG
    Disabled, Separator bool
    OnClick func()
}
func NewDropdownMenu(c *ui.Context, open *bool, trigger func() ui.Element, items []MenuItem) ui.Element
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
