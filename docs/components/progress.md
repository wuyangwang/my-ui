# Progress 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

提供 0..100 量纲、Type/Variant/Status/ShowInfo/Format、`AnimateTo` + `MotionNormal`、勾/叉。

## API 契约（逐 prop）

### 类型

| 符号 | 作用 |
|---|---|
| `type ProgressType uint8` | 形状 |
| `ProgressLine` / `ProgressCircle` | 取值，零值为 Line |
| `type ProgressVariant uint8` | 线的画法 |
| `ProgressBar` / `ProgressSegments` / `ProgressInset` | 取值，零值为 Bar |
| `type ProgressStatus uint8` | 状态 |
| `ProgressNormal` / `ProgressActive` / `ProgressSuccess` / `ProgressException` | 取值，零值为 Normal |

### 构造与方法

| 方法 | 作用 |
|---|---|
| `NewProgress(c, percent float64) *Progress` | 创建；0..100（`(0,1]` 按小数），负数不确定 |
| `Type(v ProgressType) *Progress` | Line 或 Circle |
| `Variant(v ProgressVariant) *Progress` | Bar / Segments / Inset |
| `Status(s ProgressStatus) *Progress` | 状态色 + 勾/叉 |
| `ShowInfo(on bool) *Progress` | 显示百分数或状态图标 |
| `Format(fn func(float32) string) *Progress` | 自定义百分数文案；nil 恢复 `40%` |
| `Color(c ui.Color) *Progress` | 覆盖填充色；零值清除 |
| `Size(s Size) *Progress` | 线厚度 / 圆环直径（Xs/Sm/Md/Lg） |
| `Segments(n int) *Progress` | Segments 的块数；≤0 恢复 20 |
| `Percent(v float32) *Progress` | 直接设置 0..100 |
| `Value(current, max float64) *Progress` | 由当前值/最大值计算 |
| `Indeterminate() *Progress` | 不确定时长活动条 |
| `Height(h float32) *Progress` | 线的厚度（DIP） |
| `Label(label string) *Progress` | 无障碍名称 |

用法：

```go
display.NewProgress(c, 40).ShowInfo(true).Status(ProgressSuccess)
display.NewProgress(c, 0).Type(ProgressCircle).Percent(65)
```

## 限制

- `(0,1]` 的小数兼容规则意味着 `1` 被读作 100%，而不是 1%；需要 1% 请用 `Percent(1)`（该方法直接收 0..100，不做小数换算）。
- 活动态微光用绘制时帧号驱动，Reduce Motion 下框架不推进动画帧；颜色仍随主题变化。
