# DecimalInput 规格书

> 状态：**已实现**。包 mygo/ui 的同名控件，补 my-ui 的主题外观。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `DecimalInput` | 行为、键盘与无障碍 |
| shadcn/ui | 边框、圆角、状态色 |

## 2. 能否实现
- 包一层并把 `Radius/Background/Border/Transition` 来自主题 ✓。

## 3. 实现到何种程度
- 见 Go doc 的 `NewDecimalInput` 及其 fluent 方法。

## 4. 项目结构
- forms/ 子包。

## 5. 状态
- [x] 规格书 / 实现 / 测试 / showcase
