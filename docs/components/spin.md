# Spin 规格书

> 状态：**已实现**。加载指示器：小尺寸的旋转圆圈。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Spinner | `LoaderCircle` 图标 + 旋转动画 |
| Element Plus Loading | 语义色与尺寸档 |

## 2. 能否实现
- `Icon.Spin()` 用 `ui.Loop` + `Rotate`，由框架处理 Reduce Motion ✓；不需要计时器。

## 3. 实现到何种程度
- `NewSpin(c)`；`.Size(sz)`、`.Variant(v)`；默认 `IconSizeMd`、强调色。

## 4. 项目结构
- `display/` 子包。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewSpin(c *ui.Context) *Spin` | 创建加载指示器，默认文字档尺寸、强调色 |
| `Size(sz Size) *Spin` | 尺寸档：`SizeXs/Sm/Md/Lg`（默认 `IconSizeMd` 档） |
| `Variant(v Variant) *Spin` | 语义色：语义变体取主题色，默认强调色 |

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
