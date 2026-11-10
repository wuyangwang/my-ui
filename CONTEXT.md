# my-ui

MyGo 原生 UI(`package ui`)之上的组件库：设计系统与预设包装，不重造控件。

## Language

### 设计系统

**刻度 (Scale)**:
一组带名字的倍率（TextSm/SpacingLg/RadiusMd…），全部从构建时的 `*ui.Theme` 解出实际值。
_Avoid_：token 字面量、常量表（泛称）

**色板 (Palette)**:
一套强调色的"主色"声明。只声明基色；三态中的 hover/pressed 是派生的。
_Avoid_：主题（theme 是语义槽全集）、配色

**取色 (Shades)**:
从"一个基色 + 明暗"得出完整一套语义色（accent/hover/pressed/文字底色）的唯一入口。
派生规则只住在这里。
_Avoid_：调色盘、颜色工具

**外观 (Appearance)**:
浅色 / 深色 / 跟随系统，三者之一。
_Avoid_：夜间模式、暗黑模式

### 控件

**预设 (Preset)**:
一个控件的一段可命名的完整外观定义（变体 Variant、尺寸 Size 属于预设的两根轴），由一行构造 + 链式方法选择。
_Avoid_：样式类、皮肤

**变体 (Variant)**:
预设中"长什么样"的轴（Native/Primary/Outline/Ghost/Destructive/Link），共享根包词表。零值 `Native` 不涂预设，保留底层控件原生观感；`Secondary` 为 `Native` 的 shadcn 别名。它只决定颜色三态与边框；尺寸是另一根轴。
_Avoid_：类型（type 与 Go 类型系统撞词）
