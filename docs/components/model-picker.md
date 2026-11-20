# ModelPicker 规格书

> 位置：组件在 [`ai/`](../../ai) 子包；主题/tokens 与文案表在根包 `myui`。
>

## API 契约

| API | 作用 |
|---|---|
| `type ModelOption struct` | 一个模型：ID、名称、分组、说明、图标、上下文窗/速度/价格徽标、禁用 |
| `func NewModelPicker(c *ui.Context, text *string, options []ModelOption) *ModelPicker` | 构建可搜索的模型下拉；`text` 由 app 持有 |
| `func (p *ModelPicker) Chosen() string` | 上次选中的模型 ID，任意帧可读 |
| `func (p *ModelPicker) Selected() *ModelOption` | 上次选中的模型，未选中返回 nil |
| `func (p *ModelPicker) OnChanged(fn func(id string)) *ModelPicker` | 本帧选中新模型时回调 |
| `func (p *ModelPicker) Changed() bool` | 本帧是否选中了新模型 |
| `func (p *ModelPicker) Filter(fn func(ModelOption, string) bool) *ModelPicker` | 自定义过滤 |
| `func (p *ModelPicker) Variant/Size/Shape/Width/Placeholder/Label/Disabled/Invalid/EmptyText/Open/SetOpen/Control` | 与 `forms.ComboBox` 一致的外观与状态方法 |

用法：

```go
var query string
p := ai.NewModelPicker(c, &query, []ai.ModelOption{
    {ID: "gpt-5", Name: "GPT-5", Group: "OpenAI", Context: "256K", Speed: "快", Price: "$$"},
    {ID: "claude", Name: "Claude", Group: "Anthropic", Context: "200K", Speed: "中", Price: "$$"},
}).Outline()
if p.Changed() { app.model = p.Chosen() }
```

## 状态与限制

### 已知限制

- 思维强度不塞进同一个弹出层：它是另一个独立状态，由 `Reasoning` 负责，两个组件并排即可组成 shadcn 的双菜单。
- 价格/上下文窗只作展示字符串，不做排序或过滤逻辑；app 可自己排好 `options`。
