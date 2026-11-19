# ActionIconGroup 规格书

> 位置：`controls/actionicongroup.go`。

> 状态：**已实现**。参考 Lobe UI 的 ActionIconGroup：一组共用底面的图标动作，返回被点中的 Key。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [Lobe UI ActionIconGroup](https://mygo.lobehub.com/components/action-icon-group) | Items/Key、Filled/Outlined/Borderless、Vertical、Disabled |
| `controls/iconaction.go` | 纯图标按钮的共享命中区与悬停/按下墨色 |
| `controls/actionicon.go` | 三种面的取色规则 |

## 2. 能否实现

- 每项的图标用 `NewIconAction` 建，禁用项 `Disabled(true)` 且点击回调内再判一次 ✓。
- 组的底面是 `ui.Row`/`ui.Column`，每帧从 `c.Theme()` 取色并经 `StateTransition()` 淡入 ✓。
- 点击写入 app 持有的 `*string`，`Chosen()` 在同一帧可读，符合 ui 的绑定约定 ✓。
- Lobe 的 More 菜单需要浮层，`controls` 不能导入 `overlays`；本组件不内置菜单，由应用自行在组后追加 `overlays.NewDropdownMenu` ✓。

## 3. 实现到何种程度

三种面、横排/竖排、每项禁用与整组禁用、点击 Key 回写。没有阴影预设（库统一用页面卡片的阴影），没有内置 More 菜单。

## 4. 项目结构

- `controls/actionicongroup.go`；`controls/actionicongroup_test.go` 覆盖名字、点击回写与禁用项不改值。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewActionIconGroup(c *ui.Context, chosen *string, items []ActionIconGroupItem) *ActionIconGroup` | 建组，点击的 Key 写入 `chosen` |
| `ActionIconGroupItem{Key, Icon, Label, Danger, Disabled}` | 一项的键、图标、名字、危险色、禁用 |
| `Variant(v ActionIconGroupVariant)` · `Filled()` · `Outlined()` · `Borderless()` | 三种面 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 图标的共享尺寸 |
| `Vertical(on bool)` | 竖排 |
| `Disabled(on bool)` | 整组禁用 |
| `Chosen() string` | 本帧点中的 Key，未点中为 `""` |

## 6. 限制

- 不内置 More 菜单与阴影预设；需要时用 `overlays` 的菜单组件组合。
- `Chosen()` 只反映当前帧，跨帧读取应看传入的 `chosen` 指针。
