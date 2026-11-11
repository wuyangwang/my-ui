# icons（二级包）

Lucide 图标集，给 MyGo 原生 UI 应用提供图标。**只依赖 `github.com/egoist/mygo/ui`，不依赖根包 `myui`**——不 import 这个包的应用，二进制里没有它的任何字节。

## 用法

每个图标一个函数，返回解析好的 `*ui.SVG`：

```go
import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

ui.Icon(c, icons.Save()).FontSize(20)
```

在根包用 `myui.NewIcon(c, icons.Save()).Md()` 可以得到带尺寸档、语义色与无障碍名的图标。

## 三个包

| 包 | 内容 | 体积 |
|---|---|---|
| `icons` | Lucide 2,318 个图标，函数式 API | 只带用到的图标 ✓ |
| `icons/brand` | Simple Icons 3,462 个品牌标，函数式 API | 同上 ✓ |
| `icons/all` | 整套图标的**名字清单**与按名字取图标的函数 | 带上整套图标 ✗（只有图标选择器才用） |

```go
import "github.com/wuyangwang/my-ui/icons/all"

all.Names             // 全部图标名
all.ByName("Save")    // 按名字取图标
```

## 名字

- Lucide：上游 kebab 名字转 Go Pascal 名（`arrow-up-right.svg` → `ArrowUpRight()`）；同义的两个名字只留一个；另有常用别名（`icons.Close()` → `x`、`icons.Trash()` → `trash-2`），见 `gen/aliases.go`。
- 品牌标：取品牌自己的拼写（`GitHub`、`OpenAI`、`Anthropic`）；同名的只留一个。

## 体积

图标 SVG 内嵌在各自函数里的变量中，**程序没调用的图标不会进入二进制**（链接器连同函数一起丢弃）。实测（Windows amd64，与只导入 `mygo/ui` 的对照程序相比）：

| 用到的图标数 | 二进制增加 |
|---|---|
| 0（import 但未调用） | +512 B |
| 1 | +4.1 KB |
| 30 | +27.6 KB |

约 850 B 一个图标。整套 2,318 个都在二进制里时约 2 MB。SVG 只在第一次调用时解析，之后走缓存。

## 生成

图标函数由生成器写出，不要手改生成的文件。

```sh
go generate ./icons        # Lucide
go generate ./icons/brand  # 品牌标
```

生成器读 `svg/` 下的 SVG，写出 `icons_gen.go`（每个图标一个函数）、`aliases_gen.go`（别名）与 `all/all_gen.go`（名字清单）。新增或升级图标集的步骤：

1. 从上游取新的 `icons/*.svg`（原样入库，不改内容，便于比对与升级）到 `svg/`；
2. 在 `gen/aliases.go` 里增删别名；
3. 跑 `go generate ./icons`；
4. 在本文件的「引入记录」里记一笔。

## 资产与许可

- 图标：[Lucide](https://lucide.dev) 1.70.0，`svg/` 下为其原文件（ISC 许可）。
- 品牌标：[Simple Icons](https://simpleicons.org) 16.37.0，`brand/svg/` 下为其原文件（CC0 许可）。
- 许可全文见 [LICENSE](LICENSE)（Lucide，含源自 Feather 的图标所用的 MIT 段落）与 [brand/LICENSE.md](brand/LICENSE.md)（CC0 1.0）。

## 引入记录

| 日期 | 版本 | 内容 | 许可 |
|---|---|---|---|
| 2026-11-10 | 1.70.0 | 首次引入，2,322 个上游 SVG 文件（生成后 2,318 个图标函数，4 个同义名合并），另加 14 个常用别名 | ISC（含 Feather MIT 段落） |
| 2026-11-10 | 16.37.0 | 品牌标：3,470 个上游 SVG 文件（生成后 3,462 个图标函数，8 个同名合并） | CC0 1.0 |
