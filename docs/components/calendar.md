# Calendar 规格书

> 状态：**已实现**。my-ui 在我的 mygo/ui 底座上包了一层 shadcn 风格外观。

## 1. 参考实现
| 来源 | 可借鉴之处 |
|---|---|
| mygo/ui `Calendar` | 值指针、键盘与无障碍由框架负责 |
| shadcn/ui | 共享圆角、边框、底色与状态过渡 |

## 2. 能否实现
- 包 mygo 构造器，补 `Radius/Background/Border/Transition` 与共享 Variant/Size/Width ✓。

## 3. 实现到何种程度
- 月视图日历，包 mygo Calendar + 卡片外观

## 4. 项目结构
- forms/ 子包。

## 5. API 契约
见 Go doc（`NewCalendar` 及其 fluent 方法）。

## 6. 状态
- [x] 规格书 / 实现 / 测试 / showcase
