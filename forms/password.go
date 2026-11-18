package forms

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// PasswordInput is a text field that hides what it holds, with a button
// that shows it:
//
//	myui.NewPasswordInput(c, &password).Outline().Placeholder("密码")
//
// MyGo's Password() hides a text input in the frames that call it: an
// input that stops calling it shows its text again. The button therefore
// only flips a flag; the same editor is kept, so the caret, the selection
// and the undo history survive a show/hide.
type PasswordInput struct {
	Frame[*PasswordInput]
	shown    *bool
	eyeOn    *ui.SVG // the icon of the button while the text shows
	eyeOff   *ui.SVG // the icon of the button while the text is hidden
	eyeColor *ui.Color
}

// NewPasswordInput creates a password field editing *value with the
// native preset.
func NewPasswordInput(c *ui.Context, value *string) *PasswordInput {
	p := &PasswordInput{}
	p.Frame = *NewFrame[*PasswordInput](c, value, func(c *ui.Context, value *string, host ui.Element) ui.Element {
		// The state of the button: whether the text shows. The host
		// keeps it from frame to frame. The key is the same in both
		// states, so the editor — and with it the caret and the undo —
		// is not rebuilt when the button is pressed.
		p.shown = ui.Local(host, "password-shown", func() bool { return false })
		ed := ui.TextInput(c, value)
		if !*p.shown {
			ed = ed.Password()
		}
		return ed
	})
	p.Self = p

	// The button that shows and hides, at the end of the field.
	p.Suffix(func() {
		t := c.Theme()
		b := ui.ButtonBase(c).Size(18, 18).KeepFocus().Label(p.passwordLabel())
		// The icon is drawn when the frame is painted, so that the app's
		// own icon and color are read then.
		b.Draw(func(painter *ui.Painter, r ui.Rect) {
			icon := p.eyeIcon()
			color := p.eyeStateColor(t)
			if b.Hovered() {
				color = t.Text
			}
			painter.Icon(icon, ui.Rect{X: r.X + 2, Y: r.Y + 2, W: r.W - 4, H: r.H - 4}, color)
		})
		if b.Clicked() {
			*p.shown = !*p.shown
		}
	})
	return p
}

// Visible reports whether the text currently shows.
func (p *PasswordInput) Visible() bool { return p.shown != nil && *p.shown }

// EyeIcons sets the icons of the button: the one it shows while the
// text shows, and the one it shows while the text is hidden. It returns
// the field, for chaining. The icons of the set are the default.
func (p *PasswordInput) EyeIcons(shown, hidden *ui.SVG) *PasswordInput {
	p.eyeOn = shown
	p.eyeOff = hidden
	return p
}

// passwordLabel is the name of the button: what it does now, in the
// language of the frame.
func (p *PasswordInput) passwordLabel() string {
	if p.shown != nil && *p.shown {
		return Message(p.Context(), "hidePassword")
	}
	return Message(p.Context(), "showPassword")
}

// EyeColor sets the color of the button's icon, which is the muted
// tone of the theme by default. It returns the field, for chaining.
func (p *PasswordInput) EyeColor(c ui.Color) *PasswordInput {
	p.eyeColor = &c
	return p
}

// eyeIcon is the icon of the button in the state the field is in.
func (p *PasswordInput) eyeIcon() *ui.SVG {
	if *p.shown {
		if p.eyeOn != nil {
			return p.eyeOn
		}
		return icons.EyeOff()
	}
	if p.eyeOff != nil {
		return p.eyeOff
	}
	return icons.Eye()
}

// eyeStateColor is the color of the button's icon.
func (p *PasswordInput) eyeStateColor(t *ui.Theme) ui.Color {
	if p.eyeColor != nil {
		return *p.eyeColor
	}
	return t.TextMuted
}
