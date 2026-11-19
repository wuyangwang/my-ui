# Upload 规格书

> 状态：**已实现**。上传：一片可点击、可拖入文件的区域，回调把文件路径交给 app。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| shadcn/ui + dropzone | 虚线边框、图标 + 短说明、hover 高亮 |
| mygo/ui `Element.DroppedFiles()` / `FileDragOver()` | 文件拖入由框架给出路径，库不碰文件系统 |

## 2. 能否实现
- 虚线容器 + `DroppedFiles()` 回调 ✓。
- 点击打开文件选择器需要 mygo 根包的 `OpenFile`，按 AGENTS 硬规则本库不导入根包——因此库内不直接调用系统选择器，改为注入：app 提供 `OnPick` 回调（在 app 代码里调用 `mygo.OpenFile`），库只接收路径。设置后多渲染一个「打开」按钮；未设置则保持纯拖放区。

## 3. 实现到何种程度
- `UploadOptions{Accept func(path string) bool, OnFiles func(paths []string), OnPick func() (path string, ok bool)}`；`NewUpload(c, o)`。
- 拖入的路径与 `OnPick` 选中的路径走同一条 `Accept` → `OnFiles` 通道；`ok == false`（取消）不回调。
- hover/拖拽经过时用 Accent 淡底与强调色边框。

## 4. 项目结构
- `forms/` 子包；按钮用 `controls.NewButton`（本库子包，无循环依赖）。

## 5. API 契约
```go
type UploadOptions struct {
    Accept  func(path string) bool          // 过滤拖入与选中的路径，nil 全收
    OnFiles func(paths []string)            // 接收通过过滤的路径（两条路径共用）
    OnPick  func() (path string, ok bool)   // 可选；提供时渲染「打开」按钮，app 在此调用 mygo.OpenFile
}
func NewUpload(c *ui.Context, o UploadOptions) ui.Element
```

「打开」按钮文案取根 i18n 键 `open`（en/zh-Hans/zh-Hant 三语齐全）。`mygo.OpenFile` 只存在于 app 代码，不属于本库。

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
