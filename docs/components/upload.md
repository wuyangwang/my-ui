# Upload 规格书

> 状态：**已实现**。上传：一片可点击、可拖入文件的区域，回调把文件路径交给 app。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui + dropzone | 虚线边框、图标 + 短说明、hover 高亮 |
| mygo/ui `Element.DroppedFiles()` / `FileDragOver()` | 文件拖入由框架给出路径，库不碰文件系统 |

## 2. 能否实现
- 虚线容器 + `DroppedFiles()` 回调 ✓；点击打开文件选择器需要 mygo 根包的 `OpenFile`，按 AGENTS 硬规则本包不导入它——因此 v1 只做拖放区，不提供点击选文件。
- 限制写进文档，不假装支持。

## 3. 实现到何种程度
- `UploadOptions{Accept func(path string) bool, OnFiles func(paths []string)}`；`NewUpload(c, o)`。
- hover/拖拽经过时用 Accent 淡底与强调色边框。

## 4. 项目结构
- `forms/` 子包。

## 5. API 契约
```go
type UploadOptions struct {
    Accept func(path string) bool
    OnFiles func(paths []string)
}
func NewUpload(c *ui.Context, o UploadOptions) ui.Element
```

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
