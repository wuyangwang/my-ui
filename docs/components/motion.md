# 动效规格书

> 状态：**已实现第一轮**（状态色渐变、聚焦描边、出现即渐显）。跨组件的共用约定，不是一个组件。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| mygo 的 `ElementTransition`（`ui/transition.go:21-51`） | CSS 式的过渡：`Duration`/`Ease`/`Animation`（弹簧）+ `Position`/`Size`/`Colors`/`Shadows` 选哪些属性动，`ExitDuration`/`ExitEase` 管退场 |
| mygo 的 `AnimatePose` + `Pose`（`ui/pose.go:11-37`） | 合成器级动作（位移/缩放/旋转/透明度），只重绘不重跑视图——最廉价的动效；`Fade` 即渐隐渐显 |
| mygo 的 `AnimateColor`（`ui/api_gen.go:3487-3493`） | 颜色插值（含 alpha），随透明度混合保持色相 |
| mygo 的 `ui.Local[T]`（`ui/api_gen.go:1666-1682`） | 元素自己留一份跨帧状态（出现/消失的记号） |
| mygo 的弹簧 `ui.Snappy`/`Smooth`/`Bouncy`（`ui/motion.go:49`） | 弹簧参数（时长 + 回弹），比曲线更有物理感 |
| [shadcn/ui 的 CSS transition](https://ui.shadcn.com/docs/components/button) | 状态色的 `transition-colors`；按钮的 hover 渐变是默认观感的一部分 |
| [Element Plus](https://element-plus.org/zh-CN/component/button) | 按钮/输入的状态过渡时长约 0.2–0.3s；加载态的转圈 |
| [Lobe MyGo Motion](https://mygo.lobehub.com/skills/components/motion) | 同一套 token 化的动效：时长、弹簧、弹层/提示/对话框的进出曲线 |

要点：Reduce Motion 由框架兜底——`Animate`/`AnimateTo` 直接到目标，`AnimateColor` 仍平滑换色但 `DisableAnimations` 时立刻换；**库里不写任何定时器**。

## 2. 能否实现

| 需要 | 公开 API | 结论 |
|---|---|---|
| 状态色渐变 | `Element.Transition(ui.ElementTransition{…})`，`Colors: true` 只动颜色 | ✅ 按钮 hover/pressed、输入聚焦描边、徽标变体 |
| 出现即渐显 | `Element.AnimatePose(ui.Pose{Fade: 1}, ui.Snappy)` | ❌ **不可用**：实测该姿态的 Fade 通道会让元素停在不可见状态（元素在无障碍树里、也在版面上，但一个像素也没画）。故本库不用姿态动画 |
| 退场渐隐 | 元素离开父节点时框架会走 `Transition` 的 Exit；但 Exit 只覆盖位置/尺寸/颜色/阴影，**不含透明度** | ⚠️ 退场渐隐需自行保留元素并把 `Fade` 动画到 0（占位问题），本轮不做 |
| 数值动画 | `Animate`/`AnimateTo`（含弹簧、延迟） | ✅ 后续组件（如分段控件的指示条） |
| 逐帧状态 | `ui.Local[T](e, key, init)` | ✅ 需要「出现/消失」记号时 |

> 证据：`AnimatePose` 的问题在 `ui.Tester` 与真实窗口里都复现（`ui/pose.go:92-116` 的 `samplePose`：姿态的通道在元素不是「上一帧摆过」时直接取目标值，而摆过之后由通道推进——推进未发生时元素停在 Fade=0，即不可见）。本库因此只用**颜色过渡**，不用姿态动画。

## 3. 实现到何种程度

**共用约定（`motion.go`）**
- `stateTransition()`：状态色渐变的过渡（`MotionFast`，只动颜色）。
- 姿态动画（`AnimatePose`）**不用**：实测会让元素停在不可见状态（见 §2）。

**各组件**
| 组件 | 动效 |
|---|---|
| Button | hover/pressed 的底色与文字色渐变（`stateTransition`） |
| Input | 外框描边渐变（聚焦时转强调色、错误态转 danger 色）；清除按钮与错误文案**直接出现**（姿态动画不可用） |
| Badge | 变体换色渐变 |
| Icon | 只有转圈（`Spin`），其余不动 |

**后置**：退场渐隐（保留元素 + 颜色过渡）、弹层类组件的进出曲线（Popover/Dialog/Drawer 等）、列表重排动画、骨架屏的脉冲/微光（可用数值动画，不依赖姿态）。

## 4. 项目结构是否合理

- 位置：根包 L1（tokens 之下的动效约定），与 `tokens.go` 的动效刻度同层；被各组件引用。
- 依赖：只用 `mygo/ui` 的动效 API，无定时器、无 goroutine ✓（AGENTS.md 硬规则）。

## 5. API 契约

```go
// stateTransition is the transition the presets use: the colors of a
// state change fade, and the rest happens at once.
func stateTransition() ui.ElementTransition
```

（包内约定，不导出；对外的动效入口是各组件自己的方法与 `ctrlpkg.Icon.Spin`。）

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 共用约定（`motion.go`）
- [x] Button / Input / Badge 的状态色渐变
- [x] Input 的聚焦描边渐变
- [ ] 退场渐隐、弹层曲线、列表重排、骨架屏脉冲（后置）
- [ ] 姿态动画（`AnimatePose`）——实测不可用，待上游确认后再评估
