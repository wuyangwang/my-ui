# GuideCard

可关闭的提示卡，用于功能介绍：封面、标题、说明与右上角关闭按钮。

```go
display.NewGuideCard(c, display.GuideCardOptions{
    Title: "认识你的智能体", Desc: "它们记得上下文，接着上次继续。",
    Variant: myui.Outline, Shadow: true,
    Closed: &app.guideClosed, OnClose: func() {},
}).Width(220)
```

## 行为

- `Variant` 用共享词表：`Default` 是柔和渐变底、`Outline` 是带边框的面、`Ghost` 无面。
- `Cover` 是顶部图片，`CoverWidth`/`CoverHeight` 定尺寸（宽度为 0 时占满卡宽）。
- 关闭状态由 `Closed` 指针持有；为 nil 时组件用 `ui.Local` 自持。点击关闭会写指针/内部状态并调用 `OnClose`，下一帧卡片隐藏。
- 关闭按钮是 `controls.NewIconAction`，危险色悬停，带上无障碍名称。
- 颜色、圆角与间距每帧从 `c.Theme()` 取得，无计时器与 goroutine。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewGuideCard(c, o)` | `*ui.Context, GuideCardOptions` | 构造提示卡 |
| `GuideCardOptions.Title` / `.Desc` | `string` | 标题与说明 |
| `GuideCardOptions.Cover` | `ui.ImageSource` | 顶部图片 |
| `GuideCardOptions.CoverWidth` / `.CoverHeight` | `float32` | 图片尺寸 |
| `GuideCardOptions.Variant` | `myui.Variant` | 共享变体词表 |
| `GuideCardOptions.Shadow` | `bool` | 柔和阴影 |
| `GuideCardOptions.NoClose` | `bool` | 去掉关闭按钮 |
| `GuideCardOptions.Closed` | `*bool` | 可见性指针，nil 时组件自持 |
| `GuideCardOptions.OnClose` | `func()` | 关闭时回调 |

## 限制

- 关闭是永久的：组件没有“再次打开”的接口，app 需自行把 `Closed` 指针改回 false 并重建。
- 渐变只有两档主题色，不提供 Lobe 的多层光泽。
