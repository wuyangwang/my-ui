# Spin 规格书

> 状态：**已实现**。按 Lobe MyGo 审查修复：Percent/Network/Tip/Content 选项、环与节点网络、`RoleStatus`。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| Lobe MyGo `feedback/spin.go` | `SpinProps` 的 `Variant`（Loader/Network）、`Percent`、`Tip`、`Spinning`/content overlay；环用 `Path.Circle` + 折线弧；网络用 3×3 节点 + 连线 + `Loop` 脉冲；`RoleStatus` |
| shadcn/ui Spinner | `LoaderCircle` 图标 + 旋转动画 |

## 2. 能否实现（mygo v0.10.0 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 旋转 loader | `ctrlpkg.Icon.Spin()`（内部 `Loop`）✓ | ✅ |
| 百分比环 | `Box.Draw` + `Painter.StrokePath` + `Path.Circle` ✓ | ✅ |
| 节点网络 | `Box.Draw` + `Path.MoveTo/LineTo/Circle` + `Loop` ✓ | ✅ |
| 内容遮罩 | 绝对定位 `Fill` + 半透明底 ✓ | ✅ |
| 无障碍 | `Role(ui.RoleStatus)` ✓ | ✅ |

## 3. 实现到何种程度

- 构造时建两套字形（独立显示一套、内容遮罩一套），每套含 loader、环、网络、文案槽；模式切换只切显隐，不重建。
- `SpinPercent(p)`：`Painter` 画一条轨道环 + 折线圆弧，填充到 `clamp(p,0,100)%`。
- `SpinNetwork()`：`20s` 周期的 `Loop` 驱动 3×3 节点亮度与连线，节点按序脉冲。
- `SpinTip(s)`：指示器下的文字；`SpinContent(content, spinning)`：内容照常显示，`spinning` 为真时盖上半透明遮罩与指示器。
- 所有动画用 `Element.Loop`，无计时器/goroutine；Reduce Motion 由框架兜底。
- 根元素带 `RoleStatus`，读屏把它当建议性 live region。

## 4. 项目结构

- `display/` 子包。

## 5. API 契约（逐 prop）

### 类型

| 符号 | 作用 |
|---|---|
| `type SpinVariant uint8` | 指示器形态标记 |
| `SpinLoader` / `SpinNetwork` / `SpinPercentKind` | `SpinNetwork` 由 `SpinNetwork()` 选中；`SpinPercentKind` 由 `SpinPercent()` 选中 |

### 构造与方法

| 方法 | 作用 |
|---|---|
| `NewSpin(c *ui.Context) *Spin` | 创建 loader；默认文字档尺寸 |
| `Size(sz Size) *Spin` | 尺寸档（Xs/Sm/Md/Lg） |
| `Variant(v Variant) *Spin` | 语义色；默认 muted |
| `SpinPercent(p float32) *Spin` | 显示填充到 p%（0..100）的环 |
| `SpinNetwork() *Spin` | 显示节点网络 |
| `SpinTip(s string) *Spin` | 指示器下的文字；空串清除 |
| `SpinContent(content func(), spinning *bool) *Spin` | 内容 + 遮罩；`spinning==nil` 表示一直转 |

用法：

```go
display.NewSpin(c).SpinPercent(40).SpinTip("四成")
display.NewSpin(c).SpinContent(func() { ui.Text(c, "内容") }, &loading).SpinTip("加载中")
```

## 6. 限制

- 网络图是 Lobe 节点的简化版：保留 3×3 节点和连线，去掉了外圈虚线旋转与粒子流，避免额外绘制复杂度。
- `SpinContent` 的遮罩底色是 `t.Background.Alpha(0.7)`，不模糊背景（框架无现成模糊接口需额外关注性能）。

## 7. 状态

- [x] 规格书 / 实现 / 测试（尺寸/变体、Percent/Network/Tip、Content 遮罩）
