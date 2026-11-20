# Layout 规格书

> 位置：`layout/` 子包（`layout/layout.go`）。

## API 契约

| 符号 | 作用 |
|---|---|
| `LayoutProps.Header / Sidebar / Toc / Footer` | 版式槽位 |
| `LayoutProps.TocWidth` | 目录列宽，默认 `Space(50)` |
| `LayoutProps.Mobile` | 隐藏侧栏、把目录上移 |
| `NewLayout(c, props, build)` | 构造；`build` 画正文 |

## 状态与限制

**限制**

- 目录不跟踪滚动位置（没有"当前章节"高亮）：mygo 未提供公开的段落可见性事件，应用可在自己的 `Toc` 槽位里实现。
- 侧栏宽度固定为 `Space(72)`；需要可拖动侧栏时请组合 `DraggablePanel`。
