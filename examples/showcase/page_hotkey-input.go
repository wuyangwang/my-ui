package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// hotkeyInputPage demonstrates the HotkeyInput component.
func (s *showcase) hotkeyInputPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "HotkeyInput", "聚焦后按键录制快捷键，修饰键与冲突会被检查", func() {
		s.card(c, t, func() {
			root := ui.Column(c).Gap(t.Space(3)).FillWidth()
			value := ui.Local(root, "hotkey-value", func() string { return "mod+k" })
			root.Children(func() {
				ctrlpkg.NewHotkeyInput(c, value).Reset("mod+k").AllowClear(true).
					Conflicts("mod+s").
					Messages("需要修饰键加一个按键", "这个快捷键已被占用")
				ui.Text(c, "当前快捷键："+*value).
					FontSize(myui.TextSize(t, myui.TextSm)).TextColor(t.TextMuted)
			})
		})
	})
}
