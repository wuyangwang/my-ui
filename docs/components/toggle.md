# Toggle 规格书

> 位置：控件在 [`controls/`](../../controls) 子包；主题/tokens 与文案表在根包 `myui`。
>

## API 契约

| API | 作用 |
|---|---|
| `func NewToggle(c *ui.Context, on *bool, label string) *Toggle` | 构建绑定 `*on` 的切换按钮 |
| `func (t *Toggle) Variant(v Variant) *Toggle` | 应用共享变体 |
| `func (t *Toggle) Primary/Secondary/Outline/Ghost/Destructive() *Toggle` | 预设变体快捷方法 |
| `func (t *Toggle) Size(s Size) *Toggle` | 应用共享尺寸 |
| `func (t *Toggle) Xs/Sm/Md/Lg() *Toggle` | 尺寸快捷方法 |
| `func (t *Toggle) Icon(svg *ui.SVG) *Toggle` | 在文字前放图标（icon-only 时传空 label + `Label`） |
| `func (t *Toggle) Label(s string) *Toggle` | 无障碍名称 |
| `func (t *Toggle) Disabled(on bool) *Toggle` | 禁用或启用 |
| `func (t *Toggle) Block() *Toggle` | 铺满所在行 |
| `func (t *Toggle) Changed() bool` | 本帧是否翻转 |
| `func (t *Toggle) OnChanged(fn func(on bool)) *Toggle` | 本帧翻转时立即回调（必须同帧链式调用） |
| `func (t *Toggle) Handle(fn func(on bool)) *Toggle` | 注册常驻回调，可在任意帧调用 |

用法：

```go
var bold bool
ctrlpkg.NewToggle(c, &bold, "粗体").Primary().Icon(icons.Bold()).
    Handle(func(on bool) { app.status = fmt.Sprintf("bold=%v", on) })

if t.Changed() { /* 本帧翻转 */ }
```

## 状态与限制

### 已知限制

- `OnChanged` 是即时链式调用，晚一帧 attach 会丢事件；这种场景用 `Handle`。
- 图标与文字的间距取主题 `Space(1)`，不提供单独的图标位置/大小选项（用共享 IconSizeSm）。
