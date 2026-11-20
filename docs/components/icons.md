# icons（二级包）规格书

## API 契约

```go
package icons

// Save returns the Lucide "save" icon.
func Save() *ui.SVG

// Close is an alias of X.
func Close() *ui.SVG
```

```go
package myui

// NewIcon creates an icon with the native preset.
func NewIcon(c *ui.Context, svg *ui.SVG) *Icon

func (i *Icon) Variant(v Variant) *Icon   // Primary/Destructive/Link 给语义色
func (i *Icon) Size(s Size) *Icon         // SizeSm/Md/Lg → 14/20/24 DIP
func (i *Icon) Sm() *Icon
func (i *Icon) Md() *Icon
func (i *Icon) Lg() *Icon
func (i *Icon) Label(s string) *Icon      // 无障碍名
func (i *Icon) Spin() *Icon               // 旋转（ui 的 Loop/Rotate）
```

用法：

```go
ctrlpkg.NewIcon(c, icons.Save()).Md().Label("保存")
```
