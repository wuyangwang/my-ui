# icons/brand（二级包）

Simple Icons 品牌标集，给 MyGo 原生 UI 应用提供公司与项目的 logo。**只依赖 `github.com/egoist/mygo/ui`，不依赖根包 `myui`**——不 import 这个包的应用，二进制里没有它的任何字节。

## 用法

每个品牌标一个函数，返回解析好的 `*ui.SVG`：

```go
import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons/brand"
)

ui.Icon(c, brand.GitHub())
```

函数名取品牌自己的拼写（`GitHub`、`OpenAI`、`Anthropic`、`GitLab`），同名的只留一个。

## 体积

SVG 内嵌在各自函数里的变量中，**程序没调用的品牌标不会进入二进制**。约 1.4 KB 一个品牌标。

## 生成

```sh
go generate ./icons/brand
```

生成器读 `svg/` 下的 SVG，写出 `icons_gen.go`；`data.json` 是上游的元数据（名字、颜色），生成器据此取函数名。

## 资产与许可

- 品牌标：[Simple Icons](https://simpleicons.org) 16.37.0，`svg/` 下为其原文件。
- 许可：[CC0 1.0](LICENSE.md)（公有领域奉献）。
- 品牌标各自仍受其商标与品牌规范约束；Simple Icons 只保证图标本身的著作权已按 CC0 放弃。
