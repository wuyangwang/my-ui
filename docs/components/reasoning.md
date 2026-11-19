# Reasoning 规格书

> 位置：组件在 [`ai/`](../../ai) 子包；主题/tokens 与文案表在根包 `myui`。
>
> 状态：**已实现**（`ai/reasoning.go`）。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| shadcn AI Elements `Reasoning` / `Model Selector` | 「思考强度」是低/中/高（或更多档）的一排选择；温度与 top_p 是连续参数，常配精确/均衡/创意预设 |
| assistant-ui `Model Selector` | 强度档位支持自定义集合，选中模型不支持时不显示 |
| mygo `ui.Segmented` / `ui.Slider` | 单选档位与连续取值都有公开底座，键盘、拖动与无障碍归框架 |
| my-ui `controls.NewSegmented` / `controls.NewSlider` | 已把底座套上共享 Variant/Size、主题色与动效，直接复用 |

## 2. 能否实现

- 强度档位：`controls.NewSegmented` 绑定 `*int`，档位标签来自根包 i18n 文案表 ✓
- 温度/top_p：`controls.NewSlider` 绑定 `*float64`，显示当前值 ✓
- 预设：按钮点击写回 app 持有的三个指针，下一帧控件读到新值 ✓
- 思考过程展示：用 `markdown/` + `display.NewCollapsible` 自行组合，本组件不做（见「已知限制」） ✓

## 3. 实现到何种程度

- `ReasoningOptions`：`Effort *int`（档位索引）、`Levels []string`（缺省为国际化的「低/中/高」）、`Temperature *float64`（0..2）、`TopP *float64`（0..1）。
- `NewReasoning(c, o)`：从上到下排档位、温度、top_p 与三个语义预设；任一指针为 nil 即省略对应行。
- `Compact(true)` 只留档位；`Label(name)` 给档位组命名；`OnChanged(fn)` 在档位或滑杆本帧变化时回调。
- 所有颜色、间距、字号每帧取 `c.Theme()`；滑杆、分段控件走共享 `StateTransition()` 与 motion tokens，无计时器/goroutine。

## 4. 项目结构

- `ai/` 子包；具名导入 `controls/`，dot-import 根包。根包不导入它。
- 只依赖 mygo 公开 API；文案通过根包 `Message(c, key)` 取，新增的键写在 `i18n.go` 的三种语言表里。

## 5. API 契约

| API | 作用 |
|---|---|
| `type ReasoningOptions struct` | 绑定的强度索引、档位标签、温度、top_p |
| `func NewReasoning(c *ui.Context, o ReasoningOptions) *Reasoning` | 构建设置面板 |
| `func (r *Reasoning) Compact(on bool) *Reasoning` | 只显示强度档位 |
| `func (r *Reasoning) Label(name string) *Reasoning` | 给档位组命名（无障碍） |
| `func (r *Reasoning) OnChanged(fn func()) *Reasoning` | 档位或滑杆变化时回调 |
| `func (r *Reasoning) Changed() bool` | 本帧是否有控件变化 |

用法：

```go
effort, temp, topP := 1, 0.7, 0.95
ai.NewReasoning(c, ai.ReasoningOptions{
    Effort: &effort, Temperature: &temp, TopP: &topP,
}).Label("推理设置")
```

## 6. 状态

- [x] 调研与规格书
- [x] 实现（`reasoning.go`）
- [x] 测试（`reasoning_test.go`）
- [x] showcase「Reasoning」页

### 已知限制

- 档位数量与含义由 app 定义；组件只管索引，不把「高」绑定到任何具体模型行为。
- 思考过程（推理文本）的折叠展示不是本组件的事：它只是另一段 Markdown，用现有 `markdown/` 与 `display/` 即可。
