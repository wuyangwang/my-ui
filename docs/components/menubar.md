# Menubar 规格书

> 状态：**已实现**。应用内菜单栏：横排多个下拉菜单，每个共用 DropdownMenu 的外观。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Menubar | 横排菜单按钮，面板风格与 DropdownMenu 一致 |
| mygo/ui `MenuBase` | 每个菜单一个 trigger + Popup，item 的 Clicked 即选择 |

## 2. 能否实现
- 多个 `ui.MenuBase` + 现有 `menu_common.go` 的面板/条目构建器 ✓。

## 3. 实现到何种程度
- `NewMenubar(c, menus []MenubarMenu)`；`MenubarMenu{Label, Open *bool, Items []MenuItem}`；Open 可为 nil（库用局部状态）。
- 沿用 DropdownMenu 的 shadcn 面板；后置：键盘左右方向键在菜单间移动。

## 4. 项目结构
- `overlays/` 子包。

## 5. API 契约
```go
type MenubarMenu struct {
    Label string
    Open *bool
    Items []MenuItem
}
func NewMenubar(c *ui.Context, menus []MenubarMenu) *Menubar
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
