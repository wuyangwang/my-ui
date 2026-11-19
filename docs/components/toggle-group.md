# ToggleGroup 规格书

> 位置：`controls/togglegroup.go`。

> 状态：**已实现**。参考 Lobe UI 的 ToggleGroup：Outlined / Borderless / Segmented 三种分组方式；签名兼容 mygo 的 `ui.ToggleGroup`。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [Lobe UI ToggleGroup](https://mygo.lobehub.com/components/toggle-group) | 三种分组外观、单项状态、键盘 |
| mygo `ui.ToggleGroup` | 一个 Tab 停、方向键在组内移动、组角色 |
| `controls/toggle_color.go` | 单个 Toggle 的 my-ui 外观 |

## 2. 能否实现

- `ui.ToggleGroup(c, fn)` 公开可用，`fn` 里的开关各自持有状态 ✓。
- 容器用 `Padding`/`Gap`/`Background`/`Border`/`Radius` 每帧从 `c.Theme()` 重涂，经 `StateTransition()` 淡入 ✓。
- 签名保持 `NewToggleGroup(c, fn func())`，与 mygo 原样一致 ✓。
- Lobe 的 `Items`/`Multiple`/`selected *[]string` 模型不在本构造器内：mygo 的组按各开关自己的 `*bool` 工作，多项同时按下由各开关的指针决定 ✓。

## 3. 实现到何种程度

三种分组外观、共享尺寸词表、`Label` 命名整组、沿用 mygo 的键盘与无障碍。`fn` 可放 `controls.NewToggle` 或 `ui.ToggleBase`。

## 4. 项目结构

- `controls/togglegroup.go`；`controls/togglegroup_test.go` 覆盖组内开关各自绑定、角色与预设。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewToggleGroup(c *ui.Context, fn func()) *ToggleGroup` | 建组，`fn` 构建其中的开关 |
| `Variant(v ToggleGroupVariant)` · `Outlined()` · `Borderless()` · `Segmented()` | 三种分组外观 |
| `Size(s Size)` | 共享尺寸（组内开关仍由自身构造器定尺寸） |
| `Label(name string)` | 整组无障碍名 |

## 6. 限制

- 不接受 Lobe 式的 items 列表与 `Multiple` 开关；要那套模型应自行以 `ui.ToggleBase` 组织。
- 本组件只管容器外观，组内开关的样式由 `fn` 里的控件负责。
