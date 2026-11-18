package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/forms"
)

// inputOTPage demonstrates a one-character-per-box code input.
func (s *showcase) inputOTPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "验证码", "输入一格自动跳到下一格，退格回上一格", func() {
		s.card(c, t, func() {
			forms.NewInputOTP(c, &s.otpValue, 6)
			ui.Text(c, "当前值："+s.otpValue).TextColor(t.TextMuted)
		})
	})
}
