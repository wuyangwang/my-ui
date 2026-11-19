# MultiSelect 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`multiselect.go`）。多选：选中的项以芯片（tag）呈现。

## 1. 参考实现

| 来源 | 具体条目 | 我们怎么用 |
|---|---|---|
| mygo 的 `ui.TokenField` | 「a field of tokens, as of tags or the recipients of a mail: `*tokens` shows as chips, each with a button taking it out, before an input where Enter or a comma adds what was typed, and Backspace in it while empty takes out the last. The suggestions containing what was typed show below it, those not already tokens」；`Changed()` 报告变化 ✓ | 芯片与建议列表直接用它 ✓（这是框架里最接近多选的东西 ✓） |
| [Element Plus el-select multiple](https://element-plus.org/en-US/component/select.html) | `multiple`：值变数组；选中项默认以 **Tag** 形态显示；`collapse-tags`/`max-collapse-tags`/`tag-type`；`multiple-limit` | 芯片形态 ✓（TokenField ✓）；折叠与上限 ✗（后置 ✗） |
| [Ant Design Select mode=multiple](https://ant.design/components/select-cn) | `maxTagCount`（默认 responsive）、`tagRender`、`maxCount`（超出后该项不可选） | 后置 ✗ |
| [shadcn/ui Combobox 多选](https://ui.shadcn.com/docs/components/combobox) | 选中项以 badge 呈现，可单独移除 | 同上 ✓ |

## 2. 能否实现（mygo v0.8.3 公开 API）

| 能力 | 公开 API | 结论 |
|---|---|---|
| 芯片字段 | `ui.TokenField(c, tokens *[]string, suggestions []string)` ✓ | ✅ |
| 增删选项 | Enter/逗号加入、退格删最后一个、芯片上的按钮删单个 ✓（框架内建 ✓） | ✅ |
| 变化报告 | `Element.Changed()` ✓ | ✅ |
| 标签与值不同 | TokenField 的 token 就是显示的字符串 ✗ | ⚠️ 我们的选项本可 label≠value（Select ✓），但芯片显示的就是 token ✗ → **多选的值即标签**（框架限制 ✗） |
| 折叠标签 | 无公开 API ✗ | ❌ 后置 |
| 选择上限 | TokenField 无开关，但可在构造器里截断 `*tokens` ✓ | ✅ `Max(n)` 强制上限 |

## 3. 实现到何种程度

**已实现**
- `NewMultiSelect(c, tokens *[]string, options []SelectOption)`：选中的项以芯片呈现（各带一个移除按钮），下面按输入过滤出还没选的选项 ✓。
- 变体与尺寸：与 Input 同一套 `Variant`/`Size`/`Shape` ✓（外观画在字段上 ✓）。
- 事件：`OnChanged(fn func(values []string))` ✓；查询式 `Values()`/`Changed()`。
- 无障碍名 `Label(string)` ✓、占位文字 `Placeholder(string)` ✓。
- `Max(n int)`：上限达到后不再给建议，并且**真正阻止继续加入** ✓。框架 TokenField 没有上限开关，Enter/逗号/粘贴/点建议都会直接追加到 `*tokens`；构造器在读取上限后、建芯片与建议之前把切片截到 `n`，于是所有加入路径都受限，移除芯片不受影响。上限记在字段宿主元素的 `ui.Local` 上，由下一帧构造器读取；`Max` 本身不重建 TokenField，所以同帧链式调用不会让一次点击被两个字段重复消费 ✓。

**已知限制**
- **值即标签**：芯片显示的就是 app 切片里的字符串 ✗ → 多选无法像 Select 那样「值与标签不同」 ✗（框架的 TokenField 只存字符串 ✗）。若要 label≠value，得自己维护两份（芯片用标签、app 存值）✗ → 后置 ✗。
- 折叠标签（`collapse-tags`）没有公开 API ✗ → 后置 ✗。`Max` 现在强制阻止 Enter/逗号/粘贴/点建议加入（在构造器里截断切片），但移除芯片与退格删除始终可用——上限只管「加」，不管「减」。
- 上限在下一帧才生效：构造时先按当前上限建字段，`Max` 把新值写入 `ui.Local`，下一帧读取。实时应用中帧帧连续，无感；按需重建的测试里要点一次触发重建才看到新上限。

**后置**
- 值与标签分离（自己维护映射 ✓）、折叠标签、芯片的自定义渲染。

## 4. 项目结构是否合理

- 位置：根包 L3，与 `Select`/`ComboBox` 同层；**包 mygo 的 `ui.TokenField`** ✓（芯片与建议列表是框架内建的 ✓，自己重写不值 ✓）。
- 依赖：`mygo/ui` ✓。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewMultiSelect(c *ui.Context, tokens *[]string, options []SelectOption) *MultiSelect` | 创建多选字段；`tokens` 是选中的值（app 持有） |
| `Variant(v Variant)` · `Size(s Size)` · `Sm()` · `Md()` · `Lg()` · `Xs()` | 变体与尺寸（与 Input 同义） |
| `Shape(s Shape)` | 形态（与 Input 同义） |
| `Width(w float32)` | 宽度 |
| `Placeholder(s string)` · `Label(s string)` | 占位文字 / 无障碍名 |
| `OnChanged(fn func(values []string))` | 选择变化时回调 |
| `Max(n int)` | 选择上限；达到后不再给建议，且 Enter/逗号/粘贴/点建议都无法超过（`n<=0` 表示不限），移除不受限；下一帧生效，不在当前帧重建字段 |
| `Values() []string` · `Changed() bool` | 当前选中的值 / 本帧是否变化 |

用法：

```go
h := bindpkg.Bind(c)
h.MultiSelect(&app.langs, []formspkg.SelectOption{
	{Label: "简体中文", Value: "zh-Hans"},
	{Label: "English", Value: "en"},
}).Outline().Label("语言").
	OnChanged(func(v []string) { app.langs = v })
```

## 6. 状态

- [x] 调研与可行性结论（本文）
- [x] 实现（`multiselect.go`，包 mygo 的 `ui.TokenField`）
- [x] 测试（加入与移除、变化回调、同帧 `Max` 不重复消费点击、`Max(2)` 强制阻止第三项）
- [x] showcase「Select」页中的多选示例
- [x] 强制上限（已实现，见 `Max`）
- [ ] 值与标签分离、折叠标签（后置）
