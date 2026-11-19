# FloatingPanel 规格书

> 位置：`layout/` 子包，导入路径 `github.com/wuyangwang/my-ui/layout`。

> 状态：**已实现**（`layout/floating_panel.go`）。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| shadcn/ui Popover/Card | 四角定位、宽高限制、Esc/关闭按钮、自由边缩放 | 四角定位 ✓；宽高限制 ✓；Esc/关闭 ✓；双向缩放 ✓ |
| mygo | `ui.Overlay`、`Element.OverlayShortcut`、`Element.Dragged` | 浮层、Esc、拖拽都走公开 API ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 浮在窗口之上 | `ui.Overlay` | ✅ |
| Escape 关闭 | `Element.OverlayShortcut` | ✅ |
| 边缘缩放 | `Element.Dragged` | ✅ |
| 点页面关闭且页面仍可用 | 需 overlay 的命中透传 | ❌ 见限制 |

## 3. 实现到何种程度

- `NewFloatingPanel(c, &open, props, build)`：按 `Placement` 定位在窗口角，`Offset` 内缩，宽度默认主题 `Space(90)` 且不超过窗口减一小段边距。
- 卡片面用 shadcn zinc token：`t.Surface`、1px `t.Border`、`Radius(t, RadiusXl)` 与轻阴影；标题/页脚分隔线为 `t.Border` 细线。
- 标题行含标题、`Actions` 与关闭图标；`Footer` 为底部动作行。
- 背向角落的两条边可拖动缩放，范围受 `MinWidth/MaxWidth/MinHeight/MaxHeight` 约束。
- 内置淡入（`OverlayEnter`），颜色随主题每帧更新。

**已知限制**
- 我的 mygo 公开 API 没有"overlay 命中透传"：透明全窗遮罩会吃掉页面点击，与"页面仍可用"冲突。因此点击页面外部关闭不实现；关闭走 Escape 与关闭按钮。
- 缩放状态存在浮层的 `ui.Local` 中，不跨窗口重建保留。

## 4. 项目结构

- 位置：`layout/`，属反馈/浮层件 ✓；只依赖 `mygo/ui`、根包与 `controls/icons` ✓。

## 5. API 契约

| 符号 | 作用 |
|---|---|
| `FloatingPlacement` | 角落枚举：`FloatBottomRight`（默认）/ `FloatBottomLeft` / `FloatTopRight` / `FloatTopLeft` |
| `FloatingPanelProps.Title` | 标题行文字 |
| `FloatingPanelProps.Placement` | 角落；零值为右下 |
| `FloatingPanelProps.Width` | 初始宽度，零为主题 `Space(90)` |
| `FloatingPanelProps.Offset` | 距角落的缩进，零为主题 `Space(4)` |
| `FloatingPanelProps.MinWidth/MaxWidth/MinHeight/MaxHeight` | 缩放边界 |
| `FloatingPanelProps.Actions` | 标题行右侧控件 |
| `FloatingPanelProps.Footer` | 底部动作行 |
| `FloatingPanelProps.NoKeyboard` | 设为 true 后 Escape 不再关闭 |
| `NewFloatingPanel(c, open, props, build)` | 构造；`open` 为 app 的指针，面板只读 |
| `(*FloatingPanel).Open()` | 本帧是否显示 |

## 6. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（打开/关闭/标题/内容）
