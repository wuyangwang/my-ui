# Skeleton 规格书

> 状态：**已实现**。按 Lobe MyGo 审查修复：`Animation`（Fade/Sweep/Static）、末行宽度改为 66% 跟随容器。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| Lobe MyGo `display/skeleton.go` | `SkeletonAnimation`（Fade/Sweep/Static）；末行 `WidthPercent(66)`；Fade 用 `Loop` + `Bounce(EaseInOut)`，Sweep 用 `Loop` + `LinearGradient` + `TranslatePercent` |
| shadcn/ui Skeleton | 用 `Surface` 色块替代文字与图片 |
| Ant Design Skeleton | 行数、头像圆点 |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 脉冲/微光 | `Element.Loop` + `Opacity` / `LinearGradient` + `TranslatePercent` ✓ | ✅ |
| 末行 66% | `WidthPercent(66)` ✓ | ✅ |
| Reduce Motion | 框架兜底，组件无计时器 ✓ | ✅ |

## 3. 实现到何种程度

- `NewSkeleton(c, lines)` 默认 3 行；`Circle()` 头像圆；`Width(w)` 固定宽度（默认占满）；`Size(sz)` 条高（10/12/16/20）。
- `Animation(a)`：`Fade` 用 `1 - 0.5*Loop("skeleton-fade",1600ms,Bounce(EaseInOut))`；`Sweep` 用一条渐变色块按 `TranslatePercent(200t-100,0)` 扫过；`Static` 不动。
- 多行时最后一行是 `WidthPercent(66)`，随容器缩放；不再用定长 `t.Space(24)`。
- 所有动效走 `Loop`，无计时器/goroutine。

## 4. 项目结构

- `display/` 子包，展示件分组。

## 5. API 契约（逐 prop）

### 类型

| 符号 | 作用 |
|---|---|
| `type SkeletonAnimation uint8` | 占位动画 |
| `SkeletonFade` / `SkeletonSweep` / `SkeletonStatic` | 三个取值，零值为 `SkeletonFade` |

### 构造与方法

| 方法 | 作用 |
|---|---|
| `NewSkeleton(c, lines int) *Skeleton` | 创建；`lines<1` 表示 3 行 |
| `Animation(a SkeletonAnimation) *Skeleton` | 设置动画 |
| `Fade()` / `Sweep()` / `Static()` | `Animation(...)` 的快捷方法 |
| `Circle() *Skeleton` | 头像圆块 |
| `Width(w float32) *Skeleton` | 固定宽度；0 表示占满父容器 |
| `Size(sz Size) *Skeleton` | 条高：Xs=10、Sm=12、Md=16、Lg=20 |

## 6. 限制

- 默认动画是 Fade（与旧版静态不同）；需要静止请显式 `Static()`。
- Sweep 的高亮色取 `t.SurfaceHover`，与条形底 `t.Surface` 形成轻微对比。

## 7. 状态

- [x] 规格书 / 实现 / 测试（占位绘制、Circle/默认值、三种动画）
