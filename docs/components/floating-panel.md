# FloatingPanel 规格书

> 位置：`layout/` 子包，导入路径 `github.com/wuyangwang/my-ui/layout`。

## API 契约

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

## 状态与限制

**限制**

- 我的 mygo 公开 API 没有"overlay 命中透传"：透明全窗遮罩会吃掉页面点击，与"页面仍可用"冲突。因此点击页面外部关闭不实现；关闭走 Escape 与关闭按钮。
- 缩放状态存在浮层的 `ui.Local` 中，不跨窗口重建保留。
