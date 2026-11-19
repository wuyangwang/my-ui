# Progress 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`progress.go`）。按 Lobe MyGo 审查修复：0..100 量纲、Type/Variant/Status/ShowInfo/Format、`AnimateTo` + `MotionNormal`、勾/叉。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| Lobe MyGo `feedback/progress.go` | `ProgressType`（Line/Circle）、`ProgressVariant`（Bar/Segments/Inset）、`ProgressStatus`（Normal/Active/Success/Exception）、`ShowInfo`、`Format`、`AnimateTo`、Success/Exception 显示勾/叉 | 全部照做 ✓ |
| [shadcn/ui Progress](https://ui.shadcn.com/docs/components/progress) | 细圆角轨 + 强调色填充 | Bar 形态 ✓ |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 线性/分段/内嵌/圆环 | `Box`/`Row`/`WidthPercent`/`Grow`/`Painter.StrokePath`/`Path.Circle` ✓ | ✅ |
| 数值动画 | `AnimateTo` + `ui.Animation{Duration: MotionNormal}` ✓ | ✅ |
| 勾/叉 | `ui.Icon(icons.Check/X())` ✓ | ✅ |
| 无障碍数值 | `Range(0,100,value)` + `Role(ui.RoleProgress)` ✓ | ✅ |
| 活动态微光 | `Painter.AnimationFrame` + `Now` ✓ | ✅ |

## 3. 实现到何种程度

- **量纲**：构造参数是 0..100 的百分数；为兼容旧调用，`(0,1]` 的值按小数读取（`0.4` = 40%），负数表示不确定时长。传 `40` 画 40%，不再画满。`Value(current,max)` 从一对值计算百分数。
- `Type`：`ProgressLine`（轨 + 填充）或 `ProgressCircle`（圆环 + 弧）。
- `Variant`（仅 Line）：`Bar` 连续填充、`Segments` 离散块、`Inset` 填充 + 剩余两段。
- `Status`：`Normal` 强调色、`Active` info 色 + 微光、`Success` 成功色 + 勾、`Exception` 危险色 + 叉。
- `ShowInfo` 显示百分数；Success/Exception 时改为显示勾/叉。`Format` 自定义百分数文案。
- 填充值通过 `root.AnimateTo("progress-pct", percent, Animation{Duration: MotionNormal})` 渐变；Reduce Motion 下框架直接到目标。
- 环/弧用 `Painter.StrokePath` + `Path.Circle` + 折线弧，无任何计时器。

## 4. 项目结构

- `display/` 子包；只用公开 `ui` API。

## 5. API 契约（逐 prop）

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

## 6. 限制

- `(0,1]` 的小数兼容规则意味着 `1` 被读作 100%，而不是 1%；需要 1% 请用 `Percent(1)`（该方法直接收 0..100，不做小数换算）。
- 活动态微光用绘制时帧号驱动，Reduce Motion 下框架不推进动画帧；颜色仍随主题变化。

## 7. 状态

- [x] 调研与实现（`progress.go`）
- [x] 测试：在屏、量纲（40 / 0.4 / Value）、各形态与状态、`Format`
