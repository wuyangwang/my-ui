# ActionIcon 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题与 tokens 仍在根包 `myui`。

> 状态：**已实现**（`controls/actionicon.go`）。shadcn 式中「只有图标的按钮」，参考 Lobe UI 的 ActionIcon 三种面与尺寸档。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [Lobe UI ActionIcon](https://mygo.lobehub.com/components/action-icon) | Borderless / Filled / Outlined 三种面，Small / Middle / Large 三档，Active 与 Danger |
| [shadcn/ui Button](https://ui.shadcn.com/docs/components/button) | ghost / secondary / outline 的着色方式，状态只改观感 |
| `controls/button.go` | 共享 `Variant`/`Size`/`Shape`，每帧从 `c.Theme()` 重涂 |

## 2. 能否实现

- `ui.Box` + `Role(ui.RoleButton)` + `TabStop(true)` 得到可聚焦的按钮角色 ✓。
- `Hovered()`/`Pressed()`/`IsDisabled()` 每帧可读，颜色经 `Transition(StateTransition())` 淡入 ✓。
- 图标走根包 `IconSizeXs/Sm/Md/Lg`，圆角走 `Radius(t, RadiusMd)` ✓。
- 危险色走 `Shades(t.Danger, t.Dark)`，与库内其他危险态同源 ✓。
- 提示气泡（Lobe 的 Title 兼 tooltip）不放进本组件：`controls` 不能导入 `overlays`（反向依赖会成环），由应用自行用 `overlays.NewTooltip` 包裹；组件只提供无障碍名 ✓。

## 3. 实现到何种程度

三种面、四档尺寸、圆形形态、Active 常驻悬停底、Danger 危险色、`Clicked()` 沿用 ui。禁用时保持静止底色，由主题暗淡。

## 4. 项目结构

- `controls/actionicon.go`：组件本体；`controls/actionicon_test.go`：名字、尺寸、点击与随主题换色。
- 图标按钮的语义由 `controls.NewIconAction` 承载；本组件在其之上增加「面」。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewActionIcon(c *ui.Context, svg *ui.SVG, title string) *ActionIcon` | 建图标按钮，`title` 为无障碍名 |
| `Variant(v ActionIconVariant)` · `Borderless()` · `Filled()` · `Outlined()` | 三种面 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 共享尺寸（28/32/36/40 DIP 方框） |
| `Shape(s Shape)` | 圆角或圆形 |
| `Active(on bool)` | 保持悬停底，用于已切换或已展开状态 |
| `Danger(on bool)` | 图标用危险色 |
| `Label(s string)` | 无障碍名 |
| `Clicked()` | 沿用 `ui.Element` 的点击事件 |

## 6. 限制

- 不自带 tooltip；需要提示时用 `overlays.NewTooltip(c, icon.Element, "…")`。
- 尺寸档决定方框与字形，字形取共享图标尺寸，不提供任意 DIP 字形。
