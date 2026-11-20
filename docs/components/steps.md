# Steps 规格书

流程步骤：编号圆圈或圆点由连线串起，当前步与已完成步用强调色区分。

## API 契约

| 方法 / 字段 | 作用 |
|---|---|
| `NewSteps(c, p)` | 构建步骤条，`p.Current` 是 app 持有的当前步索引 |
| `StepsProps.Items` | 步骤列表（`Title/Description/Icon/Status`） |
| `StepsProps.Current` | 当前步；之前的为完成，当前为进行中 |
| `StepsProps.Vertical` | 竖排 |
| `StepsProps.Dot` | 用圆点代替编号 |
| `(*Steps).Vertical(on)` | 切换竖排 |
| `(*Steps).Dot(on)` | 切换圆点 |
