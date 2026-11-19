# Steps 规格书

> 状态：**已实现**。流程步骤：编号圆圈或圆点由连线串起，当前步与已完成步用强调色区分。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui 步骤条 | `StepItem`/`StepsProps`、`Current` 指针语义、Dot/Vertical 两形态 |

## 2. 能否实现
- mygo `ui.Box/Row/Column` + 主题色即可画出节点与连线 ✓。

## 3. 实现到何种程度
- `NewSteps(c, StepsProps)`；`.Vertical(on)`、`.Dot(on)`；每步可带 `Icon` 与 `Status` 覆盖。
- 颜色与尺寸收敛到 shadcn token：完成/选中面为 `t.Accent.Alpha(0.10)` 配 `t.Accent`，连线与未达面用 `t.Border` / `t.Surface`，节点圆角 `Radius(t, RadiusMd)`，图标与编号字号用 `IconSizeXs` / `TextSize(t, TextXs)`。

## 4. 项目结构
- `layout/` 子包。

## 5. API 契约

| 方法 / 字段 | 作用 |
|---|---|
| `NewSteps(c, p)` | 构建步骤条，`p.Current` 是 app 持有的当前步索引 |
| `StepsProps.Items` | 步骤列表（`Title/Description/Icon/Status`） |
| `StepsProps.Current` | 当前步；之前的为完成，当前为进行中 |
| `StepsProps.Vertical` | 竖排 |
| `StepsProps.Dot` | 用圆点代替编号 |
| `(*Steps).Vertical(on)` | 切换竖排 |
| `(*Steps).Dot(on)` | 切换圆点 |

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
