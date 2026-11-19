# Tokens 规格书

> 状态：**已实现**。my-ui 的共享设计刻度：主题色、字号、间距、圆角、图标尺寸与动效时长，全部从 `*ui.Theme` 解析，逐帧跟随主题与桌面偏好。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui CSS 变量 | 一套 token 驱动全部控件 |
| Lobe MyGo Tokens | Go 常量 + 主题派生色 |

## 2. 能否实现
- 根包 `tokens.go`/`theme.go`/`palettes.go`/`shades.go`/`motion.go` 已实现；无包级颜色缓存 ✓。

## 3. 实现到何种程度
- 字号 `TextXs..Text2Xl` + `TextSize(t,n)`；间距 `SpacingXs..Spacing2Xl` + `Space`；圆角 `RadiusSm..RadiusXl` + `Radius`（对齐 shadcn 4/6/8/12）；图标 `IconSizeXs..Lg`（12/14/16/24）；动效 `MotionFast/Normal/Slow`。
- 颜色一律取 `t.Background/Surface/Border/Text/Accent/...`；主题留空的 `Inverse/InverseText` 用 `InverseColors(t)` 兜底。

## 4. 项目结构
- 根包 `myui`（核心底座，不导出组件）。

## 5. API 契约

| 符号 | 作用 |
|---|---|
| `TextSize(t,n)` | 主题字号的 n 倍 |
| `Space(t,n)` | 主题间距的 n 单位 |
| `Radius(t,s)` | 主题圆角的 s 倍 |
| `IconSizeXs/Sm/Md/Lg` | 12/14/16/24 DIP 图标刻度 |
| `MotionFast/Normal/Slow` | 120/200/320ms 动效时长 |
| `Apply(c,o)` / `NewTheme(c,o)` | 把主题装进帧/子树 |
| `InverseColors(t)` | 主题留空时的反色兜底 |

## 6. 状态
- [x] 规格书 / 实现 / 测试
