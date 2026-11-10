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
一个控件的一段可命名的完整外观定义（变体 Variant、尺寸 Size、形态 Shape 属于预设的三根轴），由一行构造 + 链式方法选择。
_Avoid_：样式类、皮肤

**变体 (Variant)**:
预设中"长什么样"的轴（Default/Secondary/Primary/Outline/Ghost/Destructive/Link/Success/Info/Warning），共享根包词表。零值 `Default` 是控件自带的观感；`Secondary` 是弱填充面；`Native` 是 `Default` 的旧名。Success/Info/Warning 是语义色。
_Avoid_：类型（type 与 Go 类型系统撞词）

**尺寸 (Size)**:
预设中"多大"的轴（SizeMd 零值、SizeSm、SizeLg、SizeXs），走主题的间距与字号刻度。
_Avoid_：宽度（宽度是单个控件的事，不是档位）

**形态 (Shape)**:
预设中"什么形状"的轴（ShapeRect 零值、ShapeRound 胶囊、ShapeCircle 圆形）。
_Avoid_：圆角（圆角是主题的事，形态是控件的事）

**语义色 (Semantic)**:
主题层的一类颜色（success/warning/info），明暗各一套；像 shadcn 的主题变量那样，将来可由 app 自定义。
_Avoid_：状态色（状态是控件的事，语义色是主题的事）

### 弹层

**弹层 (Overlay)**:
浮在窗口之上的一层内容：提示（Tooltip）、浮层（Popover）、对话框（Modal）。库给外观、位置与关闭时机，开合的 flag 是 app 的。
_Avoid_：浮窗、弹出层

**触发器 (Trigger)**:
打开弹层的那个控件（提示的「锚点」、浮层的锚点、对话框的按钮）。库读它的 `Clicked()` 代为打开 flag，所以触发器要建在弹层之前——点击当帧就生效。
_Avoid_：开关按钮（开关是 flag，触发器是控件）

**覆盖层 (Layer)**:
弹层所在的那一层。弹层的内容要建在弹层里（`panel.Children(...)`），否则内容落到窗口角落。
_Avoid_：遮罩（遮罩是背景那一层）
