# Burger 规格书

> 位置：`layout/` 子包（`layout/burger.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `BurgerItem.Key / Label / Icon / Divider` | 菜单行或分隔线 |
| `BurgerProps.Items` | 行列表 |
| `BurgerProps.Selected` | 选中键的 app 指针 |
| `BurgerProps.OnSelect` | 选中回调 |
| `BurgerProps.Fullscreen` | 切换全屏菜单 |
| `BurgerProps.HeaderHeight` | 下拉面板的额外顶部间距 |
| `BurgerProps.Footer` | 菜单页脚 |
| `NewBurger(c, open, props)` | 构造；`open` 为 app 的指针 |

## 状态与限制

**限制**

- 这是**桌面版菜单按钮**，不是移动端原生控件：mygo v0.10.0 的 `ui/mobile` 只有 TopBar/Drawer/Sheet/ActionSheet/ListRow 等，没有 burger。`Fullscreen` 用模态 `DialogBase`，不是从顶栏滑下的原生抽屉。
- 下拉面板由 `PopoverBase` 定位在按钮下方；`HeaderHeight` 只是追加一段上边距，不做真正的抽屉滑入动画。
