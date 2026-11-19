# Layout 规格书

> 位置：`layout/` 子包（`layout/layout.go`）。

> 状态：**已实现**。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| shadcn/ui 文档版式 | 头栏 + 侧栏 + 正文 + 目录 + 页脚、细线分区、独立滚动 | 全部对应 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 独立滚动区 | `ui.Scroll` + `ScrollbarStyle` | ✅ |
| 分区细线 | `BorderWidth` / `BorderColor` | ✅ |
| 移动端隐藏侧栏 | `Mobile` 选项 | ✅ |

## 3. 实现到何种程度

- `Header` 置顶；侧栏在带自己滚动条的列中；正文、`Toc` 与 `Footer` 一起在头栏下滚动。
- `Mobile` 去掉侧栏并把 `Toc` 移到正文上方。
- `TocWidth` 默认 `Space(50)`；所有表面与细线来自 shadcn 风格主题 token（`t.Background` / `t.Surface` / `t.Border`）。

**已知限制**
- 目录不跟踪滚动位置（没有"当前章节"高亮）：mygo 未提供公开的段落可见性事件，应用可在自己的 `Toc` 槽位里实现。
- 侧栏宽度固定为 `Space(72)`；需要可拖动侧栏时请组合 `DraggablePanel`。

## 4. API 契约

| 符号 | 作用 |
|---|---|
| `LayoutProps.Header / Sidebar / Toc / Footer` | 版式槽位 |
| `LayoutProps.TocWidth` | 目录列宽，默认 `Space(50)` |
| `LayoutProps.Mobile` | 隐藏侧栏、把目录上移 |
| `NewLayout(c, props, build)` | 构造；`build` 画正文 |

## 5. 状态

- [x] 调研与实现
- [x] `ui.Tester` 测试（各槽位绘制）
