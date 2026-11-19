# GroupAvatar

用成员头像的网格表示群聊，落在主题色的底面上；单个成员时放大显示。

```go
display.NewGroupAvatar(c, display.GroupAvatarOptions{
    Size: 40, Grid: display.GroupGridAuto,
    Avatars: []display.GroupAvatarItem{
        {Avatar: "Ada", Fallback: display.AvatarText},
        {Avatar: "🤯"}, {Avatar: "🚀"},
    },
})
```

## 行为

- 两列最多 4 人、三列最多 9 人；`GroupGridAuto` 在超过 4 人时用三列。
- 圆角方形或圆形由 `Corner` 决定，成员头像由 `NewAvatar` 绘制，复用其纹理、首字母与 emoji 逻辑。
- 成员坐标按网格居中排布，单个成员占八成边长。

## API

| 属性 | 类型 | 说明 |
| --- | --- | --- |
| `NewGroupAvatar(c, o)` | `*ui.Context, GroupAvatarOptions` | 构造群头像 |
| `GroupAvatarOptions.Avatars` | `[]GroupAvatarItem` | 成员，按顺序；超出网格的省略 |
| `GroupAvatarOptions.Size` | `float32` | 边长（DIP），0 用 32 |
| `GroupAvatarOptions.Grid` | `GroupGrid` | `Two`/`Three`/`Auto` |
| `GroupAvatarOptions.Corner` | `GroupCorner` | `Square` 或 `Circle` |
| `GroupAvatarOptions.Shape` | `myui.Shape` | 每个成员头像的形状 |
| `GroupAvatarItem.Avatar` | `string` | 成员的姓名或 emoji |
| `GroupAvatarItem.Image` | `ui.ImageSource` | 成员图片 |
| `GroupAvatarItem.Background` | `ui.Color` | 成员底色 |
| `GroupAvatarItem.Fallback` | `AvatarFallback` | 成员回退方式 |

## 限制

- Lobe 的 squircle 超椭圆需要路径蒙版，mygo 没有该能力；这里只有圆角方形与圆形两种轮廓。
- 成员数超过网格容量时被截断，不提供“+N”角标。
