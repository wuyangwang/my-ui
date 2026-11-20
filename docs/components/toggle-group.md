# ToggleGroup 规格书

> 位置：`controls/togglegroup.go`。

参考 Lobe UI 的 ToggleGroup：Outlined / Borderless / Segmented 三种分组方式；签名兼容 mygo 的 `ui.ToggleGroup`。

## shadcn 风格

容器是 shadcn ToggleGroup 的弱底发丝边：Outlined 用 1px `t.Border` 且裁切共享边，
Borderless 无面，Segmented 用 `t.Surface` + 1px `t.Border` 的轨道。内边距与间距用
`t.Space`，圆角 `Radius(t, RadiusMd)`，状态变化经 `StateTransition()` 淡入。组内开关的
选中/悬停墨色由各自的控件（如 `NewToggle`）按同一套主题槽位绘制。没有 Lobe 固定色板
与自造尺寸。

## API 契约

| 方法 | 作用 |
|---|---|
| `NewToggleGroup(c *ui.Context, fn func()) *ToggleGroup` | 建组，`fn` 构建其中的开关 |
| `Variant(v ToggleGroupVariant)` · `Outlined()` · `Borderless()` · `Segmented()` | 三种分组外观 |
| `Size(s Size)` | 共享尺寸（组内开关仍由自身构造器定尺寸） |
| `Label(name string)` | 整组无障碍名 |

## 限制

- 不接受 Lobe 式的 items 列表与 `Multiple` 开关；要那套模型应自行以 `ui.ToggleBase` 组织。
- 本组件只管容器外观，组内开关的样式由 `fn` 里的控件负责。
