# ScrollArea 规格书

> 位置：`layout/` 子包（`layout/scroll.go`）。

## API 契约

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

## 状态与限制

**限制**

- 淡出渐变使用主题的 `t.Background` 作终点，内容自带背景时需自行匹配。
- mygo 的滚动条宽度在悬停时不变，只变颜色；与部分控件库"悬停变粗"的描述不同。
