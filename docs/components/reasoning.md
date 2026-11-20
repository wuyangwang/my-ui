# Reasoning 规格书

> 位置：组件在 [`ai/`](../../ai) 子包；主题/tokens 与文案表在根包 `myui`。
>

## API 契约

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

## 状态与限制

### 已知限制

- 档位数量与含义由 app 定义；组件只管索引，不把「高」绑定到任何具体模型行为。
- 思考过程（推理文本）的折叠展示不是本组件的事：它只是另一段 Markdown，用现有 `markdown/` 与 `display/` 即可。
