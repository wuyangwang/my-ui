# 动效规格书

跨组件的共用约定，不是一个组件。

## API 契约

```go
// stateTransition is the transition the presets use: the colors of a
// state change fade, and the rest happens at once.
func stateTransition() ui.ElementTransition
```

（包内约定，不导出；对外的动效入口是各组件自己的方法与 `ctrlpkg.Icon.Spin`。）
