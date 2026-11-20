# Kbd 规格书

> 位置：组件在 [`display/`](../../display) 子包；主题/tokens 仍在根包 `myui`。

## API 契约

```go
func NewKbd(c *ui.Context, keys ...string) *Kbd
```

## 状态与限制

**限制**

- 键帽之间没有「+」之类的连接符（app 自己写）。
