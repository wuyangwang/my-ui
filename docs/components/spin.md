# Spin 规格书

提供 Percent/Network/Tip/Content 选项、环与节点网络、`RoleStatus`。

## API 契约（逐 prop）

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

## 限制

- 网络图是 Lobe 节点的简化版：保留 3×3 节点和连线，去掉了外圈虚线旋转与粒子流，避免额外绘制复杂度。
- `SpinContent` 的遮罩底色是 `t.Background.Alpha(0.7)`，不模糊背景（框架无现成模糊接口需额外关注性能）。
