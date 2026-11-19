# CopyButton 规格书

> 位置：`controls/copybutton.go`。

> 状态：**已实现**。参考 Lobe UI 的 CopyButton：点击复制文本，随后短暂显示对勾与「已复制」。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [Lobe UI CopyButton](https://mygo.lobehub.com/components/copy-button) | 点击写剪贴板、Jump/AnimateTo 的短暂反馈、对勾替换复制图标 |
| `controls/iconaction.go` | 纯图标按钮的共享命中区与状态 |
| `controls/actionicon.go` | 三种面的取色规则 |

## 2. 能否实现

- `c.WriteClipboard(content)` 是 ui 的公开 API ✓。
- 反馈时长由 `Element.Jump` + `AnimateTo` 的框架动画承担，组件没有计时器与 goroutine；`NonMotion: true` 让减少动态效果时仍保留语义反馈 ✓。
- 图标用 `NewIconAction` 建，复制态与已复制态分别用 `icons.Copy()` / `icons.Check()` ✓。
- 无障碍名随状态切换，可由 `Labels` 本地化 ✓。

## 3. 实现到何种程度

三种面、四档尺寸、点击复制、两秒对勾反馈、可替换的双态名字。反馈时长是框架动画的时长，不是组件自持的定时器。

## 4. 项目结构

- `controls/copybutton.go`；`controls/copybutton_test.go` 覆盖剪贴板内容、Copied 状态、状态名切换与本地化。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewCopyButton(c *ui.Context, content string) *CopyButton` | 建按钮，点击复制 `content` |
| `Variant(v ActionIconVariant)` · `Borderless()` · `Filled()` · `Outlined()` | 三种面 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 共享尺寸 |
| `Labels(copy, copied string)` | 复制态与已复制态的无障碍名 |
| `Copied() bool` | 对勾是否还在显示 |

## 6. 限制

- 不展示浮动提示；需要时用 `overlays.NewTooltip` 包裹。
- 两秒反馈时长固定为框架动画常量；应用若要别的时长需自行以 `AnimateTo` 组合。
