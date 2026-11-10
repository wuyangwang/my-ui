package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// PasswordInput is a text field that hides what it holds, with a button
// that shows it:
//
//	myui.NewPasswordInput(c, &password).Outline().Placeholder("密码")
//
// MyGo's text input hides what it holds for good: there is no public way
// to show it again. The button therefore shows the text in a field of
// its own — the same value, another field — and the two take turns.
// The value is the app's, so it survives the change; the state of the
// editor, the caret and the undo, does not.
type PasswordInput struct {
	field[*PasswordInput]
	shown *bool
}

// NewPasswordInput creates a password field editing *value with the
// native preset.
func NewPasswordInput(c *ui.Context, value *string) *PasswordInput {
	p := &PasswordInput{}
	p.field = *newField[*PasswordInput](c, value, func(c *ui.Context, value *string, host ui.Element) ui.Element {
		// The state of the button: whether the text shows. The host
		// keeps it from frame to frame.
		p.shown = ui.Local(host, "password-shown", func() bool { return false })
		if *p.shown {
			c.Key("password-shown")
			return ui.TextInput(c, value)
		}
		c.Key("password-hidden")
		return ui.TextInput(c, value).Password()
	})
	p.self = p

	// The button that shows and hides, at the end of the field.
	p.Suffix(func() {
		t := c.Theme()
		b := ui.ButtonBase(c).Size(18, 18).KeepFocus().Label("显示密码")
		b.Children(func() {
			ic := ui.Icon(c, icons.Eye())
			if *p.shown {
				ic = ui.Icon(c, icons.EyeOff())
			}
			ic = ic.Size(14, 14).TextColor(t.TextMuted)
			if b.Hovered() {
				ic = ic.TextColor(t.Text)
			}
		})
		if b.Clicked() {
			*p.shown = !*p.shown
		}
	})
	return p
}
