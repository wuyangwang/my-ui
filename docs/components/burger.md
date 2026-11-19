# Burger 规格书

> 位置：`layout/` 子包（`layout/burger.go`）。

> 状态：**已实现（桌面替代方案）**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| shadcn/ui 导航菜单 | 顶栏菜单按钮、下拉抽屉、全屏菜单、选中行、页脚 | 对应 ✓ |
| mygo `ui/mobile` | TopBar / Drawer / Sheet / ActionSheet 等 | 可用，但**没有 burger/menu button** |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 图标按钮 | `controls.NewIconAction` | ✅ |
| 下拉菜单面板 | `ui.PopoverBase`，按 AGENTS.md 用 `panel.Children` | ✅ |
| 全屏菜单 | `ui.DialogBase` | ✅ |
| 移动端原生 burger | `ui/mobile` 无此控件 | ❌ 见限制 |

## 3. 实现到何种程度

- 触发器先建、面板后建：点击触发器与打开面板落在同一帧；`open` 是 app 的指针，组件只读。
- 下拉模式用 `PopoverBase`，点击外部/Escape 自动关闭；全屏模式用 `DialogBase`，自带大号行与关闭图标。
- 面板按 shadcn Card/Popover token 绘制：`t.Background` / `t.Border`（1px）、`Radius(t, RadiusLg)` 与轻阴影；悬停行用 `t.SurfaceHover`，选中行用 `t.Accent.Alpha(0.10)` 配 `t.Accent`。
- `Selected *string` 指针与 `OnSelect` 回调任选；点击行写回选中键并关闭。

**已知限制**
- 这是**桌面版菜单按钮**，不是移动端原生控件：mygo v0.10.0 的 `ui/mobile` 只有 TopBar/Drawer/Sheet/ActionSheet/ListRow 等，没有 burger。`Fullscreen` 用模态 `DialogBase`，不是从顶栏滑下的原生抽屉。
- 下拉面板由 `PopoverBase` 定位在按钮下方；`HeaderHeight` 只是追加一段上边距，不做真正的抽屉滑入动画。

## 4. API 契约

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

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（打开、选行、写回指针）
