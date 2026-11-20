# Upload 规格书

上传：一片可点击、可拖入文件的区域，回调把文件路径交给 app。

## API 契约

```go
type UploadOptions struct {
    Accept  func(path string) bool          // 过滤拖入与选中的路径，nil 全收
    OnFiles func(paths []string)            // 接收通过过滤的路径（两条路径共用）
    OnPick  func() (path string, ok bool)   // 可选；提供时渲染「打开」按钮，app 在此调用 mygo.OpenFile
}
func NewUpload(c *ui.Context, o UploadOptions) ui.Element
```

「打开」按钮文案取根 i18n 键 `open`（en/zh-Hans/zh-Hant 三语齐全）。`mygo.OpenFile` 只存在于 app 代码，不属于本库。
