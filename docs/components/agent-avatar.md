# AgentAvatar

把智能体画成扁平剪影加一对大眼睛。形状、颜色与倾斜角由种子决定；状态改变眼睛与角标。

```go
display.NewAgentAvatar(c, display.AgentAvatarOptions{
    Seed: "lobe", Name: "助手", State: display.AgentWorking,
})
```

## 行为

- `Seed`（为空时用 `Name`）经 FNV-1a 得到稳定的形状、颜色与倾斜角，同名总是同一形象。
- 形状包含 Face、Bowl、Cloud、Tile、Hex、Shield、Pudding、Cookie、Star、Capsule、Ghost、Cat；多边形用 `ui.Path` 绘制，圆角形用 `Painter.Fill`。
- 颜色来自每帧 `c.Theme()`（强调色与语义色），可经 `CustomColor` 覆盖，不使用固定调色板。
- `State` 改变视线与眼睛：思考时上望、工作时下视、完成时闭眼、阻塞时睁大并带危险色角标。
- 绘制是静态的：没有计时器或 goroutine，因此不产生动画，也就不违反桌面“减少动态效果”设置。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewAgentAvatar(c, o)` | `*ui.Context, AgentAvatarOptions` | 构造智能体头像 |
| `AgentAvatarOptions.Seed` | `string` | 决定形象、颜色与倾斜的种子 |
| `AgentAvatarOptions.Name` | `string` | 无障碍名称 |
| `AgentAvatarOptions.Shape` | `AgentShape` | `AgentShapeSeeded` 或具体形状 |
| `AgentAvatarOptions.Color` | `AgentColor` | `AgentColorSeeded` 或具体颜色 |
| `AgentAvatarOptions.CustomColor` | `ui.Color` | 覆盖颜色 |
| `AgentAvatarOptions.Size` | `float32` | 边长（DIP），0 用 40 |
| `AgentAvatarOptions.State` | `AgentState` | `AgentIdle`/`Working`/`Waiting`/`Blocked`/`Thinking`/`Done` |
| `AgentAvatarOptions.NoTilt` | `bool` | 直立不倾斜 |

## 限制

- 与 Lobe 不同，这里不实现眨眼、扫视与摆动动画：组件不带计时器，状态是静态姿态。
- Cloud 由几个圆叠加近似；超椭圆蒙版在 mygo 没有对应能力，改用最接近的圆角。
- 小于或等于 24 DIP 时省略眼睛以保持辨识度。
