# HoverCard 规格书

> 状态：**已实现**。悬停卡片：指针停在触发元素上时在浮层里显示详情。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui HoverCard | trigger + content，浮层由 Radix 定位 |
| Element Plus Popover `trigger="hover"` | 悬停打开，指针移到面板上保持打开 |

## 2. 能否实现
- mygo `ui.PopoverBase(c, anchor, open, fn)` 提供浮层定位与层级 ✓；组件只负责 hover 判定和面板外观。
- 组件不带计时器：用帧计数给指针从触发区移到面板留几帧宽限（Reduce Motion 无关）。

## 3. 实现到何种程度
- `NewHoverCard(c, open *bool, anchor func() ui.Element, build func(panel ui.Element))`；面板主题底色、边框、圆角、阴影。
- 触发区或面板 `Hovered()` 时保持打开；两者都离开后数帧再关闭。

## 4. 项目结构
- `overlays/` 子包。

## 5. API 契约
```go
func NewHoverCard(c *ui.Context, open *bool, anchor func() ui.Element, build func(panel ui.Element)) *HoverCard
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
