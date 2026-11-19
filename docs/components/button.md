# Button 规格书

> 位置：组件在 [`controls/`](../../controls) 子包；主题/tokens 仍在根包 `myui`。

> 状态：**已实现**（`button.go`）。shadcn 式变体词汇的样板：后来的控件照它的形状做。

## 1. 参考实现

| 来源 | 可借鉴之处 |
|---|---|
| [shadcn/ui Button](https://ui.shadcn.com/docs/components/button) | 六个变体（default/secondary/outline/ghost/destructive/link）与三档尺寸；变体只管「长什么样」，尺寸是另一根轴 |
| [Element Plus el-button](https://element-plus.org/zh-CN/component/button) | `type`（primary/success/warning/danger/info）、`plain`/`round`/`circle`/`text`、尺寸档；事件 `click` |
| mygo 的 `ui.Button` / `ui.PrimaryButton` | 立即模式的绑定约定：`Clicked()` 在下一句可读；主题自带的观感（native look） |

## 2. 能否实现

- 构造：`ui.Button(c, label)` ✓；主题槽位（Accent/Surface/Border/Radius/Space/Rem）齐全 ✓。
- 状态：`ui.Element.Pressed()/Hovered()/IsDisabled()` ✓，每帧在链上重涂即可跟随主题变化。
- 颜色：走 `shades.go`（hover/pressed 的唯一派生入口），不自造颜色 ✓。
- **图标**：mygo 的按钮本来就是「行 + 文字子节点」（`ui/widgets.go:140-147`：`b.Children(func(){ coreText(c, label).SingleLine() })`），因此图标可作为兄弟子节点并排；`ui.Button(c, "")` + `Children` 是官方文档给出的「空标签 + 自定义内容」写法 ✓。图标颜色随按钮文字色（`currentColor`）✓，尺寸随字号 ✓。
- **忙碌态**：spinner 用 `ui.Icon(c, svg)` + `Loop("spin", …) * 360` + `Rotate`（框架文档的旋转图标写法）✓。
- 无障碍：按钮的文字即其无障碍名（`access.go:884` 按内部文字取名）✓；只有图标时须显式 `Label("…")` ✓。
- 不做依赖 `Element.Access` 的定制（internal 泄漏）。

## 3. 实现到何种程度

**已实现**：`Native`（零值，保留内置观感）/`Secondary`（别名）/`Primary`/`Outline`/`Ghost`/`Destructive`/`Link` × `SizeSm/Md/Lg`；hover/pressed 三态；禁用态沿用主题的暗淡；`.Clicked()` 沿用 ui。
**图标与忙碌态**：`Icon(svg)`（文字之前）、`EndIcon(svg)`（文字之后）、只有图标的按钮（`NewButton(c, "")` + `Label`）、`Busy(on)`（转圈且不响应点击，`on` 为 false 时恢复）、`Loading()`（`Busy(true)` 的简写）。图标与文字之间的距离由图标自身的 margin 承担，故没有图标的按钮不留空隙。

**成组与双态**：`NewButtonGroup(c, build)` 把按钮排成共享边的一组；`Block()` 占满整行；`Toggle(pressed)` 做双态按钮（点击翻转 `*pressed`，无障碍名带按下态）。互斥/分段按钮用 `controls.NewToggleGroup` 与 `controls.NewSegmented`（mygo 签名 + my-ui 外观）。加载态的骨架（spinner 之外的形态）暂不做。

## 4. 项目结构是否合理

- 位置：根包 L3；`Variant`/`Size` 为**根包共享词表**（`input.go`/`badge.go`/`icon.go` 共用），组件特有的值留在组件文件里。
- 样式归属：边框与底色画在按钮自身（输入框是复合件，边框在外框——见 [input.md](input.md)）。

## 5. API 契约

| 方法 | 作用 |
|---|---|
| `NewButton(c *ui.Context, label string) *Button` | 创建按钮（`label` 留空即只有图标） |
| `Variant(v Variant)` · `Default()` · `Native()` · `Secondary()` · `Primary()` · `Outline()` · `Ghost()` · `Destructive()` · `Link()` · `Success()` · `Info()` · `Warning()` | 变体：底色、边框、文字色与 hover/pressed 三态 |
| `Size(s Size)` · `Xs()` · `Sm()` · `Md()` · `Lg()` | 尺寸：内衬与字号 |
| `Shape(s Shape)` | 形态：`ShapeRect`（默认）/ `ShapeRound`（胶囊）/ `ShapeCircle`（圆形） |
| `Icon(svg *ui.SVG)` | 文字之前的图标（取文字色与字号） |
| `EndIcon(svg *ui.SVG)` | 文字之后的图标 |
| `Busy(on bool)` | 忙碌态：`on` 时转圈并停止响应点击，传 false 恢复图标与点击 |
| `Loading()` | 忙碌态简写，等同 `Busy(true)` |
| `Toggle(pressed *bool)` | 双态按钮（shadcn Toggle）：点击翻转 `*pressed`，无障碍名带按下态 |
| `NewButtonGroup(c *ui.Context, build func()) ui.Element` | 按钮组：把若干按钮排成共享边的一行 |
| `NewToggleGroup` · `NewSegmented` | mygo 签名的按钮组与分段控件，现由 [toggle-group.md](toggle-group.md) / [segmented.md](segmented.md) 各自带 my-ui 外观 |
| `Label(s string)` | 无障碍名（只有图标时必须给） |
| `Width(w float32)` | 宽度（DIP） |
| `Block()` | 占满整行 |

变体词表（与 shadcn 对齐）：

| 变体 | 观感 |
|---|---|
| `Default`（零值） | 控件自带的观感：主题的面 + 描边（mygo 的按钮外观） |
| `Secondary` | 弱填充面：主题的面，无描边 |
| `Primary` | 强调色填充 |
| `Outline` | 背景上的描边面 |
| `Ghost` | 平时无底，悬停才现 |
| `Destructive` | 危险色填充 |
| `Link` | 纯文字，强调色 |
| `Success`/`Info`/`Warning` | 语义色填充（主题层的三种语义色，明暗各一套） |

`Native` 是 `Default` 的旧名，保留以免旧代码失效。

用法：

```go
ctrlpkg.NewButton(c, "保存").Primary()
ctrlpkg.NewButton(c, "删除").Destructive().Sm()
ctrlpkg.NewButton(c, "保存").Icon(icons.Save()).Primary()
ctrlpkg.NewButton(c, "").Icon(icons.Settings()).Label("设置")   // 只有图标
ctrlpkg.NewButton(c, "筛选").Outline().Shape(myui.ShapeRound)
```

## 6. 状态

- [x] 调研与实现
- [x] 测试（`button_test.go`：变体绘制与 hover/pressed、尺寸、禁用、图标在文字之前、只有图标的按钮有名字、忙碌态不响应点击）
- [x] showcase「Button」页（含图标与忙碌态）
- [x] 按钮组、块级按钮、双态按钮（Toggle）、忙碌态（Busy）
