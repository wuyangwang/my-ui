# QRCode

卡片上的二维码，可加中央图标与加载/过期状态。编码在库内完成，不引入新依赖。

```go
display.NewQRCode(c, "https://lobehub.com", display.QRCodeOptions{
    Size: 120, Status: display.QRActive,
})
```

## 行为

- 编码为 QR Model 2、字节模式、纠错级别 L、版本 1–5（单块）。模块按最优的 8 种掩码之一绘制。
- `Color`/`Background` 为零时取当帧主题的正文墨色与背景；圆角取 `Radius(t, RadiusXl)`。
- `Icon` 在中央留洞并画 logo；洞的大小约为边长的 22%。
- `Status` 覆盖加载（旋转图标）或过期（文字与刷新按钮）状态；刷新调用 `OnRefresh`。
- 模块路径按像素对齐绘制，避免缩放时的缝隙；矩阵用 `ui.Local` 按值缓存，不跨帧持有颜色。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewQRCode(c, value, o)` | `*ui.Context, string, QRCodeOptions` | 构造二维码 |
| `QRCodeOptions.Size` | `float32` | 码边长（DIP，不含内边距），0 用 160 |
| `QRCodeOptions.Color` | `ui.Color` | 模块颜色，零值用主题正文墨色 |
| `QRCodeOptions.Background` | `ui.Color` | 卡片底色，零值用主题背景 |
| `QRCodeOptions.NoBorder` | `bool` | 去掉卡片边框 |
| `QRCodeOptions.Icon` | `func(c *ui.Context)` | 中央 logo，留洞 |
| `QRCodeOptions.Status` | `QRStatus` | `QRActive`/`QRLoading`/`QRExpired` |
| `QRCodeOptions.OnRefresh` | `func()` | 过期时的刷新按钮回调 |
| `QRCodeOptions.Label` | `string` | 无障碍名称，默认 `value` |

## 限制

- 不支持 Lobe 的纠错级别 L/M/Q/H 切换：本实现固定为级别 L，图标留洞会削弱纠错能力而非自动提升到 H。
- 只支持版本 1–5（约 100 字节以内的 UTF-8 值）；更长的值返回空矩阵，卡片显示为空。
- 加载状态只显示静态图标，不做旋转动画（无计时器）。
