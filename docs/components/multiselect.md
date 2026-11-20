# MultiSelect 规格书

> 位置：组件在 [`forms/`](../../forms) 子包；主题/tokens 仍在根包 `myui`。

多选：选中的项以芯片（tag）呈现。

## API 契约

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

## 状态与限制

**限制**

- **值即标签**：芯片显示的就是 app 切片里的字符串，多选无法像 Select 那样「值与标签不同」（框架的 TokenField 只存字符串）。若要 label≠value，得自己维护两份（芯片用标签、app 存值）。
- 折叠标签（`collapse-tags`）没有公开 API。`Max` 现在强制阻止 Enter/逗号/粘贴/点建议加入（在构造器里截断切片），但移除芯片与退格删除始终可用——上限只管「加」，不管「减」。
- 上限在下一帧才生效：构造时先按当前上限建字段，`Max` 把新值写入 `ui.Local`，下一帧读取。实时应用中帧帧连续，无感；按需重建的测试里要点一次触发重建才看到新上限。
