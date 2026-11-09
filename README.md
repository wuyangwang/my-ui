# my-ui

MyGo 原生 UI（[`package ui`](https://github.com/egoist/mygo)）的组件库，独立发布的第三方 module —— 不依赖 webview，只面向 MyGo 自绘的原生 UI（macOS / Linux / Windows，同一份代码还能跑进浏览器的 wasm 构建）。

定位：**主题包 + 语义常量 + 变体化控件预设 + 框架没有的组合组件**，而不是重造控件。`mygo` 的 `ui` 已经是完整的立即模式控件层；my-ui 在其公开 API 之上补设计系统与桌面/AI 应用组件。

## 特性目标（与 shadcn 对照的六项决定）

| | |
|---|---|
| 多种主题 | 内置主题 preset 集（衍生自系统外观/强调色/高对比/文本缩放，自动跟随），配合 `ui.Themed` 做子树级主题；详见 [docs/architecture.md](docs/architecture.md) |
| 动画与过渡 | 全部基于框架自带的弹簧/曲线/AnimatePose/compositor 级动画体系，库只出动效刻度与雅致默认值 |
| 易用 | 一行构造器 + Variant/Size 参数，沿用 ui 的"输入即返回值"绑定约定，无需手动接回调 |
| 常量体系 | 常见字体（可选内嵌）、颜色调色板、padding/margin/圆角/动效/滚动内衬刻度 |
| 细粒度桌面组件 | Card、Kbd、HoverCard、Drawer、Command Palette、Skeleton、EmptyState、Pagination、表格排序等 `ui` 未封装的常见件 |
| AI 组件 | ChatList（虚拟化流式气泡）、模型选择器、推理档位等 AI 应用常见控件 |

已明确**不做**瀑布流（`ui` 无此布局原语，性价比不足）。

## 包结构

- 根包 `myui`（import path 含 `-`，Go 包名取 `myui`）：tokens、主题预设、变体化控件、Page 页面骨架。
- 二级包按需引入，不 import 不链接（遵循 mygo "apps pay for what they use"）：
  - `icons/` —— 内嵌 Lucide（ISC）与品牌标（simple-icons, CC0），`currentColor` 机制天然支持变色；
  - `fonts/` —— 内嵌常见开源字体（OFL），通过 `ui.RegisterFont` 注册字族；
  - `ai/` —— ChatList、模型选择器、推理档位等重组件。

## 状态

脚手架与设计文档阶段：**尚未编写组件实现**。工程结论、证据与上游注意事项见 [docs/research.md](docs/research.md)，分层与发布策略见 [docs/architecture.md](docs/architecture.md)，排期见 [docs/roadmap.md](docs/roadmap.md)。给代理/协作者的开发规则见 [AGENTS.md](AGENTS.md)。

## 版本约定

锁定所发布 mygo 版本（当前目标 **v0.8.3**，见 [docs/architecture.md](docs/architecture.md) §7；首个组件代码落地时写入 `go.mod`）。`ui` 的 element API 上游有过破坏性迁移（官方提供 `mygo migrate-ui`），my-ui 发版与其一一对应；消费者应直接使用与该版 mygo 匹配的 my-ui 版本。

## 许可

MIT。内嵌资源按其各自许可再分发（Lucide ISC、simple-icons CC0、字体 OFL）。
