# ai（二级包）

> **状态：规划中，尚无代码。**（不参与 `go build ./...` 的包集合；实现后补充 API 与示例）

AI 应用常见组件。mygo 独有的底层能力：`ui/markdown` 专为 LLM 逐 token 流式回复构建（增量解析、半成品语法按终态显示、文本淡入），`ui.List` 为锚点式虚拟化列表（行高动态、行可增删）。本库把两者的组合包装成可直接用的应用组件。

## 组件

- **ChatList**：虚拟化消息气泡布局（含头像、系统/用户角色），流式 `markdown.Stream` / 完结转 `markdown.View`，typing 指示，贴底跟随（`ScrollState` `Y ≥ MaxY → MaxFloat32`）、上滚即停、弹簧回底；底部输入条经 Scaffold 的底部预留机制不遮挡最后一行。
- **ModelPicker**：分组 `Combobox` + 厂商图标（可选联动 `icons/`）+ 上下文窗/速度/价格徽标 + `ComboboxParts.Chosen` 选中态。
- **Reasoning**：推理档位（low/medium/high/…）基于 `Segmented`；temperature/top_p 基于 `RangeSlider`/`Stepper`，带语义预设（精确/均衡/创意）。

## 约定

- 无 `go:embed` 资产；仅依赖 `myui` 根包与 mygo 公开 API，不改 mygo 行为。
- 中英文文案（如有）计划遵守根包后置的 i18n 文案表与回退链（见 roadmap）。
- 不持有的状态：消息数据由 app 持有（立即模式），本包只描述如何渲染与滚动。
