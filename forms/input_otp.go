package forms

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// InputOTP is one box per digit. The value pointer is the app's; cells
// are kept on the frame's context, so typing advances focus without a
// component-level timer or goroutine.
type InputOTP struct {
	ui.Element
}

// NewInputOTP builds digits boxes editing *value. A cell takes one
// character; entering one moves the focus to the next, and Backspace in
// an empty cell moves back.
func NewInputOTP(c *ui.Context, value *string, digits int) *InputOTP {
	if digits < 1 {
		digits = 1
	}
	t := c.Theme()
	host := ui.Row(c.Key("input otp")).Gap(t.Space(2)).AlignItems(ui.Center)
	cells := ui.Local(host, "cells", func() []string { return make([]string, digits) })
	next := ui.Local(host, "next", func() int { return 0 })
	if len(*cells) != digits {
		*cells = make([]string, digits)
	}

	host.Children(func() {
		for i := range *cells {
			dest := &(*cells)[i]
			input := ui.TextInput(c, dest).MaxLength(1).
				Width(t.Space(10)).Height(t.Space(10)).
				FontWeight(600).
				Border(1, t.Border).Radius(Radius(t, RadiusMd))
			input.OnTextKey(func(k ui.TextKey) bool {
				if k.Key == ui.KeyBackspace && *dest == "" && i > 0 {
					*next = i - 1
				}
				return false
			})
			if *dest != "" && i < digits-1 {
				*next = i + 1
			}
			if *next == i {
				input = input.Focus()
			}
		}
	})

	// The app's value is the digits joined, without the box boundaries.
	joined := ""
	for _, cell := range *cells {
		joined += cell
	}
	if value != nil {
		*value = joined
	}
	return &InputOTP{Element: host}
}

// Width overrides the width of each box. The next frame applies it.
func (o *InputOTP) Width(w float32) *InputOTP {
	o.Element = o.Element.Width(w)
	return o
}
