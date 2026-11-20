# Skeleton 规格书

提供 `Animation`（Fade/Sweep/Static）、末行宽度改为 66% 跟随容器。

## API 契约（逐 prop）

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

## 限制

- 默认动画是 Fade（与旧版静态不同）；需要静止请显式 `Static()`。
- Sweep 的高亮色取 `t.SurfaceHover`，与条形底 `t.Surface` 形成轻微对比。
