# Avatar

用户或智能体的头像：图片、emoji、姓名首字母，或按姓名生成的纹理由此呈现。同名的纹理在同一主题下始终一致。

```go
display.NewAvatar(c, "Ada Lovelace", display.AvatarOptions{}).Sm()
display.NewAvatar(c, "🤯", display.AvatarOptions{Shape: myui.ShapeCircle})
display.NewAvatar(c, "Grace", display.AvatarOptions{
    Fallback: display.AvatarText, Background: ui.Hex("#7858f9"),
})
```

## 行为

- 有 `Image` 时铺满头像；否则先识别 emoji，再按 `Fallback` 画纹理或首字母。
- 纹理的色相取自**当前帧** `c.Theme()`（强调色、成功/警告/信息色、次要墨色与按压面），不是包级固定色表，因此明暗与自定义配色会跟着变化。
- 首字母取大写的头两个 rune，`Full` 可以显示全部；文字墨色按底色亮度取主题墨色或白色。
- 圆角按尺寸比例计算：`ShapeCircle` 为圆，小于 24 DIP 时约为三分之一，其余约六分之一。
- 无计时器、无 goroutine；颜色与尺寸每帧从主题重算。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewAvatar(c, name, o)` | `*ui.Context, string, AvatarOptions` | 构造头像；`c` 为帧上下文，`name` 为姓名或 emoji |
| `AvatarOptions.Size` | `float32` | 边长（DIP），0 用 48 |
| `AvatarOptions.Shape` | `myui.Shape` | `ShapeRect`（默认圆角）或 `ShapeCircle` |
| `AvatarOptions.Image` | `ui.ImageSource` | 覆盖显示的图片 |
| `AvatarOptions.Fallback` | `AvatarFallback` | `AvatarMarble`（纹理）或 `AvatarText`（首字母） |
| `AvatarOptions.Full` | `bool` | 文本回退显示全部字母，而非前两个 |
| `AvatarOptions.Background` | `ui.Color` | 底色；文字按亮度取可读墨色 |
| `AvatarOptions.Label` | `string` | 无障碍名称，默认 `name` |
| `(*Avatar).Pixels(px)` | `float32` | 固定边长 |
| `(*Avatar).Size(s)` | `myui.Size` | 共享尺寸词表：Xs 24、Sm 32、Md 48、Lg 64 |
| `(*Avatar).Circle()` | — | 变成圆形 |

`Fallback` 与 `Full` 在 `AvatarOptions` 中构造时设置；流式方法只调整尺寸与形状，不重建元素。

## 限制

- Lobe 的纹理使用固定色表；这里按 AGENTS.md 的要求改为主题派生色，同一姓名在不同主题下颜色不同，但仍保持可复现。
- 图片解码由调用方完成，组件只接收 `ui.ImageSource`。
