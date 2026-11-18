# SearchBar 规格书

> 状态：**已实现**。搜索框：放大镜图标 + 输入框 + 清除，用于筛选与搜索。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui Input with icon | 前置搜索图标、有值时显示清除 |
| Element Plus Input `prefix-icon` / `clearable` | 同样结构 |

## 2. 能否实现
- 复用 `forms` 的 Frame 基座：`PrefixIcon(icons.Search())` + `Clearable()`，值指针绑定 ✓。

## 3. 实现到何种程度
- `NewSearchBar(c, query *string)`；`.Placeholder(s)`、`.Width(w)`（默认占满）、`.Variant`、`.Size`。
- 清除图标用共享 `ctrlpkg.NewIconAction`（替换式 hover，不是按钮壳）。

## 4. 项目结构
- `forms/` 子包。

## 5. API 契约
```go
func NewSearchBar(c *ui.Context, query *string) *SearchBar
func (s *SearchBar) Placeholder(v string) *SearchBar
func (s *SearchBar) Width(w float32) *SearchBar
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
