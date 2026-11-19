# ScrollArea 规格书

> 位置：`layout/` 子包（`layout/scroll.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| shadcn/ui `ScrollArea` | 细滚动条、可滚方向淡出、横竖两向 | 全部对应 ✓ |
| mygo | `ui.Scroll` / `ScrollHorizontal` / `ScrollBoth`、`ScrollbarStyle`、`ScrollState` | 公开 API 足够 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 滚动容器 | `ui.Scroll` 系列 | ✅ |
| 细滚动条 | `Element.ScrollbarStyle` | ✅ |
| 判定边缘淡出 | `ScrollState.X/Y/MaxX/MaxY` | ✅ |
| 指针移入才显示滚动条 | 框架内建 | ✅ |

## 3. 实现到何种程度

- `ShadcnScrollbar(t)` 给出细条（thumb 6 DIP、track 12 DIP、圆角 3），颜色留零由主题每帧填充；命名为 shadcn 术语，不再是某个库的专用名。边缘淡出用当帧的 `t.Background` 派生。
- `NewScroll`/`NewScrollHorizontal` 保留 mygo 签名，仅套用细条；`ScrollArea` 在可滚方向画渐变淡出。
- `Horizontal`/`Both` 切换方向，`Fade` 开关淡出，`Width`/`Height` 定尺。

**已知限制**
- 淡出渐变使用主题的 `t.Background` 作终点，内容自带背景时需自行匹配。
- mygo 的滚动条宽度在悬停时不变，只变颜色；与部分控件库"悬停变粗"的描述不同。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `ShadcnScrollbar(t)` | 返回主题派生的 shadcn 风格细滚动条样式 |
| `LobeScrollbar(t)`（deprecated） | 旧名，等同 `ShadcnScrollbar`，仅为兼容保留 |
| `NewScroll(c)` | 带细条的纵向滚动（兼容 `ui.Scroll` 签名） |
| `NewScrollHorizontal(c)` | 横向版本 |
| `NewScrollPanel(c, state)` | 结构化滚动容器 |
| `NewScrollArea(c, state, build)` | 带边缘淡出的滚动区 |
| `(*ScrollArea).Horizontal() / Both() / Fade(on) / Width(w) / Height(h)` | 方向与尺寸选项 |
| `AutoGrid` 也复用 `ShadcnScrollbar` | 见 grid.md |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（内容绘制）
