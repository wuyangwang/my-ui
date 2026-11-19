# Segmented 规格书

> 位置：`controls/segmented.go`。

> 状态：**已实现**。参考 Lobe UI 的 Segmented：几选一的轨道，指示器滑向选中的那一段；签名兼容 mygo 的 `ui.Segmented`。

## shadcn 风格

轨道是 shadcn Tabs 的弱底发丝边：Filled 用 `t.Surface` + 1px `t.Border`，Outlined
保留窗口底色 + 1px `t.Border`；选中的那一段用 `t.Accent.Alpha(0.10)` 的弱强调底，
其余文字 `t.TextMuted`、选中与悬停转 `t.Text`。圆角用 `Radius(t, RadiusMd/Sm)`，
内边距与间距用 `t.Space`，字号 `TextSize(t, TextXs…TextLg)`，滑行时长走
`MotionNormal`。没有 Lobe 固定色板与自造像素尺寸。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [Lobe UI Segmented](https://mygo.lobehub.com/components/segmented) | Filled/Outlined、Small/Middle/Large、Block、滑动指示器 |
| mygo `ui.SegmentedBase` | 键盘、radio 组角色、`Changed`、段索引回写 |
| `myui/motion.go` | 指示器滑行用 `MotionNormal` |

## 2. 能否实现

- `SegmentedBase(c, selected, n)` 公开可用，`Track` 与 `Segment(i)` 可各自涂色 ✓。
- 指示器是段内绝对定位的 `ui.Box`，用 `ElementTransition{Position: true, Match: id}` 在帧间滑行；`id` 取 `ui.Local(track, …)` 的独有指针，避免不同控件互相借位 ✓。
- 按下时的缩放（Lobe 的 `AnimatePose`）不做：库约定不用 `AnimatePose`，它会让被姿态化的元素不可见 ✓。
- 签名保持 `NewSegmented(c, selected *int, labels ...string)`，与 mygo 原样一致，只是换上 my-ui 外观 ✓。

## 3. 实现到何种程度

两种轨道、四档尺寸、Block 平分宽度、选中/悬停墨色、滑动指示器、radio 组无障碍与 `Changed`。`Label` 保留为 fluent 方法以便命名整组。

## 4. 项目结构

- `controls/segmented.go`；`controls/segmented_test.go` 覆盖索引回写、Changed 提交、角色与预设。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewSegmented(c *ui.Context, selected *int, labels ...string) *Segmented` | 建分段控件，选中的下标写入 `*selected` |
| `Variant(v SegmentedVariant)` · `Filled()` · `Outlined()` | 两种轨道 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 共享尺寸 |
| `Block(on bool)` | 占满整行并平分宽度 |
| `Label(name string)` | 整组无障碍名 |
| `Changed() bool` | 本帧是否换了段 |

## 6. 限制

- 只接受文字标签；图标段可用 `ui.SegmentedBase` 自行构建，本构造器不支持。
- 滑行来自框架 transition，减少动态效果时按框架约定瞬间到位。
