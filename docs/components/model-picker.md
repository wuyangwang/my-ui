# ModelPicker 规格书

> 位置：组件在 [`ai/`](../../ai) 子包；主题/tokens 与文案表在根包 `myui`。
>
> 状态：**已实现**（`ai/modelpicker.go`）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| shadcn AI Elements `Model Selector` | 触发器是输入框，弹出层按厂商/系列分组，条目带图标、说明与能力徽标 |
| assistant-ui `Model Selector` | 模型与思维强度分两级；选中项回显在触发器上 |
| mygo `ui.ComboboxBase` | 可搜索下拉底座：`Popup`/`Item`/`Chosen`/`Open` 全部公开 |
| my-ui `forms.NewComboBox` | `Chosen` 只活当帧，需存进 `ui.Local` 槽位才能跨帧读；`Item` 写回文本 |

## 2. 能否实现

- 分组搜索下拉：`ui.ComboboxBase` ✓；`Item(value)` 的 value 用模型 ID，回显用模型名 ✓
- 厂商图标与徽标：`icons/`（品牌标可选）+ `ui.Icon`、主题色小胶囊 ✓
- 选中态：`ComboboxParts.Chosen()` 存进 `ui.Local`，`Chosen()`/`Selected()` 跨帧可读 ✓
- 思维强度：放在 `Reasoning` 组件里，由 app 组合（见「已知限制」） ✓

## 3. 实现到何种程度

- `ModelOption`：`ID`、`Name`、`Group`、`Description`、`Icon`、`Context`、`Speed`、`Price`、`Disabled`。
- `NewModelPicker(c, text, options)`：`text` 是 app 持有的搜索文本；弹出层分组渲染，行内显示图标、名称、说明与能力徽标，选中项打勾。
- 与 `forms.ComboBox` 同一套 `Variant`/`Size`/`Shape`/`Width`/`Placeholder`/`Label`/`Disabled`/`Invalid`/`Filter`/`EmptyText`/`Open`/`SetOpen`/`OnChanged`/`Chosen`/`Selected`/`Control`。
- 默认过滤按名称大小写不敏感匹配，可被 `Filter` 覆盖；`Chosen` 只活一帧，`Chosen()` 读持久槽。
- 每帧重建外观：字段背景、边框、圆角、字号取 `FrameLookOf`/`FramePadding`/`FrameFontSize`/`Radius`；`StateTransition()` 让焦点/禁用换色渐变。无包级颜色状态。

## 4. 项目结构

- `ai/` 子包，与 `ChatList`、`Reasoning` 同层；dot-import 根包，具名导入 `icons/`。
- 只依赖 mygo 公开 `ui` 与根包、`icons/`，不碰 internal，不带资产。

## 5. API 契约

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

## 6. 状态

- [x] 调研与规格书
- [x] 实现（`modelpicker.go`）
- [x] 测试（`modelpicker_test.go`）
- [x] showcase「ModelPicker」页

### 已知限制

- 思维强度不塞进同一个弹出层：它是另一个独立状态，由 `Reasoning` 负责，两个组件并排即可组成 shadcn 的双菜单。
- 价格/上下文窗只作展示字符串，不做排序或过滤逻辑；app 可自己排好 `options`。
