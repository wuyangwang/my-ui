# Badge 规格书

> 状态：**已实现**（`badge.go`）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [shadcn/ui Badge](https://ui.shadcn.com/docs/components/badge) | 小标签：`default`/`secondary`/`destructive`/`outline` 四变体 × 三尺寸；圆角胶囊 |
| [Element Plus el-tag](https://element-plus.org/zh-CN/component/tag) | 语义色（success/info/warning/danger）、可关闭、尺寸档 |
| [Lobe MyGo Badge](https://mygo.lobehub.com/skills/components/badge) | 两种形态：**文字标签**与**角落计数/状态点**（挂在图标或头像的角上） |
| mygo 的 `ui.Badge(c, text)` | 框架自带胶囊小字（未读计数用） |

## 2. 能否实现

- 构造：`ui.Badge(c, text)` ✓（框架自带胶囊观感）。
- 尺寸/颜色：走主题槽位与 `Size` 刻度 ✓。
- 角落计数（角标 + 状态点）：需要把徽标叠在图标/头像的角上 → mygo 有绝对定位（`Absolute()`）与 `Overflow` 语义可用，但**框架自带的 Badge 是独立元素**，叠放要靠布局组合，后置。

## 3. 实现到何种程度

**已实现**：`Native`（零值，保留内置观感）/`Secondary`（别名）/`Primary`/`Outline`/`Ghost`/`Destructive`/`Link` × `SizeSm/Md/Lg`；填充色/描边/文字色按预设。

**后置**：角标与状态点（挂在 `myui.Icon`/头像的角上）、可关闭标签（关闭按钮）、语义色别名（success/info/warning）。

## 4. 项目结构是否合理

- 位置：根包 L3，与 Button/Input/Icon 同层；共享 `Variant`/`Size` 词表。

## 5. API 契约

```go
func NewBadge(c *ui.Context, text string) *Badge

func (b *Badge) Variant(v Variant) *Badge
func (b *Badge) Primary() *Badge
func (b *Badge) Secondary() *Badge
func (b *Badge) Native() *Badge
func (b *Badge) Outline() *Badge
func (b *Badge) Ghost() *Badge
func (b *Badge) Destructive() *Badge
func (b *Badge) Link() *Badge

func (b *Badge) Size(s Size) *Badge  // SizeSm/Md/Lg
```

用法：

```go
myui.NewBadge(c, "12").Primary().Sm()
```

## 6. 状态

- [x] 调研与实现
- [x] 测试（`badge_test.go`：变体填充色、尺寸、暗色主题）
- [x] showcase「Badge」页
- [ ] 角标/状态点、可关闭标签（后置）
