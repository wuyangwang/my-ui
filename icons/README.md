# icons（二级包）

内嵌图标集，导出 `*ui.SVG` 常量（typed 注册表：`icons.Save` 等，编译期保证名字存在，不用字符串查表）。

## 变色机制（框架原生）

- `stroke="currentColor"` 的图标随 `TextColor` 自动着色：`ui.Icon(c, icons.Save).TextColor(c.Theme().Danger)`，甚至 `AnimateColor` 悬停变色；
- 多色 SVG（品牌标）保留原文件的颜色；
- 尺寸走 `FontSize`/`Size`，旋转走 `Rotate`/`Loop`（旋转指示器不用另做）。

## 资产与许可

- Lucide（ISC）：全集或子集，实现为 `go generate` 从 SVG 目录产出 Go 常量文件，生成产物提交入库；
- simple-icons（CC0）：常见科技厂商标，多色场景下用于填充色封面/列表。
- 记录：每加入一批图标，在此登记来源目录与许可。

## 边界

- 不含 SF Symbols（macOS 专属），如需由 app 直接引入 `plugins/sfsymbols` 的公开 API，本库可选适配层不承诺跨平台。
- `go:embed` 只出现在本包；根包永不进入本包。
