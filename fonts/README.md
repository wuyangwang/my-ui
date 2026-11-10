# fonts（二级包）

> **状态：规划中，尚无代码。**（不参与 `go build ./...` 的包集合；实现后补充 API 与示例）

内嵌常见开源字体，`init` 时经 `ui.RegisterFont(data []byte, family string)` 注册字族，注册后 app 直接按字族名使用（`ui.Font{Family: ...}`、`Theme.Font`）。

## 资产与许可

只收录 OFL（SIL Open Font License）字体——允许再分发、允许内嵌；每引入一款在此登记字体名称、版本、来源与许可副本。

候选清单（待一期定稿）：Inter（正文）、JetBrains Mono（等宽）、Noto Sans SC / Noto Sans TC（中文场景）等。

## 体积原则

- 全部 `go:embed` 在本包内，不 import 不链接：主打系统字体或其他字体的 app 不承担任何字体字节。
- 中文字体体积较大时提供按字体文件拆分的子目录（如 `fonts/notosc/`），进一步按需引入。

## 常量

字族名常量（`fonts.Sans`、`fonts.Mono` 等）与字号刻度（配合 `Theme.Rem`）随代码落地归于本包或根包 tokens，具体在实现时定。
